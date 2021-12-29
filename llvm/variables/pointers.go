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

// type Pointer interface{}

// type globalPointer struct {
// 	Pointer
// 	p *ir.InstAlloca
// }

// func NewGlobalPointer(point *ir.InstAlloca) *globalPointer {
// 	return &globalPointer{
// 		p: point,
// 	}
// }

// func (gp *globalPointer) get(id []string) *ir.InstAlloca {
// 	return gp.p
// }

// type structPointer struct {
// 	Pointer
// 	p map[string]map[string]*ir.InstAlloca
// }

// func NewStructPointer(id []string, point *ir.InstAlloca) *structPointer {
// 	s := &structPointer{
// 		p: make(map[string]map[string]*ir.InstAlloca),
// 	}
// 	s.p[id[1]] = make(map[string]*ir.InstAlloca)
// 	s.p[id[1]][id[2]] = point
// 	return s
// }

// func (sp *structPointer) get(id []string) *ir.InstAlloca {
// 	return sp.p[id[1]][id[1]]
// }

// type instancePointer struct {
// 	Pointer
// 	p map[string]map[string]map[string]*ir.InstAlloca
// }

// func NewInstancePointer(id []string, point *ir.InstAlloca) *instancePointer {
// 	i := &instancePointer{
// 		p: make(map[string]map[string]map[string]*ir.InstAlloca),
// 	}
// 	i.p[id[1]] = make(map[string]map[string]*ir.InstAlloca)
// 	i.p[id[1]][id[2]] = make(map[string]*ir.InstAlloca)
// 	i.p[id[1]][id[2]][id[3]] = point
// 	return i
// }

// func (ip *instancePointer) get(id []string) *ir.InstAlloca {
// 	return ip.p[id[1]][id[2]][id[3]]
// }
