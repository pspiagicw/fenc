package code

type Instruction struct {
	OP       int
	Argument int
}

const (
	PUSH = iota
	PEEK
)
