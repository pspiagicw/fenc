package vm

import (
	"testing"

	"github.com/alecthomas/assert/v2"
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
		func(e *emitter.Emitter) {
			e.PushBool(true)
		},
		func(e *emitter.Emitter) {
			e.PushInt(10)
		},
		func(e *emitter.Emitter) {
			e.PushInt(99)
		},
	)

	testVM(t, e, object.CreateInt(10))
}

func TestIf_FalseBranch(t *testing.T) {
	e := emitter.NewEmitter()

	e.If(
		func(e *emitter.Emitter) { // condition
			e.PushBool(false)
		},
		func(e *emitter.Emitter) { // then
			e.PushInt(10)
		},
		func(e *emitter.Emitter) { // else
			e.PushInt(99)
		},
	)

	testVM(t, e, object.CreateInt(99))
}
func TestIf_True_NoElse(t *testing.T) {
	e := emitter.NewEmitter()

	e.If(
		func(e *emitter.Emitter) { // condition
			e.PushBool(true)
		},
		func(e *emitter.Emitter) { // then
			e.PushInt(42)
		},
		nil, // no else
	)

	testVM(t, e, object.CreateInt(42))
}

func TestIf_Nested(t *testing.T) {
	e := emitter.NewEmitter()

	e.If(
		func(e *emitter.Emitter) { // outer condition
			e.PushBool(true)
		},
		func(e *emitter.Emitter) { // outer then
			e.If(
				func(e *emitter.Emitter) { // inner condition
					e.PushBool(false)
				},
				func(e *emitter.Emitter) { // inner then
					e.PushInt(111)
				},
				func(e *emitter.Emitter) { // inner else
					e.PushInt(222)
				},
			)
		},
		func(e *emitter.Emitter) { // outer else
			e.PushInt(333)
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
	e.Function("test", []string{}, func(e *emitter.Emitter) {
		e.PushInt(2)
		e.ReturnValue()
	})
	e.Load("test")
	e.Call(0)

	expected := object.CreateInt(2)

	testVM(t, e, expected)

}

func TestFunctionWithArgs(t *testing.T) {
	e := emitter.NewEmitter()
	e.Function("test", []string{"x", "y"}, func(e *emitter.Emitter) {
		e.Load("x")
		e.Load("y")
		e.AddInt()
		e.ReturnValue()
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
	e.Lambda([]string{}, func(e *emitter.Emitter) {
		e.PushInt(1)
		e.PushInt(1)
		e.Return()
	})
	e.Call(0)

	expected := object.CreateInt(2)

	testVMStackEmpty(t, e, expected)
}
func TestLambdaWithReturn(t *testing.T) {
	e := emitter.NewEmitter()
	e.PushInt(1)
	e.PushInt(1)
	e.Lambda([]string{"x", "y"}, func(e *emitter.Emitter) {
		e.Load("x")
		e.Load("y")
		e.AddInt()
		e.ReturnValue()
	})
	e.Call(2)
	expected := object.CreateInt(2)

	testVM(t, e, expected)
}

func TestFunctionWithReturn(t *testing.T) {
}

func testVM(t *testing.T, e *emitter.Emitter, expected object.Object) {
	vm := NewVM(e)

	vm.Run()

	o := vm.Peek()
	assert.Equal(t, o, expected, "Result not equal!")
}

func testVMStackEmpty(t *testing.T, e *emitter.Emitter, expected object.Object) {
	vm := NewVM(e)

	vm.Run()

	assert.Equal(t, vm.stackPointer, 0, "Stack not empty!")

}
