package emitter

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/pspiagicw/fenc/code"
	"github.com/pspiagicw/fenc/object"
)

func TestPush(t *testing.T) {

	e := getEmitter()
	e.PushInt(1)

	constants := []object.Object{
		object.CreateInt(1),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
	}

	testEmitter(t, e, expected, constants)
}

func TestMultiplePush(t *testing.T) {
	e := getEmitter()
	e.PushInt(1)
	e.PushInt(222)
	e.PushInt(1)

	constants := []object.Object{
		object.CreateInt(1),
		object.CreateInt(222),
		object.CreateInt(1),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
		{OpCode: code.PUSH, Args: createArgs(1)},
		{OpCode: code.PUSH, Args: createArgs(2)},
	}

	testEmitter(t, e, expected, constants)
}

func TestAdd(t *testing.T) {
	e := getEmitter()
	e.PushInt(1)
	e.PushInt(2)
	e.AddInt()

	constants := []object.Object{
		object.CreateInt(1),
		object.CreateInt(2),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
		{OpCode: code.PUSH, Args: createArgs(1)},
		{OpCode: code.ADD_INT},
	}

	testEmitter(t, e, expected, constants)
}

func TestPushFloatStringBool(t *testing.T) {
	e := getEmitter()
	e.PushFloat(1)
	e.PushBool(true)
	e.PushString("string")

	constants := []object.Object{
		object.CreateFloat(1),
		object.CreateBool(true),
		object.CreateString("string"),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
		{OpCode: code.PUSH, Args: createArgs(1)},
		{OpCode: code.PUSH, Args: createArgs(2)},
	}

	testEmitter(t, e, expected, constants)
}

func TestSubInt(t *testing.T) {
	e := getEmitter()
	e.PushInt(5)
	e.PushInt(3)
	e.SubInt()

	constants := []object.Object{
		object.CreateInt(5),
		object.CreateInt(3),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
		{OpCode: code.PUSH, Args: createArgs(1)},
		{OpCode: code.SUB_INT},
	}

	testEmitter(t, e, expected, constants)
}

func TestMulInt(t *testing.T) {
	e := getEmitter()
	e.PushInt(4)
	e.PushInt(2)
	e.MulInt()

	constants := []object.Object{
		object.CreateInt(4),
		object.CreateInt(2),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
		{OpCode: code.PUSH, Args: createArgs(1)},
		{OpCode: code.MUL_INT},
	}

	testEmitter(t, e, expected, constants)
}

func TestDivInt(t *testing.T) {
	e := getEmitter()
	e.PushInt(10)
	e.PushInt(4)
	e.DivInt()

	constants := []object.Object{
		object.CreateInt(10),
		object.CreateInt(4),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
		{OpCode: code.PUSH, Args: createArgs(1)},
		{OpCode: code.DIV_INT},
	}

	testEmitter(t, e, expected, constants)
}

func TestLtInt(t *testing.T) {
	e := getEmitter()
	e.PushInt(2)
	e.PushInt(5)
	e.LtInt()

	constants := []object.Object{
		object.CreateInt(2),
		object.CreateInt(5),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
		{OpCode: code.PUSH, Args: createArgs(1)},
		{OpCode: code.LT_INT},
	}

	testEmitter(t, e, expected, constants)
}

func TestLteInt(t *testing.T) {
	e := getEmitter()
	e.PushInt(2)
	e.PushInt(2)
	e.LteInt()

	constants := []object.Object{
		object.CreateInt(2),
		object.CreateInt(2),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
		{OpCode: code.PUSH, Args: createArgs(1)},
		{OpCode: code.LTE_INT},
	}

	testEmitter(t, e, expected, constants)
}

func TestGtInt(t *testing.T) {
	e := getEmitter()
	e.PushInt(10)
	e.PushInt(2)
	e.GtInt()

	constants := []object.Object{
		object.CreateInt(10),
		object.CreateInt(2),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
		{OpCode: code.PUSH, Args: createArgs(1)},
		{OpCode: code.GT_INT},
	}

	testEmitter(t, e, expected, constants)
}

func TestGteInt(t *testing.T) {
	e := getEmitter()
	e.PushFloat(10)
	e.PushFloat(10)
	e.GteInt()

	constants := []object.Object{
		object.CreateFloat(10),
		object.CreateFloat(10),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
		{OpCode: code.PUSH, Args: createArgs(1)},
		{OpCode: code.GTE_INT},
	}

	testEmitter(t, e, expected, constants)
}

