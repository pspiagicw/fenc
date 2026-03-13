# `fenc`

`fenc` is a bytecode compiler and VM runtime for a stack-based, typed virtual machine.
In simple words, it helps you emit bytecode and run it without dealing with the low-level instruction patching yourself.

The API is designed to be used by a language frontend after lexing, parsing, and typechecking.

## Quick Example

To add two integers:

```go
package main

import (
	"github.com/pspiagicw/fenc/emitter"
	"github.com/pspiagicw/fenc/object"
)

func main() {
	builtins := map[string]object.Builtin{}
	e := emitter.NewEmitter(builtins)
	e.PushInt(1)
	e.PushInt(3)
	e.AddInt()

	// The stack will contain 4 on top after execution.
}
```

Similarly, an `if` statement is as simple as:

```go
// API implementation for: if true then 5 else 4 end
e.If(
	func(e *emitter.Emitter) error {
		e.PushBool(true)
		return nil
	},
	func(e *emitter.Emitter) error {
		e.PushInt(5)
		return nil
	},
	func(e *emitter.Emitter) error {
		e.PushInt(4)
		return nil
	},
)
```

You do not need to worry about jump instructions, offset calculations, or back-patching branches manually.

`fenc` abstracts bytecode emission for you and exposes a predictable, stack-oriented API.

## Usage

This library is meant to be used along with a language frontend.
You still perform lexing, parsing, and typechecking yourself.

