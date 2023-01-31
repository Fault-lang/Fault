package smt

import (
	"fault/ast"
	"fault/util"
	"fmt"
)

func (g *Generator) parseAssert(a *ast.AssertionStatement) []string {
	stateRange := a.Constraint.Operator == "then"
	left := g.parseInvariantNode(a.Constraint.Left, stateRange)
	right := g.parseInvariantNode(a.Constraint.Right, stateRange)
	
	if stateRange && a.Assume{

	}

	if stateRange {
		
	}

	return []string{}
}

func (g *Generator) parseInvariantNode(exp ast.Expression, stateRange bool) *stateGroup {
	switch e := exp.(type) {
	case *ast.InfixExpression:
		operator := smtlibOperators(e.Operator)
		left := g.parseInvariantNode(e.Left, stateRange)
		right := g.parseInvariantNode(e.Right, stateRange)

		return g.mergeInvariantInfix(left, right, operator)

		// i := &invariant{
		// 	left:     left,
		// 	operator: smtlibOperators(e.Operator),
		// 	right:    right,
		// }
		// if e.Operator == "!=" { //Not valid in SMTLib
		// 	return &invariant{operator: "not",
		// 		right: i}
		// }
		// return i

	case *ast.AssertVar:
		if len(e.Instances) == 1 {
			st, _, c := captureState(e.Instances[0])
			vr := g.varRounds(e.Instances[0], st)
			return &stateGroup{wraps: []*states{{
				base:     e.Instances[0],
				states:   vr,
				constant: c,
			}}}

		}
		var wg = &stateGroup{}
		for _, v := range e.Instances {
			st, _, c := captureState(v)
			vr := g.varRounds(v, st)
			wg.wraps = append(wg.wraps, &states{base: v,
				states:   vr,
				constant: c,
			})
		}
		return wg
	case *ast.IntegerLiteral:
		s := make(map[int][]string)
		s[0] = []string{fmt.Sprint(e.Value)}
		return &stateGroup{wraps: []*states{{
			base:     "",
			states:   s,
			constant: true,
		}}}
	case *ast.FloatLiteral:
		s := make(map[int][]string)
		s[0] = []string{fmt.Sprint(e.Value)}
		return &stateGroup{wraps: []*states{{
			base:     "",
			states:   s,
			constant: true,
		}}}
	case *ast.Boolean:
		s := make(map[int][]string)
		s[0] = []string{fmt.Sprint(e.Value)}
		return &stateGroup{wraps: []*states{{
			base:     "",
			states:   s,
			constant: true,
		}}}
	case *ast.StringLiteral:
		s := make(map[int][]string)
		s[0] = []string{fmt.Sprint(e.Value)}
		return &stateGroup{wraps: []*states{{
			base:     "",
			states:   s,
			constant: true,
		}}}
	case *ast.PrefixExpression:
		var operator string
		right := g.parseInvariant(e.Right)
		if e.Operator == "!" { //Not valid in SMTLib
			operator = "not"
		} else {
			operator = smtlibOperators(e.Operator)
		}

		if r, ok := right.(*states); ok {
			return g.mergeInvariantPrefix([]*states{r}, operator)
		} else {
			return g.mergeInvariantPrefix(right.(*stateGroup).wraps, operator)
		}

	case *ast.Nil:
	case *ast.IndexExpression:
		vr := g.varRounds(e.IdString(), e.Index.String())
		return &stateGroup{wraps: []*states{{
			base:     e.IdString(),
			states:   vr,
			constant: true,
		}}}
		// return &wrap{value: g.convertIndexExpr(e),
		// 	state:    "",
		// 	all:      false,
		// 	constant: true,
		// }
	default:
		pos := e.Position()
		panic(fmt.Sprintf("illegal node %T in assert or assume line: %d, col: %d", e, pos[0], pos[1]))
	}
	return nil
}

func (g *Generator) mergeInvariantPrefix(right []*states, operator string) *stateGroup {
	var ret []*states
	for _, r := range right {
		states := make(map[int][]string)
		for i, s := range r.states {
			states[i] = append(states[i], fmt.Sprintf("(%s %s)", operator, s))
		}
		r.states = states
		ret = append(ret, r)
	}
	return &stateGroup{
		wraps: ret,
	}
}

func (g *Generator) mergeInvariantInfix(left *stateGroup, right *stateGroup, operator string) *stateGroup {
	var states []*states
	for _, l := range left.wraps {
		for _, r := range right.wraps {
			state := g.mergeByRound(l, r, operator)
			states = append(states, state)
		}
	}

	return &stateGroup{wraps: states}

}

func (g *Generator) mergeByRound(left *states, right *states, operator string) *states {
	ret := &states{}

	st := make(map[int][]string)
	if left.constant && right.constant {
		combos := util.Combinations([][]string{left.states[0], right.states[0]}, 2)
		st[0] = packageStateGraph(combos[0], operator)
		ret.states = st
		return ret
	}

	if left.constant {
		st := g.balance(right, left, operator)
		ret.states = st
		return ret
	}

	if right.constant {
		st := g.balance(left, right, operator)
		ret.states = st
		return ret
	}

	for i := 0; i < g.Rounds; i++ {
		var l, llast, r, rlast []string
		var ok bool
		if l, ok = left.states[i]; !ok {
			if llast == nil {
				l = []string{fmt.Sprintf("%s_%s", left.base, "0")}
			} else {
				l = llast
			}
		}

		if r, ok = right.states[i]; !ok {
			if rlast == nil {
				r = []string{fmt.Sprintf("%s_%s", right.base, "0")}
			} else {
				r = rlast
			}
		}

		combos := util.PairCombinations(l, r)
		st[i] = packageStateGraph(combos, operator)
	}
	ret.states = st
	return ret
}

func (g *Generator) balance(vr *states, con *states, operator string) map[int][]string {
	ret := make(map[int][]string)
	for i, v := range vr.states {
		combos := util.PairCombinations(v, con.states[0])
		ret[i] = packageStateGraph(combos, operator)
	}
	return ret
}
