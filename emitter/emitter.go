package emitter

import (
	"github.com/pspiagicw/fenc/code"
	"github.com/pspiagicw/fenc/object"
)

type Emitter struct {
	tape          []code.Instruction
	constants     []object.Object
	constantIndex int
}

func NewEmitter() *Emitter {
	return &Emitter{
		constants:     []object.Object{},
		constantIndex: 0,
		tape:          []code.Instruction{},
	}
}
func (e *Emitter) Emit(Op code.Op, Args ...int) error {
	return nil
}

func (e *Emitter) Constant(o object.Object) int {
	e.constants = append(e.constants, o)
	// alternative to return constantIndex++
	e.constantIndex += 1
	return e.constantIndex - 1
}

func (e *Emitter) PushInt(value int) {
	o := object.CreateInt(value)
	index := e.Constant(o)
	e.Emit(code.PUSH, index)
}
