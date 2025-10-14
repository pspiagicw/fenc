package convert

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/pspiagicw/fenc/code"
	"github.com/pspiagicw/fenc/object"
)

func TestSimple(t *testing.T) {
	ins := []code.Instruction{
		{OpCode: code.ADD_INT},
	}

	constants := []object.Object{}

	bytecode := Convert(ins, constants)

	expected := []byte{2}

	assert.Equal(t, bytecode, expected, "Converted bytecode not matching.")
}

func TestPush(t *testing.T) {
	ins := []code.Instruction{
		{OpCode: code.PUSH, Args: []int{1}},
		{OpCode: code.PUSH, Args: []int{65535}},
	}
	constants := []object.Object{}
	bytecode := Convert(ins, constants)

	expected := []byte{1, 0, 1, 1, 255, 255}

	assert.Equal(t, bytecode, expected, "Converted bytecode not matching.")
}

func TestClosure(t *testing.T) {
	ins := []code.Instruction{
		{OpCode: code.CLOSURE, Args: []int{65535, 65535}},
	}
	constants := []object.Object{}
	bytecode := Convert(ins, constants)

	expected := []byte{33, 255, 255, 255, 255}

	assert.Equal(t, bytecode, expected, "Converted bytecode not matching.")

}
