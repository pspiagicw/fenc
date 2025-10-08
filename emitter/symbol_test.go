package emitter

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestFree(t *testing.T) {
	global := NewSymbolTable()
	global.Define("a")
	global.Define("b")

	firstLocal := NewEnclosedSymbolTable(global)
	firstLocal.Define("c")
	firstLocal.Define("d")

	secondLocal := NewEnclosedSymbolTable(firstLocal)
	secondLocal.Define("e")
	secondLocal.Define("f")

	tests := []struct {
		table               *SymbolTable
		expectedSymbols     []Symbol
		expectedFreeSymbols []Symbol
	}{
		{
			firstLocal,
			[]Symbol{
				Symbol{Name: "a", Scope: GLOBAL_SCOPE, Index: 0},
				Symbol{Name: "b", Scope: GLOBAL_SCOPE, Index: 1},
				Symbol{Name: "c", Scope: LOCAL_SCOPE, Index: 0},
				Symbol{Name: "d", Scope: LOCAL_SCOPE, Index: 1},
			},
			[]Symbol{},
		},
		{
			secondLocal,
			[]Symbol{
				Symbol{Name: "a", Scope: GLOBAL_SCOPE, Index: 0},
				Symbol{Name: "b", Scope: GLOBAL_SCOPE, Index: 1},
				Symbol{Name: "c", Scope: FREE_SCOPE, Index: 0},
				Symbol{Name: "d", Scope: FREE_SCOPE, Index: 1},
				Symbol{Name: "e", Scope: LOCAL_SCOPE, Index: 0},
				Symbol{Name: "f", Scope: LOCAL_SCOPE, Index: 1},
			},
			[]Symbol{
				Symbol{Name: "c", Scope: LOCAL_SCOPE, Index: 0},
				Symbol{Name: "d", Scope: LOCAL_SCOPE, Index: 1},
			},
		},
	}

	for _, tt := range tests {
		for _, sym := range tt.expectedSymbols {
			result, ok := tt.table.Resolve(sym.Name)
			if !ok {
				t.Errorf("name %s not resolvable", sym.Name)
			}
			assert.Equal(t, result, sym, "Symbols not matching")
		}

		assert.Equal(t, tt.expectedFreeSymbols, tt.table.Free)
	}
}
