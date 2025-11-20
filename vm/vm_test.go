package vm

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/pspiagicw/fenc/dump"
	"github.com/pspiagicw/fenc/emitter"
	"github.com/pspiagicw/fenc/object"
)

func TestPush(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushInt(1)

	expected := object.CreateInt(1)

	testVM(t, e, expected)
}

func TestAdd(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushInt(1)
	e.PushInt(1)
	e.AddInt()

	expected := object.CreateInt(2)
	testVM(t, e, expected)
}

// ==========================
// Integer arithmetic
// ==========================

func TestAddInt(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushInt(2)
	e.PushInt(3)
	e.AddInt()
	testVM(t, e, object.CreateInt(5))
}

func TestSubInt(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushInt(5)
	e.PushInt(3)
	e.SubInt()
	testVM(t, e, object.CreateInt(2))
}

func TestMulInt(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushInt(4)
	e.PushInt(3)
	e.MulInt()
	testVM(t, e, object.CreateInt(12))
}

func TestDivInt(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushInt(10)
	e.PushInt(2)
	e.DivInt()
	testVM(t, e, object.CreateInt(5))
}

// ==========================
// Integer comparisons
// ==========================

func TestLtInt(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushInt(2)
	e.PushInt(5)
	e.LtInt()
	testVM(t, e, object.CreateBool(true))
}

func TestLteInt(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushInt(5)
	e.PushInt(5)
	e.LteInt()
	testVM(t, e, object.CreateBool(true))
}

func TestGtInt(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushInt(6)
	e.PushInt(4)
	e.GtInt()
	testVM(t, e, object.CreateBool(true))
}

func TestGteInt(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushInt(5)
	e.PushInt(5)
	e.GteInt()
	testVM(t, e, object.CreateBool(true))
}

func TestEqInt(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushInt(7)
	e.PushInt(7)
	e.Eq()
	testVM(t, e, object.CreateBool(true))
}

func TestNeqInt(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushInt(8)
	e.PushInt(9)
	e.Neq()
	testVM(t, e, object.CreateBool(true))
}

// ==========================
// Float arithmetic
// ==========================

func TestAddFloat(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushFloat(1.5)
	e.PushFloat(2.5)
	e.AddFloat()
	testVM(t, e, object.CreateFloat(4.0))
}

func TestSubFloat(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushFloat(5.5)
	e.PushFloat(2.5)
	e.SubFloat()
	testVM(t, e, object.CreateFloat(3.0))
}

func TestMulFloat(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushFloat(2.0)
	e.PushFloat(4.0)
	e.MulFloat()
	testVM(t, e, object.CreateFloat(8.0))
}

func TestDivFloat(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushFloat(27.0)
	e.PushFloat(9.0)
	e.DivFloat()
	testVM(t, e, object.CreateFloat(3.0))
}

// ==========================
// Float comparisons
// ==========================

func TestLtFloat(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushFloat(1.2)
	e.PushFloat(2.4)
	e.LtFloat()
	testVM(t, e, object.CreateBool(true))
}

func TestLteFloat(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushFloat(3.3)
	e.PushFloat(3.3)
	e.LteFloat()
	testVM(t, e, object.CreateBool(true))
}

func TestGtFloat(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushFloat(4.5)
	e.PushFloat(3.2)
	e.GtFloat()
	testVM(t, e, object.CreateBool(true))
}

func TestGteFloat(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushFloat(3.5)
	e.PushFloat(3.5)
	e.GteFloat()
	testVM(t, e, object.CreateBool(true))
}

func TestEqFloat(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushFloat(2.5)
	e.PushFloat(2.5)
	e.Eq()
	testVM(t, e, object.CreateBool(true))
}

func TestNeqFloat(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushFloat(2.5)
	e.PushFloat(3.5)
	e.Neq()
	testVM(t, e, object.CreateBool(true))
}

// ==========================
// Boolean logic
// ==========================

func TestAndBool(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushBool(false)
	e.PushBool(true)
	e.AndBool()
	testVM(t, e, object.CreateBool(false))
}

func TestOrBool(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushBool(true)
	e.PushBool(false)
	e.OrBool()
	testVM(t, e, object.CreateBool(true))
}

func TestEqBool(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushBool(false)
	e.PushBool(true)
	e.Eq()
	testVM(t, e, object.CreateBool(false))
}

func TestNeqBool(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushBool(true)
	e.PushBool(false)
	e.Neq()
	testVM(t, e, object.CreateBool(true))
}

// ==========================
// String operations
// ==========================

func TestAddString(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushString("Hello, ")
	e.PushString("World!")
	e.AddString()
	testVM(t, e, object.CreateString("Hello, World!"))
}

