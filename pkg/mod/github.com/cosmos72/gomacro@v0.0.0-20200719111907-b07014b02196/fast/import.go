/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * import.go
 *
 *  Created on Apr 02, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	r "reflect"
	"strconv"
	"strings"
	"unsafe"

	"github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/base/genimport"
	"github.com/cosmos72/gomacro/base/output"
	"github.com/cosmos72/gomacro/base/paths"
	"github.com/cosmos72/gomacro/base/reflect"
	"github.com/cosmos72/gomacro/base/untyped"
	xr "github.com/cosmos72/gomacro/xreflect"
)

// =========================== forget package ==================================

// remove package 'path' from the list of known packages.
// later attempts to import it again will trigger a recompile.
func (cg *CompGlobals) UnloadPackage(path string) {
	cg.Globals.UnloadPackage(path)
	delete(cg.KnownImports, path)
}

// ========================== switch to package ================================

func (ir *Interp) ChangePackage(name, path string) {
	if len(path) == 0 {
		path = name
	} else {
		name = paths.FileName(path)
	}
	c := ir.Comp
	if path == c.Path {
		return
	}
	// load requested package if it exists, but do not define any binding in current one
	newp, err := c.ImportPackageOrError("_", path)
	if err != nil {
		c.Debugf("%v", err)
	}
	oldp := ir.asImport()

	c.CompGlobals.KnownImports[oldp.Path] = oldp // overwrite any cached import with same path as current Interp

	trace := c.Globals.Options&base.OptShowPrompt != 0
	top := &Interp{c.TopComp(), ir.env.Top()}
	if newp != nil {
		newp.Name = name
		*ir = newp.asInterpreter(top)
		if trace {
			c.Debugf("switched to package %v", newp)
		}
	} else {
		// requested package does not exist - create an empty one
		ir.Comp = NewComp(top.Comp, nil)
		ir.env = NewEnv(top.env, 0, 0)
		if c.Globals.Options&base.OptDebugger != 0 {
			ir.env.DebugComp = ir.Comp
		}
		ir.Comp.Name = name
		ir.Comp.Path = path
		if trace {
			c.Debugf("switched to new package %v", path)
		}
	}
	// env is at file/package level => its FileEnv is itself
	ir.env.FileEnv = ir.env
	ir.env.Run.Globals.PackagePath = path
}

// convert *Interp to *Import. used to change package from 'ir'
func (ir *Interp) asImport() *Import {
	env := ir.env
	env.MarkUsedByClosure() // do not try to recycle this Env
	return &Import{
		CompBinds: ir.Comp.CompBinds,
		EnvBinds:  &ir.env.EnvBinds,
		env:       env,
	}
}

// convert *Import to *Interp. used to change package to 'imp'
func (imp *Import) asInterpreter(outer *Interp) Interp {
	c := NewComp(outer.Comp, nil)
	c.CompBinds = imp.CompBinds
	env := imp.env
	// preserve env.IP, env.Code[], env.DebugPos[]
	if env.Outer == nil {
		env.Outer = outer.env
	}
	env.Run = outer.env.Run
	return Interp{c, env}
}

// =========================== import package =================================

// ImportPackage imports a package. Panics if the import fails.
// If alias is the empty string, it defaults to the identifier
// specified in the package clause of the imported package
func (ir *Interp) ImportPackage(alias, path string) *Import {
	return ir.Comp.ImportPackage(alias, path)
}

// ImportPackageOrError imports a package.
// If alias is the empty string, it defaults to the identifier
// specified in the package clause of the imported package
func (ir *Interp) ImportPackageOrError(alias, path string) (*Import, error) {
	return ir.Comp.ImportPackageOrError(alias, path)
}

// ImportPackage imports a package. Panics if the import fails.
// Usually invoked as Comp.FileComp().ImportPackage(alias, path)
// because imports are usually top-level statements in a source file.
// But we also support local imports, i.e. import statements inside a function or block.
func (c *Comp) ImportPackage(alias, path string) *Import {
	imp, err := c.ImportPackageOrError(alias, path)
	if err != nil {
		panic(err)
	}
	return imp
}

