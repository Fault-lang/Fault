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

func (t *Token) GetPosition() []int {
	if len(t.Position) != 0 {
		return t.Position
	} else {
		return []int{0, 0, 0, 0}
	}
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
	"UNKNOWN":   6,
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
	Type() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Operand interface {
	Expression
	operandNode()
}

type Spec struct {
	Statements []Statement
	Ext        string // Is this a fspec or fsystem?
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

func (s *Spec) Type() string {
	var out bytes.Buffer

	for _, st := range s.Statements {
		out.WriteString(st.Type())
	}
	return out.String()
}

type SpecDeclStatement struct {
	Token Token
	Name  *Identifier
}

func (sd *SpecDeclStatement) statementNode()       {}
func (sd *SpecDeclStatement) TokenLiteral() string { return sd.Token.Literal }
func (sd *SpecDeclStatement) Position() []int      { return sd.Token.GetPosition() }
func (sd *SpecDeclStatement) String() string {
	var out bytes.Buffer

	out.WriteString(sd.TokenLiteral() + " ")
	out.WriteString(sd.Name.String())
	out.WriteString(";")
	return out.String()
}
func (sd *SpecDeclStatement) Type() string {
	return ""
}

type SysDeclStatement struct {
	Token Token
	Name  *Identifier
}

func (sd *SysDeclStatement) statementNode()       {}
func (sd *SysDeclStatement) TokenLiteral() string { return sd.Token.Literal }
func (sd *SysDeclStatement) Position() []int      { return sd.Token.GetPosition() }
func (sd *SysDeclStatement) String() string {
	var out bytes.Buffer

	out.WriteString(sd.TokenLiteral() + " ")
	out.WriteString(sd.Name.String())
	out.WriteString(";")
	return out.String()
}
func (sd *SysDeclStatement) Type() string {
	return ""
}

type ImportStatement struct {
	Token Token
	Name  *Identifier
	Path  *StringLiteral
	Tree  *Spec
}

func (is *ImportStatement) statementNode()       {}
func (is *ImportStatement) TokenLiteral() string { return is.Token.Literal }
func (is *ImportStatement) Position() []int      { return is.Token.GetPosition() }
func (is *ImportStatement) String() string {
	var out bytes.Buffer

	out.WriteString(is.TokenLiteral() + " ")
	out.WriteString(is.Name.String())
	out.WriteString(" = ")

	out.WriteString(is.Path.String())

	out.WriteString(";")
	return out.String()
}
func (is *ImportStatement) Type() string {
	return ""
}

type ConstantStatement struct {
	Token        Token
	Name         *Identifier
	Value        Expression
	InferredType *Type
}

func (cs *ConstantStatement) statementNode()       {}
func (cs *ConstantStatement) TokenLiteral() string { return cs.Token.Literal }
func (cs *ConstantStatement) Position() []int      { return cs.Token.GetPosition() }
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
func (cs *ConstantStatement) Type() string {
	return cs.Value.Type()
}

type DefStatement struct {
	Token Token
	Name  *Identifier
	Value Expression
}

func (ds *DefStatement) statementNode()       {}
func (ds *DefStatement) TokenLiteral() string { return ds.Token.Literal }
func (ds *DefStatement) Position() []int      { return ds.Token.GetPosition() }
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
func (ds *DefStatement) Type() string {
	return ds.Token.Literal
}

type AssertionStatement struct {
	Token          Token
	Variables      []Expression
	Constraints    *InvariantClause
	Temporal       string
	TemporalFilter string
	TemporalN      int
}

func (as *AssertionStatement) statementNode()       {}
func (as *AssertionStatement) TokenLiteral() string { return as.Token.Literal }
func (as *AssertionStatement) Position() []int      { return as.Token.GetPosition() }
func (as *AssertionStatement) String() string {
	var out bytes.Buffer

	out.WriteString(as.TokenLiteral() + " ")
	out.WriteString(as.Constraints.Left.String())
	out.WriteString(as.Constraints.Operator)
	out.WriteString(as.Constraints.Right.String())
	out.WriteString(";")
	return out.String()
}
func (as *AssertionStatement) Type() string {
	return as.Constraints.Right.Type()
}

type AssumptionStatement struct {
	Token          Token
	Variables      []Expression
	Constraints    *InvariantClause
	Temporal       string
	TemporalFilter string
	TemporalN      int
}

func (as *AssumptionStatement) statementNode()       {}
func (as *AssumptionStatement) TokenLiteral() string { return as.Token.Literal }
func (as *AssumptionStatement) Position() []int      { return as.Token.GetPosition() }
func (as *AssumptionStatement) String() string {
	var out bytes.Buffer

	out.WriteString(as.TokenLiteral() + " ")
	out.WriteString(as.Constraints.Left.String())
	out.WriteString(as.Constraints.Operator)
	out.WriteString(as.Constraints.Right.String())
	out.WriteString(";")
	return out.String()
}
func (as *AssumptionStatement) Type() string {
	return as.Constraints.Right.Type()
}

type InvariantClause struct {
	Token        Token
	Left         Expression
	Operator     string
	Right        Expression
	InferredType *Type
}

func (i *InvariantClause) expressionNode()      {}
func (i *InvariantClause) TokenLiteral() string { return i.Token.Literal }
func (i *InvariantClause) Position() []int      { return i.Token.GetPosition() }
func (i *InvariantClause) String() string {
	var out bytes.Buffer

	out.WriteString(i.TokenLiteral() + "assert ")
	out.WriteString(i.Left.String())
	out.WriteString(i.Operator)
	out.WriteString(i.Right.String())

	out.WriteString(";")
	return out.String()
}
func (i *InvariantClause) Type() string {
	return i.Right.Type()
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
func (i *Invariant) Position() []int      { return i.Token.GetPosition() }
func (i *Invariant) String() string {
	var out bytes.Buffer

	out.WriteString(i.TokenLiteral() + "assert ")
	out.WriteString(i.Variable.String())
	out.WriteString(i.Conjuction)
	out.WriteString(i.Expression.String())

	out.WriteString(";")
	return out.String()
}
func (i *Invariant) Type() string {
	return i.Variable.Type()
}

type ForStatement struct {
	Token  Token
	Rounds *IntegerLiteral
	Body   *BlockStatement
}

func (fs *ForStatement) statementNode()       {}
func (fs *ForStatement) TokenLiteral() string { return fs.Token.Literal }
func (fs *ForStatement) Position() []int      { return fs.Token.GetPosition() }
func (fs *ForStatement) String() string {
	var out bytes.Buffer

	out.WriteString(fs.TokenLiteral() + " ")
	out.WriteString(fs.Rounds.String())
	out.WriteString(fs.Body.String())

	out.WriteString(";")
	return out.String()
}
func (fs *ForStatement) Type() string {
	return ""
}

type StartStatement struct {
	Token Token
	Pairs [][]string
}

func (ss *StartStatement) statementNode()       {}
func (ss *StartStatement) TokenLiteral() string { return ss.Token.Literal }
func (ss *StartStatement) Position() []int      { return ss.Token.GetPosition() }
func (ss *StartStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ss.TokenLiteral() + " ")
	pairs := []string{}
	for _, value := range ss.Pairs {
		pairs = append(pairs, strings.Join(value, " : "))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	out.WriteString(";")
	return out.String()
}
func (ss *StartStatement) Type() string {
	return ""
}

type Identifier struct {
	Token        Token
	InferredType *Type
	Spec         string
	Value        string
}

func (i *Identifier) operandNode()         {}
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) Position() []int      { return i.Token.GetPosition() }
func (i *Identifier) String() string       { return i.Value }
func (i *Identifier) Type() string {
	t := i.InferredType
	if t != nil {
		return t.Type
	} else {
		return ""
	}
}

type ParameterCall struct {
	Token        Token
	InferredType *Type
	Value        []string
}

func (p *ParameterCall) operandNode()         {}
func (p *ParameterCall) expressionNode()      {}
func (p *ParameterCall) TokenLiteral() string { return p.Token.Literal }
func (p *ParameterCall) Position() []int      { return p.Token.GetPosition() }
func (p *ParameterCall) String() string {
	var out bytes.Buffer

	for _, s := range p.Value {
		out.WriteString(s)
	}

	return out.String()
}
func (p *ParameterCall) Type() string {
	t := p.InferredType
	if t != nil {
		return t.Type
	} else {
		return ""
	}
}

type AssertVar struct {
	Token        Token
	InferredType *Type
	Spec         string
	Instances    []string
}

func (av *AssertVar) expressionNode()      {}
func (av *AssertVar) TokenLiteral() string { return av.Token.Literal }
func (av *AssertVar) Position() []int      { return av.Token.GetPosition() }
func (av *AssertVar) String() string       { return strings.Join(av.Instances, " ") }
func (av *AssertVar) Type() string {
	t := av.InferredType
	if t != nil {
		return t.Type
	} else {
		return ""
	}
}

type Instance struct {
	Token        Token
	InferredType *Type
	Value        *Identifier
	Name         string
	Complex      bool //If stock does this stock contain another stock?
}

func (i *Instance) expressionNode()      {}
func (i *Instance) TokenLiteral() string { return i.Token.Literal }
func (i *Instance) Position() []int      { return i.Token.GetPosition() }
func (i *Instance) String() string {
	var out bytes.Buffer
	out.WriteString(i.Name)
	out.WriteString("= new ")
	out.WriteString(i.Value.String())

	return out.String()
}
func (i *Instance) Type() string {
	return i.Value.Type()
}

type ExpressionStatement struct {
	Token        Token
	InferredType *Type
	Expression   Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) Position() []int      { return es.Token.GetPosition() }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}
func (es *ExpressionStatement) Type() string {
	if es.Expression != nil {
		return es.Expression.Type()
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
func (il *IntegerLiteral) Position() []int      { return il.Token.GetPosition() }
func (il *IntegerLiteral) Type() string {
	ty := il.InferredType
	if ty != nil {
		return il.InferredType.Type
	} else {
		return "INT"
	}
}

type FloatLiteral struct {
	Token        Token
	InferredType *Type
	Value        float64
}

func (fl *FloatLiteral) expressionNode()      {}
func (fl *FloatLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FloatLiteral) Position() []int      { return fl.Token.GetPosition() }
func (fl *FloatLiteral) String() string       { return fmt.Sprint(fl.Value) }
func (fl *FloatLiteral) Type() string {
	ty := fl.InferredType
	if ty != nil {
		return fl.InferredType.Type
	} else {
		return "FLOAT"
	}
}

type Natural struct {
	Token        Token
	InferredType *Type
	Value        int64
}

func (n *Natural) expressionNode()      {}
func (n *Natural) TokenLiteral() string { return n.Token.Literal }
func (n *Natural) String() string       { return strconv.FormatInt(n.Value, 10) }
func (n *Natural) Position() []int      { return n.Token.GetPosition() }
func (n *Natural) Type() string {
	t := n.InferredType
	if t != nil {
		return t.Type
	} else {
		return "NATURAL"
	}
}

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
func (u *Uncertain) Type() string {
	t := u.InferredType
	if t != nil {
		return t.Type
	} else {
		return "UNCERTAIN"
	}
}
func (u *Uncertain) Position() []int { return u.Token.GetPosition() }

type Unknown struct {
	Token        Token
	InferredType *Type
	Name         *Identifier
}

func (u *Unknown) expressionNode()      {}
func (u *Unknown) TokenLiteral() string { return u.Token.Literal }
func (u *Unknown) String() string {
	var out bytes.Buffer
	out.WriteString("unknown(")
	if u.Name != nil { //This sometimes is set further up the tree and might be nil
		out.WriteString(u.Name.Value)
	}
	out.WriteString(")")
	return out.String()
}
func (u *Unknown) Type() string {
	t := u.InferredType
	if t != nil {
		return t.Type
	} else {
		return "UNKNOWN"
	}
}
func (u *Unknown) Position() []int { return u.Token.GetPosition() }

type PrefixExpression struct {
	Token        Token
	InferredType *Type
	Operator     string
	Right        Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) Position() []int      { return pe.Token.GetPosition() }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}
func (pe *PrefixExpression) Type() string { return pe.Right.Type() }

type InfixExpression struct {
	Token        Token
	InferredType *Type
	Left         Expression
	Operator     string
	Right        Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) Position() []int      { return ie.Token.GetPosition() }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}
func (ie *InfixExpression) Type() string { return ie.Right.Type() }

type Boolean struct {
	Token        Token
	InferredType *Type
	Value        bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) Position() []int      { return b.Token.GetPosition() }
func (b *Boolean) String() string       { return b.Token.Literal }
func (b *Boolean) Type() string {
	ty := b.InferredType
	if ty != nil {
		return b.InferredType.Type
	} else {
		return "BOOL"
	}
}

type This struct {
	Token        Token
	InferredType *Type
	Value        []string
}

func (t *This) expressionNode()      {}
func (t *This) TokenLiteral() string { return t.Token.Literal }
func (t *This) Position() []int      { return t.Token.GetPosition() }
func (t *This) String() string       { return t.Token.Literal }
func (t *This) Type() string {
	t2 := t.InferredType
	if t2 != nil {
		return t2.Type
	} else {
		return ""
	}
}

type Clock struct {
	Token        Token
	InferredType *Type
	Value        string
}

func (c *Clock) expressionNode()      {}
func (c *Clock) TokenLiteral() string { return c.Token.Literal }
func (c *Clock) Position() []int      { return c.Token.GetPosition() }
func (c *Clock) String() string       { return c.Token.Literal }
func (c *Clock) Type() string {
	t := c.InferredType
	if t != nil {
		return t.Type
	} else {
		return ""
	}
}

type Nil struct {
	Token        Token
	InferredType *Type
}

func (n *Nil) expressionNode()      {}
func (n *Nil) TokenLiteral() string { return n.Token.Literal }
func (n *Nil) Position() []int      { return n.Token.GetPosition() }
func (n *Nil) String() string       { return n.Token.Literal }
func (n *Nil) Type() string {
	t := n.InferredType
	if t != nil {
		return t.Type
	} else {
		return ""
	}
}

type BlockStatement struct {
	Token        Token
	InferredType *Type
	Statements   []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) Position() []int      { return bs.Token.GetPosition() }
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}
func (bs *BlockStatement) Type() string { return "" }

