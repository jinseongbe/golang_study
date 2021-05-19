// this file was generated by gomacro command: import _i "github.com/cosmos72/gomacro/classic"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package classic

import (
	r "reflect"

	"github.com/cosmos72/gomacro/imports"
)

// reflection: allow interpreted code to import "github.com/cosmos72/gomacro/classic"
func init() {
	imports.Packages["github.com/cosmos72/gomacro/classic"] = imports.Package{
		Binds: map[string]r.Value{
			"MultiThread":      r.ValueOf(MultiThread),
			"New":              r.ValueOf(New),
			"NewEnv":           r.ValueOf(NewEnv),
			"NewThreadGlobals": r.ValueOf(NewThreadGlobals),
			"NilREnv":           r.ValueOf(&NilREnv).Elem(),
		}, Types: map[string]r.Type{
			"BindMap":       r.TypeOf((*BindMap)(nil)).Elem(),
			"CallFrame":     r.TypeOf((*CallFrame)(nil)).Elem(),
			"CallStack":     r.TypeOf((*CallStack)(nil)).Elem(),
			"Constructor":   r.TypeOf((*Constructor)(nil)).Elem(),
			"Env":           r.TypeOf((*Env)(nil)).Elem(),
			"Error_builtin": r.TypeOf((*Error_builtin)(nil)).Elem(),
			"Function":      r.TypeOf((*Function)(nil)).Elem(),
			"Interp":        r.TypeOf((*Interp)(nil)).Elem(),
			"Macro":         r.TypeOf((*Macro)(nil)).Elem(),
			"Methods":       r.TypeOf((*Methods)(nil)).Elem(),
			"ThreadGlobals": r.TypeOf((*ThreadGlobals)(nil)).Elem(),
			"TypeMap":       r.TypeOf((*TypeMap)(nil)).Elem(),
			"TypedValue":    r.TypeOf((*TypedValue)(nil)).Elem(),
		}, Untypeds: map[string]string{
			"MultiThread": "bool:true",
		}, Wrappers: map[string][]string{
			"Env":           []string{"CollectAst", "CollectNode", "CollectPackageImports", "Copy", "Debugf", "Error", "Errorf", "Fprintf", "Gensym", "GensymAnonymous", "GensymPrivate", "ImportPackage", "IncLine", "IncLineBytes", "Init", "LookupPackage", "ObjMethodByName", "ParseBytes", "Position", "ShowHelp", "Sprintf", "ToString", "WarnExtraValues", "Warnf", "WriteDeclsToFile", "WriteDeclsToStream"},
			"Interp":        []string{"AsPackage", "CallerFrame", "ChangePackage", "ClassicEval", "CollectAst", "CollectNode", "CollectPackageImports", "Copy", "CurrentFrame", "Debugf", "DefineConst", "DefineFunc", "DefineVar", "Error", "Errorf", "Eval", "Eval1", "EvalAst", "EvalAst1", "EvalNode", "EvalNode1", "FastEval", "FileEnv", "Fprintf", "Gensym", "GensymAnonymous", "GensymPrivate", "ImportPackage", "IncLine", "IncLineBytes", "Init", "Inspect", "LookupPackage", "MacroExpand", "MacroExpand1", "MacroExpandAstCodewalk", "MacroExpandCodewalk", "MergePackage", "ObjMethodByName", "Parse", "ParseBytes", "ParseOnly", "Position", "ReadMultiline", "ShowHelp", "ShowPackage", "Sprintf", "ToString", "TopEnv", "ValueOf", "WarnExtraValues", "Warnf", "WriteDeclsToFile", "WriteDeclsToStream"},
			"ThreadGlobals": []string{"CollectAst", "CollectNode", "CollectPackageImports", "Copy", "Debugf", "Error", "Errorf", "Fprintf", "Gensym", "GensymAnonymous", "GensymPrivate", "ImportPackage", "IncLine", "IncLineBytes", "Init", "LookupPackage", "ParseBytes", "Position", "Sprintf", "ToString", "WarnExtraValues", "Warnf", "WriteDeclsToFile", "WriteDeclsToStream"},
		},
	}
}