func TestEqInt(t *testing.T) {
	e := getEmitter()
	e.PushFloat(10)
	e.PushFloat(10)
	e.Eq()

	constants := []object.Object{
		object.CreateFloat(10),
		object.CreateFloat(10),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
		{OpCode: code.PUSH, Args: createArgs(1)},
		{OpCode: code.EQ},
	}

	testEmitter(t, e, expected, constants)
}

func TestNeqInt(t *testing.T) {
	e := getEmitter()
	e.PushFloat(10)
	e.PushFloat(10)
	e.Neq()

	constants := []object.Object{
		object.CreateFloat(10),
		object.CreateFloat(10),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
		{OpCode: code.PUSH, Args: createArgs(1)},
		{OpCode: code.NEQ},
	}

	testEmitter(t, e, expected, constants)
}

// ---- (Float) ----

func TestAddFloat(t *testing.T) {
	e := getEmitter()
	e.PushFloat(1.5)
	e.PushFloat(2.5)
	e.AddFloat()

	constants := []object.Object{
		object.CreateFloat(1.5),
		object.CreateFloat(2.5),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
		{OpCode: code.PUSH, Args: createArgs(1)},
		{OpCode: code.ADD_FLOAT},
	}

	testEmitter(t, e, expected, constants)
}

func TestSubFloat(t *testing.T) {
	e := getEmitter()
	e.PushFloat(5.5)
	e.PushFloat(3.25)
	e.SubFloat()

	constants := []object.Object{
		object.CreateFloat(5.5),
		object.CreateFloat(3.25),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
		{OpCode: code.PUSH, Args: createArgs(1)},
		{OpCode: code.SUB_FLOAT},
	}

	testEmitter(t, e, expected, constants)
}

func TestMulFloat(t *testing.T) {
	e := getEmitter()
	e.PushFloat(2.5)
	e.PushFloat(4)
	e.MulFloat()

	constants := []object.Object{
		object.CreateFloat(2.5),
		object.CreateFloat(4),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
		{OpCode: code.PUSH, Args: createArgs(1)},
		{OpCode: code.MUL_FLOAT},
	}

	testEmitter(t, e, expected, constants)
}

// ---- Comparison (Float) ----

func TestLtFloat(t *testing.T) {
	e := getEmitter()
	e.PushFloat(1.2)
	e.PushFloat(3.4)
	e.LtFloat()

	constants := []object.Object{
		object.CreateFloat(1.2),
		object.CreateFloat(3.4),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
		{OpCode: code.PUSH, Args: createArgs(1)},
		{OpCode: code.LT_FLOAT},
	}

	testEmitter(t, e, expected, constants)
}

func TestLteFloat(t *testing.T) {
	e := getEmitter()
	e.PushFloat(1.2)
	e.PushFloat(3.4)
	e.LteFloat()

	constants := []object.Object{
		object.CreateFloat(1.2),
		object.CreateFloat(3.4),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
		{OpCode: code.PUSH, Args: createArgs(1)},
		{OpCode: code.LTE_FLOAT},
	}

	testEmitter(t, e, expected, constants)
}

func TestGtFloat(t *testing.T) {
	e := getEmitter()
	e.PushFloat(5.0)
	e.PushFloat(2.0)
	e.GtFloat()

	constants := []object.Object{
		object.CreateFloat(5.0),
		object.CreateFloat(2.0),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
		{OpCode: code.PUSH, Args: createArgs(1)},
		{OpCode: code.GT_FLOAT},
	}

	testEmitter(t, e, expected, constants)
}

func TestGteFloat(t *testing.T) {
	e := getEmitter()
	e.PushFloat(5.0)
	e.PushFloat(2.0)
	e.GteFloat()

	constants := []object.Object{
		object.CreateFloat(5.0),
		object.CreateFloat(2.0),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
		{OpCode: code.PUSH, Args: createArgs(1)},
		{OpCode: code.GTE_FLOAT},
	}

	testEmitter(t, e, expected, constants)
}

func TestEqFloat(t *testing.T) {
	e := getEmitter()
	e.PushFloat(1.0)
	e.PushFloat(1.0)
	e.Eq()

	constants := []object.Object{
		object.CreateFloat(1.0),
		object.CreateFloat(1.0),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
		{OpCode: code.PUSH, Args: createArgs(1)},
		{OpCode: code.EQ},
	}

	testEmitter(t, e, expected, constants)
}