// ImportPackageOrError imports a package.
// If name is the empty string, it defaults to the identifier
// specified in the package clause of the imported package
func (c *Comp) ImportPackageOrError(alias, path string) (*Import, error) {
	g := c.CompGlobals
	imp := g.KnownImports[path]
	if imp == nil {
		pkgref, err := g.Importer.ImportPackageOrError(
			alias, path, g.Options&base.OptModuleImport != 0)
		if err != nil {
			return nil, err
		}
		imp = g.NewImport(pkgref)
	}
	if alias == "." {
		c.declDotImport0(imp)
	} else if alias != "_" {
		// https://golang.org/ref/spec#Package_clause states:
		// If the PackageName is omitted, it defaults to the identifier
		// specified in the package clause of the imported package
		if len(alias) == 0 {
			alias = imp.Name
		}
		c.declImport0(alias, imp)
	}
	g.KnownImports[path] = imp
	return imp, nil
}

// Import compiles an import statement
func (c *Comp) Import(node ast.Spec) {
	switch node := node.(type) {
	case *ast.ImportSpec:
		str := node.Path.Value
		path, err := strconv.Unquote(str)
		if err != nil {
			c.Errorf("error unescaping import path %q: %v", str, err)
		}
		path = c.sanitizeImportPath(path)
		var name string
		if node.Name != nil {
			name = node.Name.Name
		}
		// yes, we support local imports
		// i.e. a function or block can import packages
		c.ImportPackage(name, path)
	default:
		c.Errorf("unimplemented import: %v", node)
	}
}

func (g *CompGlobals) sanitizeImportPath(path string) string {
	path = strings.Replace(path, "\\", "/", -1)
	l := len(path)
	if path == ".." || l >= 3 && (path[:3] == "../" || path[l-3:] == "/..") || strings.Contains(path, "/../") {
		g.Errorf("invalid import %q: contains \"..\"", path)
	}
	if path == "." || l >= 2 && (path[:2] == "./" || path[l-2:] == "/.") || strings.Contains(path, "/./") {
		g.Errorf("invalid import %q: contains \".\"", path)
	}
	return path
}

// declDotImport0 compiles an import declaration.
// Note: does not loads proxies, use ImportPackage for that
func (c *Comp) declImport0(name string, imp *Import) {
	// treat imported package as a constant,
	// because to compile code we need the declarations it contains:
	// importing them at runtime would be too late.
	bind := c.NewBind(name, ConstBind, c.TypeOfPtrImport())
	bind.Value = imp // Comp.Binds[] is a map[string]*Bind => changes to *Bind propagate to the map
}

// declDotImport0 compiles an import . "path" declaration, i.e. a dot-import.
// Note: does not loads proxies, use ImportPackage for that
func (c *Comp) declDotImport0(imp *Import) {
	// Note 2: looking at the difference between the above Comp.declImport0() and this ugly monster,
	// shows one more reason why dot-imports are dirty and discouraged.
	if c.Types == nil {
		c.Types = make(map[string]xr.Type)
	}
	for name, typ := range imp.Types {
		if t, exists := c.Types[name]; exists {
			c.Warnf("redefined type: %v", t)
		}
		c.Types[name] = typ
	}

	var indexv, cindexv []int // mapping between Import.Vals[index] and Env.Vals[cindex]

	var funv []func(*Env) xr.Value
	var findexv []int

	for name, bind := range imp.Binds {
		// use c.CompBinds.NewBind() to prevent optimization VarBind -> IntBind
		// also, if class == IntBind, we must preserve the address of impenv.Ints[idx]
		// thus we must convert it into a VarBind (argh!)
		class := bind.Desc.Class()
		if class == IntBind {
			class = VarBind
		}
		cbind := c.CompBinds.NewBind(&c.Output, name, class, bind.Type)
		cidx := cbind.Desc.Index()
		switch bind.Desc.Class() {
		case ConstBind:
			cbind.Value = bind.Value
		case IntBind:
			if cidx == NoIndex {
				continue
			}
			// this is painful. and slow
			fun := imp.intPlace(c, bind, PlaceSettable).Fun
			funv = append(funv, fun)
			findexv = append(findexv, cidx)
		default:
			if cidx == NoIndex {
				continue
			}
			indexv = append(indexv, bind.Desc.Index())
			cindexv = append(cindexv, cidx)
		}
	}
	if len(indexv) != 0 || len(funv) != 0 {
		impvals := imp.Vals
		c.append(func(env *Env) (Stmt, *Env) {
			for i, index := range indexv {
				env.Vals[cindexv[i]] = impvals[index]
			}
			for i, fun := range funv {
				env.Vals[findexv[i]] = fun(nil) // fun(env) is unnecessary
			}
			env.IP++
			return env.Code[env.IP], env
		})
	}
}

