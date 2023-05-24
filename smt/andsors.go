package smt

import (
	"fault/smt/rules"
	"fmt"

	"github.com/llir/llvm/ir/value"
)

func (g *Generator) storeStateChange(inst value.Value) {
	sc := &rules.StateChange{
		Ors:  []value.Value{},
		Ands: []value.Value{},
	}
	andAd, _ := g.parseChoice(inst, sc)
	id := inst.Ident()
	refname := fmt.Sprintf("%s-%s", g.currentFunction, id)
	g.variables.Loads[refname] = inst
	g.storedChoice[refname] = andAd
}

func (g *Generator) mergeStateChange(refname string, v value.Value, sc *rules.StateChange, op string) *rules.StateChange {
	if v == nil {
		return sc
	}

	if g.peek(v) != "infix" {
		sc, ret2 := g.parseChoice(v, sc)
		if op == "or" {
			sc.Ors = append(sc.Ors, ret2...)
		}
		if op == "and" {
			sc.Ands = append(sc.Ands, ret2...)
		}
	} else {
		sc2 := g.storedChoice[refname].(*rules.StateChange)
		if sc.Empty() {
			return sc2
		}

		sc.Ands = append(sc.Ands, sc2.Ands...)
		sc.Ors = append(sc.Ors, sc2.Ors...)
	}
	return sc
}
