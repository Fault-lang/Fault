package unroll

import (
	"fault/ast"
	"fault/generator/rules"
	"fault/llvm"
	"fault/util"
	"fmt"
	"runtime"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/value"
)

// Step 1 in the Generation Process: Unroll
// the LL IR into LLUnits that reflect the
// various branches of the state graph

type Env struct {
	RawInputs        *llvm.RawInputs
	VarLoads         map[string]value.Value
	VarTypes         map[string]string
	CurrentFunction  string
	CurrentRound     int
	returnVoid       *PhiState
	ParallelGrouping string
	WhensThens       map[string]map[string][]string // map[variable_name][assert_id][]string{other_variables in the assert...}
}

func NewEnv(ri *llvm.RawInputs) *Env {
	return &Env{
		RawInputs:    ri,
		VarLoads:     make(map[string]value.Value),
		VarTypes:     make(map[string]string),
		CurrentRound: 0,
		returnVoid:   NewPhiState(),
		WhensThens:   make(map[string]map[string][]string),
	}
}

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

type LLUnit interface {
	Unroll()
	AddRules([]rules.Rule)
	AddBlock(LLUnit)
	ExecuteCallstack() []rules.Rule
	GenerateCallstack([]string) []rules.Rule
	String() string
}

type LLFunc struct {
	Ident          string
	Env            *Env
	Rules          []rules.Rule
	Start          *LLBlock
	localCallstack []string
	functions      map[string]*LLFunc
	rawFunctions   map[string]*ir.Func
	rawIR          *ir.Func
}

func NewLLFunc(e *Env, rawFunc map[string]*ir.Func, irf *ir.Func) *LLFunc {
	return &LLFunc{
		Ident:        irf.Ident(),
		Env:          e,
		Rules:        []rules.Rule{},
		functions:    make(map[string]*LLFunc),
		rawFunctions: rawFunc,
		rawIR:        irf,
	}
}

func (f *LLFunc) Unroll() {

	if isBuiltIn(f.Ident) {
		return
	}

	f.Env.CurrentFunction = util.FormatIdent(f.Ident)

	for _, block := range f.rawIR.Blocks {
		if !f.Env.returnVoid.Check() {
			b := NewLLBlock(f.Env, f.rawFunctions, block)
			b.ParentFunction = util.FormatIdent(f.Ident)
			b.Unroll()
			f.AddBlock(b)
		}
	}

	f.Env.returnVoid.Out()
}

func (f *LLFunc) GetAllRules(enter rules.Rule, exit rules.Rule) []rules.Rule {
	var r []rules.Rule
	if enter != nil {
		r = []rules.Rule{enter}
	}

	r = append(r, f.Rules...)
	if f.Start != nil {
		r = append(r, f.Start.GetAllRules(nil, nil)...)
	}

	if exit != nil {
		r = append(r, exit)
	}
	return r
}

func (f *LLFunc) String() string {
	var sb strings.Builder
	sb.WriteString("LLFunc:\n")
	sb.WriteString("Env:\n")
	sb.WriteString(fmt.Sprintf("VarLoads: %v\n", f.Env.VarLoads))
	sb.WriteString(fmt.Sprintf("VarTypes: %v\n", f.Env.VarTypes))
	sb.WriteString(fmt.Sprintf("CurrentFunction: %s\n", f.Env.CurrentFunction))
	sb.WriteString("Rules:\n")
	for _, rule := range f.Rules {
		sb.WriteString(fmt.Sprintf("- %s\n", rule.String()))
	}
	if f.Start != nil {
		sb.WriteString("Start:\n")
		sb.WriteString(f.Start.String())
	}
	return sb.String()
}

func (f *LLFunc) AddRules(r []rules.Rule) {
	f.Rules = append(f.Rules, r...)
}

func (f *LLFunc) AddBlock(b LLUnit) {
	if f.Start == nil {
		f.Start = b.(*LLBlock)
		return
	}
	for block := f.Start; block != nil; block = block.After {
		if block.After == nil {
			block.After = b.(*LLBlock)
			return
		}
	}
}

