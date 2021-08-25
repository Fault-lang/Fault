package variables

import (
	"github.com/llir/llvm/ir/value"
)

type Entry interface{}

type globalEntry struct {
	Entry
	value value.Value
}

func NewGlobalEntry(val value.Value) *globalEntry {
	return &globalEntry{
		value: val,
	}
}

func (ge *globalEntry) get(id []string) value.Value {
	return ge.value
}

type structEntry struct {
	Entry
	value map[string]map[string][]value.Value
}

func NewStructEntry(id []string, val value.Value) *structEntry {
	s := &structEntry{
		value: make(map[string]map[string][]value.Value),
	}
	s.value[id[1]] = make(map[string][]value.Value)
	s.value[id[1]][id[2]] = []value.Value{val}
	return s
}

func (se *structEntry) get(id []string) value.Value {
	if len(se.value[id[1]][id[2]]) == 0 {
		return nil
	}
	return se.value[id[1]][id[2]][len(se.value[id[1]][id[2]])-1]
}

func (se *structEntry) update(id []string, val value.Value) {
	se.value[id[1]][id[2]] = append(se.value[id[1]][id[2]], val)
}

type instanceEntry struct {
	Entry
	value map[string]map[string]map[string][]value.Value
}

func NewInstanceEntry(id []string, val value.Value) *instanceEntry {
	s := &instanceEntry{
		value: make(map[string]map[string]map[string][]value.Value),
	}
	s.value[id[1]] = make(map[string]map[string][]value.Value)
	s.value[id[1]][id[2]] = make(map[string][]value.Value)
	s.value[id[1]][id[2]][id[3]] = []value.Value{val}
	return s
}

func (ie *instanceEntry) get(id []string) value.Value {
	if ie.value[id[1]][id[2]] == nil {
		return nil
	}
	if len(ie.value[id[1]][id[2]][id[3]]) == 0 {
		return nil
	}
	return ie.value[id[1]][id[2]][id[3]][len(ie.value[id[1]][id[2]][id[3]])-1]
}

func (ie *instanceEntry) update(id []string, val value.Value) {
	ie.value[id[1]][id[2]][id[3]] = append(ie.value[id[1]][id[2]][id[3]], val)
}
