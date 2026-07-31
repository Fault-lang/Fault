package variables

import "github.com/llir/llvm/ir/value"

type Pointers struct {
	p map[string]value.Value
}

func NewPointers() *Pointers {
	return &Pointers{
		p: make(map[string]value.Value),
	}
}

func (p *Pointers) get(name string) value.Value {
	return p.p[name]
}

func (p *Pointers) store(name string, point value.Value) {
	p.p[name] = point
}
