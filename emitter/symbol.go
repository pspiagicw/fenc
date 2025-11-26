package emitter

import (
	"fmt"

	"github.com/pspiagicw/fenc/object"
)

const (
	_ int = iota
	PRINT
	STRI
)

var BuiltinMap = map[int]object.Builtin{
	PRINT: {
		Internal: func(args []object.Object) object.Object {
			formatString := args[0]
			fmt.Println(formatString.String())

			return object.Null{}
		},
	},
	STRI: {
		Internal: func(args []object.Object) object.Object {
			value := args[0].String()

			return object.String{
				Value: value,
			}
		},
	},
}

type SymbolTable struct {
	Outer *SymbolTable

	store      map[string]Symbol
	storeIndex int
	Free       []Symbol
}

func NewSymbolTable() *SymbolTable {
	s := &SymbolTable{
		Outer:      nil,
		store:      map[string]Symbol{},
		storeIndex: 0,
		Free:       []Symbol{},
	}
	s.DefineBuiltin("print", PRINT)
	s.DefineBuiltin("stri", STRI)

	return s
}
func (s *SymbolTable) DefineBuiltin(name string, bid int) {
	b := Symbol{Name: name, Index: bid, Scope: BUILTIN_SCOPE}
	s.store[name] = b
}

func NewEnclosedSymbolTable(s *SymbolTable) *SymbolTable {
	table := NewSymbolTable()
	table.Outer = s
	return table
}

type Symbol struct {
	Name  string
	Index int
	Scope SymbolScope
}

type SymbolScope string

const (
	GLOBAL_SCOPE  SymbolScope = "GLOBAL"
	LOCAL_SCOPE   SymbolScope = "LOCAL"
	FREE_SCOPE    SymbolScope = "FREE"
	BUILTIN_SCOPE SymbolScope = "BUILTIN"
)

func (s *SymbolTable) Define(name string) Symbol {
	if existing, ok := s.Resolve(name); ok {
		// Use existing variable symbol
		return existing
	}
	symbol := Symbol{Name: name, Index: s.storeIndex, Scope: GLOBAL_SCOPE}
	if s.Outer != nil {
		symbol.Scope = LOCAL_SCOPE
	}

	s.store[name] = symbol
	s.storeIndex++
	return symbol
}

func (s *SymbolTable) Resolve(name string) (Symbol, bool) {
	obj, ok := s.store[name]
	if !ok && s.Outer != nil {
		obj, ok = s.Outer.Resolve(name)
		if !ok {
			return obj, ok
		}

		if obj.Scope == GLOBAL_SCOPE {
			return obj, ok
		}

		free := s.defineFree(obj)
		return free, true
	}
	return obj, ok
}
func (s *SymbolTable) defineFree(original Symbol) Symbol {
	s.Free = append(s.Free, original)

	symbol := Symbol{Name: original.Name, Index: len(s.Free) - 1}
	symbol.Scope = FREE_SCOPE

	s.store[original.Name] = symbol
	return symbol
}