func TestNeqFloat(t *testing.T) {
	e := getEmitter()
	e.PushFloat(2.0)
	e.PushFloat(3.0)
	e.Neq()

	constants := []object.Object{
		object.CreateFloat(2.0),
		object.CreateFloat(3.0),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
		{OpCode: code.PUSH, Args: createArgs(1)},
		{OpCode: code.NEQ},
	}

	testEmitter(t, e, expected, constants)
}

// ---- (bool) ----

func TestNeqBool(t *testing.T) {
	e := getEmitter()
	e.PushBool(true)
	e.PushBool(false)
	e.Neq()

	constants := []object.Object{
		object.CreateBool(true),
		object.CreateBool(false),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
		{OpCode: code.PUSH, Args: createArgs(1)},
		{OpCode: code.NEQ},
	}

	testEmitter(t, e, expected, constants)
}

func TestAndBool(t *testing.T) {
	e := getEmitter()
	e.PushBool(true)
	e.PushBool(false)
	e.AndBool()

	constants := []object.Object{
		object.CreateBool(true),
		object.CreateBool(false),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
		{OpCode: code.PUSH, Args: createArgs(1)},
		{OpCode: code.AND_BOOL},
	}

	testEmitter(t, e, expected, constants)
}

func TestOrBool(t *testing.T) {
	e := getEmitter()
	e.PushBool(true)
	e.PushBool(true)
	e.OrBool()

	constants := []object.Object{
		object.CreateBool(true),
		object.CreateBool(true),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
		{OpCode: code.PUSH, Args: createArgs(1)},
		{OpCode: code.OR_BOOL},
	}

	testEmitter(t, e, expected, constants)
}

// --- (string) ---
func TestAddString(t *testing.T) {
	e := getEmitter()
	e.PushString("foo")
	e.PushString("bar")
	e.AddString()

	constants := []object.Object{
		object.CreateString("foo"),
		object.CreateString("bar"),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
		{OpCode: code.PUSH, Args: createArgs(1)},
		{OpCode: code.ADD_STRING},
	}

	testEmitter(t, e, expected, constants)
}

func TestEqString(t *testing.T) {
	e := getEmitter()
	e.PushString("a")
	e.PushString("a")
	e.Eq()

	constants := []object.Object{
		object.CreateString("a"),
		object.CreateString("a"),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
		{OpCode: code.PUSH, Args: createArgs(1)},
		{OpCode: code.EQ},
	}

	testEmitter(t, e, expected, constants)
}

func TestNEqString(t *testing.T) {
	e := getEmitter()
	e.PushString("a")
	e.PushString("a")
	e.Neq()

	constants := []object.Object{
		object.CreateString("a"),
		object.CreateString("a"),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
		{OpCode: code.PUSH, Args: createArgs(1)},
		{OpCode: code.NEQ},
	}

	testEmitter(t, e, expected, constants)
}

func TestIfStatement(t *testing.T) {
	e := getEmitter()
	e.If(
		func(e *Emitter) error {
			e.PushInt(1)
			e.PushInt(2)
			e.LtInt()
			return nil
		},
		func(e *Emitter) error {
			e.PushInt(10)
			return nil
		},
		func(e *Emitter) error {
			e.PushInt(20)
			return nil
		},
	)

	constants := []object.Object{
		object.CreateInt(1),
		object.CreateInt(2),
		object.CreateInt(10),
		object.CreateInt(20),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},       // 0000
		{OpCode: code.PUSH, Args: createArgs(1)},       // 0001
		{OpCode: code.LT_INT},                          // 0002
		{OpCode: code.JUMP_FALSE, Args: createArgs(6)}, // 0003

		// --- Consequence --
		{OpCode: code.PUSH, Args: createArgs(2)}, // 0004
		{OpCode: code.JUMP, Args: createArgs(7)}, // 0005

		// -- Alternaitve --
		{OpCode: code.PUSH, Args: createArgs(3)}, // 0006
	}

	testEmitter(t, e, expected, constants)

}
func TestReturn(t *testing.T) {
	e := getEmitter()
	e.Return()

	constants := []object.Object{}

	expected := []code.Instruction{
		{OpCode: code.RETURN, Args: createArgs()},
	}

	testEmitter(t, e, expected, constants)

}
func TestReturnValue(t *testing.T) {
	e := getEmitter()
	e.ReturnValue()

	constants := []object.Object{}

	expected := []code.Instruction{
		{OpCode: code.RETURN_VALUE, Args: createArgs()},
	}

	testEmitter(t, e, expected, constants)
}

