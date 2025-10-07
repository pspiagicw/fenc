package vm

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/pspiagicw/fenc/emitter"
	"github.com/pspiagicw/fenc/object"
)

func TestPush(t *testing.T) {
	t.Skip()
	e := emitter.NewEmitter()
	e.PushInt(1)

	expected := object.CreateInt(1)

	testVM(t, e, expected)
}
func testVM(t *testing.T, e *emitter.Emitter, expected object.Object) {
	vm := NewVM(e)

	vm.Run()

	o := vm.Peek()
	assert.Equal(t, o, expected, "Result not equal!")
}
