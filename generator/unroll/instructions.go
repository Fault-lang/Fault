package unroll

import (
	"fault/generator/rules"
	"fault/util"
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/llir/llvm/ir"
	irtypes "github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (b *LLBlock) parseInstruct(inst ir.Instruction) []rules.Rule {
	switch inst := inst.(type) {
	case *ir.InstAlloca:
		//Do nothing
	case *ir.InstStore:
		return b.parseStore(inst)
	case *ir.InstLoad:
		return b.parseLoad(inst)
	case *ir.InstFAdd:
		return b.parseAdd(inst)
	case *ir.InstFSub:
		return b.parseSub(inst)
	case *ir.InstFMul:
		return b.parseMul(inst)
	case *ir.InstFDiv:
		return b.parseDiv(inst)
	case *ir.InstFRem:
		//Cannot be implemented because SMT solvers do poorly with modulo
	case *ir.InstICmp:
		return b.parseICmp(inst)
	case *ir.InstFCmp:
		return b.parseFCmp(inst)
	case *ir.InstCall:
		return b.parseCall(inst)
	// case *ir.InstPhi:
	// 	return b.parsePhi(inst)
	// case *ir.InstGetElementPtr:
	// 	return b.parseGetElementPtr(inst)
	case *ir.InstXor:
		return b.parseXor(inst)
	case *ir.InstAnd:
		return b.parseAnd(inst)
	case *ir.InstOr:
		return b.parseOr(inst)
	case *ir.InstBitCast:
		//Do nothing
	case *ir.InstFNeg:
		return b.parseFNeg(inst)
	default:
		panic(fmt.Sprintf("unrecognized instruction: %T", inst))
	}
	return []rules.Rule{}
}

func (b *LLBlock) parseStore(inst *ir.InstStore) []rules.Rule {
	var ru []rules.Rule
	vname := inst.Dst.Ident()
	if vname == "@__rounds" {
		//Clear the callstack first
		r := b.ExecuteCallstack()
		b.setRuleRounds(r)
		b.AddRules(r)
		round, err := strconv.Atoi(inst.Src.Ident())
		if err != nil {
			panic(fmt.Sprintf("failed to parse round value '%s': %v", inst.Src.Ident(), err))
		}
		b.Env.CurrentRound = round
		return ru
	}

	if vname == "@__parallelGroup" {
		return ru
	}

	switch inst.Src.Type().(type) {
	case *irtypes.ArrayType:
		refname := fmt.Sprintf("%s-%s", b.ParentFunction, inst.Dst.Ident())
		b.Env.VarLoads[refname] = inst.Src
	default:
		base := util.FormatIdent(inst.Dst.Ident())
		if IsTemp(inst.Src.Ident()) {
			refname := fmt.Sprintf("%s-%s", b.ParentFunction, inst.Src.Ident())
			if val, ok := b.Env.VarLoads[refname]; ok {
				ty := LookupType(refname, val)
				b.Env.VarTypes[refname] = ty

				v := FormatValue(val)
				if !IsBoolean(v) && !IsNumeric(v) {
					v = util.FormatIdent(v)
				}
				ru = append(ru, b.createRule(base, v, ty, "="))
			} else if ref, ok := b.irRefs[refname]; ok {
				switch r := ref.(type) {
				case *rules.Infix:
					r.X = b.tempToIdent(r.X)
					r.Y = b.tempToIdent(r.Y)
					xIs := IsIndexed(base)
					_, file, line, _ := runtime.Caller(1)
					wid := rules.NewWrap(base, "", true, file, line, true, xIs)

					if IsStaticValue(r.X.String()) {
						wid.Variable = false
					}

					if IsBoolean(r.Y.String()) {
						wid.Type = "Bool"
						ru = append(ru, &rules.Infix{X: wid, Ty: "Bool", Y: r, Op: "="})
					} else if IsNumeric(r.Y.String()) {
						wid.Type = "Real"
						ru = append(ru, &rules.Infix{X: wid, Ty: "Real", Y: r, Op: "="})
					} else if isASolvable(r.X.String(), b.Env.RawInputs) {
						wid.Type = "Real"
						ru = append(ru, &rules.Infix{X: wid, Ty: "Real", Y: r, Op: "="})
					} else {
						wid.Type = "Real"
						ru = append(ru, &rules.Infix{X: wid, Ty: "Real", Y: r, Op: "="})
					}
					b.Env.VarTypes[base] = wid.Type
				default:
					ty := LookupType(base, nil)
					b.Env.VarTypes[base] = ty
					xIs := IsIndexed(base)
					_, file, line, _ := runtime.Caller(1)
					wid := rules.NewWrap(base, ty, true, file, line, true, xIs)
					ru = append(ru, &rules.Infix{X: wid, Ty: ty, Y: r})
				}
			} else {
				panic(fmt.Sprintf("smt generation error, value for %s not found", base))
			}
		} else {
			ty := LookupType(base, inst.Src)
			b.Env.VarTypes[base] = ty

			ru = append(ru, b.createRule(base, inst.Src.Ident(), ty, "="))
		}
		return ru
	}
	return ru
}

