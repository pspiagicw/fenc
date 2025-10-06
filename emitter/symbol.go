package emitter

type SymbolTable struct {
	Outer *SymbolTable

	store      map[string]Symbol
	storeIndex int
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		Outer:      nil,
		store:      map[string]Symbol{},
		storeIndex: 0,
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
	}
	return obj, ok
}
