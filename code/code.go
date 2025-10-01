package code

type Op int

const (
	_ = iota
	PUSH
)

type Instruction struct {
	OpCode  Op
	Args    []int
	Comment string
}
