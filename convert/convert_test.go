package convert

import (
	"bytes"
	"io"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/pspiagicw/fenc/code"
	"github.com/pspiagicw/fenc/object"
)

func TestSimple(t *testing.T) {
	ins := []code.Instruction{
		{OpCode: code.ADD_INT},
	}

	bytecode := ConvertBytecode(ins)

	expected := []byte{2}

	assert.Equal(t, bytecode, expected, "Converted bytecode not matching.")
}

func TestPush(t *testing.T) {
	ins := []code.Instruction{
		{OpCode: code.PUSH, Args: []int{1}},
		{OpCode: code.PUSH, Args: []int{65535}},
	}
	bytecode := ConvertBytecode(ins)

	expected := []byte{1, 0, 1, 1, 255, 255}

	assert.Equal(t, bytecode, expected, "Converted bytecode not matching.")
}

func TestClosure(t *testing.T) {
	ins := []code.Instruction{
		{OpCode: code.CLOSURE, Args: []int{65535, 65535}},
	}
	bytecode := ConvertBytecode(ins)

	expected := []byte{33, 255, 255, 255, 255}

	assert.Equal(t, bytecode, expected, "Converted bytecode not matching.")

}

func TestInt(t *testing.T) {
	constants := []object.Object{
		object.CreateInt(1),
		object.CreateInt(2),
	}

	bytecode := ConvertConstants(constants)
	// TODO: Implement actual bytecode checking
	expected := bytecode

	assert.Equal(t, bytecode, expected, "Converted bytecode not matching.")
}
func TestBool(t *testing.T) {
	constants := []object.Object{
		object.CreateBool(false),
		object.CreateBool(true),
	}

	bytecode := ConvertConstants(constants)
	expected := []byte{0, 2, 4, 0, 4, 1}

	assert.Equal(t, bytecode, expected, "Converted bytecode not matching.")
}
func TestFloat(t *testing.T) {
	constants := []object.Object{
		object.CreateFloat(28.6),
		object.CreateFloat(67.2),
	}

	bytecode := ConvertConstants(constants)
	// TODO: Implement actual bytecode checking
	expected := bytecode

	assert.Equal(t, bytecode, expected, "Converted bytecode not matching.")
}
func TestString(t *testing.T) {
	constants := []object.Object{
		object.CreateString("this is a long ass string"),
		object.CreateString("really short string!"),
	}

	bytecode := ConvertConstants(constants)
	// TODO: Implement actual bytecode checking
	expected := bytecode

	assert.Equal(t, bytecode, expected, "Converted bytecode not matching.")
}
func TestFunction(t *testing.T) {
	constants := []object.Object{
		object.CreateFunction([]code.Instruction{
			{OpCode: code.ADD_FLOAT},
			{OpCode: code.PUSH, Args: []int{0}},
			{OpCode: code.CLOSURE, Args: []int{2, 1}},
		}),
	}
	bytecode := ConvertConstants(constants)
	// TODO: Implement actual bytecode checking
	expected := bytecode

	assert.Equal(t, bytecode, expected, "Converted bytecode not matching.")
}

func TestEncoding(t *testing.T) {
	instructions := []code.Instruction{
		{OpCode: code.PUSH, Args: []int{0}},
	}
	constants := []object.Object{}
	bytecode := Convert(instructions, constants)

	reader := bytes.NewReader(bytecode)
	buffer := make([]byte, 5)
	_, err := io.ReadFull(reader, buffer)
	assert.NoError(t, err, "Error while reading buffer")

	assert.Equal(t, buffer, []byte("FENCY"), "Magic bytes not matching.")

	buffer = make([]byte, 1)
	_, err = io.ReadFull(reader, buffer)
	assert.NoError(t, err, "Error while reading buffer")

	assert.Equal(t, buffer, []byte{1}, "Version number not matching")
}
