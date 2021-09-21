package smt

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	irtypes "github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type rule struct {
	x  interface{}
	y  interface{}
	ty string
	op string
}

type Generator struct {
	callgraph       string
	smt             []string
	inits           []string
	constants       []string
	rules           []string
	asserts         []string
	ssa             map[string]int16
	loads           map[string]value.Value
	ref             map[string]*rule
	call            int
	parallel        string
	parallelEnds    map[string][]int
	callstack       map[int][]string
	functions       map[string]*ir.Func
	currentFunction string
}

func NewGenerator() *Generator {
	return &Generator{
		ssa:             make(map[string]int16),
		loads:           make(map[string]value.Value),
		ref:             make(map[string]*rule),
		parallelEnds:    make(map[string][]int),
		callstack:       make(map[int][]string),
		functions:       make(map[string]*ir.Func),
		currentFunction: "@__run",
	}
}

func (g *Generator) SMT() string {
	var out bytes.Buffer

	out.WriteString(strings.Join(g.inits, "\n"))
	out.WriteString(strings.Join(g.constants, "\n"))
	out.WriteString(strings.Join(g.rules, "\n"))
	out.WriteString(strings.Join(g.asserts, "\n"))
	out.WriteString("(check-sat)\n(get-model)")

	//fmt.Println(out.String())
	return out.String()
}

func (g *Generator) Run(llopt string) {
	m, err := asm.ParseString("", llopt) //"" because ParseString has an
	if err != nil {                      // optional path parameter
		panic(err)
	}
	g.newCallgraph(m)

}

func (g *Generator) newCallgraph(m *ir.Module) {
	for _, gl := range m.Globals {
		id := g.formatIdent(gl.GlobalIdent.Ident())
		g.constants = append(g.constants, g.constantRule(id, gl.Init))
	}
	for _, f := range m.Funcs {
		// Get function name.
		g.currentFunction = f.Ident()
		if g.currentFunction != "@__run" {
			g.functions[g.currentFunction] = f
			continue
		}
		// Rules that are in the run block.
		run := g.parseFunction(f, nil)
		g.rules = append(g.rules, g.generateRules(run)...)
	}
	for i := 1; i < len(g.callstack); i++ {
		var raw []*rule
		if len(g.callstack[i]) > 1 {
			//Generate parallel runs
			perm := g.parallelPermutations(g.callstack[i])
			startVars := make(map[string]string)
			startVars = g.gatherStarts(g.callstack[i], startVars)
			g.runParallel(perm, startVars)

		} else {
			fname := g.callstack[i][0]
			v := g.functions[fname]
			raw = g.parseFunction(v, nil)

			for _, v := range raw {
				g.rules = append(g.rules, g.writeRule(v))
			}
		}
	}
}

func (g *Generator) parseFunction(f *ir.Func, startVars map[string]string) []*rule {
	var rules []*rule
	g.currentFunction = f.Ident()
	for _, block := range f.Blocks {
		// For each non-branching instruction of the basic block.
		r, sv := g.parseInstruct(block, startVars)
		rules = append(rules, r...)
		startVars = sv
	}
	return rules
}

func (g *Generator) parseInstruct(block *ir.Block, startVars map[string]string) ([]*rule, map[string]string) {
	var rules []*rule
	for _, inst := range block.Insts {
		// Type switch on instruction to find call instructions.
		switch inst := inst.(type) {
		case *ir.InstAlloca:
			//Do nothing
		case *ir.InstLoad:
			id := inst.Ident()
			g.loads[id] = inst.Src
		case *ir.InstStore:
			id := g.formatIdent(inst.Dst.Ident())
			if g.isTemp(inst.Src.Ident()) {
				srcId := inst.Src.Ident()
				if val, ok := g.loads[srcId]; ok {
					ty := g.getType(val)
					id = g.advanceSSA(id)
					rules = append(rules, g.parseRule(id, g.formatValue(val), ty, ""))
				} else if ref, ok := g.ref[srcId]; ok {
					id = g.advanceSSA(id)
					rules = append(rules, &rule{x: id, ty: "Real", y: ref})
				} else {
					panic(fmt.Sprintf("smt generation error, value for %s not found", id))
				}
			} else {
				ty := g.getType(inst.Src)
				id = g.advanceSSA(id)
				rules = append(rules, g.parseRule(id, inst.Src.Ident(), ty, ""))
			}
		case *ir.InstFAdd:
			startVars = g.parseInfix(inst.Ident(),
				inst.X.Ident(), inst.Y.Ident(), "+", startVars)
		case *ir.InstFSub:
			startVars = g.parseInfix(inst.Ident(),
				inst.X.Ident(), inst.Y.Ident(), "-", startVars)
		case *ir.InstFMul:
			startVars = g.parseInfix(inst.Ident(),
				inst.X.Ident(), inst.Y.Ident(), "*", startVars)
		case *ir.InstFDiv:
			startVars = g.parseInfix(inst.Ident(),
				inst.X.Ident(), inst.Y.Ident(), "/", startVars)
		case *ir.InstFRem:
		case *ir.InstFCmp:
		case *ir.InstCall:
			callee := inst.Callee.Ident()
			meta := inst.Metadata
			if len(meta) == 0 || g.parallel != meta[0].Name {
				g.call = g.call + 1
			}
			for _, v := range meta {
				g.parallel = v.Name
			}

			g.callstack[g.call] = append(g.callstack[g.call], callee)
		}
	}
	return rules, startVars
}

