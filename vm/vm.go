package vm

import (
	"reflect"

	"github.com/pspiagicw/fenc/code"
	"github.com/pspiagicw/fenc/emitter"
	"github.com/pspiagicw/fenc/object"
	"github.com/pspiagicw/goreland"
)

const StackSize = 2048
const MaxFrames = 256

type Frame struct {
	tape       []code.Instruction
	ip         int
	locals     []object.Object
	oldPointer int
}

type VM struct {
	// ip           int
	// tape         []code.Instruction
	frames       []*Frame
	framePointer int
	stack        []object.Object
	stackPointer int
	constants    []object.Object
	globals      map[int]object.Object
}

func NewFrame(tape []code.Instruction) *Frame {
	return &Frame{
		tape: tape,
		ip:   0,
	}
}

func (vm *VM) currentFrame() *Frame {
	return vm.frames[vm.framePointer-1]
}

func (vm *VM) popFrame() *Frame {
	vm.framePointer -= 1
	return vm.frames[vm.framePointer]
}
func (vm *VM) pushFrame(f *Frame) {
	vm.frames[vm.framePointer] = f
	vm.framePointer++
}

func NewVM(e *emitter.Emitter) *VM {
	instructions, constants := e.Bytecode()
	frames := make([]*Frame, MaxFrames)
	frames[0] = NewFrame(instructions)
	return &VM{
		frames:       frames,
		stack:        make([]object.Object, StackSize),
		stackPointer: 0,
		constants:    constants,
		globals:      map[int]object.Object{},
		framePointer: 1,
	}
}
func (vm *VM) Run() {
	for vm.currentFrame().ip < len(vm.currentFrame().tape) {
		ins := vm.currentFrame().tape[vm.currentFrame().ip]
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
		case code.STORE_GLOBAL:
			vm.Store(ins.Args[0])
		case code.LOAD_GLOBAL:
			vm.Load(ins.Args[0])
		case code.CLOSURE:
			vm.Closure(ins.Args[0], ins.Args[1])
		case code.CALL:
			vm.Call(ins.Args[0])
		case code.RETURN:
			vm.Return()
		case code.RETURN_VALUE:
			vm.ReturnValue()
		case code.LOAD_LOCAL:
			vm.LoadLocal(ins.Args[0])
		default:
			goreland.LogFatal("Invalid Op: %s", ins.OpCode)
		}
		vm.currentFrame().ip += 1
	}
}
func (vm *VM) LoadLocal(id int) {
	o := vm.currentFrame().locals[id]
	vm.Push(o)
}
func (vm *VM) ReturnValue() {
	value := vm.Pop()
	vm.Return()
	vm.Push(value)
}
func (vm *VM) Return() {
	f := vm.popFrame()
	vm.stackPointer = f.oldPointer
}
func (vm *VM) Call(numArgs int) {
	fn := vm.PopClosure()

	args := make([]object.Object, numArgs)
	for i := 0; i < numArgs; i++ {
		args[i] = vm.stack[vm.stackPointer-numArgs+i]
	}

	newFrame := NewFrame(fn.Value.Value)
	newFrame.locals = args
	// The ip will be incremented automatically, thus we need to set it to -1, thus it will be incremented to 0.
	newFrame.ip = -1
	newFrame.oldPointer = vm.stackPointer

	vm.pushFrame(newFrame)

}
func (vm *VM) Closure(constId int, numFree int) {
	fn := vm.getConstant(constId)
	v, ok := fn.(object.Function)
	if !ok {
		goreland.LogFatal("Unable to resolve function from constant pool.")
	}

	// Load the variables in reverse.
	// Copy the variables
	// And reset the stackPointer to treat as if those variables didn't exist.
	free := make([]object.Object, numFree)
	for i := 0; i < numFree; i++ {
		free[i] = vm.stack[vm.stackPointer-numFree+i]
	}

	// Reset the stackPointer.
	vm.stackPointer = vm.stackPointer - numFree

	closure := object.Closure{
		Value: v,
		Free:  free,
	}
	vm.Push(closure)

}
func (vm *VM) Store(id int) {
	o := vm.Pop()
	vm.globals[id] = o
}
func (vm *VM) Load(id int) {
	val := vm.globals[id]
	vm.Push(val)
}
func (vm *VM) Jump(pos int) {
	vm.currentFrame().ip = pos - 1
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
func (vm *VM) PopClosure() object.Closure {
	o := vm.Pop()

	v, ok := o.(object.Closure)
	if !ok {
		goreland.LogFatal("Can't cast object to closure.")
	}

	return v
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
	if vm.stackPointer < 0 {
		goreland.LogFatal("Stack underflow!")
	}
	o := vm.stack[vm.stackPointer-1]
	vm.stackPointer -= 1
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
	if vm.stackPointer > StackSize {
		goreland.LogFatal("Stack Overflow!")
	}

	vm.stack[vm.stackPointer] = o
	vm.stackPointer += 1
}
func (vm *VM) PushConstant(index int) {
	o := vm.getConstant(index)

	vm.Push(o)

}
func (vm *VM) Peek() object.Object {
	return vm.stack[vm.stackPointer-1]
}
