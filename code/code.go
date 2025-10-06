package code

//go:generate stringer -type=Op

type Op int

const (
	_ Op = iota
	PUSH
	ADD_INT
	SUB_INT
	MUL_INT
	DIV_INT

	LT_INT
	LTE_INT
	GT_INT
	GTE_INT

	// EQ_INT
	// NEQ_INT

	ADD_FLOAT
	SUB_FLOAT
	MUL_FLOAT
	DIV_FLOAT

	// EQ_FLOAT
	// NEQ_FLOAT

	// EQ_BOOL
	// NEQ_BOOL

	AND_BOOL
	OR_BOOL

	EQ
	NEQ

	LT_FLOAT
	LTE_FLOAT
	GT_FLOAT
	GTE_FLOAT

	ADD_STRING

	JUMP
	JUMP_FALSE
)

type Instruction struct {
	OpCode  Op
	Args    []int
	Comment string
}
