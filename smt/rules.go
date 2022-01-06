package smt

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/value"
)

func (g *Generator) constantRule(id string, c constant.Constant) string {
	switch val := c.(type) {
	case *constant.Float:
		id = g.advanceSSA(id)
		return g.writeInitRule(id, "Real", val.String())
	}
	return ""
}

func (g *Generator) loadsRule(inst *ir.InstLoad) {
	id := inst.Ident()
	g.loads[id] = inst.Src
}

// func (g *Generator) trackRounds(id string, inst *ir.InstStore) {
// 	for _, v := range inst.Metadata {
// 		if v.Name[0:5] == "rounds-" {
// 			i := g.ssa[id]
// 			if g.rounds[v.Name[5:]] == nil {
// 				g.rounds[v.Name[5:]] = map[string][]int16{id: []int16{i}}
// 			} else {
// 				g.rounds[v.Name[5:]][id] = append(g.rounds[v.Name[5:]][id], i)
// 			}
// 		}
// 	}
// }

func (g *Generator) storeRule(inst *ir.InstStore, rules []rule) []rule {
	id := g.formatIdent(inst.Dst.Ident())
	if g.isTemp(inst.Src.Ident()) {
		srcId := inst.Src.Ident()
		if val, ok := g.loads[srcId]; ok {
			ty := g.getType(val)
			id = g.advanceSSA(id)
			//g.trackRounds(id, inst)
			rules = append(rules, g.parseRule(id, g.formatValue(val), ty, ""))
		} else if ref, ok := g.ref[srcId]; ok {
			switch r := ref.(type) {
			case *infix:
				r.x = g.tempToIdent(r.x)
				r.y = g.tempToIdent(r.y)
				id = g.advanceSSA(id)
				//g.trackRounds(id, inst)
				wid := &wrap{value: id}
				rules = append(rules, &infix{x: wid, ty: "Real", y: r})
			default:
				id = g.advanceSSA(id)
				//g.trackRounds(id, inst)
				wid := &wrap{value: id}
				rules = append(rules, &infix{x: wid, ty: "Real", y: r})
			}
		} else {
			panic(fmt.Sprintf("smt generation error, value for %s not found", id))
		}
	} else {
		ty := g.getType(inst.Src)
		id = g.advanceSSA(id)
		//g.trackRounds(id, inst)
		rules = append(rules, g.parseRule(id, inst.Src.Ident(), ty, ""))
	}
	g.last = rules[len(rules)-1]
	return rules
}

func (g *Generator) parallelMeta(parallelGroup string, meta ir.Metadata) string {
	for _, v := range meta {
		if v.Name != parallelGroup {
			g.call = g.call + 1
		}
		if v.Name[0:5] != "round-" {
			g.parallel = v.Name
		}
	}
	return g.parallel
}

func (g *Generator) callRule(inst *ir.InstCall) string {
	callee := inst.Callee.Ident()
	meta := inst.Metadata
	g.parallelMeta(g.parallel, meta)
	g.last = nil
	return callee
}

func (g *Generator) tempRule(inst value.Value, r rule) {
	// If infix rule is stored in a temp variable
	id := inst.Ident()
	if g.isTemp(id) {
		g.ref[id] = r
	}
}

func (g *Generator) tempToIdent(ru rule) rule {
	switch r := ru.(type) {
	case *wrap:
		return g.fetchIdent(r.value, r)
	case *infix:
		r.x = g.tempToIdent(r.x)
		r.y = g.tempToIdent(r.y)
		return r
	}
	return ru
}

func (g *Generator) fetchIdent(id string, r rule) rule {
	if g.isTemp(id) {
		if _, ok := g.loads[id]; ok {
			id = g.advanceSSA(id)
			wid := &wrap{value: id}
			return wid
		} else if ref, ok := g.ref[id]; ok {
			switch r := ref.(type) {
			case *infix:
				r.x = g.tempToIdent(r.x)
				r.y = g.tempToIdent(r.y)
				return r
			}
		} else {
			panic(fmt.Sprintf("smt generation error, value for %s not found", id))
		}
	}
	return r
}