func (b *LLBlock) createRule(id string, val string, ty string, op string) rules.Rule {
	xIs := IsIndexed(id)
	_, file, line, _ := runtime.Caller(1)
	wid := rules.NewWrap(id, ty, true, file, line, true, xIs)
	var wval *rules.Wrap

	if IsBoolean(val) {
		_, file, line, _ := runtime.Caller(1)
		wval = rules.NewWrap(val, "Bool", false, file, line, false, false)
	} else if IsNumeric(val) {
		_, file, line, _ := runtime.Caller(1)
		wval = rules.NewWrap(val, ty, false, file, line, false, false)
	} else {
		_, file, line, _ := runtime.Caller(1)
		wval = rules.NewWrap(val, ty, true, file, line, false, false)
	}
	return &rules.Infix{X: wid, Ty: ty, Y: wval, Op: op}
}

func (b *LLBlock) tempRule(inst value.Value, r rules.Rule) {
	// If infix rule is stored in a temp variable
	id := inst.Ident()
	if IsTemp(id) {
		refname := fmt.Sprintf("%s-%s", b.Env.CurrentFunction, id)
		b.irRefs[refname] = r
	}
}

func (b *LLBlock) parseLoad(inst *ir.InstLoad) []rules.Rule {
	refname := fmt.Sprintf("%s-%s", b.Env.CurrentFunction, inst.Ident())
	b.Env.VarLoads[refname] = inst.Src
	b.Env.VarTypes[refname] = inst.Src.Type().String()
	return []rules.Rule{}
}

func (b *LLBlock) parseAdd(inst *ir.InstFAdd) []rules.Rule {
	r := b.createInfixRule(inst.Ident(), inst.X.Ident(), inst.Y.Ident(), "+")
	b.tempRule(inst, r)
	return []rules.Rule{}
}

func (b *LLBlock) parseSub(inst *ir.InstFSub) []rules.Rule {
	r := b.createInfixRule(inst.Ident(), inst.X.Ident(), inst.Y.Ident(), "-")
	b.tempRule(inst, r)
	return []rules.Rule{}
}

func (b *LLBlock) parseMul(inst *ir.InstFMul) []rules.Rule {
	r := b.createInfixRule(inst.Ident(), inst.X.Ident(), inst.Y.Ident(), "*")
	b.tempRule(inst, r)
	return []rules.Rule{}
}

func (b *LLBlock) parseDiv(inst *ir.InstFDiv) []rules.Rule {
	r := b.createInfixRule(inst.Ident(), inst.X.Ident(), inst.Y.Ident(), "/")
	b.tempRule(inst, r)
	return []rules.Rule{}
}

func (b *LLBlock) parseICmp(inst *ir.InstICmp) []rules.Rule {
	var r rules.Rule
	op, y := b.createCompareRule(inst.Pred.String())
	if op == "true" || op == "false" {
		r = b.createInfixRule(inst.Ident(),
			inst.X.Ident(), y.(*rules.Wrap).Value, op)
	} else {
		r = b.createInfixRule(inst.Ident(),
			inst.X.Ident(), inst.Y.Ident(), op)
	}

	// If LLVM is storing this is a temp var
	// Happens in conditionals
	if IsTemp(inst.Ident()) {
		refname := fmt.Sprintf("%s-%s", b.Env.CurrentFunction, inst.Ident())
		b.irRefs[refname] = r
		return []rules.Rule{}
	}

	return []rules.Rule{r}
}
func (b *LLBlock) parseFCmp(inst *ir.InstFCmp) []rules.Rule {
	var r rules.Rule
	op, y := b.createCompareRule(inst.Pred.String())
	if op == "true" || op == "false" {
		r = b.createInfixRule(inst.Ident(),
			inst.X.Ident(), y.(*rules.Wrap).Value, op)
	} else {
		r = b.createInfixRule(inst.Ident(),
			inst.X.Ident(), inst.Y.Ident(), op)
	}

	// If LLVM is storing this is a temp var
	// Happens in conditionals
	if IsTemp(inst.Ident()) {
		refname := fmt.Sprintf("%s-%s", b.Env.CurrentFunction, inst.Ident())
		b.irRefs[refname] = r
		return []rules.Rule{}
	}
	return []rules.Rule{r}
}