func (f *LLFunc) ExecuteCallstack() []rules.Rule {
	stack := util.Copy(f.localCallstack)
	f.localCallstack = []string{}
	r := f.GenerateCallstack(stack)
	return r
}
func (f *LLFunc) GenerateCallstack(callstack []string) []rules.Rule {
	if len(callstack) == 0 {
		return nil
	}

	p := rules.NewParallels(parallelPermutations(callstack))

	for _, fname := range callstack {
		var enter, exit *rules.FuncCall
		var v *LLFunc
		var ok bool

		enter = rules.NewFuncCall(fname, "Enter", f.Env.CurrentRound)
		exit = rules.NewFuncCall(fname, "Exit", f.Env.CurrentRound)
		p.Round = f.Env.CurrentRound
		v, ok = f.functions[fname]
		if !ok {
			v = NewLLFunc(f.Env, f.rawFunctions, f.rawFunctions[fname])
			v.Unroll()
		}

		if len(callstack) == 1 {
			return v.GetAllRules(enter, exit)
		}
		p.Calls[fname] = v.GetAllRules(enter, exit)
	}
	return []rules.Rule{p}
}

type LLBlock struct {
	Ident          string
	Env            *Env
	ParentFunction string
	Round          int
	Rules          []rules.Rule
	After          *LLBlock
	localCallstack []string
	functions      map[string]*LLFunc
	rawFunctions   map[string]*ir.Func
	rawIR          *ir.Block
	irRefs         map[string]rules.Rule //Happens in conditionals, LLVM reference to a value
	irTemps        map[string]int        //Number of times a temp variable is referenced in the block. Used to identify when the generate a rule from a multi-clause conditional (Ors primarily)
}

func NewLLBlock(e *Env, rawFunc map[string]*ir.Func, irb *ir.Block) *LLBlock {
	round := e.CurrentRound
	return &LLBlock{
		Ident:        irb.Ident(),
		Env:          e,
		Rules:        []rules.Rule{},
		Round:        round,
		functions:    make(map[string]*LLFunc),
		rawFunctions: rawFunc,
		rawIR:        irb,
		irRefs:       make(map[string]rules.Rule),
		irTemps:      make(map[string]int),
	}
}

func (b *LLBlock) setRuleRounds(ru []rules.Rule) {
	for _, r := range ru {
		if r == nil {
			continue
		}
		r.SetRound(b.Env.CurrentRound)
	}
}

func (b *LLBlock) scanTemps() {
	for _, inst := range b.rawIR.Insts {
		switch v := inst.(type) {
		case *ir.InstOr:
			b.irTemps[v.Ident()] = b.irTemps[v.Ident()] + 1
			if IsTemp(v.X.Ident()) {
				b.irTemps[v.X.Ident()] = b.irTemps[v.X.Ident()] + 1
			}
			if IsTemp(v.Y.Ident()) {
				b.irTemps[v.Y.Ident()] = b.irTemps[v.Y.Ident()] + 1
			}
		case *ir.InstAnd:
			b.irTemps[v.Ident()] = b.irTemps[v.Ident()] + 1
			if IsTemp(v.X.Ident()) {
				b.irTemps[v.X.Ident()] = b.irTemps[v.X.Ident()] + 1
			}
			if IsTemp(v.Y.Ident()) {
				b.irTemps[v.Y.Ident()] = b.irTemps[v.Y.Ident()] + 1
			}
		case *ir.InstICmp:
			b.irTemps[v.Ident()] = b.irTemps[v.Ident()] + 1
			if IsTemp(v.X.Ident()) {
				b.irTemps[v.X.Ident()] = b.irTemps[v.X.Ident()] + 1
			}
			if IsTemp(v.Y.Ident()) {
				b.irTemps[v.Y.Ident()] = b.irTemps[v.Y.Ident()] + 1
			}
		case *ir.InstFCmp:
			b.irTemps[v.Ident()] = b.irTemps[v.Ident()] + 1
			if IsTemp(v.X.Ident()) {
				b.irTemps[v.X.Ident()] = b.irTemps[v.X.Ident()] + 1
			}
			if IsTemp(v.Y.Ident()) {
				b.irTemps[v.Y.Ident()] = b.irTemps[v.Y.Ident()] + 1
			}
		case *ir.InstXor:
			b.irTemps[v.Ident()] = b.irTemps[v.Ident()] + 1
			if IsTemp(v.X.Ident()) {
				b.irTemps[v.X.Ident()] = b.irTemps[v.X.Ident()] + 1
			}
			if IsTemp(v.Y.Ident()) {
				b.irTemps[v.Y.Ident()] = b.irTemps[v.Y.Ident()] + 1
			}
		}
	}
}