func TestEqString(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushString("foo")
	e.PushString("foo")
	e.Eq()
	testVM(t, e, object.CreateBool(true))
}

func TestNeqString(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushString("foo")
	e.PushString("bar")
	e.Neq()
	testVM(t, e, object.CreateBool(true))
}

func TestIf(t *testing.T) {
	e := emitter.NewEmitter()

	e.If(
		func(e *emitter.Emitter) error {
			e.PushBool(true)
			return nil
		},
		func(e *emitter.Emitter) error {
			e.PushInt(10)
			return nil
		},
		func(e *emitter.Emitter) error {
			e.PushInt(99)
			return nil
		},
	)

	testVM(t, e, object.CreateInt(10))
}

func TestIf_FalseBranch(t *testing.T) {
	e := emitter.NewEmitter()

	e.If(
		func(e *emitter.Emitter) error { // condition
			e.PushBool(false)
			return nil
		},
		func(e *emitter.Emitter) error { // then
			e.PushInt(10)
			return nil
		},
		func(e *emitter.Emitter) error { // else
			e.PushInt(99)
			return nil
		},
	)

	testVM(t, e, object.CreateInt(99))
}
func TestIf_True_NoElse(t *testing.T) {
	e := emitter.NewEmitter()

	e.If(
		func(e *emitter.Emitter) error { // condition
			e.PushBool(true)
			return nil
		},
		func(e *emitter.Emitter) error { // then
			e.PushInt(42)
			return nil
		},
		nil, // no else
	)

	testVM(t, e, object.CreateInt(42))
}

func TestIf_Nested(t *testing.T) {
	e := emitter.NewEmitter()

	e.If(
		func(e *emitter.Emitter) error { // outer condition
			e.PushBool(true)
			return nil
		},
		func(e *emitter.Emitter) error { // outer then
			return e.If(
				func(e *emitter.Emitter) error { // inner condition
					e.PushBool(false)
					return nil
				},
				func(e *emitter.Emitter) error { // inner then
					e.PushInt(111)
					return nil
				},
				func(e *emitter.Emitter) error { // inner else
					e.PushInt(222)
					return nil
				},
			)
		},
		func(e *emitter.Emitter) error { // outer else
			e.PushInt(333)
			return nil
		},
	)

	testVM(t, e, object.CreateInt(222))
}

func TestGlobalIntStoreLoad(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushInt(42)
	e.Store("x")
	e.PushInt(10)
	e.Load("x")

	expected := object.CreateInt(42)
	testVM(t, e, expected)
}

// update existing variable
func TestGlobalVariableOverwrite(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushInt(10)
	e.Store("x")
	e.PushInt(99)
	e.Store("x")
	e.PushInt(10)
	e.Load("x")

	expected := object.CreateInt(99)
	testVM(t, e, expected)
}

// multiple variables at once
func TestGlobalMultipleVariables(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushInt(7)
	e.Store("a")
	e.PushInt(8)
	e.Store("b")

	e.PushInt(2)
	e.Load("a")
	e.Load("b")
	e.AddInt()

	expected := object.CreateInt(15)
	testVM(t, e, expected)
}

// variable reuse and reassign after computation
func TestGlobalReuseAfterComputation(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushInt(5)
	e.Store("x")
	e.PushInt(3)

	e.Load("x")
	e.PushInt(3)
	e.AddInt()
	e.Store("x")
	e.PushInt(999)

	e.Load("x")

	expected := object.CreateInt(8)
	testVM(t, e, expected)
}

func TestFunction(t *testing.T) {
	e := emitter.NewEmitter()
	e.Function("test", []string{}, func(e *emitter.Emitter) error {
		e.PushInt(2)
		e.ReturnValue()
		return nil
	})
	e.Load("test")
	e.Call(0)

	expected := object.CreateInt(2)

	testVM(t, e, expected)

}

func TestFunctionWithArgs(t *testing.T) {
	e := emitter.NewEmitter()
	e.Function("test", []string{"x", "y"}, func(e *emitter.Emitter) error {
		e.Load("x")
		e.Load("y")
		e.AddInt()
		e.ReturnValue()
		return nil
	})
	e.PushInt(2)
	e.PushInt(2)

	e.Load("test")
	e.Call(2)

	expected := object.CreateInt(4)
	testVM(t, e, expected)
}

func TestLambdaWithNoReturn(t *testing.T) {
	e := emitter.NewEmitter()
	e.Lambda([]string{}, func(e *emitter.Emitter) error {
		e.PushInt(1)
		e.PushInt(1)
		e.Return()
		return nil
	})
	e.Call(0)

	testVMStackEmpty(t, e)
}
func TestLambdaWithReturn(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushInt(1)
	e.PushInt(1)
	e.Lambda([]string{"x", "y"}, func(e *emitter.Emitter) error {
		e.Load("x")
		e.Load("y")
		e.AddInt()
		e.ReturnValue()
		return nil
	})
	e.Call(2)
	expected := object.CreateInt(2)

	testVM(t, e, expected)
}

