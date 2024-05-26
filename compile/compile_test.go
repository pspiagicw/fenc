package compile

import (
	"slices"
	"testing"

	"github.com/pspiagicw/fenc/code"
)

func TestCompile(t *testing.T) {
	tt := []struct {
		input    string
		expected []*code.Instruction
	}{
		{
			"push 2",
			makeInstructionArr(makePush(2)),
		},
		{
			"push 65534",
			makeInstructionArr(makePush(65534)),
		},
		{
			"peek",
			makeInstructionArr(makePeek()),
		},
	}

	for _, tt := range tt {
		actual := Compile(tt.input)
		expected := tt.expected

		checkInstructions(t, actual, expected)
	}
}
func makeInstructionArr(ins ...*code.Instruction) []*code.Instruction {
	return ins
}
func checkInstructions(t *testing.T, actual, expected []*code.Instruction) {

	t.Helper()

	if len(actual) != len(expected) {
		t.Errorf("Length of instructions don't match, expected '%d', got '%d'", len(expected), len(actual))
	}

	for i, a := range actual {
		e := expected[i]

		if a.OP != e.OP {
			t.Errorf("OPCODE for instruction doesn't match, expected '%d', got '%d'", e.OP, a.OP)
		}

		if !slices.Equal(a.Arguments, e.Arguments) {
			t.Errorf("Arguments of instruction doesn't match, expected '%v', got '%v'", e.Arguments, a.Arguments)

		}

	}
}