func TestCall(t *testing.T) {
	e := getEmitter()
	e.Call(0)

	constants := []object.Object{}

	expected := []code.Instruction{
		{OpCode: code.CALL, Args: createArgs(0)},
	}

	testEmitter(t, e, expected, constants)
}

func TestGlobals(t *testing.T) {
	e := getEmitter()

	e.PushInt(2)
	e.Store("x")
	e.Load("x")

	e.PushString("pspiagicw")
	e.Store("name")
	e.Load("name")

	constants := []object.Object{
		object.CreateInt(2),
		object.CreateString("pspiagicw"),
	}

	expected := []code.Instruction{
		{OpCode: code.PUSH, Args: createArgs(0)},
		{OpCode: code.STORE_GLOBAL, Args: createArgs(0)},
		{OpCode: code.LOAD_GLOBAL, Args: createArgs(0)},

		{OpCode: code.PUSH, Args: createArgs(1)},
		{OpCode: code.STORE_GLOBAL, Args: createArgs(1)},
		{OpCode: code.LOAD_GLOBAL, Args: createArgs(1)},
	}

	testEmitter(t, e, expected, constants)
}

func TestFunctionSimple(t *testing.T) {
	e := getEmitter()
	e.Function("test", []string{}, func(e *Emitter) error {
		e.PushInt(2)
		return nil
	})
	e.Load("test")
	e.Call(0)

	constants := []object.Object{
		object.CreateInt(2),
		object.CreateFunction([]code.Instruction{
			{OpCode: code.PUSH, Args: createArgs(0)},
		}),
	}

	expected := []code.Instruction{
		{OpCode: code.CLOSURE, Args: createArgs(1, 0)},
		{OpCode: code.STORE_GLOBAL, Args: createArgs(0)},
		{OpCode: code.LOAD_GLOBAL, Args: createArgs(0)},
		{OpCode: code.CALL, Args: createArgs(0)},
	}

	testEmitter(t, e, expected, constants)

}
func TestFunctionWithArg(t *testing.T) {
	e := getEmitter()
	e.Function("add", []string{"x", "y"}, func(e *Emitter) error {
		e.Load("x")
		e.Load("y")
		e.AddInt()
		e.Store("z")
		e.Load("z")
		e.ReturnValue()
		return nil
	})

	constants := []object.Object{
		object.CreateFunction([]code.Instruction{
			{OpCode: code.LOAD_LOCAL, Args: createArgs(0)},
			{OpCode: code.LOAD_LOCAL, Args: createArgs(1)},
			{OpCode: code.ADD_INT},
			{OpCode: code.STORE_LOCAL, Args: createArgs(2)},
			{OpCode: code.LOAD_LOCAL, Args: createArgs(2)},
			{OpCode: code.RETURN_VALUE},
		}),
	}

	expected := []code.Instruction{
		{OpCode: code.CLOSURE, Args: createArgs(0, 0)},
		{OpCode: code.STORE_GLOBAL, Args: createArgs(0)},
	}

	testEmitter(t, e, expected, constants)
}
func TestLambda(t *testing.T) {
	e := getEmitter()
	e.Lambda([]string{}, func(e *Emitter) error {
		e.PushInt(1)
		return nil
	})

	constants := []object.Object{
		object.CreateInt(1),
		object.CreateFunction([]code.Instruction{
			{OpCode: code.PUSH, Args: createArgs(0)},
		}),
	}

	expected := []code.Instruction{
		{OpCode: code.CLOSURE, Args: createArgs(1, 0)},
	}

	testEmitter(t, e, expected, constants)

}

// TODO: Find out why is this not working!
// func TestFunctionWithLambdaReturn(t *testing.T) {
// 	e := getEmitter()
// 	e.Function("adder", []string{"x"}, func(e *Emitter) {
// 		e.Lambda([]string{"y"}, func(e *Emitter) {
// 			e.Load("x")
// 			e.Load("y")
// 			e.AddInt()
// 			e.Return()
// 		})
// 		e.Return()
// 	})
// 	fmt.Println("CONSTANTS")
// 	dump.Constants(e.constants.constants)
// 	fmt.Println("TAPE")
// 	dump.Dump(e.tape)
// }

