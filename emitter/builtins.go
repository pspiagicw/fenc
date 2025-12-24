package emitter

import (
	"fmt"

	"github.com/pspiagicw/fenc/object"
	"github.com/pspiagicw/goreland"
)

const (
	_ int = iota
	PRINT
	STRI
	SLEN
	PUSH
)

// TODO: Check number of arguments in builtins and normal functions.
var BuiltinMap = map[int]object.Builtin{
	PRINT: {Internal: PrintFunc},
	STRI:  {Internal: StringFunc},
	SLEN:  {Internal: LenFunc},
	PUSH:  {Internal: PushFunc},
}

func registerBuiltins(s *SymbolTable) {
	s.DefineBuiltin("print", PRINT)
	s.DefineBuiltin("stri", STRI)
	s.DefineBuiltin("len", SLEN)
	s.DefineBuiltin("push", PUSH)
}

func PrintFunc(args []object.Object) object.Object {
	formatString := args[0]
	fmt.Println(formatString.String())

	return object.Null{}
}
func StringFunc(args []object.Object) object.Object {
	value := args[0].String()

	return object.String{
		Value: value,
	}
}

func LenFunc(args []object.Object) object.Object {
	value := args[0]

	o, ok := value.(object.String)
	if !ok {
		goreland.LogFatal("Expected string, got %s", value.Type())
	}

	return object.Int{
		Value: len(o.Value),
	}
}

func PushFunc(args []object.Object) object.Object {
	arr := args[0]

	a, ok := arr.(object.Array)
	if !ok {
		goreland.LogFatal("Expected array, got %s", arr.Type())
	}

	a.Values = append(a.Values, a)

	return a
}
