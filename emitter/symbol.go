package emitter

type SymbolTable struct {
	Outer *SymbolTable

	store      map[string]Symbol
	storeIndex int
	Free       []Symbol
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		Outer:      nil,
		store:      map[string]Symbol{},
		storeIndex: 0,
		Free:       []Symbol{},
	}
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
	GLOBAL_SCOPE SymbolScope = "GLOBAL"
	LOCAL_SCOPE  SymbolScope = "LOCAL"
	FREE_SCOPE   SymbolScope = "FREE"
)

func (s *SymbolTable) Define(name string) Symbol {
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