func (g *CompGlobals) NewImport(pkgref *genimport.PackageRef) *Import {
	env := &Env{
		UsedByClosure: true, // do not try to recycle this Env
	}
	imp := &Import{
		EnvBinds: &env.EnvBinds,
		env:      env,
	}
	if pkgref != nil {
		imp.Name = pkgref.Name
		imp.Path = pkgref.Path
		imp.loadTypes(g, pkgref)
		imp.loadBinds(g, pkgref)
		g.loadProxies(pkgref.Proxies, imp.Types)
	}
	return imp
}

func (imp *Import) loadBinds(g *CompGlobals, pkgref *genimport.PackageRef) {
	vals := make([]xr.Value, len(pkgref.Binds))
	untypeds := pkgref.Untypeds
	o := &g.Output
	for name, val := range pkgref.Binds {
		if untyped, ok := untypeds[name]; ok {
			untypedlit, typ := g.parseUntyped(untyped)
			if typ != nil {
				bind := imp.CompBinds.NewBind(o, name, ConstBind, typ)
				bind.Value = untypedlit
				continue
			}
		}
		k := val.Kind()
		class := FuncBind
		// distinguish typed constants, variables and functions
		if val.IsValid() && val.CanAddr() && val.CanSet() {
			class = VarBind
		} else if k == r.Invalid || (reflect.IsOptimizedKind(k) && val.CanInterface()) {
			class = ConstBind
		}
		typ := g.Universe.FromReflectType(val.Type())
		bind := imp.CompBinds.NewBind(o, name, class, typ)
		if class == ConstBind && k != r.Invalid {
			bind.Value = val.Interface()
		}
		idx := bind.Desc.Index()
		if idx == NoIndex {
			continue
		}
		if len(vals) <= idx {
			tmp := make([]xr.Value, idx*2)
			copy(tmp, vals)
			vals = tmp
		}
		vals[idx] = xr.MakeValue(val)
	}
	imp.Vals = vals
}

func (g *CompGlobals) parseUntyped(untypedstr string) (UntypedLit, xr.Type) {
	kind, value := untyped.Unmarshal(untypedstr)
	if kind == untyped.None {
		return UntypedLit{}, nil
	}
	lit := untyped.MakeLit(kind, value, &g.Universe.BasicTypes)
	return lit, g.TypeOfUntypedLit()
}

func (imp *Import) loadTypes(g *CompGlobals, pkgref *genimport.PackageRef) {
	v := g.Universe
	types := make(map[string]xr.Type)
	wrappers := pkgref.Wrappers
	for name, rtype := range pkgref.Types {
		// Universe.FromReflectType uses cached *types.Package if possible
		t := v.FromReflectType(rtype)
		if twrappers := wrappers[name]; len(twrappers) != 0 {
			t.RemoveMethods(twrappers, "")
		}
		types[name] = t
	}
	imp.Types = types
}

// loadProxies adds to thread-global maps the proxies found in import
func (g *CompGlobals) loadProxies(proxies map[string]r.Type, types map[string]xr.Type) {
	for name, proxy := range proxies {
		g.loadProxy(name, proxy, types[name])
	}
}