func TestClosure(t *testing.T) {
	e := getEmitter()
	e.Lambda([]string{"a"}, func(e *Emitter) error {
		e.Lambda([]string{"b"}, func(e *Emitter) error {
			e.Load("a")
			e.Load("b")
			e.AddInt()
			e.ReturnValue()
			return nil
		})
		e.Return()
		return nil
	})

	constants := []object.Object{
		object.CreateFunction([]code.Instruction{
			{OpCode: code.LOAD_FREE, Args: createArgs(0)},
			{OpCode: code.LOAD_LOCAL, Args: createArgs(0)},
			{OpCode: code.ADD_INT},
			{OpCode: code.RETURN_VALUE},
		}),
		object.CreateFunction([]code.Instruction{
			{OpCode: code.LOAD_LOCAL, Args: createArgs(0)},
			{OpCode: code.CLOSURE, Args: createArgs(0, 1)},
			{OpCode: code.RETURN},
		}),
	}

	expected := []code.Instruction{
		createInstruction(code.CLOSURE, 1, 0),
	}

	testEmitter(t, e, expected, constants)
}

func TestNestedClosure(t *testing.T) {
	e := getEmitter()
	e.Lambda([]string{"a"}, func(e *Emitter) error {
		e.Lambda([]string{"b"}, func(e *Emitter) error {
			e.Lambda([]string{"c"}, func(e *Emitter) error {
				e.Load("a")
				e.Load("b")
				e.AddInt()
				e.Load("c")
				e.AddInt()
				e.ReturnValue()
				return nil
			})
			e.ReturnValue()
			return nil
		})
		e.ReturnValue()
		return nil
	})

	constants := []object.Object{
		object.CreateFunction([]code.Instruction{
			createInstruction(code.LOAD_FREE, 0),
			createInstruction(code.LOAD_FREE, 1),
			createInstruction(code.ADD_INT),
			createInstruction(code.LOAD_LOCAL, 0),
			createInstruction(code.ADD_INT),
			createInstruction(code.RETURN_VALUE),
		}),
		object.CreateFunction([]code.Instruction{
			createInstruction(code.LOAD_FREE, 0),
			createInstruction(code.LOAD_LOCAL, 0),
			createInstruction(code.CLOSURE, 0, 2),
			createInstruction(code.RETURN_VALUE),
		}),
		object.CreateFunction([]code.Instruction{
			createInstruction(code.LOAD_LOCAL, 0),
			createInstruction(code.CLOSURE, 1, 1),
			createInstruction(code.RETURN_VALUE),
		}),
	}

	expected := []code.Instruction{
		createInstruction(code.CLOSURE, 2, 0),
	}

	testEmitter(t, e, expected, constants)
}

func TestClosureComplex(t *testing.T) {
	e := getEmitter()
	e.PushInt(55)
	e.Store("global")
	e.Lambda([]string{}, func(e *Emitter) error {
		e.PushInt(66)
		e.Store("a")
		e.Lambda([]string{}, func(e *Emitter) error {
			e.PushInt(77)
			e.Store("b")
			e.Lambda([]string{}, func(e *Emitter) error {
				e.PushInt(88)
				e.Store("c")
				e.Load("global")
				e.Load("a")
				e.AddInt()
				e.Load("b")
				e.AddInt()
				e.Load("c")
				e.AddInt()
				e.ReturnValue()
				return nil
			})
			e.ReturnValue()
			return nil
		})
		e.ReturnValue()
		return nil
	})

	constants := []object.Object{
		object.CreateInt(55),
		object.CreateInt(66),
		object.CreateInt(77),
		object.CreateInt(88),
		object.CreateFunction([]code.Instruction{
			createInstruction(code.PUSH, 3),
			createInstruction(code.STORE_LOCAL, 0),
			createInstruction(code.LOAD_GLOBAL, 0),
			createInstruction(code.LOAD_FREE, 0),
			createInstruction(code.ADD_INT),
			createInstruction(code.LOAD_FREE, 1),
			createInstruction(code.ADD_INT),
			createInstruction(code.LOAD_LOCAL, 0),
			createInstruction(code.ADD_INT),
			createInstruction(code.RETURN_VALUE),
		}),
		object.CreateFunction([]code.Instruction{
			createInstruction(code.PUSH, 2),
			createInstruction(code.STORE_LOCAL, 0),
			createInstruction(code.LOAD_FREE, 0),
			createInstruction(code.LOAD_LOCAL, 0),
			createInstruction(code.CLOSURE, 4, 2),
			createInstruction(code.RETURN_VALUE),
		}),
		object.CreateFunction([]code.Instruction{
			createInstruction(code.PUSH, 1),
			createInstruction(code.STORE_LOCAL, 0),
			createInstruction(code.LOAD_LOCAL, 0),
			createInstruction(code.CLOSURE, 5, 1),
			createInstruction(code.RETURN_VALUE),
		}),
	}

	expected := []code.Instruction{
		createInstruction(code.PUSH, 0),
		createInstruction(code.STORE_GLOBAL, 0),
		createInstruction(code.CLOSURE, 6, 0),
	}

	testEmitter(t, e, expected, constants)
}
func TestArrays(t *testing.T) {
	e := getEmitter()
	e.PushInt(2)
	e.PushInt(3)
	e.Array(2)

	expected := []code.Instruction{
		createInstruction(code.PUSH, 0),
		createInstruction(code.PUSH, 1),
		createInstruction(code.ARRAY, 2),
	}

	constants := []object.Object{
		object.CreateInt(2),
		object.CreateInt(3),
	}

	testEmitter(t, e, expected, constants)
}

