# `fenc`

`fenc` is a bytecode-compiler and runner. 
In simple words, it helps you write emit bytecode and run it, without worrying about internal implementation.

The API (and the bytecode) is for a stack-based typed VM.



For example, to add 2 integers. You can use the `fenc` API as follows

```go
import "github.com/pspiagicw/fenc/emitter"

e := emitter.NewEmitter()
e.PushInt(1)
e.PuhsInt(3)
e.AddInt()

// The stack should contain 4 on top.
```

Similarly a if-statement is as simple as

```go
// API implementation for 'if true then 5 else 4 end'
e.If(
    func(e *emitter.Emitter) error {
        // The condition
        e.PushBool(true)
    },
    func(e *emitter.Emitter) error {
        // The consequence
        e.PushInt(5)
    },
    func(e *emitter.Emitter) error {
        // The alternative
        e.PushInt(4)
    },
)

```

You don't have to worry about jmp instructions, nor do you have to perform calculations or batch-patch any bytecode.

`fenc` abstracts the bytecode emission for you, providing you with a simple, predictable functional API.

### Usage

This library is to be used along with a language-frontend. 
You will have to perform the lexing, parsing and typechecking yourself.

A robust, first-party implementation using `fenc` is [`tremor`](https://github.com/pspiagicw/tremor)

## Features

`fenc` being a stack-based bytecode, has public functions which work on the stack.

You can push objects (currently only primitive-types and functions/closures/builtins) onto the stack.

Any operation will pop the items on the stack and push the result back.

### Basics

You have 4 basic function calls to push values onto the stack.

- `PushInt()`
- `PushFloat()`
- `PushBool()`
- `PushString()`

```go
e.PushInt(4)
e.PushString("Hello, World")
e.PushBool(false)
e.PushFloat(3.1)
```

### Arithmetic

You have 4 types of artihmetic, on two types of data-types (int and float)

- `AddInt()/AddFloat()`
- `SubInt()/SubFloat()`
- `MulInt()/MulFloat()`
- `DivInt()/DivFloat()`

### Comparison

You again have 4 types of comparison on two data types.

- `LtInt()/LtFloat()`
- `LteInt()/LteFloat()`
- `GtInt()/GtFloat()`
- `GteInt()/GteFloat()`

### Logical

Logical operations only work on booleans.

- `AndBool()`
- `OrBool()`

### Special Comparison

These comparison operators directly compare the items on the stack.

- `Eq()`
- `Neq()`

### Misc

These are functions that are expected to be used by the library user, but not in regular operations.

For example

- `ToFloat()` 

It's only used to convert the integer value on the stack to a float.
Useful when evaluating mixed-type expressions.

- `AddString()`

Adding 2 strings together.

### Control Flow

Currently we support only `if` statements with/without else branches.

The function header is as follows:

`func (e *Emitter) If(cond, consequence, alternative CompileFunc) error`

Here `CompileFunc` is of type `func(*Emitter) error`
It accepts a emitter where you can emit/compile the condition, consequence and alternative.

The function will calculate the number of instructions to jump and insert the necessary jump instructions.

We plan to support the following conditionals in the future:
- `cond` (from lisp)
- `while` 
- `for`

Few examples:

- Standard if statement.
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

- If Without Else Branch

```go
e.If(
    func(e *emitter.Emitter) error { // condition
        e.PushBool(true)
        return nil
    },
    func(e *emitter.Emitter) error { // then
        e.PushInt(42)
        return nil
    },
    nil, // no else
)
```

- Nested If
```go
e.If(
    func(e *emitter.Emitter) error { // outer condition
        e.PushBool(true)
        return nil
    },
    func(e *emitter.Emitter) error { // outer then
        return e.If(
            func(e *emitter.Emitter) error { // inner condition
                e.PushBool(false)
                return nil
            },
            func(e *emitter.Emitter) error { // inner then
                e.PushInt(111)
                return nil
            },
            func(e *emitter.Emitter) error { // inner else
                e.PushInt(222)
                return nil
            },
        )
    },
    func(e *emitter.Emitter) error { // outer else
        e.PushInt(333)
        return nil
    },
)
```

### Variables

`fenc` abstract local and global variables, you simply invoke the `Store()` and `Load()` functions.
It takes care of using the local store or global store.

It even takes care of closures and free variables :) (Example shown later)

No need of implementing a symbol-table and tracking free-variables.

```go
e.PushInt(10)
e.Store("x")
e.Load("x")
```

### Functions

### Lambda