// loadProxy adds to thread-global maps the specified proxy that allows interpreted types
// to implement an interface
func (g *CompGlobals) loadProxy(name string, proxy r.Type, xtype xr.Type) {
	if proxy == nil && xtype == nil {
		g.Errorf("cannot load nil proxy")
		return
	}
	if xtype == nil {
		g.Warnf("import %q: type not found for proxy <%v>", proxy.PkgPath(), proxy)
		return
	}
	if xtype.Kind() != r.Interface {
		g.Warnf("import %q: type for proxy <%v> is not an interface: %v", proxy.PkgPath(), proxy, xtype)
		return
	}
	if proxy == nil {
		g.Errorf("import %q: nil proxy for type <%v>", xtype.PkgPath(), xtype)
		return
	}
	rtype := xtype.ReflectType()
	g.interf2proxy[rtype] = proxy
	g.proxy2interf[proxy] = xtype
}

// ======================== use package symbols ===============================

// selectorPlace compiles pkgname.varname returning a settable and/or addressable Place
func (imp *Import) selectorPlace(c *Comp, name string, opt PlaceOption) *Place {
	bind, ok := imp.Binds[name]
	if !ok {
		c.Errorf("package %v %q has no symbol %s", imp.Name, imp.Path, name)
	}
	class := bind.Desc.Class()
	if bind.Desc.Index() != NoIndex {
		switch class {
		case IntBind:
			return imp.intPlace(c, bind, opt)
		case VarBind:
			// optimization: read imp.Vals[] at compile time:
			// val remains valid even if imp.Vals[] is reallocated
			val := imp.Vals[bind.Desc.Index()]
			// a settable reflect.Value is always addressable.
			// the converse is not guaranteed: unexported fields can be addressed but not set.
			// see implementation of reflect.Value.CanAddr() and reflect.Value.CanSet() for details
			if val.IsValid() && val.CanAddr() && val.CanSet() {
				return &Place{
					Var: Var{Type: bind.Type},
					Fun: func(*Env) xr.Value {
						return val
					},
					Addr: func(*Env) xr.Value {
						return val.Addr()
					},
				}
			}
		}
	}
	c.Errorf("%v %v %v.%v", opt, class, bind.Type.Kind(), imp.Name, name)
	return nil
}

// selector compiles foo.bar where 'foo' is an imported package
func (imp *Import) selector(name string, st *output.Stringer) *Expr {
	bind, ok := imp.Binds[name]
	if !ok {
		st.Errorf("package %v %q has no symbol %s", imp.Name, imp.Path, name)
	}
	switch bind.Desc.Class() {
	case ConstBind:
		return exprLit(bind.Lit, bind.AsSymbol(0))
	case FuncBind, VarBind:
		return imp.symbol(bind, st)
	case IntBind:
		return imp.intSymbol(bind, st)
	default:
		st.Errorf("package symbol %s.%s has unknown class %s", imp.Name, name, bind.Desc.Class())
		return nil
	}
}

// create an expression that will return the value of imported variable described by bind.
//
// mandatory optimization: for basic kinds, unwrap reflect.Value
func (imp *Import) symbol(bind *Bind, st *output.Stringer) *Expr {
	idx := bind.Desc.Index()
	if idx == NoIndex {
		st.Errorf("undefined identifier %s._", imp.Name)
	}
	// optimization: read imp.Vals[] at compile time:
	// v remains valid even if imp.Vals[] is reallocated
	v := imp.Vals[idx]
	t := bind.Type
	if !v.IsValid() {
		return exprValue(t, xr.Zero(t).Interface())
	}
	var fun I
	switch t.Kind() {
	case xr.Bool:
		fun = func(*Env) bool {
			return v.Bool()
		}
	case xr.Int:
		fun = func(*Env) int {
			return int(v.Int())
		}
	case xr.Int8:
		fun = func(*Env) int8 {
			return int8(v.Int())
		}
	case xr.Int16:
		fun = func(*Env) int16 {
			return int16(v.Int())
		}
	case xr.Int32:
		fun = func(*Env) int32 {
			return int32(v.Int())
		}
	case xr.Int64:
		fun = func(*Env) int64 {
			return v.Int()
		}
	case xr.Uint:
		fun = func(*Env) uint {
			return uint(v.Uint())
		}
	case xr.Uint8:
		fun = func(*Env) uint8 {
			return uint8(v.Uint())
		}
	case xr.Uint16:
		fun = func(*Env) uint16 {
			return uint16(v.Uint())
		}
	case xr.Uint32:
		fun = func(*Env) uint32 {
			return uint32(v.Uint())
		}
	case xr.Uint64:
		fun = func(*Env) uint64 {
			return v.Uint()
		}
	case xr.Uintptr:
		fun = func(*Env) uintptr {
			return uintptr(v.Uint())
		}
	case xr.Float32:
		fun = func(*Env) float32 {
			return float32(v.Float())
		}
	case xr.Float64:
		fun = func(*Env) float64 {
			return v.Float()
		}
	case xr.Complex64:
		fun = func(*Env) complex64 {
			return complex64(v.Complex())
		}
	case xr.Complex128:
		fun = func(*Env) complex128 {
			return v.Complex()
		}
	case xr.String:
		fun = func(*Env) string {
			return v.String()
		}
	default:
		fun = func(*Env) xr.Value {
			return v
		}
	}
	// v is an imported variable. do NOT store its value in *Expr,
	// because that's how constants are represented:
	// fast interpreter will then (incorrectly) perform constant propagation.
	return exprFun(t, fun)
}