A robust, first-party implementation using `fenc` is [`tremor`](https://github.com/pspiagicw/tremor).

## Builtins

Builtins are registered by name and passed into both the emitter and the VM:

```go
builtins := map[string]object.Builtin{
	"print": {
		Internal: func(args ...object.Object) object.Object {
			// side effect here
			return object.Null{}
		},
	},
}

e := emitter.NewEmitter(builtins)
vm := vm.NewVM(e.Bytecode(), builtins)
```

This matters because builtin names are assigned indexes during compilation and those same indexes are used again during execution.
In practice, use the same builtin map for both `NewEmitter(...)` and `vm.NewVM(...)`.

## Mental Model

`fenc` is stack-based.

- Push values onto the stack.
- Emit operations that consume values from the stack.
- The operation pushes its result back onto the stack.

For example:

```go
e.PushInt(4)
e.PushInt(2)
e.DivInt()
```

This emits bytecode for `4 / 2`, and the result is pushed back onto the stack.

## API Overview

Instead of thinking about the API as a long flat list of functions, it is easier to read it as a small set of operation families.

### Pushing Constants

Use these to put primitive values or compiled functions onto the stack.

| Category | Methods |
| --- | --- |
| Primitive values | `PushInt`, `PushFloat`, `PushBool`, `PushString` |
| Functions | `PushFunction` |

```go
e.PushInt(4)
e.PushString("Hello, World")
e.PushBool(false)
e.PushFloat(3.1)
```

### Arithmetic

Arithmetic follows a consistent naming scheme:

- `Add*`, `Sub*`, `Mul*`, `Div*`
- Type suffix: `Int` or `Float`

So the full set is:

- `AddInt`, `SubInt`, `MulInt`, `DivInt`
- `AddFloat`, `SubFloat`, `MulFloat`, `DivFloat`

### Comparison

Comparisons also follow the same pattern:

- `Lt*`, `Lte*`, `Gt*`, `Gte*`
- Type suffix: `Int` or `Float`

So the full set is:

- `LtInt`, `LteInt`, `GtInt`, `GteInt`
- `LtFloat`, `LteFloat`, `GtFloat`, `GteFloat`

### Equality, Boolean, and Unary Operations

These operations work directly on the current stack values.

| Category | Methods |
| --- | --- |
| Equality | `Eq`, `Neq` |
| Boolean logic | `AndBool`, `OrBool`, `Not` |
| Unary numeric ops | `NegateInt`, `NegateFloat` |
| Conversion and string ops | `ToFloat`, `AddString` |

`ToFloat` is useful when compiling mixed-type expressions.

## Control Flow

Currently, `fenc` supports `if` expressions with or without an `else` branch.

Function signature:

```go
func (e *Emitter) If(cond, consequence, alternative CompileFunc) error
```

Where `CompileFunc` is:

```go
type CompileFunc func(*Emitter) error
```

Each callback receives an emitter and emits the corresponding bytecode block.
`fenc` calculates jump positions and inserts the required branch instructions automatically.

Examples:

Standard `if`:

```go
e.If(
	func(e *emitter.Emitter) error {
		e.PushBool(true)
		return nil
	},
	func(e *emitter.Emitter) error {
		e.PushInt(10)
		return nil
	},
	func(e *emitter.Emitter) error {
		e.PushInt(99)
		return nil
	},
)
```

`if` without `else`:

```go
e.If(
	func(e *emitter.Emitter) error {
		e.PushBool(true)
		return nil
	},
	func(e *emitter.Emitter) error {
		e.PushInt(42)
		return nil
	},
	nil,
)
```

Nested `if`:

```go
e.If(
	func(e *emitter.Emitter) error {
		e.PushBool(true)
		return nil
	},
	func(e *emitter.Emitter) error {
		return e.If(
			func(e *emitter.Emitter) error {
				e.PushBool(false)
				return nil
			},
			func(e *emitter.Emitter) error {
				e.PushInt(111)
				return nil
			},
			func(e *emitter.Emitter) error {
				e.PushInt(222)
				return nil
			},
		)
	},
	func(e *emitter.Emitter) error {
		e.PushInt(333)
		return nil
	},
)
```

Planned future control-flow helpers:

- `cond`
- `while`
- `for`

## Variables

`fenc` abstracts local and global variables behind `Store` and `Load`.
It also handles closures and free variables internally.

You do not need to manually manage symbol lookup logic in normal use.

```go
e.PushInt(10)
e.Store("x")
e.Load("x")
```

## Functions and Calls

Function support is also grouped around a few core operations:

| Purpose | Methods |
| --- | --- |
| Named functions | `Function` |
| Anonymous functions | `Lambda` |
| Call site | `Call`, `Load` |
| Returning | `Return`, `ReturnValue` |

Typical pattern:

```go
e.Function("addOne", []string{"x"}, func(e *emitter.Emitter) error {
	e.Load("x")
	e.PushInt(1)
	e.AddInt()
	e.ReturnValue()
	return nil
})

e.Load("addOne")
e.PushInt(41)
e.Call(1)
```

Builtin calls follow the same pattern:

```go
e.Load("print")
e.PushString("hello")
e.Call(1)
```

## Collections and Object-Like Values

`fenc` also supports collection-building and lookup instructions.

| Category | Methods |
| --- | --- |
| Arrays | `Array`, `Index` |
| Hash maps | `Hash`, `Access` |
| Classes | `Class` |

These are lower-level building blocks intended for frontends that need structured runtime values.

Note: `classes` are WIP.

## Bytecode and Execution

Once emission is complete:

- `Bytecode()` returns the instruction tape and constant pool.
- `Errors()` returns any emitter errors collected during compilation.

You can then execute the bytecode with the VM package.

```go
bytecode := e.Bytecode()
machine := vm.NewVM(bytecode, builtins)
machine.Run()
```

## API Reference by Workflow

If you are integrating `fenc` into a compiler, the most commonly used methods are:

- Value emission: `PushInt`, `PushFloat`, `PushBool`, `PushString`
- Variables: `Store`, `Load`
- Arithmetic and comparison: the `*Int` and `*Float` operator families
- Control flow: `If`
- Functions: `Function`, `Lambda`, `Call`, `ReturnValue`
- Containers: `Array`, `Hash`, `Index`, `Access`
- Finalization: `Bytecode`, `Errors`

That is usually enough to compile a small language frontend without thinking in terms of raw opcodes.
