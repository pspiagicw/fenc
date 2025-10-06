package vm

import (
	"github.com/pspiagicw/fenc/code"
	"github.com/pspiagicw/fenc/emitter"
	"github.com/pspiagicw/fenc/object"
)

type VM struct {
	ip        int
	tape      []code.Instruction
	stack     []object.Object
	constants []object.Object
}

func NewVM(e *emitter.Emitter) *VM {
	return &VM{
		tape:  e.Bytecode(),
		ip:    0,
		stack: []object.Object{},
		// constants: e.Constant(),
	}
}
func (vm *VM) Run() {
	for vm.ip < len(vm.tape) {
		ins := vm.tape[vm.ip]
		switch ins.OpCode {
		case code.PUSH:
		}
		vm.ip += 1
	}
}
func (v *VM) Peek() object.Object {
	return nil
}