// create an expression that will return the value of imported variable described by bind.
//
// mandatory optimization: for basic kinds, do not wrap in reflect.Value
func (imp *Import) intSymbol(bind *Bind, st *output.Stringer) *Expr {
	idx := bind.Desc.Index()
	if idx == NoIndex {
		st.Errorf("undefined identifier %s._", imp.Name)
	}
	t := bind.Type
	env := imp.env
	var fun I
	switch t.Kind() {
	case xr.Bool:
		fun = func(*Env) bool {
			return *(*bool)(unsafe.Pointer(&env.Ints[idx]))
		}
	case xr.Int:
		fun = func(*Env) int {
			return *(*int)(unsafe.Pointer(&env.Ints[idx]))
		}
	case xr.Int8:
		fun = func(*Env) int8 {
			return *(*int8)(unsafe.Pointer(&env.Ints[idx]))
		}
	case xr.Int16:
		fun = func(*Env) int16 {
			return *(*int16)(unsafe.Pointer(&env.Ints[idx]))
		}
	case xr.Int32:
		fun = func(*Env) int32 {
			return *(*int32)(unsafe.Pointer(&env.Ints[idx]))
		}
	case xr.Int64:
		fun = func(*Env) int64 {
			return *(*int64)(unsafe.Pointer(&env.Ints[idx]))
		}
	case xr.Uint:
		fun = func(*Env) uint {
			return *(*uint)(unsafe.Pointer(&env.Ints[idx]))
		}
	case xr.Uint8:
		fun = func(*Env) uint8 {
			return *(*uint8)(unsafe.Pointer(&env.Ints[idx]))
		}
	case xr.Uint16:
		fun = func(*Env) uint16 {
			return *(*uint16)(unsafe.Pointer(&env.Ints[idx]))
		}
	case xr.Uint32:
		fun = func(*Env) uint32 {
			return *(*uint32)(unsafe.Pointer(&env.Ints[idx]))
		}
	case xr.Uint64:
		fun = func(*Env) uint64 {
			return env.Ints[idx]
		}
	case xr.Uintptr:
		fun = func(*Env) uintptr {
			return *(*uintptr)(unsafe.Pointer(&env.Ints[idx]))
		}
	case xr.Float32:
		fun = func(*Env) float32 {
			return *(*float32)(unsafe.Pointer(&env.Ints[idx]))
		}
	case xr.Float64:
		fun = func(*Env) float64 {
			return *(*float64)(unsafe.Pointer(&env.Ints[idx]))
		}
	case xr.Complex64:
		fun = func(*Env) complex64 {
			return *(*complex64)(unsafe.Pointer(&env.Ints[idx]))
		}
	case xr.Complex128:
		fun = func(*Env) complex128 {
			return *(*complex128)(unsafe.Pointer(&env.Ints[idx]))
		}
	default:
		st.Errorf("unsupported symbol type, cannot use for optimized read: %v %v.%v <%v>",
			bind.Desc.Class(), imp.Name, bind.Name, bind.Type)
		return nil
	}
	// Do NOT store env.Ints[idx] into *Expr, because that's how constants are represented:
	// fast interpreter will then (incorrectly) perform constant propagation.
	return exprFun(t, fun)
}

