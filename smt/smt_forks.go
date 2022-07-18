package smt

import (
	"fmt"
	"sort"
)

// Key is hash of rule that creates the fork
type fork map[string][]*choice

type choice struct {
	base   string  // What variable?
	branch string  // For conditionals, is this the true block or false block?
	values []int16 // All the versions of this variable in this branch
}

func (c *choice) addChoiceValue(n int16) {
	sort.Slice(c.values, func(i, j int) bool { return c.values[i] < c.values[j] })
	c.values = append(c.values, n)
}

func (c *choice) getEnd() int16 {
	return c.values[len(c.values)-1]
}

func (g *Generator) newFork() {
	if g.inPhiState { // a fork inside a fork (facepalm)
		g.parentFork = g.getCurrentFork()
		g.forks = append(g.forks[0:len(g.forks)-1], fork{})
	} else {
		g.parentFork = nil
		g.inPhiState = true
		g.forks = append(g.forks, fork{})
	}
}

func (g *Generator) newChoice(base string, n int16, b string) *choice {
	return &choice{
		base:   base,
		branch: b,
		values: []int16{n},
	}
}

func (g *Generator) getCurrentFork() fork {
	return g.forks[len(g.forks)-1]
}

func (g *Generator) buildForkChoice(rules []rule, b string) {
	var stateChanges []string
	fork := g.getCurrentFork()
	for _, ru := range rules {
		stateChanges = append(stateChanges, g.allStateChangesInRule(ru)...)
	}

	seenVar := make(map[string]bool)
	for _, s := range stateChanges {
		base, i := g.variables.getVarBase(s)
		n := int16(i)
		// Have we seen this variable in a previous branch of
		// this fork?
		if _, ok := fork[base]; ok {
			// Have we seen this variable before in this branch?
			if seenVar[base] {
				fork[base][len(fork[base])-1].addChoiceValue(n)
			} else {
				seenVar[base] = true
				fork[base] = append(fork[base], g.newChoice(base, n, b))
			}
		} else {
			fork[base] = []*choice{g.newChoice(base, n, b)}
		}
	}
	g.forks[len(g.forks)-1] = fork
}

func (g *Generator) allStateChangesInRule(ru rule) []string {
	var wg []string
	switch r := ru.(type) {
	case *infix:
		ch := g.allStateChangesInRule(r.x)
		wg = append(wg, ch...)
		ch = g.allStateChangesInRule(r.y)
		wg = append(wg, ch...)

	case *ite:
		for _, w := range r.t {
			ch := g.allStateChangesInRule(w)
			wg = append(wg, ch...)
		}

		for _, w := range r.f {
			ch := g.allStateChangesInRule(w)
			wg = append(wg, ch...)
		}
	case *wrapGroup:
		for _, w := range r.wraps {
			ch := g.allStateChangesInRule(w)
			wg = append(wg, ch...)
		}
	case *wrap:
		if _, ok := g.variables.ssa[r.value]; ok { // Wraps might be static values
			return []string{r.value}
		}
	}
	return wg
}

///////////////////////////////
// Logic Behind Parallel Runs
//////////////////////////////

func (g *Generator) runParallel(perm [][]string) {
	g.branchId = g.branchId + 1
	branch := fmt.Sprint("branch_", g.branchId)
	for i, calls := range perm {
		branchBlock := fmt.Sprint("option_", i)
		var opts [][]rule
		g.newFork()
		for _, c := range calls {
			v := g.functions[c]
			raw := g.parseFunction(v)
			raw = g.tagRules(raw, branch, branchBlock)
			// Pull all the variables out of the rules and
			// sort them into fork choices
			g.buildForkChoice(raw, "")
			opts = append(opts, raw)
			i += 1
		}
		//Flat the rules
		raw := g.parallelRules(opts)
		for _, v := range raw {
			g.rules = append(g.rules, g.writeRule(v))
		}
	}
	g.rules = append(g.rules, g.capParallel()...)
}

func (g *Generator) parallelRules(r [][]rule) []rule {
	var rules []rule
	for _, op := range r {
		rules = append(rules, op...) // Flatten
	}
	return rules
}

func (g *Generator) capParallel() []string {
	// Take all the end variables for the all the branches
	// and cap them with a phi value
	// writes OR nodes to end each parallel run

	fork := g.getCurrentFork()
	var rules []string
	for k, v := range fork {
		id := g.variables.advanceSSA(k)
		g.declareVar(id, "Real")

		var nums []int16
		for _, c := range v {
			nums = append(nums, c.getEnd())
		}

		ends := g.formatEnds(k, nums, id)
		rule := g.writeAssert("or", ends)
		rules = append(rules, rule)

		n := g.variables.ssa[id]
		g.variables.storeLastState(id, n)

	}
	return rules
}

func (g *Generator) capRule(k string, nums []int16, id string) []rule {
	var e []rule
	for _, v := range nums {
		id2 := fmt.Sprint(k, "_", v)
		r := &infix{
			x:  &wrap{value: id},
			y:  &wrap{value: id2},
			op: "=",
			ty: "Real",
		}
		e = append(e, r)
	}
	return e
}

func (g *Generator) capCond(b string) []rule {
	fork := g.getCurrentFork()
	var rules []rule
	for k, v := range fork {
		id := g.variables.advanceSSA(k)
		g.declareVar(id, "Real")

		for _, c := range v {
			if c.branch == b {
				rules = append(rules, g.capRule(k, []int16{c.getEnd()}, id)...)
			}
		}
	}
	return rules
}

// func (g *Generator) capCondSyncRules() ([]rule, []rule) {
// 	// For cases where variables changed in one branch are not
// 	// present in the other, add a rule
// 	var tends []rule
// 	var fends []rule
// 	fork := g.getCurrentFork()
// 	for k, c := range fork {
// 		if len(c) == 1 {
// 			start := g.variables.getLastState(k)
// 			id := g.variables.getSSA(k)
// 			switch c[0].branch {
// 			case "true":
// 				tends = append(tends, g.capRule(k, []int16{start}, id)...)
// 			case "false":
// 				fends = append(fends, g.capRule(k, []int16{start}, id)...)
// 			}
// 			n := g.variables.ssa[k]
// 			g.variables.storeLastState(k, n)
// 		}
// 	}
// 	return tends, fends
// }

func (g *Generator) tagRules(rules []rule, branch string, block string) []rule {
	var tagged []rule
	for i := 0; i < len(rules); i++ {
		tagged = append(tagged, g.tagRule(rules[i], branch, block))
	}
	return tagged
}

func (g *Generator) tagRule(ru rule, branch string, block string) rule {
	switch r := ru.(type) {
	case *infix:
		r.x = g.tagRule(r.x, branch, block)
		r.y = g.tagRule(r.y, branch, block)
		r.Tag(branch, block)
		return r
	case *ite:
		r.cond = g.tagRule(r.cond, branch, block)
		r.t = g.tagRules(r.t, branch, block)
		r.f = g.tagRules(r.f, branch, block)
		r.Tag(branch, block)
		return r
	case *wrap:
		r.Tag(branch, block)
		return r
	case *vwrap:
		r.Tag(branch, block)
		return r
	default:
		panic(fmt.Sprintf("%T is not a valid rule type", ru))
	}
}