type ParallelFunctions struct {
	Token        Token
	InferredType *Type
	Expressions  []Expression
}

func (pf *ParallelFunctions) statementNode()       {}
func (pf *ParallelFunctions) Position() []int      { return pf.Token.GetPosition() }
func (pf *ParallelFunctions) TokenLiteral() string { return pf.Token.Literal }
func (pf *ParallelFunctions) String() string {
	var out bytes.Buffer

	for _, s := range pf.Expressions {
		out.WriteString(s.String())
	}

	return out.String()
}
func (pf *ParallelFunctions) Type() string { return "" }

type InitExpression struct {
	Token        Token
	InferredType *Type
	Expression   Expression
}

func (ie *InitExpression) expressionNode()      {}
func (ie *InitExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InitExpression) Position() []int      { return ie.Token.GetPosition() }
func (ie *InitExpression) String() string {
	var out bytes.Buffer

	out.WriteString("init ")
	out.WriteString(ie.Expression.String())
	return out.String()
}
func (ie *InitExpression) Type() string { return "" }

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
func (ie *IfExpression) Position() []int      { return ie.Token.GetPosition() }
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if (ie.Elif != &IfExpression{}) && ie.Elif != nil {
		out.WriteString("else if")
		out.WriteString(ie.Elif.String())
	}

	if (ie.Alternative != &BlockStatement{}) && ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}
