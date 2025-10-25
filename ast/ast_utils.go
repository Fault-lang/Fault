package ast

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

func GenerateToken(token string, literal string, start antlr.Token, stop antlr.Token) Token {
	if start == nil || stop == nil {
		return Token{
			Type:    TokenType(token),
			Literal: literal,
			Position: []int{0,
				0,
				0,
				0,
			},
		}
	}

	return Token{
		Type:    TokenType(token),
		Literal: literal,
		Position: []int{start.GetLine(),
			start.GetColumn(),
			stop.GetLine(),
			stop.GetColumn(),
		},
	}
}

func Preparse(pairs map[*Identifier]Expression) map[string]Node {
	properties := make(map[string]Node)
	for k, v := range pairs {
		id := strings.TrimSpace(k.String())
		properties[id] = v
	}
	return properties
}

func MergeNodeMaps(m1 map[string]Node, m2 map[string]Node) map[string]Node {
	for k, v := range m2 {
		m1[k] = v
	}
	return m1
}

func Keys(m map[string]Node) []string {
	var ret []string
	for k := range m {
		ret = append(ret, k)
	}
	return ret
}

func ExtractBranches(b map[string]*StructProperty) map[string]Node {
	ret := make(map[string]Node)
	for k, v := range b {
		ret[k] = v.Value
	}
	return ret
}

func WrapBranches(b map[string]Node) map[string]*StructProperty {
	ret := make(map[string]*StructProperty)
	for k, v := range b {
		rawid := v.(Nameable).RawId()
		ret[k] = &StructProperty{Value: v}
		ret[k].ProcessedName = rawid
		ret[k].SetType(&Type{Type: v.Type()})
		ret[k].Spec = rawid[0]
		ret[k].Name = k
	}
	return ret
}

func evalFloat(f1 float64, f2 float64, op string) float64 {
	switch op {
	case "+":
		return f1 + f2
	case "-":
		return f1 - f2
	case "*":
		return f1 * f2
	case "/":
		return f1 / f2
	default:
		panic(fmt.Sprintf("unsupported operator %s", op))
	}
}

func evalInt(i1 int64, i2 int64, op string) int64 {
	switch op {
	case "+":
		return i1 + i2
	case "-":
		return i1 - i2
	case "*":
		return i1 * i2
	default:
		panic(fmt.Sprintf("unsupported operator %s", op))
	}
}

func IsCompare(op string) bool {
	switch op {
	case ">":
		return true
	case "<":
		return true
	case "==":
		return true
	case "!=":
		return true
	case "<=":
		return true
	case ">=":
		return true
	case "&&":
		return true
	case "||":
		return true
	case "!":
		return true
	default:
		return false
	}
}

func Evaluate(n *InfixExpression) Expression {
	if IsCompare(n.Operator) {
		return n
	}
	f1, ok1 := n.Left.(*FloatLiteral)
	i1, ok2 := n.Left.(*IntegerLiteral)

	if !ok1 && !ok2 {
		return n
	}

	f2, ok1 := n.Right.(*FloatLiteral)
	i2, ok2 := n.Right.(*IntegerLiteral)

	if !ok1 && !ok2 {
		return n
	}

	if f1 != nil {
		if f2 != nil {
			v := evalFloat(f1.Value, f2.Value, n.Operator)
			return &FloatLiteral{
				Token: n.Token,
				Value: v,
			}
		} else {
			v := evalFloat(f1.Value, float64(i2.Value), n.Operator)
			return &FloatLiteral{
				Token: n.Token,
				Value: v,
			}
		}
	} else {
		if f2 != nil {
			v := evalFloat(float64(i1.Value), f2.Value, n.Operator)
			return &FloatLiteral{
				Token: n.Token,
				Value: v,
			}
		} else {
			if n.Operator == "/" {
				//Return a float in the case of division
				v := evalFloat(float64(i1.Value), float64(i2.Value), n.Operator)
				return &FloatLiteral{
					Token: n.Token,
					Value: v,
				}
			}
			v := evalInt(i1.Value, i2.Value, n.Operator)
			return &IntegerLiteral{
				Token: n.Token,
				Value: v,
			}
		}
	}
}