func (g *Generator) parseInfix(id string, x string, y string, op string, startVars map[string]string) map[string]string {
	if g.isTemp(x) {
		if v, ok := g.loads[x]; ok {
			xid := v.Ident()
			xidNoPercent := g.formatIdent(xid)
			if id, ok := startVars[xidNoPercent]; ok {
				x = g.formatIdent(id)
				delete(startVars, xidNoPercent)
			} else {
				x = g.convertIdent(xid)
			}
		}
	}

	if g.isTemp(y) {
		if v, ok := g.loads[y]; ok {
			yid := v.Ident()
			yidNoPercent := g.formatIdent(yid)
			if id, ok := startVars[yidNoPercent]; ok {
				y = g.formatIdent(id)
				delete(startVars, yidNoPercent)
			} else {
				y = g.convertIdent(yid)
			}
		}
	}
	g.ref[id] = g.parseRule(x, y, "", op)
	return startVars
}

func (g *Generator) constantRule(id string, c constant.Constant) string {
	switch val := c.(type) {
	case *constant.Float:
		id = g.advanceSSA(id)
		return g.writeInitRule(id, "Real", val.String())
	}
	return ""
}

func (g *Generator) getType(val value.Value) string {
	switch val.Type().(type) {
	case *irtypes.FloatType:
		return "Real"
	}
	return ""
}

func (g *Generator) parseRule(id string, val string, ty string, op string) *rule {
	return &rule{x: id, ty: ty, y: val, op: op}
}

// func (g *Generator) variableRule(id string, val value.Value) {
// 	switch val.Type().(type) {
// 	case *irtypes.FloatType:
// 		g.writeInitRule(id, "Real", g.formatValue(val))
// 	}
// }

func (g *Generator) convertIdent(val string) string {
	if g.isTemp(val) {
		if v, ok := g.loads[val]; ok {
			id := g.formatIdent(v.Ident())
			if v, ok := g.ssa[id]; ok {
				id = g.formatIdent(id)
				return fmt.Sprint(id, "_", v)
			} else {
				panic(fmt.Sprintf("variable %s not initialized", id))
			}

		} else {
			panic(fmt.Sprintf("variable %s not initialized", val))
		}
	} else {
		id := val
		if string(id[0]) == "%" {
			id = g.formatIdent(id)
			return fmt.Sprint(id, "_", g.ssa[id])
		}
		return id //Is a value, not in identifier
	}
}

func (g *Generator) isTemp(id string) bool {
	if string(id[0]) == "%" && g.isNumeric(string(id[1])) {
		return true
	}
	return false
}

func (g *Generator) isNumeric(char string) bool {
	if _, err := strconv.Atoi(char); err != nil {
		return false
	}
	return true
}

func (g *Generator) formatIdent(id string) string {
	//Removes LLVM IR specific leading characters
	if string(id[0]) == "@" {
		return id[1:]
	} else if string(id[0]) == "%" {
		return id[1:]
	}
	return id
}

func (g *Generator) formatValue(val value.Value) string {
	v := strings.Split(val.String(), " ")
	return v[1]
}

func (g *Generator) advanceSSA(id string) string {
	if i, ok := g.ssa[id]; ok {
		g.ssa[id] = i + 1
		return fmt.Sprint(id, "_", g.ssa[id])
	} else {
		g.ssa[id] = 0
		return fmt.Sprint(id, "_0")
	}
}

func (g *Generator) getVarBase(id string) (string, int) {
	v := strings.Split(id, "_")
	num, err := strconv.Atoi(v[len(v)-1])
	if err != nil {
		panic(fmt.Sprintf("improperly formatted variable SSA name %s", id))
	}
	return strings.Join(v[0:len(v)-1], "_"), num
}

