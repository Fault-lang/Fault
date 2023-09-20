package ast

import (
	"fault/parser"
	"fmt"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func InitNodes() []Node {
	token := Token{Literal: "test", Position: []int{1, 2, 3, 4}}
	tokenAlt1 := Token{Literal: "test2", Position: []int{1, 2, 3, 4}}
	tokenAlt2 := Token{Literal: "test3", Position: []int{1, 2, 3, 4}}
	pairs := make(map[*Identifier]Expression)
	pairs[&Identifier{Token: token, Value: "foo"}] = &IntegerLiteral{Token: token, Value: 3}
	pairs[&Identifier{Token: token, Value: "bar"}] = &IntegerLiteral{Token: token, Value: 5}
	pairs[&Identifier{Token: token, Value: "bash"}] = &IntegerLiteral{Token: token, Value: -4}
	pairOrder := []string{"foo", "bar", "bash"}
	properties := make(map[string]*StructProperty)
	properties["foo"] = &StructProperty{Token: token, Spec: "test", Name: "foo", Value: &IntegerLiteral{Token: token, Value: 3}}
	params := make(map[string]Operand)
	params["zoo"] = &ParameterCall{Token: token, Value: []string{"foo", "bar"}}

	baseType := &Type{Type: "test"}
	stringType := &Type{Type: "STRING"}
	//boolType := &Type{Type: "BOOL"}
	//floatType := &Type{Type: "FLOAT"}
	intType := &Type{Type: "INT"}

	return []Node{
		&Spec{Statements: []Statement{&SpecDeclStatement{Token: token, Name: &Identifier{InferredType: baseType, Token: token, Value: "foo"}}}},
		&SpecDeclStatement{Token: token, Name: &Identifier{InferredType: baseType, Token: token, Value: "foo"}},
		&SysDeclStatement{Token: token, Name: &Identifier{InferredType: baseType, Token: token, Value: "foo"}},
		&ImportStatement{Token: token, Name: &Identifier{InferredType: baseType, Token: token, Value: "bar"}, Path: &StringLiteral{InferredType: stringType, Token: token, Value: "foo/bar/baz"}},
		&ConstantStatement{Token: token, Name: &Identifier{InferredType: intType, Token: token, Value: "fuzz"}, Value: &IntegerLiteral{InferredType: intType, Token: token, Value: 24}},
		&DefStatement{Token: token, Name: &Identifier{InferredType: intType, Token: token, Value: "buzz"}, Value: &IntegerLiteral{InferredType: intType, Token: token, Value: 3}},
		&AssertionStatement{Token: token, Constraint: &InvariantClause{Token: token, Operator: "==", Left: &IntegerLiteral{Token: token, Value: 3}, Right: &IntegerLiteral{Token: token, Value: 3}}, Assume: false},
		&AssertionStatement{Token: token, Constraint: &InvariantClause{Token: token, Operator: "==", Left: &IntegerLiteral{Token: token, Value: 3}, Right: &IntegerLiteral{Token: token, Value: 3}}, Assume: true},
		//&Invariant{Token: token, Variable: &IntegerLiteral{Token: token, Value: 3}, Comparison: "==", Expression: &IntegerLiteral{Token: token, Value: 3}},
		//&Invariant{Token: token, Variable: &IntegerLiteral{Token: token, Value: 3}, Conjuction: "==", Expression: &IntegerLiteral{Token: token, Value: 3}},
		&InvariantClause{Token: token, Operator: "==", Left: &IntegerLiteral{Token: token, Value: 3}, Right: &IntegerLiteral{Token: token, Value: 3}},
		&ForStatement{Token: token, Rounds: &IntegerLiteral{Token: token, Value: 5}, Body: &BlockStatement{Statements: []Statement{&ConstantStatement{Token: token, Name: &Identifier{InferredType: intType, Token: token, Value: "fuzz"}, Value: &IntegerLiteral{InferredType: intType, Token: token, Value: 24}}}}},
		&ExpressionStatement{Token: token, Expression: &PrefixExpression{Token: token, Operator: "!", Right: &IntegerLiteral{Token: token, Value: 3}}},
		&Identifier{InferredType: baseType, Token: token, Value: "foo"},
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
		&BlockStatement{Token: token, Statements: []Statement{&ConstantStatement{Token: token, Name: &Identifier{InferredType: intType, Token: token, Value: "fuzz"}, Value: &IntegerLiteral{InferredType: intType, Token: token, Value: 24}}}},
		&ParallelFunctions{Token: token, Expressions: []Expression{&Boolean{Token: token, Value: true}, &Boolean{Token: token, Value: true}}},
		&InitExpression{Token: token, Expression: &Boolean{Token: token, Value: true}},
		&IfExpression{Token: token, Condition: &Boolean{Token: token, Value: true}, Consequence: &BlockStatement{
			Token:      token,
			Statements: []Statement{&ConstantStatement{Token: token, Name: &Identifier{Token: token, Value: "fuzz"}, Value: &IntegerLiteral{Token: token, Value: 24}}},
		},
			Elif: &IfExpression{Token: token, Condition: &Boolean{Token: token, Value: false}, Consequence: &BlockStatement{
				Token:      token,
				Statements: []Statement{&ConstantStatement{Token: token, Name: &Identifier{Token: token, Value: "fuzz"}, Value: &IntegerLiteral{Token: token, Value: 100}}},
			}}},
		&IfExpression{Token: tokenAlt1, Condition: &Boolean{Token: token, Value: true}, Consequence: &BlockStatement{
			Token:      token,
			Statements: []Statement{&ConstantStatement{Token: token, Name: &Identifier{Token: token, Value: "fuzz"}, Value: &IntegerLiteral{Token: token, Value: 24}}},
		}},
		&IfExpression{Token: tokenAlt2, Condition: &Boolean{Token: token, Value: true}, Consequence: &BlockStatement{
			Token:      token,
			Statements: []Statement{&ConstantStatement{Token: token, Name: &Identifier{Token: token, Value: "fuzz"}, Value: &IntegerLiteral{Token: token, Value: 24}}},
		}, Alternative: &BlockStatement{Token: token,
			Statements: []Statement{&ConstantStatement{Token: token, Name: &Identifier{Token: token, Value: "buzz"}, Value: &IntegerLiteral{Token: token, Value: 20}}},
		}},
		&FunctionLiteral{Token: token, Parameters: []*Identifier{{Token: token, Value: "foo"}}, Body: &BlockStatement{Statements: []Statement{&ConstantStatement{Token: token, Name: &Identifier{InferredType: intType, Token: token, Value: "fuzz"}, Value: &IntegerLiteral{InferredType: intType, Token: token, Value: 24}}}}},
		&StringLiteral{Token: token, Value: "test"},
		&IndexExpression{Token: token, Left: &Identifier{Token: token, Value: "foo"}, Index: &IntegerLiteral{Token: token, Value: 3}},
		&StockLiteral{Token: token, Pairs: pairs, Order: pairOrder},
		&FlowLiteral{Token: token, Pairs: pairs, Order: pairOrder},
		&ComponentLiteral{Token: token, Pairs: pairs, Order: pairOrder},
		&Unknown{Token: token, Name: &Identifier{Token: token, Value: "foo"}},
		&StructInstance{Token: token, Properties: properties},
		&BuiltIn{Token: token, Parameters: params, Function: "advance"},
		&StartStatement{Token: token, Pairs: [][]string{{"foo", "bar"}, {"hello", "world"}}},
	}
}

func TestTokenLiteral(t *testing.T) {
	nodes := InitNodes()
	for _, n := range nodes {
		if n.TokenLiteral() != "test" && n.TokenLiteral() != "test2" && n.TokenLiteral() != "test3" {
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
		case *SysDeclStatement:
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
			if t.Assume {
				want = "test assume 3==3;"
			} else {
				want = "test assert 3==3;"
			}
		// case *Invariant:
		// 	got = t.String()
		// 	want = "test assert 3==3;"
		case *InvariantClause:
			got = t.String()
			want = "testassert 3==3;"
		case *ForStatement:
			got = t.String()
			want = "test 5test fuzz = 24;;"
		case *Identifier:
			got = t.String()
			want = "foo"
		case *ParameterCall:
			got = t.String()
			want = "foo.bar"
		case *AssertVar:
			got = t.String()
			want = "(foo or bar)"
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
			want = "!3"
		case *InfixExpression:
			got = t.String()
			want = "(3 > 3)"
		case *Boolean:
			got = t.String()
			want = "test"
		case *This:
			got = t.String()
			want = "foo.bar"
		case *Clock:
			got = t.String()
			want = "now"
		case *Nil:
			got = t.String()
			want = "test"
		case *BlockStatement:
			got = t.String()
			want = "test fuzz = 24;"
		case *ExpressionStatement:
			got = t.String()
			want = "!3"
		case *ParallelFunctions:
			got = t.String()
			want = "testtest"
		case *InitExpression:
			got = t.String()
			want = "init test"
		case *IfExpression:
			got = t.String()
			if t.TokenLiteral() == "test2" {
				want = "if(test){test fuzz = 24;}"
			} else if t.TokenLiteral() == "test3" {
				want = "if(test){test fuzz = 24;}else{test buzz = 20;}"
			} else {
				want = "if(test){test fuzz = 24;}else if(test){test fuzz = 100;}"
			}
		case *FunctionLiteral:
			got = t.String()
			want = "test(foo) test fuzz = 24;"
		case *StringLiteral:
			got = t.String()
			want = "test"
		case *IndexExpression:
			got = t.String()
			want = "(foo[3])"
		case *StockLiteral:
			got = t.String()
			want = "{foo:3, bar:5, bash:-4}"
		case *FlowLiteral:
			got = t.String()
			want = "{foo:3, bar:5, bash:-4}"
		case *ComponentLiteral:
			got = t.String()
			want = "{foo:3, bar:5, bash:-4}"
		case *Unknown:
			got = t.String()
			want = "unknown(foo)"
		case *StructInstance:
			got = t.String()
			want = "__foo:3"
		case *BuiltIn:
			got = t.String()
			want = "advance(foo.bar)"
		case *StartStatement:
			got = t.String()
			want = "test {foo : bar, hello : world};"
		}
		if got != want {
			t.Fatalf("String failed for node type %T. got=%s", n, got)
		}
	}
}

func TestTypes(t *testing.T) {
	var got, want string
	nodes := InitNodes()
	for _, n := range nodes {
		switch t := n.(type) {
		case *Spec:
			got = t.Type()
			want = "SPEC"
		case *SpecDeclStatement:
			got = t.Type()
			want = "SPEC"
		case *SysDeclStatement:
			got = t.Type()
			want = "SYSTEM"
		case *ImportStatement:
			got = t.Type()
			want = "IMPORT"
		case *ConstantStatement:
			got = t.Type()
			want = "INT"
		case *DefStatement:
			got = t.Type()
			want = "test"
		case *AssertionStatement:
			got = t.Type()
			want = "INT"
		// case *Invariant:
		// 	got = t.Type()
		// 	want = "INT"
		case *InvariantClause:
			got = t.Type()
			want = "INT"
		case *ForStatement:
			got = t.Type()
			want = ""
		case *Identifier:
			got = t.Type()
			want = "test"
		case *ParameterCall:
			got = t.Type()
			want = ""
		case *AssertVar:
			got = t.Type()
			want = ""
		case *Instance:
			got = t.Type()
			want = "test"
		case *IntegerLiteral:
			got = t.Type()
			want = "INT"
		case *FloatLiteral:
			got = t.Type()
			want = "FLOAT"
		case *Natural:
			got = t.Type()
			want = "NATURAL"
		case *Uncertain:
			got = t.Type()
			want = "UNCERTAIN"
		case *PrefixExpression:
			got = t.Type()
			want = "INT"
		case *InfixExpression:
			got = t.Type()
			want = "INT"
		case *Boolean:
			got = t.Type()
			want = "BOOL"
		case *This:
			got = t.Type()
			want = ""
		case *Clock:
			got = t.Type()
			want = ""
		case *Nil:
			got = t.Type()
			want = ""
		case *BlockStatement:
			got = t.Type()
			want = ""
		case *ExpressionStatement:
			got = t.Type()
			want = "INT"
		case *ParallelFunctions:
			got = t.Type()
			want = ""
		case *InitExpression:
			got = t.Type()
			want = ""
		case *IfExpression:
			got = t.Type()
			want = ""
		case *FunctionLiteral:
			got = t.Type()
			want = "INT"
		case *StringLiteral:
			got = t.Type()
			want = "STRING"
		case *IndexExpression:
			got = t.Type()
			want = ""
		case *StockLiteral:
			got = t.Type()
			want = "STOCK"
		case *FlowLiteral:
			got = t.Type()
			want = "FLOW"
		case *ComponentLiteral:
			got = t.Type()
			want = "COMPONENT"
		case *Unknown:
			got = t.Type()
			want = "UNKNOWN"
		case *BuiltIn:
			got = t.Type()
			want = "BUILTIN"
		case *StartStatement:
			got = t.Type()
			want = "START"
		}
		if got != want {
			t.Fatalf("Type failed for node type %T. got=%s", n, got)
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

func TestPropertyIdent(t *testing.T) {
	nodes := InitNodes()
	for _, n := range nodes {
		switch s := n.(type) {
		case *StockLiteral:
			id := s.GetPropertyIdent("foo")
			v := s.Pairs[id]
			if i, ok := v.(*IntegerLiteral); !ok || i.Value != 3 {
				t.Fatal("GetPropertyIdent broken on StockLiteral")
			}
		case *FlowLiteral:
			id := s.GetPropertyIdent("foo")
			v := s.Pairs[id]
			if i, ok := v.(*IntegerLiteral); !ok || i.Value != 3 {
				t.Fatal("GetPropertyIdent broken on FlowLiteral")
			}
		case *ComponentLiteral:
			id := s.GetPropertyIdent("foo")
			v := s.Pairs[id]
			if i, ok := v.(*IntegerLiteral); !ok || i.Value != 3 {
				t.Fatal("GetPropertyIdent broken on ComponentLiteral")
			}
		}
	}
}

func TestOrder(t *testing.T) {
	nodes := InitNodes()
	for _, n := range nodes {
		switch s := n.(type) {
		case *StockLiteral:
			keys := s.Order
			if keys[0] != "foo" || keys[1] != "bar" || keys[2] != "bash" {
				t.Fatal("order broken on StockLiteral")
			}
		case *FlowLiteral:
			keys := s.Order
			if keys[0] != "foo" || keys[1] != "bar" || keys[2] != "bash" {
				t.Fatal("order broken on FlowLiteral")
			}
		case *ComponentLiteral:
			keys := s.Order
			if keys[0] != "foo" || keys[1] != "bar" || keys[2] != "bash" {
				t.Fatal("order broken on ComponentLiteral")
			}
		}
	}
}

func TestEval(t *testing.T) {
	tests := []*InfixExpression{{
		Left:  &IntegerLiteral{Value: 2},
		Right: &IntegerLiteral{Value: 2},
	},
		{
			Left:  &FloatLiteral{Value: 2.5},
			Right: &IntegerLiteral{Value: 2},
		},
		{
			Left:     &IntegerLiteral{Value: 2},
			Operator: "+",
			Right:    &FloatLiteral{Value: 2.5},
		}}

	operators := []string{"+", "-", "/", "*"}

	results := []Node{
		&IntegerLiteral{Value: 4},
		&FloatLiteral{Value: 4.5},
		&FloatLiteral{Value: 4.5},
		&IntegerLiteral{Value: 0},
		&FloatLiteral{Value: .5},
		&FloatLiteral{Value: -.5},
		&FloatLiteral{Value: 1},
		&FloatLiteral{Value: 1.25},
		&FloatLiteral{Value: .8},
		&IntegerLiteral{Value: 4},
		&FloatLiteral{Value: 5},
		&FloatLiteral{Value: 5},
	}

	i := 0
	for _, o := range operators {
		for _, n := range tests {
			n.Operator = o
			test := Evaluate(n)
			switch actual := test.(type) {
			case *IntegerLiteral:
				expected, ok := results[i].(*IntegerLiteral)
				if !ok {
					t.Fatalf("expected value a different type from actual expected=%s actual=%s", results[i], test)
				}
				if expected.Value != actual.Value {
					t.Fatalf("expected value a different from actual expected=%s actual=%s", expected, actual)
				}
			case *FloatLiteral:
				expected, ok := results[i].(*FloatLiteral)
				if !ok {
					t.Fatalf("expected value a different type from actual expected=%s actual=%s", results[i], test)
				}
				if expected.Value != actual.Value {
					t.Fatalf("expected value a different from actual expected=%s actual=%s", expected, actual)
				}
			}
			i++
		}
	}
}

func TestEvalFloat(t *testing.T) {
	test1 := evalFloat(2.1, 1.5, "+")
	if test1 != 3.6 {
		t.Fatal("evalFloat failed to eval + correctly")
	}
	test2 := evalFloat(2.5, 1.5, "-")
	if test2 != 1 {
		t.Fatal("evalFloat failed to eval - correctly")
	}
	test3 := evalFloat(2.1, 1.0, "*")
	if test3 != 2.1 {
		t.Fatal("evalFloat failed to eval * correctly")
	}
	test4 := evalFloat(2.0, 2.0, "/")
	if test4 != 1.0 {
		t.Fatal("evalFloat failed to eval / correctly")
	}
}

func TestEvalInt(t *testing.T) {
	test1 := evalInt(2, 1, "+")
	if test1 != 3 {
		t.Fatal("evalInt failed to eval + correctly")
	}
	test2 := evalInt(2, 1, "-")
	if test2 != 1 {
		t.Fatal("evalInt failed to eval - correctly")
	}
	test3 := evalInt(2, 1, "*")
	if test3 != 2 {
		t.Fatal("evalInt failed to eval * correctly")
	}
}

func TestPreparse(t *testing.T) {
	token := Token{Literal: "test", Position: []int{1, 2, 3, 4}}
	pairs := make(map[*Identifier]Expression)
	pairs[&Identifier{Token: token, Value: "foo"}] = &IntegerLiteral{Token: token, Value: 3}
	pairs[&Identifier{Token: token, Value: "bash"}] = &FunctionLiteral{Token: token,
		Parameters: []*Identifier{{Token: token, Value: "foo"}},
		Body: &BlockStatement{Token: token,
			Statements: []Statement{&ConstantStatement{Token: token, Name: &Identifier{Token: token, Value: "buzz"}, Value: &IntegerLiteral{Token: token, Value: 20}}}}}

	ret := Preparse(pairs)

	if len(ret) != 2 {
		t.Fatalf("item removed from map. got=%s", ret)
	}

	for k, v := range ret {
		if k == "foo" {
			ty, ok := v.(*IntegerLiteral)
			if !ok {
				t.Fatalf("pair type incorrect. want=IntegerLiteral got=%T", v)
			}

			if ty.Value != 3 {
				t.Fatalf("pair value incorrect. want=3 got=%d", ty.Value)
			}
		} else if k == "bash" {
			ty, ok := v.(*FunctionLiteral)
			if !ok {
				t.Fatalf("pair type incorrect. want=FunctionLiteral got=%T", v)
			}

			if ty.Body.Statements[0].(*ConstantStatement).Value.(*IntegerLiteral).Value != 20 {
				t.Fatalf("pair value incorrect. want=20 got=%d", ty.Body.Statements[0].(*ConstantStatement).Value.(*IntegerLiteral).Value)
			}
		} else {
			t.Fatalf("pair key unrecognized. got=%s", k)
		}
	}
}

func TestKeys(t *testing.T) {
	test := make(map[string]Node)
	test["here"] = &IntegerLiteral{}
	test["are"] = &IntegerLiteral{}
	test["your"] = &IntegerLiteral{}
	test["keys"] = &IntegerLiteral{}

	results := Keys(test)

	if len(results) != 4 {
		t.Fatalf("incorrect number of keys returned got=%d", len(results))
	}

}

func TestGeneratorToken(t *testing.T) {
	test := `spec test;`
	is := antlr.NewInputStream(test)
	lexer := parser.NewFaultLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	stream.GetAllText()
	tokens := stream.GetAllTokens()

	token := GenerateToken("IMPORT_DECL", "IMPORT_DECL", tokens[0], tokens[0])
	if token.Literal != "IMPORT_DECL" {
		t.Fatalf("token literal not correct. got=%s", token.Literal)
	}
	if token.Position[0] != 1 {
		t.Fatalf("token position not correct. want=1 got=%d", token.Position[0])
	}

	if token.Position[1] != 0 {
		t.Fatalf("token position not correct. want=0 got=%d", token.Position[1])
	}

	if token.Position[2] != 1 {
		t.Fatalf("token position not correct. want=1 got=%d", token.Position[2])
	}

	if token.Position[3] != 0 {
		t.Fatalf("token position not correct. want=0 got=%d", token.Position[3])
	}

}

func TestMergeNodeMaps(t *testing.T) {
	m1 := make(map[string]Node)
	m1["foo"] = &IntegerLiteral{Value: 5}
	m1["bar"] = &IntegerLiteral{Value: 15}

	m2 := make(map[string]Node)
	m2["test"] = &IntegerLiteral{Value: 2}

	m3 := MergeNodeMaps(m1, m2)

	if len(m3) != 3 {
		t.Fatalf("merged map has the wrong length got=%d", len(m3))
	}

	if m3["test"].(*IntegerLiteral).Value != 2 || m3["foo"].(*IntegerLiteral).Value != 5 {
		t.Fatalf("node map not merged correctly")
	}

}

func TestExtractBranches(t *testing.T) {
	test := make(map[string]*StructProperty)
	test["foo"] = &StructProperty{Value: &IntegerLiteral{Value: 5}}
	test["bar"] = &StructProperty{Value: &IntegerLiteral{Value: 2}}

	r := ExtractBranches(test)

	if r["foo"].(*IntegerLiteral).Value != 5 || r["bar"].(*IntegerLiteral).Value != 2 {
		t.Fatal("ExtractBranches returned the wrong result")
	}
}

func TestIsCompare(t *testing.T) {
	if IsCompare("hihi") {
		t.Fatal("first test of IsCompare has failed")
	}

	test := []string{">", "<", "==", "!=", "<=", ">=", "&&", "||", "!"}
	for i, c := range test {
		if !IsCompare(c) {
			t.Fatalf("test %d of %d IsCompare tests has failed", i, len(test))
		}
	}
}
