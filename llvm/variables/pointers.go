package variables

import "github.com/llir/llvm/ir"

type Pointers struct {
	p map[string]*ir.InstAlloca
}

func NewPointers() *Pointers {
	return &Pointers{
		p: make(map[string]*ir.InstAlloca),
	}
}

func (p *Pointers) get(name string) *ir.InstAlloca {
	return p.p[name]
}

func (p *Pointers) store(name string, point *ir.InstAlloca) {
	p.p[name] = point
}