func (b *LLBlock) Unroll() {
	//Scan the block for temp references before parsing
	b.scanTemps()

	// For each non-branching instruction of the basic block.
	for _, inst := range b.rawIR.Insts {
		r := b.parseInstruct(inst)
		b.setRuleRounds(r)
		if len(r) > 0 && r[0] != nil {
			b.AddRules(r)
		}
	}
	//Make sure call stack is clear
	r1 := b.ExecuteCallstack()
	b.setRuleRounds(r1)
	b.AddRules(r1)

	switch term := b.rawIR.Term.(type) {
	case *ir.TermCondBr:
		r := b.parseTermCon(term)
		b.setRuleRounds(r)
		b.AddRules(r)
	case *ir.TermRet:
		b.Env.returnVoid.In()
	}
}

func (b *LLBlock) GetAllRules(enter rules.Rule, exit rules.Rule) []rules.Rule {
	var ru []rules.Rule
	if enter != nil {
		ru = []rules.Rule{enter}
	}

	ru = append(ru, b.Rules...)
	if b.After != nil {
		ru = append(ru, b.After.GetAllRules(nil, nil)...)
	}

	if exit != nil {
		ru = append(ru, exit)
	}
	return ru
}

func (b *LLBlock) String() string {
	var sb strings.Builder
	sb.WriteString("LLBlock:\n")
	sb.WriteString("Env:\n")
	sb.WriteString(fmt.Sprintf("VarLoads: %v\n", b.Env.VarLoads))
	sb.WriteString(fmt.Sprintf("VarTypes: %v\n", b.Env.VarTypes))
	sb.WriteString("Rules:\n")
	for _, rule := range b.Rules {
		sb.WriteString(fmt.Sprintf("- %s\n", rule.String()))
	}
	if b.After != nil {
		sb.WriteString("After:\n")
		sb.WriteString(b.After.String())
	}
	return sb.String()
}

func (b *LLBlock) AddRules(r []rules.Rule) {
	b.Rules = append(b.Rules, r...)
}

func (b *LLBlock) AddBlock(after LLUnit) {
	if b.After == nil {
		b.After = after.(*LLBlock)
		return
	}

	for block := b.After; block != nil; block = block.After {
		if block.After == nil {
			block.After = after.(*LLBlock)
			return
		}
	}
}

func (b *LLBlock) ExecuteCallstack() []rules.Rule {
	stack := util.Copy(b.localCallstack)
	b.localCallstack = []string{}
	r := b.GenerateCallstack(stack)
	return r
}

func (b *LLBlock) GenerateCallstack(callstack []string) []rules.Rule {
	if len(callstack) == 0 {
		return nil
	}

	p := rules.NewParallels(parallelPermutations(callstack))

	for _, fname := range callstack {
		var enter, exit *rules.FuncCall
		var v *LLFunc
		var ok bool

		enter = rules.NewFuncCall(fname, "Enter", b.Env.CurrentRound)
		exit = rules.NewFuncCall(fname, "Exit", b.Env.CurrentRound)
		p.Round = b.Env.CurrentRound
		v, ok = b.functions[fname]
		if !ok {
			v = NewLLFunc(b.Env, b.rawFunctions, b.rawFunctions[fname])
			v.Unroll()

		}
		if len(callstack) == 1 {
			return v.GetAllRules(enter, exit)
		}
		p.Calls[fname] = v.GetAllRules(enter, exit)
	}
	return []rules.Rule{p}
}

func declareVar(id string, ty string, val rules.Rule, solvable bool) *rules.Init {
	var indexed bool
	if IsIndexed(id) {
		indexed = true
	}

	return &rules.Init{
		Ident:    id,
		Type:     ty,
		Value:    val,
		Solvable: solvable,
		Indexed:  indexed,
	}
}

func (b *LLBlock) tempToIdent(ru rules.Rule) rules.Rule {
	switch r := ru.(type) {
	case *rules.Wrap:
		return b.fetchIdent(r.Value, r)
	case *rules.Infix:
		r.X = b.tempToIdent(r.X)
		r.Y = b.tempToIdent(r.Y)
		return r
	case *rules.Prefix:
		r.X = b.tempToIdent(r.X)
		return r
	}
	return ru
}

