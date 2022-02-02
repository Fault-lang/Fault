package ast

import (
	"fmt"
	"testing"
)

func InitNodes() []Node {
	token := Token{Literal: "test", Position: []int{1, 2, 3, 4}}
	pairs := make(map[Expression]Expression)
	pairs[&Identifier{Token: token, Value: "foo"}] = &IntegerLiteral{Token: token, Value: 3}

	return []Node{
		&Spec{Statements: []Statement{&SpecDeclStatement{Token: token, Name: &Identifier{Token: token, Value: "foo"}}}},
		&SpecDeclStatement{Token: token, Name: &Identifier{Token: token, Value: "foo"}},
		&ImportStatement{Token: token, Name: &Identifier{Token: token, Value: "bar"}, Path: &StringLiteral{Token: token, Value: "foo/bar/baz"}},
		&ConstantStatement{Token: token, Name: &Identifier{Token: token, Value: "fuzz"}, Value: &IntegerLiteral{Token: token, Value: 24}},
		&DefStatement{Token: token, Name: &Identifier{Token: token, Value: "buzz"}, Value: &IntegerLiteral{Token: token, Value: 3}},
		&AssertionStatement{Token: token, Constraints: &Invariant{Token: token, Variable: &IntegerLiteral{Token: token, Value: 3}, Comparison: "==", Expression: &IntegerLiteral{Token: token, Value: 3}}},
		&AssumptionStatement{Token: token, Constraints: &Invariant{Token: token, Variable: &IntegerLiteral{Token: token, Value: 3}, Comparison: "==", Expression: &IntegerLiteral{Token: token, Value: 3}}},
		&Invariant{Token: token, Variable: &IntegerLiteral{Token: token, Value: 3}, Comparison: "==", Expression: &IntegerLiteral{Token: token, Value: 3}},
		&ForStatement{Token: token, Rounds: &IntegerLiteral{Token: token, Value: 5}, Body: &BlockStatement{}},
		&Identifier{Token: token, Value: "foo"},
		&ParameterCall{Token: token, Value: []string{"foo", "bar"}},
		&AssertVar{Token: token, Spec: "test", Instances: []string{"foo", "bar"}},
		&Instance{Token: token, Value: &Identifier{Token: token, Value: "foo"}, Name: "test"},
		&IntegerLiteral{Token: token, Value: 3},
		&FloatLiteral{Token: token, Value: 3.2},
		&Natural{Token: token, Value: 3},
		&Uncertain{Token: token, Mean: 2.0, Sigma: .4},
		&PrefixExpression{Token: token, Operator: "!", Right: &IntegerLiteral{Token: token, Value: 3}},
		&InfixExpression{Token: token, Left: &IntegerLiteral{Token: token, Value: 3}, Right: &IntegerLiteral{Token: token, Value: 3}, Operator: ">"},
		&Boolean{Token: token, Value: true},
		&This{Token: token, Value: []string{"foo", "bar"}},
		&Clock{Token: token, Value: "foo"},
		&Nil{Token: token},
		&BlockStatement{Token: token},
		&ParallelFunctions{Token: token, Expressions: []Expression{&Boolean{Token: token, Value: true}, &Boolean{Token: token, Value: true}}},
		&InitExpression{Token: token, Expression: &Boolean{Token: token, Value: true}},
		&IfExpression{Token: token, Condition: &Boolean{Token: token, Value: true}, Consequence: &BlockStatement{
			Token:      token,
			Statements: []Statement{&ConstantStatement{Token: token, Name: &Identifier{Token: token, Value: "fuzz"}, Value: &IntegerLiteral{Token: token, Value: 24}}},
		},
			Alternative: &BlockStatement{Token: token,
				Statements: []Statement{&ConstantStatement{Token: token, Name: &Identifier{Token: token, Value: "buzz"}, Value: &IntegerLiteral{Token: token, Value: 20}}},
			}, Elif: &IfExpression{Token: token, Condition: &Boolean{Token: token, Value: false}, Consequence: &BlockStatement{}}},
		&FunctionLiteral{Token: token, Parameters: []*Identifier{{Token: token, Value: "foo"}}, Body: &BlockStatement{}},
		&StringLiteral{Token: token, Value: "test"},
		&IndexExpression{Token: token, Left: &Identifier{Token: token, Value: "foo"}, Index: &IntegerLiteral{Token: token, Value: 3}},
		&StockLiteral{Token: token, Pairs: pairs},
		&FlowLiteral{Token: token, Pairs: pairs},
		&Unknown{Token: token, Name: &Identifier{Token: token, Value: "foo"}}}
}

