package smt

import (
	"fault/util"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/llir/llvm/ir"
)

// Key is the base variable name
type Fork map[string][]*Choice

type PhiState struct {
	levels int
}

func NewPhiState() *PhiState {
	return &PhiState{
		levels: 0,
	}
}

func (p *PhiState) Check() bool {
	return p.levels > 0
}

func (p *PhiState) Level() int {
	return p.levels
}

func (p *PhiState) In() {
	p.levels = p.levels + 1
}

func (p *PhiState) Out() {
	if p.levels != 0 {
		p.levels = p.levels - 1
	}
}

func GetForkEndPoints(c []*Choice) []int16 {
	var ends []int16
	for _, v := range c {
		e := v.Values[len(v.Values)-1]
		ends = append(ends, e)
	}
	return ends
}

type Choice struct {
	Base   string  // What variable?
	Branch string  // For conditionals, is this the true block or false block?
	Values []int16 // All the versions of this variable in this branch
}

func (c *Choice) addChoiceValue(n int16) *Choice {
	c.Values = append(c.Values, n)
	sort.Slice(c.Values, func(i, j int) bool { return c.Values[i] < c.Values[j] })
	return c
}

func (c *Choice) getEnd() int16 {
	return c.Values[len(c.Values)-1]
}

func (g *Generator) newFork() {
	if len(g.forks) == 0 {
		g.forks = append(g.forks, Fork{})
		return
	}

	if g.inPhiState.Check() {
		g.forks = append(g.forks[0:len(g.forks)-1], Fork{})
	} else {
		g.forks = append(g.forks, Fork{})
	}
}

func (g *Generator) newChoice(base string, n int16, b string) *Choice {
	return &Choice{
		Base:   base,
		Branch: b,
		Values: []int16{n},
	}
}

func (g *Generator) getCurrentFork() Fork {
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
			if seenVar[base] && // Have we seen this variable before?
				fork[base][len(fork[base])-1].Branch == b { // in this branch?
				fork[base][len(fork[base])-1] = fork[base][len(fork[base])-1].addChoiceValue(n)
			} else {
				seenVar[base] = true
				fork[base] = append(fork[base], g.newChoice(base, n, b))
			}
		} else {
			seenVar[base] = true
			fork[base] = []*Choice{g.newChoice(base, n, b)}
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
		if !g.variables.isNumeric(r.value) && !g.variables.isBolean(r.value) { // Wraps might be static values
			return []string{r.value}
		}
	case *ands:
		for _, w := range r.x {
			ch := g.allStateChangesInRule(w)
			wg = append(wg, ch...)
		}
	case *choices:
		for _, w := range r.x {
			ch := g.allStateChangesInRule(w)
			wg = append(wg, ch...)
		}
	}
	return wg
}

///////////////////////////////
// Logic Behind Parallel Runs
//////////////////////////////

func (g *Generator) parallelPermutations(p []string) (permuts [][]string) {
	var rc func([]string, int)
	rc = func(a []string, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]string{}, a...))
		} else {
			for i := k; i < len(p); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(p, 0)

	return permuts
}

func (g *Generator) runParallel(perm [][]string) []rule {
	var rules []rule
	g.branchId = g.branchId + 1
	branch := fmt.Sprint("branch_", g.branchId)
	g.newFork()
	for i, calls := range perm {
		branchBlock := fmt.Sprint("option_", i)
		var opts [][]rule
		varState := g.variables.saveState()
		for _, c := range calls {
			g.parallelRunStart = true
			g.inPhiState.Out() //Don't behave like we're in Phi inside the function
			v := g.functions[c]
			raw := g.parseFunction(v)
			g.inPhiState.In()
			raw = g.tagRules(raw, branch, branchBlock)
			opts = append(opts, raw)
		}
		//Flat the rules
		raw := g.parallelRules(opts)
		// Pull all the variables out of the rules and
		// sort them into fork choices
		g.buildForkChoice(raw, "")
		g.variables.loadState(varState)
		rules = append(rules, raw...)
	}

	rules = append(rules, g.capParallel()...)
	return rules
}

func (g *Generator) parallelRules(r [][]rule) []rule {
	var rules []rule
	for _, op := range r {
		rules = append(rules, op...) // Flatten
	}
	return rules
}

func (g *Generator) isSameParallelGroup(meta ir.Metadata) bool {
	for _, v := range meta {

		if v.Name == g.parallelGrouping {
			return true
		}

		if g.parallelGrouping == "" {
			return true
		}
	}

	return false
}

