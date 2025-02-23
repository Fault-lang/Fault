package unroll

import (
	"fault/generator/rules"
	"fault/util"
	"fmt"

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

func (b *LLBlock) parseAlloca(inst *ir.InstAlloca) []rules.Rule {
	return []rules.Rule{}
}

func (b *LLBlock) parseStore(inst *ir.InstStore) []rules.Rule {
	var ru []rules.Rule
	vname := inst.Dst.Ident()
	if vname == "@__rounds" {
		//Clear the callstack first
		r := b.ExecuteCallstack()
		b.AddRules(r)
		//Initate new round
		b.Env.CurrentRound = b.Env.CurrentRound + 1
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
					wid := rules.NewWrap(base, "", true, "instructions.go", "97", true)

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
						ru = append(ru, &rules.Infix{X: wid, Ty: "Real", Y: r.Y, Op: r.Op})
					}
					b.Env.VarTypes[base] = wid.Type
				default:
					ty := LookupType(base, nil)
					b.Env.VarTypes[base] = ty

					wid := rules.NewWrap(base, ty, true, "instructions.go", "118", true)
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
	wid := rules.NewWrap(id, ty, true, "instructions.go", "134", true)
	var wval *rules.Wrap

	if IsBoolean(val) {
		wval = rules.NewWrap(val, "Bool", false, "instructions.go", "135", false)
	} else if IsNumeric(val) {
		wval = rules.NewWrap(val, ty, false, "instructions.go", "137", false)
	} else {
		wval = rules.NewWrap(val, ty, true, "instructions.go", "139", false)
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
		r = b.ExecuteCallstack()
		r1 := GenerateCallstack(b, []string{callee})
		r = append(r, r1...)
	} else {
		r = b.ExecuteCallstack()
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
	newState := state.Ident()
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

	base2 := currentFunction[1 : len(currentFunction)-7]

	// if complex {
	// 	declareVar(base2, "Bool", "false")
	// }
	r2 := b.createRule(base2, "false", "Bool", "=")
	return []rules.Rule{r1, r2}
}

func (b *LLBlock) parseTermCon(term *ir.TermCondBr) []rules.Rule {
	var ru []rules.Rule
	return ru
	// var cond rules.Rule
	// var phis map[string]int16

	// g.InPhiState().In()
	// id := term.Cond.Ident()
	// if g.variables.IsTemp(id) {
	// 	refname := fmt.Sprintf("%s-%s", g.CurrentFunction(), id)
	// 	if v, ok := g.variables.Ref[refname]; ok {
	// 		cond = v
	// 	}
	// } else if g.variables.IsBoolean(id) ||
	// 	g.variables.IsNumeric(id) {
	// 	cond = &rules.Wrap{Value: id} //Add Variable and Type and Init
	// }
	// g.InPhiState().Out()

	// g.variables.InitPhis()

	// t, f, a := g.parseTerms(term.Succs())
	// if len(t) == 0 && len(f) == 0 { // This happens in a construction like func{stay();}
	// 	g.variables.PopPhis() // in state charts since we convert them to if state{ stay(); }
	// 	g.variables.AppendState(phis)

	// 	if a != nil {
	// 		after := g.parseAfterBlock(a)
	// 		ru = append(ru, after...)
	// 	}
	// 	return ru
	// }

	// choiceId := uuid.NewString()
	// branchT := fmt.Sprintf("%s-%s", choiceId, "true")
	// branchF := fmt.Sprintf("%s-%s", choiceId, "false")
	// g.Forks.Choices[choiceId] = []string{branchT, branchF}
	// t = rules.TagRules(t, branchT, choiceId)
	// f = rules.TagRules(f, branchF, choiceId)

	// g.buildForkChoice(t, choiceId, branchT)
	// g.buildForkChoice(f, choiceId, branchF)

	// if !g.isBranchClosed(t, f) {
	// 	var tEnds, fEnds []rules.Rule
	// 	ru = append(ru, t...)
	// 	ru = append(ru, f...)

	// 	g.InPhiState().In() //We need to step back into a Phi state to make sure multiconditionals are handling correctly
	// 	//g.buildForkChoice(t, choiceId, branchT)
	// 	//g.buildForkChoice(f, choiceId, branchF)

	// 	tEnds, phis = g.capCond(choiceId, branchT, make(map[string]int16))
	// 	fEnds, _ = g.capCond(choiceId, branchF, phis)

	// 	// Keep variable names in sync across branches
	// 	syncs := g.capCondSyncRules([]string{branchT, branchF})
	// 	tEnds = append(tEnds, syncs[branchT]...)
	// 	fEnds = append(fEnds, syncs[branchF]...)

	// 	tEnds = rules.TagRules(tEnds, branchT, choiceId)
	// 	fEnds = rules.TagRules(fEnds, branchF, choiceId)

	// 	ru = append(ru, &rules.Ite{Cond: cond, T: tEnds, F: fEnds})
	// 	g.InPhiState().Out()
	// }

	// g.variables.PopPhis()
	// g.variables.AppendState(phis)

	// if a != nil {
	// 	after := g.parseAfterBlock(a)
	// 	ru = append(ru, after...)
	// }

	// return ru
}

func (b *LLBlock) parsePhi(inst *ir.InstPhi) []rules.Rule {
	return []rules.Rule{}
}

func (b *LLBlock) parseGetElementPtr(inst *ir.InstGetElementPtr) []rules.Rule {
	return []rules.Rule{}
}

func (b *LLBlock) parseXor(inst *ir.InstXor) []rules.Rule {
	id := inst.Ident()
	x := inst.X.Ident()
	xRule := b.LookupCondPart(b.Env.CurrentFunction, x)
	if xRule == nil {
		x = b.ConvertIdent(b.Env.CurrentFunction, x)
		xRule = rules.NewWrap(x, "Bool", true, "instructions.go", "388", false)
	}
	return []rules.Rule{b.createMultiCondRule(id, xRule, rules.NewWrap("", "", false, "instructions", "390", false), "not")}
}

func (b *LLBlock) createMultiCondRule(id string, x rules.Rule, y rules.Rule, op string) rules.Rule {
	if op == "not" {
		refname := fmt.Sprintf("%s-%s", b.Env.CurrentFunction, id)
		b.irRefs[refname] = &rules.Prefix{X: x, Ty: "Bool", Op: op}
		return b.irRefs[refname]
	}

	refname := fmt.Sprintf("%s-%s", b.Env.CurrentFunction, id)
	b.irRefs[refname] = &rules.Infix{X: x, Ty: "Bool", Y: y, Op: op}
	return b.irRefs[refname]
}

func (b *LLBlock) parseAnd(inst *ir.InstAnd) []rules.Rule {
	return []rules.Rule{}
}

func (b *LLBlock) parseOr(inst *ir.InstOr) []rules.Rule {
	return []rules.Rule{}
}

func (b *LLBlock) parseBitCast(inst *ir.InstBitCast) []rules.Rule {
	return []rules.Rule{}
}

func (b *LLBlock) parseFNeg(inst *ir.InstFNeg) []rules.Rule {
	return []rules.Rule{}
}

func (b *LLBlock) createCompareRule(op string) (string, rules.Rule) {
	var y *rules.Wrap
	op = b.compareRuleOp(op)
	switch op {
	case "false":
		y = rules.NewWrap("False", "Bool", false, "instructions.go", "426", false)
	case "true":
		y = rules.NewWrap("True", "Bool", false, "instructions.go", "428", false)
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
