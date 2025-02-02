package unroll

import (
	"fault/llvm"
	"fault/smt/rules"
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
	VarLoads        map[string]value.Value
	VarTypes        map[string]string
	CurrentFunction string
}

func NewEnv() *Env {
	return &Env{
		VarLoads: make(map[string]value.Value),
		VarTypes: make(map[string]string),
	}
}

type LLUnit interface {
	Unroll()
	AddRules([]rules.Rule)
	AddBlock(LLUnit)
	ExecuteCallstack() *LLFunc
	String() string
}

type LLFunc struct {
	Env            *Env
	Rules          []rules.Rule
	Start          *LLBlock
	localCallstack []string
	functions      map[string]*LLFunc
	rawFunctions   map[string]*ir.Func
	rawIR          *ir.Func
}

func NewLLFunc(e *Env, irf *ir.Func) *LLFunc {
	return &LLFunc{
		Env:          e,
		Rules:        []rules.Rule{},
		functions:    make(map[string]*LLFunc),
		rawFunctions: make(map[string]*ir.Func),
		rawIR:        irf,
	}
}

func (f *LLFunc) Unroll() {}

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

func (f *LLFunc) ExecuteCallstack() *LLFunc {
	stack := util.Copy(f.localCallstack)
	f.localCallstack = []string{}
	r := GenerateCallstack(f, stack)
	return r
}

type LLBlock struct {
	Env            *Env
	Rules          []rules.Rule
	After          *LLBlock
	localCallstack []string
	functions      map[string]*LLFunc
	rawFunctions   map[string]*ir.Func
	rawIR          *ir.Block
}

func NewLLBlock(e *Env, irb *ir.Block) *LLBlock {
	return &LLBlock{
		Env:          e,
		Rules:        []rules.Rule{},
		functions:    make(map[string]*LLFunc),
		rawFunctions: make(map[string]*ir.Func),
		rawIR:        irb,
	}
}

func (b *LLBlock) Unroll() {}

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

func (f *LLBlock) ExecuteCallstack() *LLFunc {
	stack := util.Copy(f.localCallstack)
	f.localCallstack = []string{}
	r := GenerateCallstack(f, stack)
	return r
}

func GenerateCallstack(llu LLUnit, callstack []string) *LLFunc {
	if len(callstack) == 0 {
		return nil
	}

	// if len(callstack) > 1 {
	// 	//Generate parallel runs
	// 	perm := g.parallelPermutations(callstack)
	// 	return g.runParallel(perm)
	// } else {
	var v *LLFunc
	var ok bool
	fname := callstack[0]

	switch u := llu.(type) {
	case *LLFunc:
		v, ok = u.functions[fname]
		if !ok {
			v = NewLLFunc(u.Env, u.rawFunctions[fname])
		}
	case *LLBlock:
		v, ok = u.functions[fname]
		if !ok {
			v = NewLLFunc(u.Env, u.rawFunctions[fname])
		}
	}
	return v
	//}
}

func declareVar(id string, ty string, val string) *rules.Init {
	return &rules.Init{
		Ident: id,
		Type:  ty,
		Value: val,
	}
}

func NewConstants(e *Env, globals []*ir.Global, RawInputs *llvm.RawInputs) []rules.Rule {
	// Constants cannot be changed and therefore don't increment
	// in SSA. So instead of return a *rule we can skip directly
	// to a set of strings
	b := NewLLBlock(e, nil)
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

func convertInfixVar(e *Env, x string) string {
	if IsTemp(x) {
		refname := fmt.Sprintf("%s-%s", e.CurrentFunction, x)
		if v, ok := e.VarLoads[refname]; ok {
			xid := v.Ident()
			return util.FormatIdent(xid)
		}
	}

	if IsGlobal(x) {
		return util.FormatIdent(x)
	}
	return x
}

func (b *LLBlock) createInfixRule(id string, x string, y string, op string) rules.Rule {
	x = convertInfixVar(b.Env, x)
	y = convertInfixVar(b.Env, y)
	return &rules.Infix{
		X:  &rules.Wrap{Value: x},
		Y:  &rules.Wrap{Value: y},
		Op: op,
	}
}

func (b *LLBlock) createPrefixRule(id string, x string, op string) rules.Rule {
	return &rules.Prefix{
		X:  &rules.Wrap{Value: x},
		Op: op,
	}
}
