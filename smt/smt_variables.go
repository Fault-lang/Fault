package smt

import (
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
}

func NewVariables() *variables {
	return &variables{
		ssa:   make(map[string]int16),
		ref:   make(map[string]rule),
		loads: make(map[string]value.Value),
		phis:  make(map[string]int16),
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
	if _, err := strconv.Atoi(char); err != nil {
		return false
	}
	return true
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

func (g *Generator) getType(val value.Value) string {
	switch val.Type().(type) {
	case *irtypes.FloatType:
		return "Real"
	}
	return ""
}

func (g *variables) convertIdent(val string) string {
	if g.isTemp(val) {
		if v, ok := g.loads[val]; ok {
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
		if v, ok := g.variables.loads[x]; ok {
			xid := v.Ident()
			xidNoPercent := g.variables.formatIdent(xid)
			x = g.variables.getSSA(xidNoPercent)
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

func (g *variables) storeLastState(id string, n int16) {
	if _, ok := g.phis[id]; ok {
		g.phis[id] = n
	} else {
		g.phis[id] = 0
	}
}