func (b *LLBlock) fetchIdent(id string, r rules.Rule) rules.Rule {
	if IsTemp(id) {
		refname := fmt.Sprintf("%s-%s", b.ParentFunction, id)
		if load, ok := b.Env.VarLoads[refname]; ok {
			wid := &rules.Vwrap{Value: load}
			return wid
		} else if ref, ok := b.irRefs[refname]; ok {
			switch r := ref.(type) {
			case *rules.Infix:
				r.X = b.tempToIdent(r.X)
				r.Y = b.tempToIdent(r.Y)
				return r
			case *rules.Prefix:
				r.X = b.tempToIdent(r.X)
				return r
			}
		} else {
			panic(fmt.Sprintf("smt generation error, value for %s not found", id))
		}
	}
	return r
}

func NewConstants(e *Env, globals []*ir.Global, RawInputs *llvm.RawInputs) []rules.Rule {
	// Constants cannot be changed and therefore don't increment
	// in SSA. So instead of return a *rule we can skip directly
	// to a set of strings
	b := &LLBlock{
		Ident:        "__constants",
		Env:          e,
		Rules:        []rules.Rule{},
		functions:    make(map[string]*LLFunc),
		rawFunctions: make(map[string]*ir.Func),
		rawIR:        nil,
		irRefs:       make(map[string]rules.Rule),
	}

	r := []rules.Rule{}
	for _, gl := range globals {
		id := util.FormatIdent(gl.GlobalIdent.Ident())
		b.Env.VarTypes[id] = gl.Type().String()

		if !IsIndexed(id) && !IsClocked(id) {
			ru := b.constantRule(id, gl.Init, RawInputs)
			if ru == nil {
				continue
			}
			r = append(r, ru)
		}
	}
	return r
}

func WhenThen(aw []*ast.AssertionStatement) map[string]map[string][]string {
	//Look for "when/then" asserts and build a map defining which
	//variables are on the right side (then) for every variable on the left side (when)
	//this allows us to build out asserts later and capture overlapping SSA transitions
	//Example: a1 is true, b1 must be true. b1 becomes b2 before a1 becomes a2 so b2 must still be true
	var whenThens = make(map[string]map[string][]string)
	for _, a := range aw {
		if a.Constraint.Operator == "then" {
			when := extractVariables(a.Constraint.Left)
			then := extractVariables(a.Constraint.Right)
			for _, w := range when {
				left := util.RemoveFromStringSlice(when, w)
				if whenThens[w] == nil {
					whenThens[w] = make(map[string][]string)
				}
				whenThens[w][a.String()] = append(left, then...)
			}

			for _, t := range then {
				right := util.RemoveFromStringSlice(then, t)
				if whenThens[t] == nil {
					whenThens[t] = make(map[string][]string)
				}
				whenThens[t][a.String()] = append(when, right...)
			}
		}
	}
	return whenThens
}

func extractVariables(e ast.Node) []string {
	var vars []string
	switch v := e.(type) {
	case *ast.Identifier:
		vars = append(vars, v.IdString())
	case *ast.PrefixExpression:
		vars = append(vars, extractVariables(v.Right)...)
	case *ast.InfixExpression:
		vars = append(vars, extractVariables(v.Left)...)
		vars = append(vars, extractVariables(v.Right)...)
	case *ast.IfExpression:
		vars = append(vars, extractVariables(v.Consequence)...)
		vars = append(vars, extractVariables(v.Alternative)...)
	case *ast.IntegerLiteral:
		return vars
	case *ast.FloatLiteral:
		return vars
	case *ast.StringLiteral:
		return vars
	case *ast.Boolean:
		return vars
	case *ast.Natural:
		return vars
	case *ast.Uncertain: //Set to dummy value for LLVM IR, catch during SMT generation
		return vars
	case *ast.Unknown:
		vars = append(vars, v.Name.IdString())
	case *ast.Nil:
		return vars
	case *ast.IndexExpression:
		vars = append(vars, extractVariables(v.Index)...)
	case *ast.AssertVar:
		vars = append(vars, v.Instances...)
	default:
		panic(fmt.Sprintf("unrecognized expression: %T", v))
	}
	return vars
}