func (g *Generator) singleParallelStep(callee string) bool {
	if len(g.localCallstack) == 0 {
		return false
	}

	if callee == g.localCallstack[len(g.localCallstack)-1] {
		return true
	}

	return false
}

func (g *Generator) updateParallelGroup(meta ir.Metadata) {
	for _, v := range meta {
		if v.Name[0:5] != "round-" {
			g.parallelGrouping = v.Name
		}
	}
}

func (g *Generator) capParallel() []rule {
	// Take all the end variables for the all the branches
	// and cap them with a phi value
	// writes OR nodes to end each parallel run

	fork := g.getCurrentFork()
	var rules []rule
	for k, v := range fork {
		id := g.variables.advanceSSA(k)
		g.addVarToRound(k, int(g.variables.ssa[k]))

		var nums []int16
		for _, c := range v {
			nums = append(nums, c.getEnd())
		}

		rule := &phi{
			baseVar:  k,
			endState: id,
			nums:     nums,
		}
		g.VarChangePhi(k, id, nums)
		rules = append(rules, rule)

		base, i := g.variables.getVarBase(id)
		n := int16(i)
		if g.inPhiState.Level() == 1 {
			g.variables.newPhi(base, n)
		} else {
			g.variables.storeLastState(base, n)
		}

	}
	return rules
}

func (g *Generator) capRule(k string, nums []int16, id string) []rule {
	var e []rule
	for _, v := range nums {
		id2 := fmt.Sprint(k, "_", v)
		g.AddNewVarChange(k, id, id2)
		ty := g.variables.lookupType(k, nil)
		if ty == "Bool" {
			r := &infix{
				x:  &wrap{value: id},
				y:  &wrap{value: id2},
				op: "=",
				ty: "Bool",
			}
			e = append(e, r)
		} else {
			r := &infix{
				x:  &wrap{value: id},
				y:  &wrap{value: id2},
				op: "=",
				ty: "Real",
			}
			e = append(e, r)
		}
	}
	return e
}

func (g *Generator) capCond(b string, phis map[string]int16) ([]rule, map[string]int16) {
	fork := g.getCurrentFork()
	var rules []rule
	for k, v := range fork {
		// Because we're looking at all the variables in
		// the true branch THEN all the variables in the
		// false branch, we only increment the variable
		// when we produce the phi value for the first time
		var id string
		if phi, ok := phis[k]; !ok {
			id = g.variables.advanceSSA(k)
			g.declareVar(id, g.variables.lookupType(k, nil))
			_, i := g.variables.getVarBase(id)
			g.addVarToRound(k, i)
			phis[k] = int16(i)
		} else {
			id = fmt.Sprintf("%s_%d", k, phi)
		}

		for _, c := range v {
			if c.Branch == b {
				rules = append(rules, g.capRule(k, []int16{c.getEnd()}, id)...)
			}
		}
	}
	return rules, phis
}

func (g *Generator) capCondSyncRules(branches []string) map[string][]rule {
	// For cases where variables changed in one branch are not
	// present in the other, add a rule
	ends := make(map[string][]rule)
	for _, b := range branches {
		var e []rule
		fork := g.getCurrentFork()
		for k, c := range fork {
			if len(c) == 1 && c[0].Branch == b {
				start := g.variables.getStartState(k)
				id := g.variables.getSSA(k)
				e = append(e, g.capRule(k, []int16{start}, id)...)
				n := g.variables.ssa[k]
				if g.inPhiState.Level() == 1 {
					g.variables.newPhi(k, n)
				} else {
					g.variables.storeLastState(k, n)
				}
			}
		}
		for _, notB := range util.Intersection(branches, []string{b}, true) {
			if _, ok := ends[notB]; !ok {
				ends[notB] = e
			} else {
				ends[notB] = append(ends[notB], e...)
			}
		}
	}
	return ends
}

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
	case *phi:
		r.Tag(branch, block)
		return r
	case *ands:
		r.x = g.tagRules(r.x, branch, block)
		r.Tag(branch, block)
		return r
	case *choices:
		var tagged []*ands
		for _, v := range r.x {
			r2 := g.tagRule(v, branch, block)
			tagged = append(tagged, r2.(*ands))
		}
		r.x = tagged
		r.Tag(branch, block)
		return r
	default:
		panic(fmt.Sprintf("%T is not a valid rule type", ru))
	}
}

func (g *Generator) formatEnds(k string, nums []int16, id string) string {
	var e []string
	for _, v := range nums {
		v := fmt.Sprint(k, "_", strconv.Itoa(int(v)))
		r := g.writeAssertlessRule("=", id, v)
		e = append(e, r)
	}
	return strings.Join(e, " ")
}
