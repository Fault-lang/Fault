package variables

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

type LookupTable struct {
	state    map[string]State //ssa position
	pointers *Pointers
	values   map[string]Entry
	params   map[string][]value.Value
}

func NewTable() *LookupTable {
	l := &LookupTable{
		state:    make(map[string]State),
		pointers: NewPointers(),
		values:   make(map[string]Entry),
		params:   make(map[string][]value.Value),
	}
	return l
}

func (l *LookupTable) List() []string {
	var vals []string
	for k, v := range l.values {
		fmt.Print(k)
		switch ty := v.(type) {
		case *globalEntry:
			vals = append(vals, k)
		case *structEntry:
			fmt.Print("~~~STRUCT~~~")
			for k1, v1 := range ty.value {
				fmt.Print(k1)
				for k2, _ := range v1 {
					fmt.Print(k2)
					vals = append(vals, fmt.Sprint(k, "_", k1, "_", k2))
				}
			}
		case *instanceEntry:
			fmt.Print("~~~INST~~~")
			for k1, v1 := range ty.value {
				fmt.Print(k1)
				for k2, _ := range v1 {
					fmt.Print(k2)
					vals = append(vals, fmt.Sprint(k, "_", k1, "_", k2))
				}
			}
		}
	}
	return vals
}

func (l *LookupTable) Add(id []string, val value.Value) {
	switch len(id) {
	case 2: // a global variable
		l.values[id[1]] = NewGlobalEntry(val)
		l.state[id[1]] = NewGlobalSEntry(0)
	case 3: // a struct param
		switch st := l.values[id[1]].(type) {
		case *structEntry:
			st.update(id, val)
			l.state[id[1]].(*structSEntry).update(id, 0)
		case *instanceEntry: // In the run block
			st.update(id, val)
			l.state[id[1]].(*instanceSEntry).update(id, 0)
		case nil:
			l.values[id[1]] = NewStructEntry(id, val)
			l.state[id[1]] = NewStructSEntry(id, 0)
		}
	case 4: // an instance of a struct
		if l.values[id[1]] != nil {
			l.values[id[1]].(*instanceEntry).update(id, val)
			l.state[id[1]].(*instanceSEntry).update(id, 0)
		} else {
			l.values[id[1]] = NewInstanceEntry(id, val)
			l.state[id[1]] = NewInstanceSEntry(id, 0)
		}
	default:
		panic(fmt.Sprintf("cannot add %s to lookup table, missing full namespace. got=%d",
			strings.Join(id, "_"), len(id)))
	}
}

func (l *LookupTable) AddParam(id []string, p value.Value) {
	l.params[id[1]] = append(l.params[id[1]], p)
}

func (l *LookupTable) Store(id []string, name string, point *ir.InstAlloca) {
	switch l.values[id[1]].(type) {
	case *globalEntry:
		l.pointers.store(name, point)
	case *structEntry:
		l.pointers.store(name, point)
	case *instanceEntry:
		l.pointers.store(name, point)
	default:
		panic(fmt.Sprintf("variable %s not in the lookup table", strings.Join(id, "_")))
	}
}

func (l *LookupTable) Get(id []string) value.Value {
	switch v := l.values[id[1]].(type) {
	case *globalEntry:
		return v.get(id)
	case *structEntry:
		return v.get(id)
	case *instanceEntry:
		return v.get(id)
	}
	return nil
}

func (l *LookupTable) GetState(id []string) int16 {
	switch v := l.state[id[1]].(type) {
	case *globalSEntry:
		return v.get(id)
	case *structSEntry:
		return v.get(id)
	case *instanceSEntry:
		return v.get(id)
	}
	fid := strings.Join(id, "_")
	panic(fmt.Sprintf("no state found for variable %s", fid))
}

func (l *LookupTable) IncrState(id []string) {
	switch v := l.state[id[1]].(type) {
	case *structSEntry:
		v.increment(id)
	case *instanceSEntry:
		v.increment(id)
	}
}

func (l *LookupTable) ResetState(id []string) {
	switch v := l.state[id[1]].(type) {
	case *structSEntry:
		v.update(id, 0)
	case *instanceSEntry:
		v.update(id, 0)
	}
}

func (l *LookupTable) GetPointer(name string) *ir.InstAlloca {
	p := l.pointers.get(name)
	if p == nil {
		panic(fmt.Sprintf("no pointer found for variable %s", name))
	}
	return p
}

func (l *LookupTable) GetParams(id []string) []value.Value {
	return l.params[id[1]]
}

func (l *LookupTable) Update(id []string, val value.Value) {
	switch v := l.values[id[1]].(type) {
	case *structEntry:
		v.update(id, val)
		l.state[id[1]].(*structSEntry).increment(id)
	case *instanceEntry:
		v.update(id, val)
		l.state[id[1]].(*instanceSEntry).increment(id)
	}
}