func TestFunctionWithReturn(t *testing.T) {
}

func TestClosure(t *testing.T) {
	e := emitter.NewEmitter()
	e.Function("newClosure", []string{"a"}, func(e *emitter.Emitter) error {
		e.Lambda([]string{}, func(e *emitter.Emitter) error {
			e.Load("a")
			e.ReturnValue()
			return nil
		})
		e.ReturnValue()
		return nil
	})
	e.PushInt(99)
	e.Load("newClosure")
	e.Call(1)
	e.Store("closure")
	e.Load("closure")
	e.Call(0)

	expected := object.CreateInt(99)

	testVM(t, e, expected)

}
func TestComplexClosure(t *testing.T) {
	e := emitter.NewEmitter()
	e.Function("newAdder", []string{"a", "b"}, func(e *emitter.Emitter) error {
		e.Lambda([]string{"c"}, func(e *emitter.Emitter) error {
			e.Load("a")
			e.Load("b")
			e.Load("c")
			e.AddInt()
			e.AddInt()
			e.ReturnValue()
			return nil
		})
		e.ReturnValue()
		return nil
	})

	e.PushInt(1)
	e.PushInt(2)
	e.Load("newAdder")
	e.Call(2)
	e.Store("adder")
	e.PushInt(8)
	e.Load("adder")
	e.Call(1)

	// b, c := e.Bytecode()
	// dump.Dump(b)
	//
	// dump.Constants(c)

	expected := object.CreateInt(11)

	testVM(t, e, expected)
}

func TestComplexClosure2(t *testing.T) {
	e := emitter.NewEmitter()
	e.Function("newAdder", []string{"a", "b"}, func(e *emitter.Emitter) error {
		e.Load("a")
		e.Load("b")
		e.AddInt()
		e.Store("c")
		e.Lambda([]string{"d"}, func(e *emitter.Emitter) error {
			e.Load("d")
			e.Load("c")
			e.AddInt()
			e.ReturnValue()
			return nil
		})
		e.ReturnValue()
		return nil
	})
	e.PushInt(1)
	e.PushInt(2)
	e.Load("newAdder")
	e.Call(2)
	e.Store("adder")
	e.PushInt(8)
	e.Load("adder")
	e.Call(1)

	expected := object.CreateInt(11)

	testVM(t, e, expected)
}