func (b *LLBlock) parseCall(inst *ir.InstCall) []rules.Rule {
	var r []rules.Rule

	callee := inst.Callee.Ident()
	if isBuiltIn(callee) {
		meta := inst.Metadata // Is this in a "b || b" construction?
		if len(meta) > 0 {
			refname := fmt.Sprintf("%s-%s", b.Env.CurrentFunction, inst.Ident())
			b.Env.VarLoads[refname] = inst
		} else {
			r := b.parseBuiltIn(inst, false)
			return r
		}
		return []rules.Rule{}
	}

	meta := inst.Metadata
	callee = util.FormatIdent(callee)

	if b.isSameParallelGroup(meta) {
		b.localCallstack = append(b.localCallstack, callee)
	} else if b.singleParallelStep(callee) {
		r0 := b.ExecuteCallstack()
		r = append(r, r0...)

		r1 := b.GenerateCallstack([]string{callee})
		r = append(r, r1...)
	} else {
		r0 := b.ExecuteCallstack()
		r = append(r, r0...)
		b.localCallstack = append(b.localCallstack, callee)
	}
	b.updateParallelGroup(meta)
	b.Env.returnVoid.Out()

	return r
}

func (b *LLBlock) parseBuiltIn(call *ir.InstCall, complex bool) []rules.Rule {
	p := call.Args
	if len(p) == 0 {
		return []rules.Rule{}
	}

	bc, ok := p[0].(*ir.InstBitCast)
	if !ok {
		panic("improper argument to built in function")
	}

	id := bc.From.Ident()
	refname := fmt.Sprintf("%s-%s", b.Env.CurrentFunction, id)
	state := b.Env.VarLoads[refname]
	newState := util.FormatIdent(state.Ident())
	// Not sure I remember/understand this. Commenting
	// it out for now.

	// if complex {
	// 	declareVar(newState, "Bool", "true")
	// }
	r1 := b.createRule(newState, "true", "Bool", "=")

	currentFunction := b.Env.CurrentFunction

	if currentFunction[len(currentFunction)-7:] != "__state" {
		panic("calling advance from outside the state chart")
	}

	//base2 := currentFunction[1:len(currentFunction)-7]
	base2 := currentFunction[:len(currentFunction)-7]

	// if complex {
	// 	declareVar(base2, "Bool", "false")
	// }
	r2 := b.createRule(base2, "false", "Bool", "=")
	return []rules.Rule{r1, r2}
}

func (b *LLBlock) parseTerms(terms []*ir.Block) ([]rules.Rule, []rules.Rule, []rules.Rule, []string) {
	var t, f, a []rules.Rule
	block_names := []string{"", "", ""}
	for _, term := range terms {
		bname := strings.Split(term.Ident(), "-")
		switch bname[len(bname)-1] {
		case "true":
			block_names[0] = term.Ident()
			b.Env.returnVoid.In()
			true_block := NewLLBlock(b.Env, b.rawFunctions, term)
			true_block.ParentFunction = b.Env.CurrentFunction
			true_block.Unroll()
			t = true_block.GetAllRules(nil, nil)
			b.Env.returnVoid.Out()
		case "false":
			block_names[1] = term.Ident()
			b.Env.returnVoid.In()
			false_block := NewLLBlock(b.Env, b.rawFunctions, term)
			false_block.ParentFunction = b.Env.CurrentFunction
			false_block.Unroll()
			f = false_block.GetAllRules(nil, nil)

			b.Env.returnVoid.Out()
		case "after":
			block_names[2] = term.Ident()
			after_block := NewLLBlock(b.Env, b.rawFunctions, term)
			after_block.ParentFunction = b.Env.CurrentFunction
			after_block.Unroll()
			a = after_block.GetAllRules(nil, nil)
		default:
			panic(fmt.Sprintf("unrecognized terminal branch: %s", term.Ident()))
		}
	}

	return t, f, a, block_names
}

