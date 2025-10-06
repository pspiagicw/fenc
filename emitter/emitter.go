package emitter

import (
	"github.com/pspiagicw/fenc/code"
	"github.com/pspiagicw/fenc/object"
	"github.com/pspiagicw/goreland"
)

type CompileFunc func(*Emitter)

type Emitter struct {
	tape      []code.Instruction
	tapeIndex int
	constants *ConstantPool
	symbols   *SymbolTable
}

func (e *Emitter) enterScope() {
	e.symbols = NewEnclosedSymbolTable(e.symbols)
}
func (e *Emitter) leaveScope() {
	e.symbols = e.symbols.Outer
}
func (e *Emitter) NewSubEmitter() *Emitter {
	n := NewEmitter()
	n.constants = e.constants
	n.symbols = e.symbols

	return n
}

func NewEmitter() *Emitter {
	return &Emitter{
		constants: NewConstantPool(),
		tape:      []code.Instruction{},
		tapeIndex: 0,
		symbols:   NewSymbolTable(),
	}
}

func (e *Emitter) Emit(op code.Op, args ...int) int {
	ins := code.Instruction{
		OpCode: op,
		Args:   args,
	}
	e.tape = append(e.tape, ins)
	e.tapeIndex += 1
	return e.tapeIndex - 1
}

func (e *Emitter) Constant(o object.Object) int {
	return e.constants.Add(o)
}

func (e *Emitter) PushInt(value int) {
	o := object.CreateInt(value)
	index := e.Constant(o)
	e.Emit(code.PUSH, index)
}

func (e *Emitter) PushFloat(value float32) {
	o := object.CreateFloat(value)
	index := e.Constant(o)
	e.Emit(code.PUSH, index)
}

func (e *Emitter) PushBool(value bool) {
	o := object.CreateBool(value)
	index := e.Constant(o)
	e.Emit(code.PUSH, index)
}

func (e *Emitter) PushString(value string) {
	o := object.CreateString(value)
	index := e.Constant(o)
	e.Emit(code.PUSH, index)
}
func (e *Emitter) PushFunction(value object.Function) {
	index := e.Constant(value)
	e.Emit(code.PUSH, index)
}

func (e *Emitter) If(cond, consequence, alternative CompileFunc) {

	// Emit the condition
	cond(e)

	condPos := e.Emit(code.JUMP_FALSE, 0)

	consequence(e)

	jumpEndPos := -1
	if alternative != nil {
		jumpEndPos = e.Emit(code.JUMP, 0)
	}

	e.Patch(condPos)

	if alternative != nil {
		alternative(e)

		e.Patch(jumpEndPos)
	}

}
func (e *Emitter) Return() {
	e.Emit(code.RETURN)
}

func (e *Emitter) Patch(jumpPos int) {
	ins := e.tape[jumpPos]
	if ins.OpCode != code.JUMP && ins.OpCode != code.JUMP_FALSE {
		goreland.LogFatal("Given instructions is not a jump instruction.")
	}

	ins.Args = []int{e.tapeIndex}

	e.tape[jumpPos] = ins
}

func (e *Emitter) Store(name string) {
	s := e.symbols.Define(name)
	switch s.Scope {
	case GLOBAL_SCOPE:
		e.Emit(code.STORE_GLOBAL, s.Index)
	case LOCAL_SCOPE:
		e.Emit(code.STORE_LOCAL, s.Index)
	}
}

func (e *Emitter) Load(name string) bool {
	s, ok := e.symbols.Resolve(name)
	if !ok {
		return ok
	}

	switch s.Scope {
	case GLOBAL_SCOPE:
		e.Emit(code.LOAD_GLOBAL, s.Index)
	case LOCAL_SCOPE:
		e.Emit(code.LOAD_LOCAL, s.Index)
	}

	return ok
}

func (e *Emitter) Function(name string, args []string, body CompileFunc) {
	e.Lambda(args, body)
	e.Store(name)
}

func (e *Emitter) Lambda(args []string, body CompileFunc) {
	funcEmitter := e.NewSubEmitter()
	funcEmitter.enterScope()

	for _, arg := range args {
		funcEmitter.symbols.Define(arg)
	}

	body(funcEmitter)

	funcEmitter.leaveScope()

	fn := object.Function{
		Value: funcEmitter.tape,
	}
	e.PushFunction(fn)
}

func (e *Emitter) Call(args int) {
	e.Emit(code.CALL, args)
}

func (e *Emitter) AddInt() {
	e.Emit(code.ADD_INT)
}
func (e *Emitter) SubInt() {
	e.Emit(code.SUB_INT)
}
func (e *Emitter) MulInt() {
	e.Emit(code.MUL_INT)
}
func (e *Emitter) DivInt() {
	e.Emit(code.DIV_INT)
}

func (e *Emitter) LtInt() {
	e.Emit(code.LT_INT)
}
func (e *Emitter) LteInt() {
	e.Emit(code.LTE_INT)
}
func (e *Emitter) GtInt() {
	e.Emit(code.GT_INT)
}
func (e *Emitter) GteInt() {
	e.Emit(code.GTE_INT)
}

func (e *Emitter) Eq() {
	e.Emit(code.EQ)
}
func (e *Emitter) Neq() {
	e.Emit(code.NEQ)
}

func (e *Emitter) AddFloat() {
	e.Emit(code.ADD_FLOAT)
}
func (e *Emitter) SubFloat() {
	e.Emit(code.SUB_FLOAT)
}
func (e *Emitter) MulFloat() {
	e.Emit(code.MUL_FLOAT)
}
func (e *Emitter) DivFloat() {
	e.Emit(code.DIV_FLOAT)
}

func (e *Emitter) LtFloat() {
	e.Emit(code.LT_FLOAT)
}
func (e *Emitter) LteFloat() {
	e.Emit(code.LTE_FLOAT)
}
func (e *Emitter) GtFloat() {
	e.Emit(code.GT_FLOAT)
}
func (e *Emitter) GteFloat() {
	e.Emit(code.GTE_FLOAT)
}

func (e *Emitter) AndBool() {
	e.Emit(code.AND_BOOL)
}
func (e *Emitter) OrBool() {
	e.Emit(code.OR_BOOL)
}

func (e *Emitter) AddString() {
	e.Emit(code.ADD_STRING)
}
