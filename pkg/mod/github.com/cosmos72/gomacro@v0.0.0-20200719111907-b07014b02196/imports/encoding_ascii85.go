// this file was generated by gomacro command: import _b "encoding/ascii85"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"encoding/ascii85"
)

// reflection: allow interpreted code to import "encoding/ascii85"
func init() {
	Packages["encoding/ascii85"] = Package{
	Binds: map[string]Value{
		"Decode":	ValueOf(ascii85.Decode),
		"Encode":	ValueOf(ascii85.Encode),
		"MaxEncodedLen":	ValueOf(ascii85.MaxEncodedLen),
		"NewDecoder":	ValueOf(ascii85.NewDecoder),
		"NewEncoder":	ValueOf(ascii85.NewEncoder),
	}, Types: map[string]Type{
		"CorruptInputError":	TypeOf((*ascii85.CorruptInputError)(nil)).Elem(),
	}, 
	}
}
