package smt

import (
	"fault/llvm"
	"fmt"
	"strconv"
	"strings"

	irtypes "github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type variables struct {
	ssa   map[string]int16
	ref   map[string]rule
	loads map[string]value.Value
	phis  map[string]int16
	types map[string]string
}

func NewVariables() *variables {
	return &variables{
		ssa:   make(map[string]int16),
		ref:   make(map[string]rule),
		loads: make(map[string]value.Value),
		phis:  make(map[string]int16),
		types: make(map[string]string),
	}
}

func (g *variables) isTemp(id string) bool {
	if string(id[0]) == "%" && g.isNumeric(string(id[1])) {
		return true
	}
	return false
}

func (g *variables) isGlobal(id string) bool {
	return string(id[0]) == "@"
}

func (g *variables) isNumeric(char string) bool {
	if _, err := strconv.ParseFloat(char, 64); err == nil {
		return true
	}
	if _, err := strconv.Atoi(char); err == nil {
		return true
	}
	return false
}

func (g *variables) isBolean(id string) bool {
	if id == "true" || id == "false" {
		return true
	}
	return false
}

func (g *Generator) isASolvable(id string) bool {
	id, _ = g.variables.getVarBase(id)
	for _, v := range g.Unknowns {
		if v == id {
			return true
		}
	}
	for k := range g.Uncertains {
		if k == id {
			return true
		}
	}
	return false
}

// func (g *Generator) getType(val value.Value) string {
// 	switch val.Type().(type) {
// 	case *irtypes.FloatType:
// 		return "Real"
// 	}
// 	return ""
// }

func (g *variables) convertIdent(f string, val string) string {
	if g.isTemp(val) {
		refname := fmt.Sprintf("%s-%s", f, val)
		if v, ok := g.loads[refname]; ok {
			id := g.formatIdent(v.Ident())
			if v, ok := g.ssa[id]; ok {
				//id = g.formatIdent(id)
				return fmt.Sprint(id, "_", v)
			} else {
				panic(fmt.Sprintf("variable %s not initialized", id))
			}

		} else {
			panic(fmt.Sprintf("variable %s not initialized", val))
		}
	} else {
		id := val
		if string(id[0]) == "%" || g.isGlobal(id) {
			id = g.formatIdent(id)
			return fmt.Sprint(id, "_", g.ssa[id])
		}
		return id //Is a value, not an identifier
	}
}

func (g *variables) formatIdent(id string) string {
	//Removes LLVM IR specific leading characters
	if string(id[0]) == "@" {
		return id[1:]
	} else if string(id[0]) == "%" {
		return id[1:]
	}
	return id
}

func (g *Generator) convertInfixVar(x string) string {
	if g.variables.isTemp(x) {
		refname := fmt.Sprintf("%s-%s", g.currentFunction, x)
		if v, ok := g.variables.loads[refname]; ok {
			xid := v.Ident()
			xidNoPercent := g.variables.formatIdent(xid)
			if g.parallelRunStart {
				n := g.variables.getLastState(xidNoPercent)
				x = fmt.Sprintf("%s_%d", xidNoPercent, n)
				g.parallelRunStart = false
			} else {
				x = g.variables.getSSA(xidNoPercent)
			}
		}
	}
	return x
}

func (g *variables) getVarBase(id string) (string, int) {
	v := strings.Split(id, "_")
	num, err := strconv.Atoi(v[len(v)-1])
	if err != nil {
		panic(fmt.Sprintf("improperly formatted variable SSA name %s", id))
	}
	return strings.Join(v[0:len(v)-1], "_"), num
}

func (v *variables) lookupType(id string, value value.Value) string {
	if cache, ok := v.types[id]; ok { //If we've seen this one before
		return cache
	}

	val := v.loads[id]
	if val == nil { // A backup method
		switch value.Type().(type) {
		case *irtypes.FloatType:
			v.types[id] = "Real"
			return "Real"
		case *irtypes.IntType: // LLVM doesn't have a bool type
			v.types[id] = "Bool" // Just int type with a bitsize 1
			return "Bool"        // since all Fault numbers are floats,
		} // ints are probably bools
	}

	if val.Type().Equal(llvm.DoubleP) {
		v.types[id] = "Real"
		return "Real"
	}
	if val.Type().Equal(llvm.I1P) {
		v.types[id] = "Bool"
		return "Bool"
	}

	panic(fmt.Sprintf("smt generation error, value for %s not found", id))
}

func (g *variables) lookupCondPart(f string, val string) rule {
	if g.isTemp(val) {
		refname := fmt.Sprintf("%s-%s", f, val)
		if v, ok := g.ref[refname]; ok {
			return v
		} else {
			panic(fmt.Sprintf("variable %s not initialized", val))
		}
	} else {
		panic(fmt.Sprintf("variable %s not valid construction", val))
	}
}

func (g *variables) formatValue(val value.Value) string {
	v := strings.Split(val.String(), " ")
	return v[1]
}

func (g *variables) getSSA(id string) string {
	if _, ok := g.ssa[id]; ok {
		return fmt.Sprint(id, "_", g.ssa[id])
	} else {
		g.ssa[id] = 0
		return fmt.Sprint(id, "_0")
	}
}

func (g *variables) advanceSSA(id string) string {
	if i, ok := g.ssa[id]; ok {
		g.ssa[id] = i + 1
		return fmt.Sprint(id, "_", g.ssa[id])
	} else {
		g.ssa[id] = 0
		return fmt.Sprint(id, "_0")
	}
}

// When we have conditionals back to back (but not if elseif else)
// we need to make sure to track the phi
func (g *variables) getLastState(id string) int16 {
	if p, ok := g.phis[id]; ok {
		return p
	}
	return 0
}

func (v *variables) saveState() map[string]int16 {
	state := make(map[string]int16)
	for k, v := range v.phis {
		state[k] = v
	}
	return state
}

func (v *variables) loadState(state map[string]int16) {
	v.phis = state
}

func (g *variables) storeLastState(id string, n int16) {
	if _, ok := g.phis[id]; ok {
		g.phis[id] = n
	} else {
		g.phis[id] = 0
	}
}

////////////////////////
// Some functions specific to variable names in rules
////////////////////////

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
	if g.variables.isTemp(id) {
		refname := fmt.Sprintf("%s-%s", g.currentFunction, id)
		if v, ok := g.variables.loads[refname]; ok {
			n := g.variables.ssa[id]
			if !g.inPhiState.Check() {
				g.variables.storeLastState(id, n+1)
			}
			id = g.variables.advanceSSA(v.Ident())
			wid := &wrap{value: id}
			return wid
		} else if ref, ok := g.variables.ref[refname]; ok {
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
