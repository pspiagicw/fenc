package main

import (
	"strings"

	"github.com/chzyer/readline"
	"github.com/pspiagicw/fenc/compile"
	"github.com/pspiagicw/fenc/decompile"
	"github.com/pspiagicw/goreland"
)

func getInput(r *readline.Instance) string {
	line, err := r.Readline()
	if err != nil {
		goreland.LogFatal("Error reading input from prompt: %v", err)
	}
	line = strings.TrimSpace(line)

	return line
}

func initREPL() *readline.Instance {
	r, err := readline.NewEx(&readline.Config{
		Prompt:          ">>> ",
		HistoryFile:     "/tmp/readline.tmp",
		InterruptPrompt: "^D",
	})

	if err != nil {
		goreland.LogFatal("Error initalizing readline: %v", err)
	}

	return r
}
func main() {

	r := initREPL()

	for {
		input := getInput(r)
		bytecode := compile.Compile(input)

		decompile.Print(bytecode)
	}

}
