package object

type CType string

const (
	INT CType = "int"
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
