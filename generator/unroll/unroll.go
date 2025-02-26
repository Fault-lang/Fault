package unroll

import (
	"fault/generator/rules"
	"fault/llvm"
	"fault/util"
	"fmt"
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
}

func NewEnv(ri *llvm.RawInputs) *Env {
	return &Env{
		RawInputs:    ri,
		VarLoads:     make(map[string]value.Value),
		VarTypes:     make(map[string]string),
		CurrentRound: 0,
		returnVoid:   NewPhiState(),
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

func (f *LLFunc) GetAllRules() []rules.Rule {
	var r []rules.Rule
	r = append(r, f.Rules...)
	if f.Start != nil {
		r = append(r, f.Start.GetAllRules()...)
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
	r := GenerateCallstack(f, stack)
	return r
}

type LLBlock struct {
	Ident          string
	Env            *Env
	ParentFunction string
	Rules          []rules.Rule
	After          *LLBlock
	localCallstack []string
	functions      map[string]*LLFunc
	rawFunctions   map[string]*ir.Func
	rawIR          *ir.Block
	irRefs         map[string]rules.Rule //Happens in conditionals, LLVM reference to a value
}

func NewLLBlock(e *Env, rawFunc map[string]*ir.Func, irb *ir.Block) *LLBlock {
	return &LLBlock{
		Ident:        irb.Ident(),
		Env:          e,
		Rules:        []rules.Rule{},
		functions:    make(map[string]*LLFunc),
		rawFunctions: rawFunc,
		rawIR:        irb,
		irRefs:       make(map[string]rules.Rule),
	}
}

func (b *LLBlock) Unroll() {

	// For each non-branching instruction of the basic block.
	for _, inst := range b.rawIR.Insts {
		r := b.parseInstruct(inst)
		if len(r) > 0 {
			b.AddRules(r)
		}
	}
	//Make sure call stack is clear
	r1 := b.ExecuteCallstack()
	b.AddRules(r1)

	switch term := b.rawIR.Term.(type) {
	case *ir.TermCondBr:
		b.parseTermCon(term)
	case *ir.TermRet:
		b.Env.returnVoid.In()
	}
}

func (b *LLBlock) GetAllRules() []rules.Rule {
	var ru []rules.Rule
	ru = append(ru, b.Rules...)
	if b.After != nil {
		ru = append(ru, b.After.GetAllRules()...)
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

func (f *LLBlock) ExecuteCallstack() []rules.Rule {
	stack := util.Copy(f.localCallstack)
	f.localCallstack = []string{}
	r := GenerateCallstack(f, stack)
	return r
}

func GenerateCallstack(llu LLUnit, callstack []string) []rules.Rule {
	if len(callstack) == 0 {
		return nil
	}

	p := rules.NewParallels(parallelPermutations(callstack))

	for _, fname := range callstack {
		var v *LLFunc
		var ok bool

		switch u := llu.(type) {
		case *LLFunc:
			p.Round = u.Env.CurrentRound
			v, ok = u.functions[fname]
			if !ok {
				v = NewLLFunc(u.Env, u.rawFunctions, u.rawFunctions[fname])
				v.Unroll()
			}
		case *LLBlock:
			p.Round = u.Env.CurrentRound
			v, ok = u.functions[fname]
			if !ok {
				v = NewLLFunc(u.Env, u.rawFunctions, u.rawFunctions[fname])
				v.Unroll()

			}
		}
		if len(callstack) == 1 {
			return v.Rules
		}
		p.Calls[fname] = v.GetAllRules()
	}
	return []rules.Rule{p}
}

func declareVar(id string, ty string, val string) *rules.Init {
	return &rules.Init{
		Ident: id,
		Type:  ty,
		Value: val,
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
		if !IsIndexed(id) && !IsClocked(id) {
			r = append(r, b.constantRule(id, gl.Init, RawInputs))
		}
	}
	return r
}

func (b *LLBlock) constantRule(id string, c constant.Constant, RawInputs *llvm.RawInputs) rules.Rule {
	if id == "__rounds" || id == "__parallelGroup" {
		return nil
	}

	switch val := c.(type) {
	case *constant.Int:
		ty := LookupType(id, val)
		return declareVar(id, ty, val.X.String())
	case *constant.ExprAnd, *constant.ExprOr, *constant.ExprFNeg:
		return b.constExpr(val)
	default:
		ty := LookupType(id, val)
		return declareVar(id, ty, val.String())
	case *constant.Float:
		ty := LookupType(id, val)
		if isASolvable(id, RawInputs) {
			return declareVar(id, ty, val.X.String())
		} else {
			v := val.X.String()
			if strings.Contains(v, ".") {
				return declareVar(id, ty, v)
			}
			return declareVar(id, ty, v+".0")
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
	return &rules.Infix{
		X:  rules.NewWrap(x, tyX, vrX, "unroll.go", "466", false),
		Y:  rules.NewWrap(y, tyY, vrY, "unroll.go", "467", false),
		Op: op,
	}
}

func (b *LLBlock) createPrefixRule(id string, x string, op string) rules.Rule {
	var vr bool
	if _, ok := b.Env.VarTypes[x]; ok {
		vr = true
	}

	return &rules.Prefix{
		X:  rules.NewWrap(x, "Bool", vr, "unroll.go", "476", false),
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
