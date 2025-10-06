package dump

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/pspiagicw/fenc/code"
	"github.com/pspiagicw/fenc/object"
	"github.com/pspiagicw/goreland"
)

var compiledStyle lipgloss.Style = lipgloss.NewStyle().Border(lipgloss.NormalBorder())
var constantStyle lipgloss.Style = lipgloss.NewStyle().Border(lipgloss.DoubleBorder()).Padding(1).MarginRight(1)
var codeStyle lipgloss.Style = lipgloss.NewStyle().Padding(1).Border(lipgloss.NormalBorder())

func Dump(bytecode []code.Instruction) {
	// code := codeStyle.Render(printInstructions(bytecode))

	fmt.Println("-----")
	fmt.Println(InstructionToString(bytecode))
	fmt.Println("-----")
	// fmt.Println(lipgloss.JoinHorizontal(lipgloss.Top, code))
}

func InstructionToString(bytecode []code.Instruction) string {
	var buffer strings.Builder
	line := 0
	for _, instruction := range bytecode {
		op := instruction.OpCode.String()

		lineNumber := getLineNumber(line)
		args := []string{}
		for _, arg := range instruction.Args {
			args = append(args, strconv.Itoa(arg))
		}
		argString := strings.Join(args, " ")
		buffer.WriteString(fmt.Sprintf("%s %s %s\t%s\n", lineNumber, op, argString, instruction.Comment))
		line++
	}
	return strings.TrimSpace(buffer.String())
}

func getLineNumber(line int) string {
	return lipgloss.NewStyle().Faint(true).Render(fmt.Sprintf("%05d", line))
}
func Constants(constants []object.Object) {
	for _, o := range constants {
		if o.Type() == object.FUNCTION {
			v, ok := o.(object.Function)
			if !ok {
				goreland.LogFatal("Can't convert constant to function!")
			}
			Dump(v.Value)
		} else {
			fmt.Println(o.String())
		}
	}

}
