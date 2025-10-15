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

	ADD_FLOAT
	SUB_FLOAT
	MUL_FLOAT
	DIV_FLOAT

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

	RETURN
	RETURN_VALUE
	CALL

	STORE_GLOBAL
	STORE_LOCAL

	LOAD_GLOBAL
	LOAD_LOCAL
	LOAD_FREE

	CLOSURE

	ARRAY
	HASH
	INDEX
	ACCESS
)

type Instruction struct {
	OpCode  Op
	Args    []int
	Comment string
}
