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
		case code.LTE_INT:
			vm.LteInt()
		case code.LT_INT:
			vm.LtInt()
		case code.SUB_FLOAT:
			vm.SubFloat()
		case code.MUL_FLOAT:
			vm.MulFloat()
		case code.LTE_FLOAT:
			vm.LteFloat()
		case code.OR_BOOL:
			vm.OrBool()
		case code.NEQ:
			vm.Neq()
		default:
			goreland.LogError("Invalid Op: %s", ins.OpCode)
		}
		vm.ip += 1
	}
}
func (vm *VM) MulInt() {
	l := vm.PopInt()
	r := vm.PopInt()

	vm.Push(object.CreateInt(l.Value * r.Value))
}
func (vm *VM) LtInt() {
	l := vm.PopInt()
	r := vm.PopInt()

	vm.Push(object.CreateBool(r.Value < l.Value))
}
func (vm *VM) LteInt() {
	l := vm.PopInt()
	r := vm.PopInt()

	vm.Push(object.CreateBool(l.Value <= r.Value))
}
func (vm *VM) Neq() {
	l := vm.Pop()
	r := vm.Pop()

	// TODO: Replace this with a more performant one
	vm.Push(object.CreateBool(!reflect.DeepEqual(l, r)))
}
func (vm *VM) Eq() {
	l := vm.Pop()
	r := vm.Pop()

	// TODO: Replace this with a more performant one
	vm.Push(object.CreateBool(reflect.DeepEqual(l, r)))
}
func (vm *VM) AndBool() {
	l := vm.PopBool()
	r := vm.PopBool()

	vm.Push(object.CreateBool(l.Value && r.Value))
}
func (vm *VM) OrBool() {
	l := vm.PopBool()
	r := vm.PopBool()

	vm.Push(object.CreateBool(l.Value || r.Value))
}
func (vm *VM) AddInt() {
	l := vm.PopInt()
	r := vm.PopInt()

	vm.Push(object.CreateInt(l.Value + r.Value))
}

func (vm *VM) SubInt() {
	l := vm.PopInt()
	r := vm.PopInt()

	vm.Push(object.CreateInt(r.Value - l.Value))
}
func (vm *VM) SubFloat() {
	l := vm.PopFloat()
	r := vm.PopFloat()

	vm.Push(object.CreateFloat(r.Value - l.Value))
}
func (vm *VM) MulFloat() {
	l := vm.PopFloat()
	r := vm.PopFloat()

	vm.Push(object.CreateFloat(r.Value * l.Value))
}
func (vm *VM) LteFloat() {
	l := vm.PopFloat()
	r := vm.PopFloat()

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
