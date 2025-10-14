package convert

import (
	"encoding/binary"

	"github.com/pspiagicw/fenc/code"
	"github.com/pspiagicw/fenc/object"
	"github.com/pspiagicw/goreland"
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

func Convert(tape []code.Instruction, constants []object.Object) []byte {
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
