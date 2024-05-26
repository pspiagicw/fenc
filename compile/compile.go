package compile

import (
	"strconv"
	"strings"

	"github.com/pspiagicw/fenc/code"
	"github.com/pspiagicw/goreland"
)

func Compile(contents string) []*code.Instruction {

	bytecode := []*code.Instruction{}

	lines := strings.Split(contents, "\n")

	for _, line := range lines {
		bytecode = append(bytecode, parseLine(line))

	}

	return bytecode
}
func parseLine(line string) *code.Instruction {

	parts := strings.Split(line, " ")

	if len(parts) == 0 {
		return nil
	}

	op := parts[0]

	switch op {
	case "push":
		return compilePush(parts)
	case "peek":
		return compilePeek()
	}

	goreland.LogError("Can't parse line %q", line)
	return nil
}
func compilePush(parts []string) *code.Instruction {
	if len(parts) != 2 {
		goreland.LogError("wrong number of arguments for push, got %v", len(parts)-1)
		return nil
	}

	value, err := strconv.ParseInt(parts[1], 10, 32)

	if err != nil {
		goreland.LogError("could not parse %q as integer", parts[1])
		return nil
	}

	return makePush(int(value))
}
func compilePeek() *code.Instruction {
	return makePeek()
}
func makePush(value int) *code.Instruction {
	return &code.Instruction{
		OP:       code.PUSH,
		Argument: value,
	}
}
func makePeek() *code.Instruction {
	return &code.Instruction{
		OP:       code.PEEK,
		Argument: -1,
	}
}