func TestTokenLiteral(t *testing.T) {
	nodes := InitNodes()
	for _, n := range nodes {
		if n.TokenLiteral() != "test" {
			t.Fatalf("TokenLiteral failed for node type %T. got=%s", n, n.TokenLiteral())
		}
	}
}

func TestString(t *testing.T) {
	var got, want string
	nodes := InitNodes()
	for _, n := range nodes {
		switch t := n.(type) {
		case *Spec:
			got = t.String()
			want = "test foo;"
		case *SpecDeclStatement:
			got = t.String()
			want = "test foo;"
		case *ImportStatement:
			got = t.String()
			want = "test bar = foo/bar/baz;"
		case *ConstantStatement:
			got = t.String()
			want = "test fuzz = 24;"
		case *DefStatement:
			got = t.String()
			want = "test buzz = 3;"
		case *AssertionStatement:
			got = t.String()
			want = "test 3==3;"
		case *AssumptionStatement:
			got = t.String()
			want = "test 3==3;"
		case *Invariant:
			got = t.String()
			want = "testassert 33;"
		case *ForStatement:
			got = t.String()
			want = "test 5;"
		case *Identifier:
			got = t.String()
			want = "foo"
		case *ParameterCall:
			got = t.String()
			want = "foobar"
		case *AssertVar:
			got = t.String()
			want = "foo bar"
		case *Instance:
			got = t.String()
			want = "test= new foo"
		case *IntegerLiteral:
			got = t.String()
			want = "3"
		case *FloatLiteral:
			got = t.String()
			want = "3.2"
		case *Natural:
			got = t.String()
			want = "3"
		case *Uncertain:
			got = t.String()
			want = "Mean: 2.000000Sigma: 0.400000;"
		case *PrefixExpression:
			got = t.String()
			want = "(!3)"
		case *InfixExpression:
			got = t.String()
			want = "(3 > 3)"
		case *Boolean:
			got = t.String()
			want = "test"
		case *This:
			got = t.String()
			want = "test"
		case *Clock:
			got = t.String()
			want = "test"
		case *Nil:
			got = t.String()
			want = "test"
		case *BlockStatement:
			got = t.String()
			want = ""
		case *ParallelFunctions:
			got = t.String()
			want = "testtest"
		case *InitExpression:
			got = t.String()
			want = "init test"
		case *IfExpression:
			got = t.String()
			want = "iftest test fuzz = 24;else ififtest else test buzz = 20;"
		case *FunctionLiteral:
			got = t.String()
			want = "test(foo) "
		case *StringLiteral:
			got = t.String()
			want = "test"
		case *IndexExpression:
			got = t.String()
			want = "(foo[3])"
		case *StockLiteral:
			got = t.String()
			want = "{foo:3}"
		case *FlowLiteral:
			got = t.String()
			want = "{foo:3}"
		case *Unknown:
			got = t.String()
			want = "unknown(foo)"
		}
		if got != want {
			t.Fatalf("String failed for node type %T. got=%s", n, got)
		}
	}
}

func TestPosition(t *testing.T) {
	nodes := InitNodes()
	for _, n := range nodes {
		pos := n.Position()
		if pos[0] != 1 || pos[1] != 2 || pos[2] != 3 || pos[3] != 4 {
			t.Fatalf("Position failed for node type %T. got=%s", n, fmt.Sprint(n.Position()))
		}
	}
}