func (b *LLBlock) constantRule(id string, c constant.Constant, RawInputs *llvm.RawInputs) rules.Rule {
	if id == "__rounds" || id == "__parallelGroup" {
		return nil
	}

	switch val := c.(type) {
	case *constant.Int:
		ty := LookupType(id, val)
		return declareVar(id, ty, &rules.Wrap{Value: val.X.String()}, false)
	case *constant.ExprAnd, *constant.ExprOr, *constant.ExprFNeg:
		ty := LookupType(id, val)
		x := b.constExpr(val)
		return declareVar(id, ty, x, false)
	default:
		ty := LookupType(id, val)
		return declareVar(id, ty, &rules.Wrap{Value: val.String()}, false)
	case *constant.Float:
		ty := LookupType(id, val)
		if isASolvable(id, RawInputs) {
			return declareVar(id, ty, &rules.Wrap{Value: val.X.String()}, true)
		} else {
			v := val.X.String()
			if strings.Contains(v, ".") {
				return declareVar(id, ty, &rules.Wrap{Value: v}, false)
			}
			return declareVar(id, ty, &rules.Wrap{Value: v + ".0"}, false)
		}
	}
}

func (b *LLBlock) constExpr(con constant.Constant) rules.Rule {
	switch inst := con.(type) {
	case *constant.ExprAnd:
		id := inst.Ident()
		x := inst.X.Ident()
		y := inst.Y.Ident()
		return b.createInfixRule(id, x, y, "and")
	case *constant.ExprOr:
		id := inst.Ident()
		x := inst.X.Ident()
		y := inst.Y.Ident()
		return b.createInfixRule(id, x, y, "or")
	case *constant.ExprFNeg:
		id := inst.Ident()
		x := inst.X.Ident()
		stmt := util.FormatIdent(x)
		return b.createPrefixRule(id, stmt, "not")
	default:
		panic(fmt.Sprintf("unrecognized constant expression: %T", inst))

	}
}

func isBuiltIn(c string) bool {
	if c == "@advance" || c == "@stay" {
		return true
	}
	return false
}

func convertInfixVar(e *Env, x string) (string, string, bool) {
	if IsTemp(x) {
		refname := fmt.Sprintf("%s-%s", e.CurrentFunction, x)
		if v, ok := e.VarLoads[refname]; ok {
			ty := LookupType(refname, v)
			xid := v.Ident()
			return util.FormatIdent(xid), ty, true
		}
	}

	if IsGlobal(x) {
		return util.FormatIdent(x), e.VarTypes[x], true
	}

	if IsInt(x) {
		return x, "Bool", false
	}

	if IsNumeric(x) {
		return x, "Real", false
	}

	if IsBoolean(x) {
		return x, "Bool", false
	}

	return x, e.VarTypes[x], false
}

func (b *LLBlock) createInfixRule(id string, x string, y string, op string) rules.Rule {
	var tyX, tyY string
	var vrX, vrY bool

	x, tyX, vrX = convertInfixVar(b.Env, x)
	y, tyY, vrY = convertInfixVar(b.Env, y)
	xIs := IsIndexed(x)
	yIs := IsIndexed(y)
	_, file, line, _ := runtime.Caller(1)

	xr := rules.NewWrap(x, tyX, vrX, file, line, false, xIs)
	xr.SetWhensThens(b.Env.WhensThens)
	yr := rules.NewWrap(y, tyY, vrY, file, line, false, yIs)
	yr.SetWhensThens(b.Env.WhensThens)

	return &rules.Infix{
		X:  xr,
		Y:  yr,
		Op: op,
	}
}

func (b *LLBlock) createPrefixRule(id string, x string, op string) rules.Rule {
	var vr bool
	if _, ok := b.Env.VarTypes[x]; ok {
		vr = true
	}
	xIs := IsIndexed(x)
	_, file, line, _ := runtime.Caller(1)

	xr := rules.NewWrap(x, "Bool", vr, file, line, false, xIs)
	xr.SetWhensThens(b.Env.WhensThens)
	return &rules.Prefix{
		X:  xr,
		Op: op,
	}
}

func parallelPermutations(p []string) (permuts [][]string) {
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

func (b *LLBlock) isSameParallelGroup(meta ir.Metadata) bool {
	for _, v := range meta {

		if v.Name == b.Env.ParallelGrouping {
			return true
		}

		if b.Env.ParallelGrouping == "" {
			return true
		}
	}

	return false
}

func (b *LLBlock) singleParallelStep(callee string) bool {
	if len(b.localCallstack) == 0 {
		return false
	}

	if callee == b.localCallstack[len(b.localCallstack)-1] {
		return true
	}

	return false
}

func (b *LLBlock) updateParallelGroup(meta ir.Metadata) {
	for _, v := range meta {
		if v.Name[0:5] != "round-" {
			b.Env.ParallelGrouping = v.Name
		}
	}
}
