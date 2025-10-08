package vm

import (
	"reflect"

	"github.com/pspiagicw/fenc/code"
	"github.com/pspiagicw/fenc/emitter"
	"github.com/pspiagicw/fenc/object"
	"github.com/pspiagicw/goreland"
)

const StackSize = 2048

type VM struct {
	ip         int
	tape       []code.Instruction
	stack      []object.Object
	stackIndex int
	constants  []object.Object
}

func NewVM(e *emitter.Emitter) *VM {
	instructions, constants := e.Bytecode()
	return &VM{
		tape:       instructions,
		ip:         0,
		stack:      make([]object.Object, StackSize),
		stackIndex: 0,
		constants:  constants,
	}
}
func (vm *VM) Run() {
	for vm.ip < len(vm.tape) {
		ins := vm.tape[vm.ip]
		switch ins.OpCode {
		case code.PUSH:
			vm.PushConstant(ins.Args[0])
		case code.ADD_INT:
			vm.AddInt()
		case code.MUL_INT:
			vm.MulInt()
		case code.SUB_INT:
			vm.SubInt()
		case code.DIV_INT:
			vm.DivInt()
		case code.LTE_INT:
			vm.LteInt()
		case code.LT_INT:
			vm.LtInt()
		case code.GTE_INT:
			vm.GteInt()
		case code.GT_INT:
			vm.GtInt()
		case code.ADD_FLOAT:
			vm.AddFloat()
		case code.SUB_FLOAT:
			vm.SubFloat()
		case code.MUL_FLOAT:
			vm.MulFloat()
		case code.DIV_FLOAT:
			vm.DivFloat()
		case code.LTE_FLOAT:
			vm.LteFloat()
		case code.LT_FLOAT:
			vm.LtFloat()
		case code.GTE_FLOAT:
			vm.GteFloat()
		case code.GT_FLOAT:
			vm.GtFloat()
		case code.OR_BOOL:
			vm.OrBool()
		case code.AND_BOOL:
			vm.AndBool()
		case code.NEQ:
			vm.Neq()
		case code.EQ:
			vm.Eq()
		case code.ADD_STRING:
			vm.AddString()
		case code.JUMP:
			vm.Jump(ins.Args[0])
		case code.JUMP_FALSE:
			vm.JumpFalse(ins.Args[0])
		default:
			goreland.LogError("Invalid Op: %s", ins.OpCode)
		}
		vm.ip += 1
	}
}
func (vm *VM) Jump(pos int) {
	vm.ip = pos - 1
}

