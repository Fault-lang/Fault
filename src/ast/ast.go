package ast

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type TokenType string

type Token struct {
	Type     TokenType
	Literal  string
	Position []int
}

var OPS = map[string]TokenType{
	"+":  "PLUS",
	"-":  "MINUS",
	"!":  "BANG",
	"^":  "CARET",
	"*":  "MULTI",
	"&":  "AMP",
	"**": "EXPO",
	"/":  "DIV",
	"%":  "PER",
	"<<": "LSHIFT",
	">>": "RSHIFT",
	"&^": "BIT_CLEAR",
	"==": "EQ",
	"!=": "NOT_EQ",
	"<":  "LT",
	"<=": "LTE",
	">":  "GT",
	">=": "GTE",
	"&&": "AND",
	"||": "OR",
	"|":  "PARA",
}

var TYPES = map[string]int{ //Convertible Types
	"STRING":    0, //Not convertible
	"BOOL":      1,
	"NATURAL":   2,
	"FLOAT":     3,
	"INT":       4,
	"UNCERTAIN": 5,
}

type Type struct {
	Type       string
	Scope      int64
	Parameters []Type
}

type Node interface {
	TokenLiteral() string
	String() string
	Position() []int
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Spec struct {
	Statements []Statement
}

func (s *Spec) TokenLiteral() string {
	if len(s.Statements) > 0 {
		return s.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (s *Spec) Position() []int {
	if len(s.Statements) > 0 {
		return s.Statements[0].Position()
	} else {
		return []int{0, 0, 0, 0}
	}
}

func (s *Spec) String() string {
	var out bytes.Buffer

	for _, st := range s.Statements {
		out.WriteString(st.String())
	}
	return out.String()
}

type SpecDeclStatement struct {
	Token Token
	Name  *Identifier
}

func (sd *SpecDeclStatement) statementNode()       {}
func (sd *SpecDeclStatement) TokenLiteral() string { return sd.Token.Literal }
func (sd *SpecDeclStatement) Position() []int      { return sd.Token.Position }
func (sd *SpecDeclStatement) String() string {
	var out bytes.Buffer

	out.WriteString(sd.TokenLiteral() + " ")
	out.WriteString(sd.Name.String())
	out.WriteString(";")
	return out.String()
}

type ImportStatement struct {
	Token Token
	Name  *Identifier
	Path  *StringLiteral
	Tree  *Spec
}

func (is *ImportStatement) statementNode()       {}
func (is *ImportStatement) TokenLiteral() string { return is.Token.Literal }
func (is *ImportStatement) Position() []int      { return is.Token.Position }
func (is *ImportStatement) String() string {
	var out bytes.Buffer

	out.WriteString(is.TokenLiteral() + " ")
	out.WriteString(is.Name.String())
	out.WriteString(" = ")

	out.WriteString(is.Path.String())

	out.WriteString(";")
	return out.String()
}

type ConstantStatement struct {
	Token        Token
	Name         *Identifier
	Value        Expression
	InferredType *Type
}

func (cs *ConstantStatement) statementNode()       {}
func (cs *ConstantStatement) TokenLiteral() string { return cs.Token.Literal }
func (cs *ConstantStatement) Position() []int      { return cs.Token.Position }
func (cs *ConstantStatement) String() string {
	var out bytes.Buffer

	out.WriteString(cs.TokenLiteral() + " ")
	out.WriteString(cs.Name.String())
	out.WriteString(" = ")

	if cs.Value != nil {
		out.WriteString(cs.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

type DefStatement struct {
	Token Token
	Name  *Identifier
	Value Expression
}

func (ds *DefStatement) statementNode()       {}
func (ds *DefStatement) TokenLiteral() string { return ds.Token.Literal }
func (ds *DefStatement) Position() []int      { return ds.Token.Position }
func (ds *DefStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ds.TokenLiteral() + " ")
	out.WriteString(ds.Name.String())
	out.WriteString(" = ")

	if ds.Value != nil {
		out.WriteString(ds.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

type AssertionStatement struct {
	Token       Token
	Constraints *Invariant
}

func (as *AssertionStatement) statementNode()       {}
func (as *AssertionStatement) TokenLiteral() string { return as.Token.Literal }
func (as *AssertionStatement) Position() []int      { return as.Token.Position }
func (as *AssertionStatement) String() string {
	var out bytes.Buffer

	out.WriteString(as.TokenLiteral() + " ")
	out.WriteString(as.Constraints.Variable.String())
	out.WriteString(as.Constraints.Comparison)
	out.WriteString(as.Constraints.Expression.String())
	out.WriteString(";")
	return out.String()
}

type AssumptionStatement struct {
	Token       Token
	Constraints *Invariant
}

func (as *AssumptionStatement) statementNode()       {}
func (as *AssumptionStatement) TokenLiteral() string { return as.Token.Literal }
func (as *AssumptionStatement) Position() []int      { return as.Token.Position }
func (as *AssumptionStatement) String() string {
	var out bytes.Buffer

	out.WriteString(as.TokenLiteral() + " ")
	out.WriteString(as.Constraints.Variable.String())
	out.WriteString(as.Constraints.Comparison)
	out.WriteString(as.Constraints.Expression.String())
	out.WriteString(";")
	return out.String()
}

type Invariant struct {
	Token        Token
	Variable     Expression
	Comparison   string
	Expression   Expression
	Conjuction   string
	InferredType *Type
}

func (i *Invariant) expressionNode()      {}
func (i *Invariant) TokenLiteral() string { return i.Token.Literal }
func (i *Invariant) Position() []int      { return i.Token.Position }
func (i *Invariant) String() string {
	var out bytes.Buffer

	out.WriteString(i.TokenLiteral() + "assert ")
	out.WriteString(i.Variable.String())
	out.WriteString(i.Conjuction)
	out.WriteString(i.Expression.String())

	out.WriteString(";")
	return out.String()
}

type ForStatement struct {
	Token  Token
	Rounds *IntegerLiteral
	Body   *BlockStatement
}

func (fs *ForStatement) statementNode()       {}
func (fs *ForStatement) TokenLiteral() string { return fs.Token.Literal }
func (fs *ForStatement) Position() []int      { return fs.Token.Position }
func (fs *ForStatement) String() string {
	var out bytes.Buffer

	out.WriteString(fs.TokenLiteral() + " ")
	out.WriteString(fs.Rounds.String())
	out.WriteString(fs.Body.String())

	out.WriteString(";")
	return out.String()
}

type Identifier struct {
	Token        Token
	InferredType *Type
	Spec         string
	Value        string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) Position() []int      { return i.Token.Position }
func (i *Identifier) String() string       { return i.Value }

type ParameterCall struct {
	Token        Token
	InferredType *Type
	Value        []string
}

func (p *ParameterCall) expressionNode()      {}
func (p *ParameterCall) TokenLiteral() string { return p.Token.Literal }
func (p *ParameterCall) Position() []int      { return p.Token.Position }
func (p *ParameterCall) String() string {
	var out bytes.Buffer

	for _, s := range p.Value {
		out.WriteString(s)
	}

	return out.String()
}

type AssertVar struct {
	Token        Token
	InferredType *Type
	Spec         string
	Instances    []string
}

func (av *AssertVar) expressionNode()      {}
func (av *AssertVar) TokenLiteral() string { return av.Token.Literal }
func (av *AssertVar) Position() []int      { return av.Token.Position }
func (av *AssertVar) String() string       { return strings.Join(av.Instances, " ") }

type Instance struct {
	Token        Token
	InferredType *Type
	Value        *Identifier
	Name         string
	Complex      bool //If stock does this stock contain another stock?
}

func (i *Instance) expressionNode()      {}
func (i *Instance) TokenLiteral() string { return i.Token.Literal }
func (i *Instance) Position() []int      { return i.Token.Position }
func (i *Instance) String() string {
	var out bytes.Buffer
	out.WriteString(i.Name)
	out.WriteString("= new ")
	out.WriteString(i.Value.String())

	return out.String()
}

type ExpressionStatement struct {
	Token        Token
	InferredType *Type
	Expression   Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) Position() []int      { return es.Token.Position }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

type IntegerLiteral struct {
	Token        Token
	InferredType *Type
	Value        int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return strconv.FormatInt(il.Value, 10) }
func (il *IntegerLiteral) Position() []int      { return il.Token.Position }

type FloatLiteral struct {
	Token        Token
	InferredType *Type
	Value        float64
}

func (fl *FloatLiteral) expressionNode()      {}
func (fl *FloatLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FloatLiteral) Position() []int      { return fl.Token.Position }
func (fl *FloatLiteral) String() string       { return fmt.Sprint(fl.Value) }

type Natural struct {
	Token        Token
	InferredType *Type
	Value        int64
}

func (n *Natural) expressionNode()      {}
func (n *Natural) TokenLiteral() string { return n.Token.Literal }
func (n *Natural) String() string       { return strconv.FormatInt(n.Value, 10) }
func (n *Natural) Position() []int      { return n.Token.Position }

type Uncertain struct {
	Token        Token
	InferredType *Type
	Mean         float64
	Sigma        float64
}

func (u *Uncertain) expressionNode()      {}
func (u *Uncertain) TokenLiteral() string { return u.Token.Literal }
func (u *Uncertain) String() string {
	var out bytes.Buffer
	out.WriteString("Mean: ")
	out.WriteString(strconv.FormatFloat(u.Mean, 'f', 6, 64))
	out.WriteString("Sigma: ")
	out.WriteString(strconv.FormatFloat(u.Sigma, 'f', 6, 64))
	out.WriteString(";")
	return out.String()
}
func (u *Uncertain) Position() []int { return u.Token.Position }

type PrefixExpression struct {
	Token        Token
	InferredType *Type
	Operator     string
	Right        Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) Position() []int      { return pe.Token.Position }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpression struct {
	Token        Token
	InferredType *Type
	Left         Expression
	Operator     string
	Right        Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) Position() []int      { return ie.Token.Position }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

type Boolean struct {
	Token        Token
	InferredType *Type
	Value        bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) Position() []int      { return b.Token.Position }
func (b *Boolean) String() string       { return b.Token.Literal }

type This struct {
	Token        Token
	InferredType *Type
	Value        []string
}

func (t *This) expressionNode()      {}
func (t *This) TokenLiteral() string { return t.Token.Literal }
func (t *This) Position() []int      { return t.Token.Position }
func (t *This) String() string       { return t.Token.Literal }

type Clock struct {
	Token        Token
	InferredType *Type
	Value        string
}

func (c *Clock) expressionNode()      {}
func (c *Clock) TokenLiteral() string { return c.Token.Literal }
func (c *Clock) Position() []int      { return c.Token.Position }
func (c *Clock) String() string       { return c.Token.Literal }

type Nil struct {
	Token        Token
	InferredType *Type
}

func (n *Nil) expressionNode()      {}
func (n *Nil) TokenLiteral() string { return n.Token.Literal }
func (n *Nil) Position() []int      { return n.Token.Position }
func (n *Nil) String() string       { return n.Token.Literal }

type BlockStatement struct {
	Token        Token
	InferredType *Type
	Statements   []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) Position() []int      { return bs.Token.Position }
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type ParallelFunctions struct {
	Token        Token
	InferredType *Type
	Expressions  []Expression
}

func (pf *ParallelFunctions) statementNode()       {}
func (pf *ParallelFunctions) Position() []int      { return pf.Token.Position }
func (pf *ParallelFunctions) TokenLiteral() string { return pf.Token.Literal }
func (pf *ParallelFunctions) String() string {
	var out bytes.Buffer

	for _, s := range pf.Expressions {
		out.WriteString(s.String())
	}

	return out.String()
}

type InitExpression struct {
	Token        Token
	InferredType *Type
	Expression   Expression
}

func (ie *InitExpression) expressionNode()      {}
func (ie *InitExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InitExpression) Position() []int      { return ie.Token.Position }
func (ie *InitExpression) String() string {
	var out bytes.Buffer

	out.WriteString("init ")
	out.WriteString(ie.Expression.String())
	return out.String()
}

type IfExpression struct {
	Token        Token
	InferredType *Type
	Condition    Expression
	Consequence  *BlockStatement
	Alternative  *BlockStatement
	Elif         *IfExpression
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) Position() []int      { return ie.Token.Position }
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Elif != nil {
		out.WriteString("else if")
		out.WriteString(ie.Elif.String())
	}

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

type FunctionLiteral struct {
	Token      Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) Position() []int      { return fl.Token.Position }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

type StringLiteral struct {
	Token        Token
	InferredType *Type
	Value        string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) Position() []int      { return sl.Token.Position }
func (sl *StringLiteral) String() string       { return sl.Value }

type IndexExpression struct {
	Token        Token
	InferredType *Type
	Left         Expression
	Index        Expression
}

func (ie *IndexExpression) expressionNode()      {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) Position() []int      { return ie.Token.Position }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}

type StockLiteral struct {
	Token        Token
	InferredType *Type
	Pairs        map[Expression]Expression
}

func (sl *StockLiteral) expressionNode()      {}
func (sl *StockLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StockLiteral) Position() []int      { return sl.Token.Position }
func (sl *StockLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for key, value := range sl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

type FlowLiteral struct {
	Token        Token
	InferredType *Type
	Pairs        map[Expression]Expression
}

func (fl *FlowLiteral) expressionNode()      {}
func (fl *FlowLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FlowLiteral) Position() []int      { return fl.Token.Position }
func (fl *FlowLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for key, value := range fl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
