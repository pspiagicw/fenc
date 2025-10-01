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
func createArgs(args ...int) []int {
	return args
}

func getEmitter() *Emitter {
	e := NewEmitter()

	return e
}

func testEmitter(t *testing.T, e *Emitter, expected []code.Instruction, constants []object.Object) {
	assert.Equal(t, constants, e.constants, "Constant pool not equal!")
	assert.Equal(t, expected, e.tape, "Instructions on tape differ!")
}