func (vm *VM) JumpFalse(pos int) {
	v := vm.PopBool()
	if v.Value == false {
		vm.Jump(pos)
	}
}
func (vm *VM) AddFloat() {
	r := vm.PopFloat()
	l := vm.PopFloat()

	vm.Push(object.CreateFloat(l.Value + r.Value))
}
func (vm *VM) AddString() {
	r := vm.PopString()
	l := vm.PopString()

	vm.Push(object.CreateString(l.Value + r.Value))
}
func (vm *VM) Andbool() {
	r := vm.PopBool()
	l := vm.PopBool()

	vm.Push(object.CreateBool(l.Value && r.Value))
}
func (vm *VM) DivFloat() {
	r := vm.PopFloat()
	l := vm.PopFloat()

	vm.Push(object.CreateFloat(l.Value / r.Value))
}
func (vm *VM) DivInt() {
	r := vm.PopInt()
	l := vm.PopInt()

	vm.Push(object.CreateInt((int)(l.Value / r.Value)))
}
func (vm *VM) GtFloat() {
	r := vm.PopFloat()
	l := vm.PopFloat()

	vm.Push(object.CreateBool(l.Value > r.Value))
}
func (vm *VM) GteFloat() {
	r := vm.PopFloat()
	l := vm.PopFloat()

	vm.Push(object.CreateBool(l.Value >= r.Value))
}
func (vm *VM) GtInt() {
	r := vm.PopInt()
	l := vm.PopInt()

	vm.Push(object.CreateBool(l.Value > r.Value))
}
func (vm *VM) GteInt() {
	r := vm.PopInt()
	l := vm.PopInt()

	vm.Push(object.CreateBool(l.Value >= r.Value))
}
func (vm *VM) LtFloat() {
	r := vm.PopFloat()
	l := vm.PopFloat()

	vm.Push(object.CreateBool(l.Value < r.Value))
}
func (vm *VM) MulInt() {
	r := vm.PopInt()
	l := vm.PopInt()

	vm.Push(object.CreateInt(l.Value * r.Value))
}
func (vm *VM) LtInt() {
	r := vm.PopInt()
	l := vm.PopInt()

	vm.Push(object.CreateBool(l.Value < r.Value))
}
func (vm *VM) LteInt() {
	r := vm.PopInt()
	l := vm.PopInt()

	vm.Push(object.CreateBool(l.Value <= r.Value))
}
func (vm *VM) Neq() {
	r := vm.Pop()
	l := vm.Pop()

	// TODO: Replace this with a more performant one
	vm.Push(object.CreateBool(!reflect.DeepEqual(l, r)))
}
func (vm *VM) Eq() {
	r := vm.Pop()
	l := vm.Pop()

	// TODO: Replace this with a more performant one
	vm.Push(object.CreateBool(reflect.DeepEqual(l, r)))
}
func (vm *VM) AndBool() {
	r := vm.PopBool()
	l := vm.PopBool()

	vm.Push(object.CreateBool(l.Value && r.Value))
}
func (vm *VM) OrBool() {
	r := vm.PopBool()
	l := vm.PopBool()

	vm.Push(object.CreateBool(l.Value || r.Value))
}
func (vm *VM) AddInt() {
	r := vm.PopInt()
	l := vm.PopInt()

	vm.Push(object.CreateInt(l.Value + r.Value))
}

func (vm *VM) SubInt() {
	r := vm.PopInt()
	l := vm.PopInt()

	vm.Push(object.CreateInt(l.Value - r.Value))
}
func (vm *VM) SubFloat() {
	r := vm.PopFloat()
	l := vm.PopFloat()

	vm.Push(object.CreateFloat(l.Value - r.Value))
}
func (vm *VM) MulFloat() {
	r := vm.PopFloat()
	l := vm.PopFloat()

	vm.Push(object.CreateFloat(r.Value * l.Value))
}
func (vm *VM) LteFloat() {
	r := vm.PopFloat()
	l := vm.PopFloat()

	vm.Push(object.CreateBool(r.Value <= l.Value))
}

func (vm *VM) Pop() object.Object {
	o := vm.stack[vm.stackIndex-1]
	vm.stackIndex -= 1
	return o
}
func (vm *VM) PopInt() object.Int {
	o := vm.Pop()
	v, ok := o.(object.Int)
	if !ok {
		goreland.LogFatal("Expected object to be Integer, got %v", o)
	}

	return v
}
func (vm *VM) PopFloat() object.Float {
	o := vm.Pop()
	v, ok := o.(object.Float)
	if !ok {
		goreland.LogFatal("Expected object to be Float, got %v", o)
	}
	return v
}
func (vm *VM) PopBool() object.Bool {
	o := vm.Pop()
	v, ok := o.(object.Bool)
	if !ok {
		goreland.LogFatal("Expected object to be Bool, got %v", o)
	}
	return v
}
func (vm *VM) PopString() object.String {
	o := vm.Pop()
	v, ok := o.(object.String)
	if !ok {
		goreland.LogFatal("Expected object to be String, got %v", o)
	}
	return v
}
func (vm *VM) getConstant(index int) object.Object {
	return vm.constants[index]
}

func (vm *VM) Push(o object.Object) {
	if vm.stackIndex > StackSize {
		goreland.LogFatal("Stack Overflow!")
	}

	vm.stack[vm.stackIndex] = o
	vm.stackIndex += 1
}
func (vm *VM) PushConstant(index int) {
	o := vm.getConstant(index)

	vm.Push(o)

}
func (vm *VM) Peek() object.Object {
	return vm.stack[vm.stackIndex-1]
}
