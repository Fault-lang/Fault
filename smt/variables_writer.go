package smt

import (
	"fault/llvm"
	"fault/smt/rules"
	"fmt"
	"strconv"
	"strings"

	irtypes "github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type VarChange struct {
	Id     string // SSA name of var
	Parent string // SSA name of proceeding var
}

func (g *Generator) AddNewVarChange(base string, id string, parent string) {
	var v *VarChange
	if id == parent {
		v = &VarChange{Id: id, Parent: ""}
	} else {
		v = &VarChange{Id: id, Parent: parent}
	}

	if len(g.Results[base]) == 0 {
		g.Results[base] = append(g.Results[base], v)
	} else {
		g.Results[base] = append(g.Results[base], v)
	}
}

func (g *Generator) VarChangePhi(base string, end string, nums []int16) {
	for _, n := range nums {
		start := fmt.Sprintf("%s_%d", base, n)
		g.AddNewVarChange(base, end, start)
	}
}

type variables struct {
	ssa   map[string]int16
	ref   map[string]rules.Rule
	loads map[string]value.Value
	phis  map[string][][]int16
	types map[string]string
}

func NewVariables() *variables {
	return &variables{
		ssa:   make(map[string]int16),
		ref:   make(map[string]rules.Rule),
		loads: make(map[string]value.Value),
		phis:  make(map[string][][]int16),
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

func (g *variables) convertIdent(f string, val string) string {
	if g.isTemp(val) {
		refname := fmt.Sprintf("%s-%s", f, val)
		if v, ok := g.loads[refname]; ok {
			id := g.formatIdent(v.Ident())
			if v, ok := g.ssa[id]; ok {
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
				n := g.variables.getStartState(xidNoPercent)
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
		// ints are probably bools
		case *irtypes.ArrayType:
			v.types[id] = "Bool"
			return "Bool"
		}
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

func (g *variables) lookupCondPart(f string, val string) rules.Rule {
	if g.isTemp(val) {
		refname := fmt.Sprintf("%s-%s", f, val)
		if v, ok := g.ref[refname]; ok {
			return v
		}
	}
	return nil
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
func (v *variables) initPhis() {
	for k := range v.phis {
		v.newPhi(k, -1)
	}
}

func (v *variables) newPhi(id string, init int16) {
	if _, ok := v.phis[id]; !ok {
		v.phis[id] = append(v.phis[id], []int16{0})
		return
	}

	if init != -1 {
		v.phis[id] = append(v.phis[id], []int16{init})
		return
	}

	init = v.getLastState(id)
	v.phis[id] = append(v.phis[id], []int16{init})
}

func (v *variables) popPhis() {
	for k := range v.phis {
		v.popPhi(k)
	}
}

func (v *variables) popPhi(id string) {
	if p, ok := v.phis[id]; ok {
		v.phis[id] = p[0 : len(p)-1]
	}
}

func (g *variables) getLastState(id string) int16 {
	if p, ok := g.phis[id]; ok {
		last := p[len(p)-1]
		return last[len(last)-1]
	}
	return 0
}

func (g *variables) getStartState(id string) int16 {
	if p, ok := g.phis[id]; ok {
		last := p[len(p)-1]
		return last[0]
	}
	return 0
}

func (v *variables) saveState() map[string]int16 {
	state := make(map[string]int16)
	for k := range v.phis {
		f := v.getStartState(k)
		state[k] = f
	}
	return state
}

func (v *variables) loadState(state map[string]int16) {
	for k, i := range state {
		v.newPhi(k, i)
	}
}

func (v *variables) appendState(state map[string]int16) {
	for k, i := range state {
		v.storeLastState(k, i)
	}
}

func (g *variables) storeLastState(id string, n int16) {
	if p, ok := g.phis[id]; ok {
		last := p[len(p)-1]
		updated := append(last, n)
		g.phis[id][len(p)-1] = updated
	} else {
		g.newPhi(id, 0) //Probably a bug but fixing it breaks a bunch of stuff haha
	}
}

////////////////////////
// Some functions specific to variable names in rules
////////////////////////

func (g *Generator) tempToIdent(ru rules.Rule) rules.Rule {
	switch r := ru.(type) {
	case *rules.Wrap:
		return g.fetchIdent(r.Value, r)
	case *rules.Infix:
		r.X = g.tempToIdent(r.X)
		r.Y = g.tempToIdent(r.Y)
		return r
	}
	return ru
}

func (g *Generator) fetchIdent(id string, r rules.Rule) rules.Rule {
	if g.variables.isTemp(id) {
		refname := fmt.Sprintf("%s-%s", g.currentFunction, id)
		if v, ok := g.variables.loads[refname]; ok {
			n := g.variables.ssa[id]
			if !g.inPhiState.Check() {
				g.variables.newPhi(id, n+1)
			} else {
				g.variables.storeLastState(id, n+1)
			}
			g.addVarToRound(id, int(n+1))
			id = g.variables.advanceSSA(v.Ident())
			wid := &rules.Wrap{Value: id}
			return wid
		} else if ref, ok := g.variables.ref[refname]; ok {
			switch r := ref.(type) {
			case *rules.Infix:
				r.X = g.tempToIdent(r.X)
				r.Y = g.tempToIdent(r.Y)
				return r
			}
		} else {
			panic(fmt.Sprintf("smt generation error, value for %s not found", id))
		}
	}
	return r
}
