package smt

// import (
// 	"fmt"

// 	"github.com/llir/llvm/ir"
// 	"github.com/llir/llvm/ir/constant"
// 	"github.com/llir/llvm/ir/value"
// )

// func (g *Generator) constantRule(id string, c constant.Constant) string {
// 	switch val := c.(type) {
// 	case *constant.Float:
// 		id = g.variables.advanceSSA(id)
// 		if g.isASolvable(id) {
// 			g.declareVar(id, "Real")
// 		} else {
// 			return g.writeInitRule(id, "Real", val.String())
// 		}
// 	}
// 	return ""
// }

// func (g *Generator) loadsRule(inst *ir.InstLoad) {
// 	id := inst.Ident()
// 	g.variables.loads[id] = inst.Src
// }

// func (g *Generator) storeRule(inst *ir.InstStore, rules []rule) []rule {
// 	id := g.variables.formatIdent(inst.Dst.Ident())
// 	if g.variables.isTemp(inst.Src.Ident()) {
// 		srcId := inst.Src.Ident()
// 		if val, ok := g.variables.loads[srcId]; ok {
// 			ty := g.getType(val)
// 			n := g.variables.ssa[id]
// 			if g.inPhiState {
// 				g.setPhiTempState(id, n+1)
// 			} else {
// 				g.variables.storeLastState(id, n+1)
// 			}
// 			id = g.variables.advanceSSA(id)
// 			rules = append(rules, g.parseRule(id, g.variables.formatValue(val), ty, ""))
// 		} else if ref, ok := g.variables.ref[srcId]; ok {
// 			switch r := ref.(type) {
// 			case *infix:
// 				r.x = g.tempToIdent(r.x)
// 				r.y = g.tempToIdent(r.y)
// 				n := g.variables.ssa[id]
// 				if g.inPhiState {
// 					g.setPhiTempState(id, n+1)
// 				} else {
// 					g.variables.storeLastState(id, n+1)
// 				}
// 				id = g.variables.advanceSSA(id)
// 				//g.trackRounds(id, inst)
// 				wid := &wrap{value: id}
// 				if g.isASolvable(r.x.String()) {
// 					rules = append(rules, &infix{x: wid, ty: "Real", y: r, declareOnly: true}) // Still need to declare the new state
// 					rules = append(rules, &infix{x: wid, ty: "Real", y: r, op: "="})
// 				} else {
// 					rules = append(rules, &infix{x: wid, ty: "Real", y: r})
// 				}
// 			default:
// 				n := g.variables.ssa[id]
// 				if g.inPhiState {
// 					g.setPhiTempState(id, n+1)
// 				} else {
// 					g.variables.storeLastState(id, n+1)
// 				}
// 				id = g.variables.advanceSSA(id)
// 				wid := &wrap{value: id}
// 				rules = append(rules, &infix{x: wid, ty: "Real", y: r})
// 			}
// 		} else {
// 			panic(fmt.Sprintf("smt generation error, value for %s not found", id))
// 		}
// 	} else {
// 		ty := g.getType(inst.Src)
// 		n := g.variables.ssa[id]
// 		if g.inPhiState {
// 			g.setPhiTempState(id, n+1)
// 		} else {
// 			g.variables.storeLastState(id, n+1)
// 		}
// 		id = g.variables.advanceSSA(id)
// 		rules = append(rules, g.parseRule(id, inst.Src.Ident(), ty, ""))
// 	}
// 	return rules
// }

// func (g *Generator) parallelMeta(parallelGroup string, meta ir.Metadata) string {
// 	for _, v := range meta {
// 		if v.Name != parallelGroup {
// 			g.call = g.call + 1
// 		}
// 		if v.Name[0:5] != "round-" {
// 			g.parallelGrouping = v.Name
// 		}
// 	}
// 	return g.parallelGrouping
// }

// func (g *Generator) callRule(inst *ir.InstCall) string {
// 	callee := inst.Callee.Ident()
// 	meta := inst.Metadata
// 	g.parallelMeta(g.parallelGrouping, meta)
// 	return callee
// }

// func (g *Generator) tempRule(inst value.Value, r rule) {
// 	// If infix rule is stored in a temp variable
// 	id := inst.Ident()
// 	if g.variables.isTemp(id) {
// 		g.variables.ref[id] = r
// 	}
// }

// func (g *Generator) tempToIdent(ru rule) rule {
// 	switch r := ru.(type) {
// 	case *wrap:
// 		return g.fetchIdent(r.value, r)
// 	case *infix:
// 		r.x = g.tempToIdent(r.x)
// 		r.y = g.tempToIdent(r.y)
// 		return r
// 	}
// 	return ru
// }

// func (g *Generator) fetchIdent(id string, r rule) rule {
// 	if g.variables.isTemp(id) {
// 		if v, ok := g.variables.loads[id]; ok {
// 			n := g.variables.ssa[id]
// 			if g.inPhiState {
// 				g.setPhiTempState(id, n+1)
// 			} else {
// 				g.variables.storeLastState(id, n+1)
// 			}
// 			id = g.variables.advanceSSA(v.Ident())
// 			wid := &wrap{value: id}
// 			return wid
// 		} else if ref, ok := g.variables.ref[id]; ok {
// 			switch r := ref.(type) {
// 			case *infix:
// 				r.x = g.tempToIdent(r.x)
// 				r.y = g.tempToIdent(r.y)
// 				return r
// 			}
// 		} else {
// 			panic(fmt.Sprintf("smt generation error, value for %s not found", id))
// 		}
// 	}
// 	return r
// }