func (b *LLBlock) parseCondNode(node value.Value) rules.Rule {
	switch cnode := node.(type) {
	case *ir.InstCall:
		if isBuiltIn(cnode.Callee.Ident()) {
			if cnode.Callee.Ident() == "@advance" {
				r := b.parseBuiltIn(cnode, true)
				if len(r) == 1 {
					return r[0]
				}
				return &rules.Ands{
					X: r,
				}
			}
		}
	case *ir.InstOr:
		id := cnode.Ident()
		refname := fmt.Sprintf("%s-%s", b.Env.CurrentFunction, id)
		if v, ok := b.irRefs[refname]; ok {
			return v
		}
		r := b.parseOr(cnode)
		// Or always returns a single rule or a nil
		if len(r) == 0 {
			panic(fmt.Sprintf("Or clause %s not found", refname))
		}
		return r[0]
	case *ir.InstAnd:
		id := cnode.Ident()
		refname := fmt.Sprintf("%s-%s", b.Env.CurrentFunction, id)
		if v, ok := b.irRefs[refname]; ok {
			return v
		}
		r := b.parseAnd(cnode)
		// And always returns a single rule or a nil
		if len(r) == 0 {
			panic(fmt.Sprintf("And clause %s not found", refname))
		}
		return r[0]
	case *ir.InstXor:
		id := cnode.Ident()
		refname := fmt.Sprintf("%s-%s", b.Env.CurrentFunction, id)
		if v, ok := b.irRefs[refname]; ok {
			return v
		}
		r := b.parseXor(cnode)
		if len(r) == 0 {
			panic(fmt.Sprintf("Xor clause %s not found", refname))
		}
		return r[0]
	default:
		n := node.Ident()
		nRule := b.LookupCondPart(b.Env.CurrentFunction, n)
		if nRule == nil {
			n = b.ConvertIdent(b.Env.CurrentFunction, n)
			nIs := IsIndexed(n)
			_, file, line, _ := runtime.Caller(1)
			nRule = rules.NewWrap(n, "Bool", true, file, line, false, nIs)
		}
		return nRule
	}
	return nil
}

func (b *LLBlock) parseTermCon(term *ir.TermCondBr) []rules.Rule {
	var cond rules.Rule
	b.Env.returnVoid.In()
	id := term.Cond.Ident()
	if IsTemp(id) {
		refname := fmt.Sprintf("%s-%s", b.Env.CurrentFunction, id)
		if v, ok := b.irRefs[refname]; ok {
			cond = v
		}
	} else if IsBoolean(id) ||
		IsNumeric(id) {
		ty := LookupType(id, nil)
		xIs := IsIndexed(id)
		_, file, line, _ := runtime.Caller(1)
		cond = rules.NewWrap(id, ty, false, file, line, true, xIs)
	}
	b.Env.returnVoid.Out()

	t, f, a, block_names := b.parseTerms(term.Succs())

	ite := rules.NewIte(cond, t, f, a, block_names)
	return []rules.Rule{ite}
}

func (b *LLBlock) parsePhi(inst *ir.InstPhi) []rules.Rule {
	return []rules.Rule{}
}

func (b *LLBlock) parseGetElementPtr(inst *ir.InstGetElementPtr) []rules.Rule {
	return []rules.Rule{}
}

func (b *LLBlock) deferRule(id string, x value.Value) bool {
	if b.irTemps[id] > 1 {
		//b.irTemps[id] = b.irTemps[id] - 1
		return false
	}

	switch node := x.(type) {
	case *ir.InstCall:
		if isBuiltIn(node.Callee.Ident()) {
			return true
		}
	case *ir.InstOr:
		return true
	case *ir.InstAnd:
		return true
	//We don't do XORs because we only use XOR for "not x"

	default:
		return false
	}
	return false
}

func (b *LLBlock) parseXor(inst *ir.InstXor) []rules.Rule {
	id := inst.Ident()
	xRule := b.parseCondNode(inst.X)
	_, file, line, _ := runtime.Caller(1)
	b.createMultiCondRule(id, xRule, rules.NewWrap("", "", false, file, line, false, false), "not")

	if b.deferRule(id, inst.X) || b.deferRule(id, inst.Y) {
		refname := fmt.Sprintf("%s-%s", b.Env.CurrentFunction, id)
		if v, ok := b.irRefs[refname]; ok {
			return []rules.Rule{v}
		}
	}
	return nil
}