func TestHashes(t *testing.T) {
	e := getEmitter()
	e.PushString("pspiagicw")
	e.PushInt(20)
	e.PushString("torvalds")
	e.PushInt(100)
	e.PushString("stallman")
	e.PushInt(80)
	e.Hash(6)

	expected := []code.Instruction{
		createInstruction(code.PUSH, 0),
		createInstruction(code.PUSH, 1),
		createInstruction(code.PUSH, 2),
		createInstruction(code.PUSH, 3),
		createInstruction(code.PUSH, 4),
		createInstruction(code.PUSH, 5),
		createInstruction(code.HASH, 6),
	}

	constants := []object.Object{
		object.CreateString("pspiagicw"),
		object.CreateInt(20),
		object.CreateString("torvalds"),
		object.CreateInt(100),
		object.CreateString("stallman"),
		object.CreateInt(80),
	}

	testEmitter(t, e, expected, constants)
}

func TestIndex(t *testing.T) {
	e := getEmitter()
	e.PushInt(2)
	e.PushInt(3)
	e.Array(2)
	e.PushInt(2)
	e.Index()

	expected := []code.Instruction{
		createInstruction(code.PUSH, 0),
		createInstruction(code.PUSH, 1),
		createInstruction(code.ARRAY, 2),
		createInstruction(code.PUSH, 2),
		createInstruction(code.INDEX),
	}

	constants := []object.Object{
		object.CreateInt(2),
		object.CreateInt(3),
		object.CreateInt(2),
	}

	testEmitter(t, e, expected, constants)

}

func TestToFloat(t *testing.T) {
	e := getEmitter()
	e.PushInt(3)
	e.ToFloat()

	expected := []code.Instruction{
		createInstruction(code.PUSH, 0),
		createInstruction(code.TO_FLOAT),
	}

	constants := []object.Object{
		object.CreateInt(3),
	}

	testEmitter(t, e, expected, constants)
}

func TestBuiltinPrint(t *testing.T) {
	e := getEmitter()
	e.Load("print")

	expected := []code.Instruction{
		createInstruction(code.BUILTIN, 1),
	}

	constants := []object.Object{}

	testEmitter(t, e, expected, constants)
}

func TestClassEmpty(t *testing.T) {
	e := getEmitter()
	e.Class("Something")

	expected := []code.Instruction{
		createInstruction(code.PUSH, 0),
		createInstruction(code.CLASS),
	}

	constants := []object.Object{
		object.CreateString("Something"),
	}

	testEmitter(t, e, expected, constants)
}

func createArgs(args ...int) []int {
	return args
}
func createInstruction(op code.Op, args ...int) code.Instruction {
	return code.Instruction{OpCode: op, Args: args}
}

func getEmitter() *Emitter {
	e := NewEmitter()

	return e
}

func testEmitter(t *testing.T, e *Emitter, expected []code.Instruction, constants []object.Object) {
	assert.Equal(t, constants, e.constants.constants, "Constant pool not equal!")
	assert.Equal(t, expected, e.tape, "Instructions on tape differ!")
}