func (g *Generator) runParallel(perm [][]string, vars map[string]string) {
	var waitGroup sync.WaitGroup
	r := make(chan [][]*rule, len(perm))
	waitGroup.Add(len(perm))

	go func() {
		waitGroup.Wait()
		close(r)
	}()
	for _, p := range perm {
		go func(calls []string) {
			defer waitGroup.Done()
			var opts [][]*rule
			startVars := make(map[string]string)
			for k, v := range vars {
				startVars[k] = v
			}
			for _, c := range calls {
				v := g.functions[c]
				raw := g.parseFunction(v, startVars)
				opts = append(opts, raw)
			}
			r <- opts
		}(p)
	}
	for opts := range r {
		raw := g.parallelRules(opts)
		for _, v := range raw {
			g.rules = append(g.rules, g.writeRule(v))
		}
	}
	g.rules = append(g.rules, g.capParallel()...)
}

func (g *Generator) parallelRules(r [][]*rule) []*rule {
	var rules []*rule
	s := g.paraStateChanges(r)
	for k, v := range s {
		if len(v) > 1 {
			g.parallelEnds[k] = append(g.parallelEnds[k], g.getEnds(v))
		}
	}
	for _, op := range r {
		rules = append(rules, op...) // Flatten
	}
	return rules
}

func (g *Generator) capParallel() []string {
	// writes OR nodes to end each parallel run
	var rules []string
	for k, v := range g.parallelEnds {
		id := g.advanceSSA(k)
		g.declareVar(id, "Real")
		ends := g.formatEnds(k, v, id)
		//(assert (or (= bathtub_drawn_water_level_10 bathtub_drawn_water_level_7) (= bathtub_drawn_water_level_10  bathtub_drawn_water_level_9)))
		rule := g.writeAssert("or", ends)
		rules = append(rules, rule)
	}
	g.parallelEnds = map[string][]int{}
	return rules
}

func (g *Generator) setStartVar(id string, startVars map[string]string) map[string]string {
	if _, ok := startVars[id]; !ok {
		startVars[id] = fmt.Sprint(id, "_", g.ssa[id])
	}
	return startVars
}

func (g *Generator) gatherStarts(p []string, startVars map[string]string) map[string]string {
	for _, fname := range p {
		f := g.functions[fname]
		for _, block := range f.Blocks {
			for _, inst := range block.Insts {
				switch inst := inst.(type) {
				case *ir.InstLoad:
					id := g.formatIdent(inst.Src.Ident())
					startVars = g.setStartVar(id, startVars)
				}
			}
		}
	}
	return startVars
}

func (g *Generator) getEnds(n map[int][]int) int {
	end := 0
	for _, v := range n {
		if v[len(v)-1] > end {
			end = v[len(v)-1]
		}
	}
	return end
}

func (g *Generator) formatEnds(k string, nums []int, id string) string {
	var e []string
	for _, v := range nums {
		v := fmt.Sprint(k, "_", strconv.Itoa(v))
		r := g.writeInfix(id, v, "=")
		e = append(e, r)
	}
	return strings.Join(e, " ")
}

func (g *Generator) paraStateChanges(r [][]*rule) map[string]map[int][]int {
	// When running functions in parallel,
	//	checks for conflicting state change
	state := make(map[string]map[int][]int)
	conflicts := make(map[string]map[int][]int)
	for i, v := range r {
		for _, r := range v {
			if r.ty != "" { //Variable assignment
				id, num := g.getVarBase(r.x.(string))
				if _, ok := state[id]; ok {
					state[id][i] = append(state[id][i], num)
				} else {
					state[id] = make(map[int][]int)
					state[id][i] = append(state[id][i], num)
				}
			}
		}

	}
	for k, s := range state {
		if len(s) > 1 {
			conflicts[k] = s
		}
	}
	return conflicts
}

func (g *Generator) writeInfix(x string, y string, op string) string {
	return fmt.Sprintf("(%s %s %s)", op, x, y)
}

func (g *Generator) declareVar(id string, t string) {
	def := fmt.Sprintf("(declare-fun %s () %s)", id, t)
	g.inits = append(g.inits, def)
}

func (g *Generator) writeAssert(op string, stmt string) string {
	return fmt.Sprintf("(assert (%s %s))", op, stmt)
}

func (g *Generator) writeRule(r *rule) string {
	x := g.unpackRule(r.x)
	y := g.unpackRule(r.y)
	if r.op != "" {
		return g.writeInfix(x, y, r.op)
	}
	return g.writeInitRule(x, r.ty, y)
}

func (g *Generator) unpackRule(x interface{}) string {
	switch r := x.(type) {
	case string:
		return r
	default:
		return g.writeRule(r.(*rule))
	}
}

func (g *Generator) writeInitRule(id string, t string, val string) string {
	// Initialize: x = Int("x")
	g.declareVar(id, t)
	// Set rule: s.add(x == 2)
	return fmt.Sprintf("(assert (= %s %s))", id, val)
}

func (g *Generator) generateRules(raw []*rule) []string {
	var rules []string
	for _, v := range raw {
		rules = append(rules, g.writeRule(v))
	}
	return rules
}

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
