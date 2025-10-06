package emitter

import "github.com/pspiagicw/fenc/object"

type ConstantPool struct {
	constants     []object.Object
	constantIndex int
}

func NewConstantPool() *ConstantPool {
	return &ConstantPool{
		constants:     []object.Object{},
		constantIndex: 0,
	}
}

func (c *ConstantPool) Add(o object.Object) int {
	c.constants = append(c.constants, o)
	c.constantIndex += 1
	return c.constantIndex - 1
}
