package object

import (
	"fmt"
	"strconv"

	"github.com/pspiagicw/fenc/code"
)

type CType string

const (
	INT      CType = "int"
	FLOAT    CType = "float"
	BOOL     CType = "bool"
	STRING   CType = "string"
	FUNCTION CType = "function"
	CLOSURE  CType = "CLOSURE"
)

type Object interface {
	Type() CType
	String() string
	Content() string
}

func CreateInt(value int) Object {
	return Int{
		Value: value,
	}
}

func CreateFloat(value float32) Object {
	return Float{
		Value: value,
	}
}

func CreateBool(value bool) Object {
	return Bool{
		Value: value,
	}
}

func CreateString(value string) Object {
	return String{
		Value: value,
	}
}

func CreateFunction(value []code.Instruction) Object {
	return Function{
		Value: value,
	}
}

type Int struct {
	Value int
}

func (i Int) Type() CType {
	return INT
}
func (i Int) String() string {
	return strconv.Itoa(i.Value)
}
func (i Int) Content() string {
	return strconv.Itoa(i.Value)
}

type Float struct {
	Value float32
}

func (f Float) Type() CType {
	return FLOAT
}

func (f Float) String() string {
	return fmt.Sprintf("%f", f.Value)
}
func (f Float) Content() string {
	return fmt.Sprintf("%f", f.Value)
}

type Bool struct {
	Value bool
}

func (f Bool) Type() CType {
	return BOOL
}

func (f Bool) String() string {
	return fmt.Sprintf("%t", f.Value)
}
func (f Bool) Content() string {
	return fmt.Sprintf("%t", f.Value)
}

type String struct {
	Value string
}

func (f String) Type() CType {
	return STRING
}

func (f String) String() string {
	return f.Value
}
func (f String) Content() string {
	return f.Value
}

type Function struct {
	Value []code.Instruction
}

func (f Function) Type() CType {
	return FUNCTION
}
func (f Function) String() string {
	return "instructions"
}
func (f Function) Content() string {
	return "instructions"
}

type Closure struct {
	Value Function
	Free  []Object
}

func (c Closure) Type() CType {
	return CLOSURE
}
func (c Closure) String() string {
	return "closure"
}
func (c Closure) Content() string {
	return "closure"
}