func (ie *IfExpression) Type() string { return "" }

type FunctionLiteral struct {
	Token      Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) Position() []int      { return fl.Token.GetPosition() }
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
func (fl *FunctionLiteral) Type() string { return fl.Body.Type() }

type StateLiteral struct {
	Token      Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (sl *StateLiteral) expressionNode()      {}
func (sl *StateLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StateLiteral) Position() []int      { return sl.Token.GetPosition() }
func (sl *StateLiteral) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range sl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(sl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(sl.Body.String())

	return out.String()
}
func (sl *StateLiteral) Type() string { return sl.Body.Type() }

type BuiltIn struct {
	Token      Token
	Parameters []Operand
	Function   string
}

func (b *BuiltIn) expressionNode()      {}
func (b *BuiltIn) TokenLiteral() string { return b.Token.Literal }
func (b *BuiltIn) Position() []int      { return b.Token.GetPosition() }
func (b *BuiltIn) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range b.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(b.Function)
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")

	return out.String()
}
func (b *BuiltIn) Type() string { return "builtin" }

type StringLiteral struct {
	Token        Token
	InferredType *Type
	Value        string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) Position() []int      { return sl.Token.GetPosition() }
func (sl *StringLiteral) String() string       { return sl.Value }
func (sl *StringLiteral) Type() string {
	t := sl.InferredType
	if t != nil {
		return t.Type
	} else {
		return "STRING"
	}
}

