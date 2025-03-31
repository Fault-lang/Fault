package smt

// import (
// 	"fault/smt/variables"
// 	"fault/util"
// 	"strings"

// 	"github.com/llir/llvm/ir"
// 	"github.com/llir/llvm/ir/constant"
// )

// func (g *Generator) newConstants(globals []*ir.Global) []string {
// 	// Constants cannot be changed and therefore don't increment
// 	// in SSA. So instead of return a *rule we can skip directly
// 	// to a set of strings
// 	r := []string{}
// 	for _, gl := range globals {
// 		id := util.FormatIdent(gl.GlobalIdent.Ident())
// 		if !variables.IsIndexed(id) && !variables.IsClocked(id) {
// 			r = append(r, g.constantRule(id, gl.Init))
// 		}
// 	}
// 	return r
// }

// func (g *Generator) constantRule(id string, c constant.Constant) string {
// 	if id == "__rounds" || id == "__parallelGroup" {
// 		return ""
// 	}

// 	switch val := c.(type) {
// 	case *constant.Int:
// 		ty := g.variables.LookupType(id, val)
// 		g.addVarToRound(id, 0)
// 		id = g.variables.AdvanceSSA(id)
// 		g.declareVar(id, ty)
// 	case *constant.ExprAnd, *constant.ExprOr, *constant.ExprFNeg:
// 		ty := g.variables.LookupType(id, val)
// 		g.addVarToRound(id, 0)
// 		id = g.variables.AdvanceSSA(id)
// 		rule := g.parseConstExpr(val)
// 		v := rule.Assertless()
// 		return g.writeInitRule(id, ty, v)
// 	default:
// 		ty := g.variables.LookupType(id, val)
// 		g.addVarToRound(id, 0)
// 		id = g.variables.AdvanceSSA(id)
// 		g.declareVar(id, ty)
// 	case *constant.Float:
// 		ty := g.variables.LookupType(id, val)
// 		g.addVarToRound(id, 0)
// 		id = g.variables.AdvanceSSA(id)
// 		if g.isASolvable(id) {
// 			g.declareVar(id, ty)
// 		} else {
// 			v := val.X.String()
// 			if strings.Contains(v, ".") {
// 				return g.writeInitRule(id, ty, v)
// 			}
// 			return g.writeInitRule(id, ty, v+".0")
// 		}
// 	}
// 	return ""
// }
