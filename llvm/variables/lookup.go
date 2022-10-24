package variables

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/ir"
	irtypes "github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type LookupTable struct {
	state    map[string]int16 //ssa position
	pointers *Pointers
	values   map[string][]value.Value
	params   map[string][]value.Value
	types    map[string]irtypes.Type
}

func NewTable() *LookupTable {
	l := &LookupTable{
		state:    make(map[string]int16),
		pointers: NewPointers(),
		values:   make(map[string][]value.Value),
		params:   make(map[string][]value.Value),
		types:    make(map[string]irtypes.Type),
	}
	return l
}

func (l *LookupTable) List() []string {
	var vals []string
	for k := range l.values {
		vals = append(vals, k)
	}
	return vals
}

func (l *LookupTable) Add(id []string, val value.Value) {
	ident := strings.Join(id, "_")
	l.values[ident] = append(l.values[ident], val)
	l.state[ident] = 0
}

func (l *LookupTable) AddParam(id []string, p value.Value) {
	l.params[id[1]] = append(l.params[id[1]], p)
}

func (l *LookupTable) AddParams(id []string, p []value.Value) {
	l.params[id[1]] = append(l.params[id[1]], p...)
}

func (l *LookupTable) Store(id []string, name string, point *ir.InstAlloca) {
	ident := strings.Join(id, "_")
	if l.values[ident] != nil {
		l.pointers.store(name, point)
	} else {
		panic(fmt.Sprintf("variable %s not in the lookup table", strings.Join(id, "_")))
	}
}

func (l *LookupTable) Get(id []string) value.Value {
	ident := strings.Join(id, "_")
	i := len(l.values[ident]) - 1
	if i == -1 {
		return nil
	}
	return l.values[ident][i]
}

func (l *LookupTable) GetState(id []string) int16 {
	ident := strings.Join(id, "_")
	s, ok := l.state[ident]
	if !ok {
		panic(fmt.Sprintf("no state found for variable %s", ident))
	}
	return s
}

func (l *LookupTable) IncrState(id []string) {
	ident := strings.Join(id, "_")
	l.state[ident] = l.state[ident] + 1
}

func (l *LookupTable) ResetState(name string) {
	l.state[name] = 0
}

func (l *LookupTable) GetPointer(name string) *ir.InstAlloca {
	return l.pointers.get(name)
}

func (l *LookupTable) GetParams(id []string) []value.Value {
	return l.params[id[1]]
}

func (l *LookupTable) Update(id []string, val value.Value) {
	ident := strings.Join(id, "_")
	l.values[ident] = append(l.values[ident], val)
	l.state[ident] = l.state[ident] + 1
}

func (l *LookupTable) Type(id []string, ty irtypes.Type) {
	ident := strings.Join(id, "_")
	l.types[ident] = ty
}

func (l *LookupTable) GetType(name string) irtypes.Type {
	return l.types[name]
}
