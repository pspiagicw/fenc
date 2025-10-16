package convert

import (
	"encoding/binary"
	"math"

	"github.com/pspiagicw/fenc/code"
	"github.com/pspiagicw/fenc/object"
	"github.com/pspiagicw/goreland"
)

type ConstKind int

const (
	None ConstKind = iota
	Int
	Float
	String
	Bool
	Function
)

var operandMap = map[code.Op]int{
	code.PUSH:         1,
	code.JUMP:         1,
	code.JUMP_FALSE:   1,
	code.CALL:         1,
	code.STORE_GLOBAL: 1,
	code.LOAD_GLOBAL:  1,
	code.LOAD_LOCAL:   1,
	code.STORE_LOCAL:  1,
	code.LOAD_FREE:    1,
	code.CLOSURE:      2,
}

func Convert(instructions []code.Instruction, constants []object.Object) []byte {

	buffer := []byte{}

	buffer = append(buffer, []byte("FENCY")...)
	buffer = append(buffer, byte(1))

	constBytecode := ConvertConstants(constants)
	actualBytecode := ConvertBytecode(instructions)
	insLength := uint32(len(actualBytecode))

	buffer = append(buffer, constBytecode...)
	buffer = binary.BigEndian.AppendUint32(buffer, insLength)
	buffer = append(buffer, actualBytecode...)

	return buffer
}

func ConvertConstants(constants []object.Object) []byte {

	buffer := []byte{}

	// Append number of constants
	buffer = binary.BigEndian.AppendUint16(buffer, uint16(len(constants)))

	for _, constant := range constants {
		switch constant := constant.(type) {
		case object.Int:
			buffer = convertInt(buffer, constant)
		case object.Bool:
			buffer = convertBool(buffer, constant)
		case object.Float:
			buffer = convertFloat(buffer, constant)
		case object.String:
			buffer = convertString(buffer, constant)
		case object.Function:
			buffer = convertFunction(buffer, constant)
		default:
			goreland.LogFatal("Unable to serialize constant: %v", constant)

		}
	}

	return buffer
}
func convertFunction(buffer []byte, constant object.Function) []byte {
	instructions := ConvertBytecode(constant.Value)
	insLength := uint16(len(instructions))

	buffer = append(buffer, byte(Function))
	buffer = binary.BigEndian.AppendUint16(buffer, insLength)
	buffer = append(buffer, instructions...)

	return buffer
}
func convertString(buffer []byte, constant object.String) []byte {
	length := uint16(len(constant.Value))
	strBytes := []byte(constant.Value)

	buffer = append(buffer, byte(String))
	buffer = binary.BigEndian.AppendUint16(buffer, length)
	buffer = append(buffer, strBytes...)

	return buffer
}
func convertFloat(buffer []byte, constant object.Float) []byte {
	val := math.Float32bits(constant.Value)
	buffer = append(buffer, byte(Float))
	buffer = binary.BigEndian.AppendUint32(buffer, val)

	return buffer
}
func convertInt(buffer []byte, constant object.Int) []byte {
	buffer = append(buffer, byte(Int))
	buffer = binary.BigEndian.AppendUint32(buffer, uint32(constant.Value))
	return buffer
}

func convertBool(buffer []byte, constant object.Bool) []byte {
	buffer = append(buffer, byte(Bool))
	if constant.Value {
		buffer = append(buffer, byte(1))
	} else {
		buffer = append(buffer, byte(0))
	}
	return buffer
}
func ConvertBytecode(tape []code.Instruction) []byte {
	buffer := []byte{}
	for _, ins := range tape {
		args, ok := operandMap[ins.OpCode]
		if !ok {
			args = 0
		}
		switch args {
		case 0:
			buffer = convertOp(buffer, ins)
		case 1:
			buffer = convertOp(buffer, ins)
			buffer = convertArg(buffer, ins.Args[0])
		case 2:
			buffer = convertOp(buffer, ins)
			buffer = convertArg(buffer, ins.Args[0])
			buffer = convertArg(buffer, ins.Args[1])
		default:
			goreland.LogFatal("Unknown number of operands!")
		}
	}

	return buffer
}
func convertOp(buffer []byte, ins code.Instruction) []byte {
	converted := int8(ins.OpCode)
	buffer = append(buffer, byte(converted))
	return buffer
}
func convertArg(buffer []byte, arg int) []byte {
	buffer = binary.BigEndian.AppendUint16(buffer, uint16(arg))
	return buffer
}