type IndexExpression struct {
	Token        Token
	InferredType *Type
	Left         Expression
	Index        Expression
}

func (ie *IndexExpression) expressionNode()      {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) Position() []int      { return ie.Token.GetPosition() }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}
func (ie *IndexExpression) Type() string {
	return ie.Left.Type()
}

type StockLiteral struct {
	Token        Token
	InferredType *Type
	Order        []string
	Pairs        map[Expression]Expression
}

func (sl *StockLiteral) expressionNode()      {}
func (sl *StockLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StockLiteral) Position() []int      { return sl.Token.GetPosition() }
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
func (sl *StockLiteral) Type() string { return "STOCK" }

type FlowLiteral struct {
	Token        Token
	InferredType *Type
	Order        []string
	Pairs        map[Expression]Expression
}

func (fl *FlowLiteral) expressionNode()      {}
func (fl *FlowLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FlowLiteral) Position() []int      { return fl.Token.GetPosition() }
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
func (fl *FlowLiteral) Type() string { return "FLOW" }

type ComponentLiteral struct {
	Token        Token
	InferredType *Type
	Order        []string
	Pairs        map[Expression]Expression
}

func (cl *ComponentLiteral) expressionNode()      {}
func (cl *ComponentLiteral) TokenLiteral() string { return cl.Token.Literal }
func (cl *ComponentLiteral) Position() []int      { return cl.Token.GetPosition() }
func (cl *ComponentLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for key, value := range cl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
func (cl *ComponentLiteral) Type() string { return "COMPONENT" }