// return a Place representing the imported variable described by bind.
//
// mandatory optimization: for basic kinds, do not wrap in reflect.Value
func (imp *Import) intPlace(c *Comp, bind *Bind, opt PlaceOption) *Place {
	idx := bind.Desc.Index()
	if idx == NoIndex {
		c.Errorf("%v %v %v.%v", opt, bind.Desc.Class(), imp.Name, bind.Name)
	}
	t := bind.Type
	var addr func(*Env) xr.Value
	impenv := imp.env
	switch t.Kind() {
	case xr.Bool:
		addr = func(env *Env) xr.Value {
			return xr.ValueOf((*bool)(unsafe.Pointer(&impenv.Ints[idx])))
		}
	case xr.Int:
		addr = func(env *Env) xr.Value {
			return xr.ValueOf((*int)(unsafe.Pointer(&impenv.Ints[idx])))
		}
	case xr.Int8:
		addr = func(env *Env) xr.Value {
			return xr.ValueOf((*int8)(unsafe.Pointer(&impenv.Ints[idx])))
		}
	case xr.Int16:
		addr = func(env *Env) xr.Value {
			return xr.ValueOf((*int16)(unsafe.Pointer(&impenv.Ints[idx])))
		}
	case xr.Int32:
		addr = func(env *Env) xr.Value {
			return xr.ValueOf((*int32)(unsafe.Pointer(&impenv.Ints[idx])))
		}
	case xr.Int64:
		addr = func(env *Env) xr.Value {
			return xr.ValueOf((*int64)(unsafe.Pointer(&impenv.Ints[idx])))
		}
	case xr.Uint:
		addr = func(env *Env) xr.Value {
			return xr.ValueOf((*uint)(unsafe.Pointer(&impenv.Ints[idx])))
		}
	case xr.Uint8:
		addr = func(env *Env) xr.Value {
			return xr.ValueOf((*uint8)(unsafe.Pointer(&impenv.Ints[idx])))
		}
	case xr.Uint16:
		addr = func(env *Env) xr.Value {
			return xr.ValueOf((*uint16)(unsafe.Pointer(&impenv.Ints[idx])))
		}
	case xr.Uint32:
		addr = func(env *Env) xr.Value {
			return xr.ValueOf((*uint32)(unsafe.Pointer(&impenv.Ints[idx])))
		}
	case xr.Uint64:
		addr = func(env *Env) xr.Value {
			return xr.ValueOf(&impenv.Ints[idx])
		}
	case xr.Uintptr:
		addr = func(env *Env) xr.Value {
			return xr.ValueOf((*uintptr)(unsafe.Pointer(&impenv.Ints[idx])))
		}
	case xr.Float32:
		addr = func(env *Env) xr.Value {
			return xr.ValueOf((*float32)(unsafe.Pointer(&impenv.Ints[idx])))
		}
	case xr.Float64:
		addr = func(env *Env) xr.Value {
			return xr.ValueOf((*float64)(unsafe.Pointer(&impenv.Ints[idx])))
		}
	case xr.Complex64:
		addr = func(env *Env) xr.Value {
			return xr.ValueOf((*complex64)(unsafe.Pointer(&impenv.Ints[idx])))
		}
	case xr.Complex128:
		addr = func(env *Env) xr.Value {
			return xr.ValueOf((*complex128)(unsafe.Pointer(&impenv.Ints[idx])))
		}
	default:
		c.Errorf("%s unsupported variable type <%v>: %s %s.%s",
			opt, t, bind.Desc.Class(), imp.Name, bind.Name)
		return nil
	}
	return &Place{
		Var: Var{Type: bind.Type, Name: bind.Name},
		Fun: func(env *Env) xr.Value {
			return addr(env).Elem()
		},
		Addr: addr,
	}
}