func TestWeirdClosure(t *testing.T) {
	e := emitter.NewEmitter()
	e.Function("newAdderOuter", []string{"a", "b"}, func(e *emitter.Emitter) error {
		e.Load("a")
		e.Load("b")
		e.AddInt()
		e.Store("c")
		e.Lambda([]string{"d"}, func(e *emitter.Emitter) error {
			e.Load("d")
			e.Load("c")
			e.AddInt()
			e.Store("e")
			e.Lambda([]string{"f"}, func(e *emitter.Emitter) error {
				e.Load("e")
				e.Load("f")
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
	e.PushInt(1)
	e.PushInt(2)
	e.Load("newAdderOuter")
	e.Call(2)
	e.Store("newAdderInner")
	e.PushInt(3)
	e.Load("newAdderInner")
	e.Call(1)
	e.Store("adder")
	e.PushInt(8)
	e.Load("adder")
	e.Call(1)

	// b, c := e.Bytecode()
	// dump.Dump(b)
	// dump.Constants(c)

	expected := object.CreateInt(14)

	testVM(t, e, expected)

}
func TestRecursion(t *testing.T) {
	t.Skip()
	e := emitter.NewEmitter()
	e.Function("fibonacci", []string{"x"}, func(e *emitter.Emitter) error {
		return e.If(func(e *emitter.Emitter) error {
			e.Load("x")
			e.PushInt(2)
			e.LtInt()
			return nil
		}, func(e *emitter.Emitter) error {
			e.Load("x")
			e.ReturnValue()
			return nil
		}, func(e *emitter.Emitter) error {
			e.Load("x")
			e.PushInt(1)
			e.SubInt()

			e.Load("fibonacci")
			e.Call(1)

			e.Load("x")
			e.PushInt(2)
			e.SubInt()

			e.Load("fibonacci")
			e.Call(1)

			e.AddInt()

			e.ReturnValue()
			return nil
		})
	})
	e.PushInt(10)
	e.Load("fibonacci")
	e.Call(1)

	expected := object.CreateInt(13)

	bytecode := e.Bytecode()

	dump.Dump(bytecode.Tape)
	dump.Constants(bytecode.Constants)

	testVM(t, e, expected)
}

// --------------------------------------------
// Array Tests
// --------------------------------------------

// simple integer array creation
func TestArrayBasic(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushInt(1)
	e.PushInt(2)
	e.PushInt(3)
	e.Array(3)

	expected := object.CreateArray([]object.Object{
		object.CreateInt(1),
		object.CreateInt(2),
		object.CreateInt(3),
	})
	testVM(t, e, expected)
}

// array indexing
func TestArrayIndex(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushInt(10)
	e.PushInt(20)
	e.PushInt(30)
	e.Array(3)

	e.PushInt(1) // index
	e.Index()

	expected := object.CreateInt(20)
	testVM(t, e, expected)
}

// nested arrays
func TestNestedArrays(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushInt(1)
	e.PushInt(2)
	e.Array(2) // [1,2]
	e.PushInt(3)
	e.Array(2) // [[1,2],3]

	expected := object.CreateArray([]object.Object{
		object.CreateArray([]object.Object{
			object.CreateInt(1),
			object.CreateInt(2),
		}),
		object.CreateInt(3),
	})
	testVM(t, e, expected)
}

// array with globals
func TestArrayWithGlobal(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushInt(42)
	e.Store("x")

	e.Load("x")
	e.PushInt(99)
	e.Array(2)
	e.Store("arr")
	e.Load("arr")

	expected := object.CreateArray([]object.Object{
		object.CreateInt(42),
		object.CreateInt(99),
	})
	testVM(t, e, expected)
}

// reuse array after mutation (conceptually re-store)
func TestArrayOverwriteGlobal(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushInt(1)
	e.PushInt(2)
	e.Array(2)
	e.Store("a")

	e.PushInt(100)
	e.PushInt(200)
	e.Array(2)
	e.Store("a")

	e.Load("a")

	expected := object.CreateArray([]object.Object{
		object.CreateInt(100),
		object.CreateInt(200),
	})
	testVM(t, e, expected)
}

// --------------------------------------------
// Hash Tests
// --------------------------------------------

// basic hash creation
func TestHashBasic(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushString("x")
	e.PushInt(10)
	e.PushString("y")
	e.PushInt(20)
	e.Hash(2)

	expected := object.CreateHash(map[object.Object]object.Object{
		object.CreateString("x"): object.CreateInt(10),
		object.CreateString("y"): object.CreateInt(20),
	})
	testVM(t, e, expected)
}

// // hash index by string key
func TestHashIndex(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushString("age")
	e.PushInt(27)
	e.Hash(1)

	e.PushString("age")
	e.Access()

	expected := object.CreateInt(27)
	testVM(t, e, expected)
}

// // nested hash inside array
func TestArrayOfHashes(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushString("a")
	e.PushInt(1)
	e.Hash(1)
	e.PushString("b")
	e.PushInt(2)
	e.Hash(1)
	e.Array(2)

	expected := object.CreateArray([]object.Object{
		object.CreateHash(
			map[object.Object]object.Object{
				object.CreateString("a"): object.CreateInt(1),
			}),
		object.CreateHash(
			map[object.Object]object.Object{
				object.CreateString("b"): object.CreateInt(2),
			}),
	})
	testVM(t, e, expected)
}

// // hash with globals and index lookup
func TestHashGlobalAndIndex(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushString("x")
	e.PushInt(5)
	e.PushString("y")
	e.PushInt(10)
	e.Hash(2)
	e.Store("h")

	e.Load("h")
	e.PushString("y")
	e.Access()

	expected := object.CreateInt(10)
	testVM(t, e, expected)
}

// // hash inside hash
func TestNestedHash(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushString("outer")
	e.PushString("inner")
	e.PushInt(123)
	e.Hash(1)

	e.Hash(1)

	expected := object.CreateHash(map[object.Object]object.Object{
		object.CreateString("outer"): object.CreateHash(map[object.Object]object.Object{
			object.CreateString("inner"): object.CreateInt(123),
		}),
	})
	testVM(t, e, expected)
}

func TestToFloat(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushInt(3)
	e.ToFloat()

	expected := object.CreateFloat(3)

	testVM(t, e, expected)
}

func TestPrint(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushString("hello, world")
	e.Builtin("print")
	e.Call(1)

	testVMStackEmpty(t, e)
}
func testVM(t *testing.T, e *emitter.Emitter, expected object.Object) {
	vm := NewVM(e.Bytecode())

	vm.Run()

	o := vm.Peek()
	assert.Equal(t, o, expected, "Result not equal!")
}

func testVMStackEmpty(t *testing.T, e *emitter.Emitter) {
	vm := NewVM(e.Bytecode())

	vm.Run()

	assert.Equal(t, vm.stackPointer, 0, "Stack not empty!")

}
