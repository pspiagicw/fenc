package decompile

import (
	"fmt"

	"github.com/pspiagicw/fenc/code"
)

func Print(bytecode []*code.Instruction) {
	line := 0
	for _, instr := range bytecode {
		if instr == nil {
			continue
		}
		op := resolveOpCode(instr.OP)
		showArgument := ""

		if instr.Argument >= 0 {
			showArgument = fmt.Sprintf("%02d", instr.Argument)
		}

		fmt.Printf("%05d %s %s\n", line, op, showArgument)
	}
}

func resolveOpCode(opcode int) string {
	switch opcode {
	case 0:
		return "PUSH"
	case 1:
		return "PEEK"
	}

	return "invalid"
}
