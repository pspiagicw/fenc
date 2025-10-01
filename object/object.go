package object

type CType string

const (
	INT    CType = "int"
	FLOAT  CType = "float"
	BOOL   CType = "bool"
	STRING CType = "string"
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

type Int struct {
	Value int
}

func (i Int) Type() CType {
	return INT
}
func (i Int) String() string {
	return "int"
}
func (i Int) Content() string {
	return "int"
}

type Float struct {
	Value float32
}

func (f Float) Type() CType {
	return FLOAT
}

func (f Float) String() string {
	return "float"
}
func (f Float) Content() string {
	return "float"
}

type Bool struct {
	Value bool
}

func (f Bool) Type() CType {
	return BOOL
}

func (f Bool) String() string {
	return "bool"
}
func (f Bool) Content() string {
	return "bool"
}

type String struct {
	Value string
}

func (f String) Type() CType {
	return STRING
}

func (f String) String() string {
	return "string"
}
func (f String) Content() string {
	return "string"
}
