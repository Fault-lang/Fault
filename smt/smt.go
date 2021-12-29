package smt

import (
	"bytes"
	"fault/ast"
	"fmt"
	"strconv"
	"strings"

	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir"
	irtypes "github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type rule interface {
	ruleNode()
	String() string
}

type assrt struct {
	rule
	variable    *wrap
	conjunction string
	assertion   rule
	tag         *branch
}

func (a *assrt) ruleNode() {}
func (a *assrt) String() string {
	return a.variable.String() + a.conjunction + a.assertion.String()
}
func (a *assrt) Tag(k1 string, k2 string) {
	a.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type infix struct {
	rule
	x   rule
	y   rule
	ty  string
	op  string
	tag *branch
}

func (i *infix) ruleNode() {}
func (i *infix) String() string {
	return fmt.Sprintf("%s %s %s", i.x.String(), i.op, i.y.String())
}
func (i *infix) Tag(k1 string, k2 string) {
	i.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type ite struct {
	rule
	cond  rule
	t     []rule
	tvars map[string]string
	f     []rule
	fvars map[string]string
	tag   *branch
}

func (it *ite) ruleNode() {}
func (it *ite) String() string {
	return fmt.Sprintf("if %s then %s else %s", it.cond.String(), it.t, it.f)
}
func (ite *ite) Tag(k1 string, k2 string) {
	ite.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type invariant struct {
	rule
	left        rule
	conjunction string
	right       rule
	tag         *branch
}

func (i *invariant) ruleNode() {}
func (i *invariant) String() string {
	return fmt.Sprint(i.left.String(), i.conjunction, i.right.String())
}
func (i *invariant) Tag(k1 string, k2 string) {
	i.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type wrap struct { //wrapper for constant values to be used in infix as rules
	rule
	value    string
	state    string //invariant only for one state
	all      bool   // invariant for all states
	constant bool   // this is a constant
	tag      *branch
}

func (w *wrap) ruleNode() {}
func (w *wrap) String() string {
	return w.value
}
func (w *wrap) Tag(k1 string, k2 string) {
	w.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type wrapGroup struct {
	rule
	wraps []*wrap
	tag   *branch
}

func (wg *wrapGroup) ruleNode() {}
func (wg *wrapGroup) String() string {
	var out bytes.Buffer
	for _, v := range wg.wraps {
		out.WriteString(v.value)
	}
	return out.String()
}
func (wg *wrapGroup) Tag(k1 string, k2 string) {
	wg.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type vwrap struct {
	rule
	value value.Value
	tag   *branch
}

func (vw *vwrap) ruleNode() {}
func (vw *vwrap) String() string {
	return vw.value.String()
}
func (vw *vwrap) Tag(k1 string, k2 string) {
	vw.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type branch struct {
	branch string
	block  string
}

func (b *branch) String() string {
	return b.branch + "." + b.block
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
	ref             map[string]rule
	call            int
	parallel        string
	parallelEnds    map[string][]int
	callstack       map[int][]string
	functions       map[string]*ir.Func
	blocks          map[string][]rule
	skipBlocks      map[string]int
	currentFunction string
	currentBlock    string
	last            rule
	branchId        int
	Branches        map[string][]string            // [varid] = branch_id
	BranchTrail     map[string]map[string][]string //[branch_id] = []string{varid}
	rawAsserts      []ast.Statement
	rounds          map[string]map[string][]int16
}

func NewGenerator() *Generator {
	return &Generator{
		ssa:             make(map[string]int16),
		loads:           make(map[string]value.Value),
		ref:             make(map[string]rule),
		parallelEnds:    make(map[string][]int),
		callstack:       make(map[int][]string),
		functions:       make(map[string]*ir.Func),
		blocks:          make(map[string][]rule),
		skipBlocks:      make(map[string]int),
		currentFunction: "@__run",
		Branches:        make(map[string][]string),
		BranchTrail:     make(map[string]map[string][]string),
		rounds:          make(map[string]map[string][]int16),
	}
}

func (g *Generator) SMT() string {
	var out bytes.Buffer

	out.WriteString(strings.Join(g.inits, "\n"))
	out.WriteString(strings.Join(g.constants, "\n"))
	out.WriteString(strings.Join(g.rules, "\n"))
	out.WriteString(strings.Join(g.asserts, "\n"))

	return out.String()
}

func (g *Generator) Run(llopt string) {
	m, err := asm.ParseString("", llopt) //"" because ParseString has an
	if err != nil {                      // optional path parameter
		panic(err)
	}
	g.newCallgraph(m)

}

func (g *Generator) getType(val value.Value) string {
	switch val.Type().(type) {
	case *irtypes.FloatType:
		return "Real"
	}
	return ""
}

func (g *Generator) newCallgraph(m *ir.Module) {
	g.constants = g.newConstants(m.Globals)
	g.rules = g.storeFuncs(m.Funcs)

	// Unroll the run block
	for i := 1; i <= len(g.callstack); i++ {
		var raw []rule
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

	for _, v := range g.rawAsserts {
		a1, a2, conj := g.parseAssert(v)
		if conj == "" {
			for _, assrt := range a1 {
				ir := g.generateAssertRules(assrt)
				or := fmt.Sprintf("(or %s)", strings.Join(ir, " "))
				g.asserts = append(g.asserts, fmt.Sprintf("(assert %s)", or))
			}
		} else {
			g.asserts = append(g.asserts, g.generateCompound(a1, a2, conj)...)
		}
	}

}

func (g *Generator) newConstants(globals []*ir.Global) []string {
	// Constants cannot be changed and therefore don't increment
	// in SSA. So instead of return a *rule we can skip directly
	// to a set of strings
	r := []string{}
	for _, gl := range globals {
		id := g.formatIdent(gl.GlobalIdent.Ident())
		r = append(r, g.constantRule(id, gl.Init))
	}
	return r
}

func (g *Generator) storeFuncs(funcs []*ir.Func) []string {
	//Iterate through all the function blocks and store them by
	// function call name.
	r := []string{}
	for _, f := range funcs {
		// Get function name.
		fname := f.Ident()

		if fname != "@__run" {
			g.functions[f.Ident()] = f
			continue
		}
		// code that is in the run block we can generate
		// rules right now.
		run := g.parseFunction(f, nil)
		r = append(r, g.generateRules(run)...)
	}
	return r
}

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

func (g *Generator) getSSA(id string) string {
	if _, ok := g.ssa[id]; ok {
		return fmt.Sprint(id, "_", g.ssa[id])
	} else {
		g.ssa[id] = 0
		return fmt.Sprint(id, "_0")
	}
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