func (b *LLBlock) createMultiCondRule(id string, x rules.Rule, y rules.Rule, op string) rules.Rule {
	refname := fmt.Sprintf("%s-%s", b.Env.CurrentFunction, id)
	if _, ok := b.irRefs[refname]; ok {
		return nil
	}

	if op == "not" {
		b.irRefs[refname] = &rules.Prefix{X: x, Ty: "Bool", Op: op}
		return nil
	}

	if op == "or" {
		// A little convoluted because we need to add phis to or clauses
		var right, left []rules.Rule

		if x_ands, ok := x.(*rules.Ands); ok {
			right = x_ands.X
		} else {
			right = []rules.Rule{x}
		}

		if _, ok := y.(*rules.Ands); ok {
			left = y.(*rules.Ands).X
		} else {
			left = []rules.Rule{y}

		}

		//Consolidate Ors instead of nesting them
		if r, ok := x.(*rules.Ors); ok {
			ors := append(r.X, left)
			b.irRefs[refname] = &rules.Ors{X: ors, BranchName: refname}
			return nil
		}

		if r, ok := y.(*rules.Ors); ok {
			ors := append([][]rules.Rule{right}, r.X...)
			b.irRefs[refname] = &rules.Ors{X: ors, BranchName: refname}
			return nil
		}

		b.irRefs[refname] = &rules.Ors{X: [][]rules.Rule{right, left}, BranchName: refname}
		return nil
	}

	if op == "and" {
		if r1, ok := x.(*rules.Ands); ok {
			if r2, ok := y.(*rules.Ands); ok {
				b.irRefs[refname] = &rules.Ands{X: append(r1.X, r2.X...)}
				return nil
			}

			b.irRefs[refname] = &rules.Ands{X: append(r1.X, y)}
			return nil
		}

		if r1, ok := y.(*rules.Ands); ok {
			b.irRefs[refname] = &rules.Ands{X: append([]rules.Rule{x}, r1.X...)}
			return nil
		}

		b.irRefs[refname] = &rules.Ands{X: []rules.Rule{x, y}}
		return nil
	}

	b.irRefs[refname] = &rules.Infix{X: x, Ty: "Bool", Y: y, Op: op}
	return nil
}

func (b *LLBlock) parseAnd(inst *ir.InstAnd) []rules.Rule {
	id := inst.Ident()
	xRule := b.parseCondNode(inst.X)
	yRule := b.parseCondNode(inst.Y)

	b.createMultiCondRule(id, xRule, yRule, "and")

	if b.deferRule(id, inst.X) || b.deferRule(id, inst.Y) {
		refname := fmt.Sprintf("%s-%s", b.Env.CurrentFunction, id)
		if v, ok := b.irRefs[refname]; ok {
			return []rules.Rule{v}
		}
	}
	return []rules.Rule{}
}

func (b *LLBlock) parseOr(inst *ir.InstOr) []rules.Rule {
	id := inst.Ident()
	xRule := b.parseCondNode(inst.X)
	yRule := b.parseCondNode(inst.Y)
	b.createMultiCondRule(id, xRule, yRule, "or")

	if b.deferRule(id, inst) { // Ors are different because we collapse them
		refname := fmt.Sprintf("%s-%s", b.Env.CurrentFunction, id)
		if v, ok := b.irRefs[refname]; ok {
			return []rules.Rule{v}
		}
	}

	return []rules.Rule{}
}

func (b *LLBlock) parseBitCast(inst *ir.InstBitCast) []rules.Rule {
	panic(fmt.Sprint("unimplemented bitcast"))
}

func (b *LLBlock) parseFNeg(inst *ir.InstFNeg) []rules.Rule {
	panic(fmt.Sprint("unimplemented FNeg"))
}

func (b *LLBlock) createCompareRule(op string) (string, rules.Rule) {
	var y *rules.Wrap
	op = b.compareRuleOp(op)
	switch op {
	case "false":
		_, file, line, _ := runtime.Caller(1)
		y = rules.NewWrap("False", "Bool", false, file, line, false, false)
	case "true":
		_, file, line, _ := runtime.Caller(1)
		y = rules.NewWrap("True", "Bool", false, file, line, false, false)
	}
	return op, y
}

func (b *LLBlock) compareRuleOp(op string) string {
	switch op {
	case "false":
		return "false"
	case "oeq":
		return "="
	case "eq":
		return "="
	case "oge":
		return ">="
	case "ogt":
		return ">"
	case "ole":
		return "<="
	case "olt":
		return "<"
	case "one":
		return "!="
	case "ne":
		return "!="
	case "true":
		return "true"
	case "ueq":
		return "="
	case "uge":
		return ">="
	case "ugt":
		return ">"
	case "ule":
		return "<="
	case "ult":
		return "<"
	case "une":
		return "!="
	default:
		return op
	}
}
