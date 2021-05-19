// this file was generated by gomacro command: import _b "strings"
// DO NOT EDIT! Any change will be lost when the file is re-generated

// +build go1.10

package go1_10

import (
	. "reflect"
	strings "strings"
)

// reflection: allow interpreted code to import "strings"
func init() {
	Packages["strings"] = Package{
	Binds: map[string]Value{
		"Compare":	ValueOf(strings.Compare),
		"Contains":	ValueOf(strings.Contains),
		"ContainsAny":	ValueOf(strings.ContainsAny),
		"ContainsRune":	ValueOf(strings.ContainsRune),
		"Count":	ValueOf(strings.Count),
		"EqualFold":	ValueOf(strings.EqualFold),
		"Fields":	ValueOf(strings.Fields),
		"FieldsFunc":	ValueOf(strings.FieldsFunc),
		"HasPrefix":	ValueOf(strings.HasPrefix),
		"HasSuffix":	ValueOf(strings.HasSuffix),
		"Index":	ValueOf(strings.Index),
		"IndexAny":	ValueOf(strings.IndexAny),
		"IndexByte":	ValueOf(strings.IndexByte),
		"IndexFunc":	ValueOf(strings.IndexFunc),
		"IndexRune":	ValueOf(strings.IndexRune),
		"Join":	ValueOf(strings.Join),
		"LastIndex":	ValueOf(strings.LastIndex),
		"LastIndexAny":	ValueOf(strings.LastIndexAny),
		"LastIndexByte":	ValueOf(strings.LastIndexByte),
		"LastIndexFunc":	ValueOf(strings.LastIndexFunc),
		"Map":	ValueOf(strings.Map),
		"NewReader":	ValueOf(strings.NewReader),
		"NewReplacer":	ValueOf(strings.NewReplacer),
		"Repeat":	ValueOf(strings.Repeat),
		"Replace":	ValueOf(strings.Replace),
		"Split":	ValueOf(strings.Split),
		"SplitAfter":	ValueOf(strings.SplitAfter),
		"SplitAfterN":	ValueOf(strings.SplitAfterN),
		"SplitN":	ValueOf(strings.SplitN),
		"Title":	ValueOf(strings.Title),
		"ToLower":	ValueOf(strings.ToLower),
		"ToLowerSpecial":	ValueOf(strings.ToLowerSpecial),
		"ToTitle":	ValueOf(strings.ToTitle),
		"ToTitleSpecial":	ValueOf(strings.ToTitleSpecial),
		"ToUpper":	ValueOf(strings.ToUpper),
		"ToUpperSpecial":	ValueOf(strings.ToUpperSpecial),
		"Trim":	ValueOf(strings.Trim),
		"TrimFunc":	ValueOf(strings.TrimFunc),
		"TrimLeft":	ValueOf(strings.TrimLeft),
		"TrimLeftFunc":	ValueOf(strings.TrimLeftFunc),
		"TrimPrefix":	ValueOf(strings.TrimPrefix),
		"TrimRight":	ValueOf(strings.TrimRight),
		"TrimRightFunc":	ValueOf(strings.TrimRightFunc),
		"TrimSpace":	ValueOf(strings.TrimSpace),
		"TrimSuffix":	ValueOf(strings.TrimSuffix),
	}, Types: map[string]Type{
		"Builder":	TypeOf((*strings.Builder)(nil)).Elem(),
		"Reader":	TypeOf((*strings.Reader)(nil)).Elem(),
		"Replacer":	TypeOf((*strings.Replacer)(nil)).Elem(),
	}, 
	}
}
