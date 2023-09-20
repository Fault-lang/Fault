package ast

import (
	"bytes"
	"fault/util"
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
	SetType(*Type)
	GetToken() Token
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
type Nameable interface {
	Expression
	SetId([]string)
	Id() []string
	IdString() string
	RawId() []string
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

func (s *Spec) GetToken() Token {
	return Token{Type: "SPEC",
		Literal:  "SPEC",
		Position: []int{0, 0, 0, 0}}
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

func (s *Spec) SetType(ty *Type) {
	//Do nothing, specs are not typeable
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

func (sd *SpecDeclStatement) GetToken() Token {
	return sd.Token
}

func (sd *SpecDeclStatement) Type() string {
	return "SPEC"
}
func (sd *SpecDeclStatement) SetType(ty *Type) {
	//Do nothing, specs are not typeable
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
func (sd *SysDeclStatement) GetToken() Token {
	return sd.Token
}
func (sd *SysDeclStatement) Type() string {
	return "SYSTEM"
}
func (sd *SysDeclStatement) SetType(ty *Type) {
	//Do nothing, specs are not typeable
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
func (is *ImportStatement) GetToken() Token {
	return is.Token
}
func (is *ImportStatement) Type() string {
	return "IMPORT"
}
func (is *ImportStatement) SetType(y *Type) {
	//Do nothing, specs are not typeable
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
func (cs *ConstantStatement) GetToken() Token {
	return cs.Token
}
func (cs *ConstantStatement) Type() string {
	return cs.Value.Type()
}
func (cs *ConstantStatement) SetType(ty *Type) {
	cs.InferredType = ty
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
func (ds *DefStatement) GetToken() Token {
	return ds.Token
}
func (ds *DefStatement) Type() string {
	return ds.Token.Literal
}
func (ds *DefStatement) SetType(ty *Type) {
	//Handling by Token Literal
}

type AssertionStatement struct {
	Token          Token
	Constraint     *InvariantClause
	Assume         bool
	Temporal       string
	TemporalFilter string
	TemporalN      int
	Violated       bool // After model checking, for output
}

func (as *AssertionStatement) statementNode()       {}
func (as *AssertionStatement) TokenLiteral() string { return as.Token.Literal }
func (as *AssertionStatement) Position() []int      { return as.Token.GetPosition() }
func (as *AssertionStatement) String() string {
	var out bytes.Buffer

	out.WriteString(as.TokenLiteral() + " ")
	if as.Assume {
		out.WriteString("assume ")
	} else {
		out.WriteString("assert ")
	}
	out.WriteString(as.Constraint.Left.String())
	out.WriteString(as.Constraint.Operator)
	out.WriteString(as.Constraint.Right.String())
	out.WriteString(";")
	return out.String()
}
func (as *AssertionStatement) EvLogString(negate bool) string {
	var out bytes.Buffer
	if as.Violated {
		out.WriteString("FAILED  ")
	} else {
		out.WriteString("OK  ")
	}

	if as.Assume {
		out.WriteString("assume ")
	} else {
		out.WriteString("assert ")
	}
	out.WriteString(as.Constraint.Left.String())
	out.WriteString(" ")
	if !as.Assume && negate {
		out.WriteString(util.OP_NEGATE[as.Constraint.Operator])
	} else {
		out.WriteString(as.Constraint.Operator)
	}
	out.WriteString(" ")
	out.WriteString(as.Constraint.Right.String())
	if as.TemporalFilter != "" {
		out.WriteString(" ")
		out.WriteString(as.TemporalFilter)
		out.WriteString(" ")
		out.WriteString(fmt.Sprintf("%v", as.TemporalN))
	} else if as.Temporal != "" {
		out.WriteString(" ")
		out.WriteString(as.Temporal)
	}
	out.WriteString(";")
	return out.String()
}
func (as *AssertionStatement) GetToken() Token {
	return as.Token
}
func (as *AssertionStatement) Type() string {
	return as.Constraint.Right.Type()
}
func (as *AssertionStatement) SetType(ty *Type) {
	//Skip
}

type InvariantClause struct {
	Token        Token
	Left         Expression
	Operator     string
	Right        Expression
	SyncedState  bool
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
func (i *InvariantClause) GetToken() Token {
	return i.Token
}
func (i *InvariantClause) Type() string {
	return i.Right.Type()
}
func (i *InvariantClause) SetType(ty *Type) {
	//Skip
}

type ForStatement struct {
	Token  Token
	Rounds *IntegerLiteral
	Body   *BlockStatement
	Inits  *BlockStatement
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
func (fs *ForStatement) GetToken() Token {
	return fs.Token
}
func (fs *ForStatement) Type() string {
	return ""
}
func (fs *ForStatement) SetType(ty *Type) {
	//Skip
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
func (ss *StartStatement) GetToken() Token {
	return ss.Token
}
func (ss *StartStatement) Type() string {
	return "START"
}
func (ss *StartStatement) SetType(ty *Type) {
	//skip
}

type Identifier struct {
	Token         Token
	InferredType  *Type
	Spec          string
	Value         string
	ProcessedName []string
}

func (i *Identifier) operandNode()         {}
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) Position() []int      { return i.Token.GetPosition() }
func (i *Identifier) String() string {
	return i.Value
}
func (i *Identifier) Type() string {
	t := i.InferredType
	if t != nil {
		return t.Type
	} else {
		return ""
	}
}
func (i *Identifier) GetToken() Token {
	return i.Token
}
func (i *Identifier) SetType(ty *Type) {
	i.InferredType = ty
}
func (i *Identifier) SetId(id []string) {
	i.ProcessedName = id
}
func (i *Identifier) Id() []string {
	return []string{i.ProcessedName[0], strings.Join(i.ProcessedName[1:], "_")}
}
func (i *Identifier) IdString() string {
	return strings.Join(i.ProcessedName, "_")
}
func (i *Identifier) RawId() []string {
	return i.ProcessedName
}

type ParameterCall struct {
	Token         Token
	InferredType  *Type
	Spec          string
	Scope         string
	Value         []string
	ProcessedName []string
}

func (p *ParameterCall) operandNode()         {}
func (p *ParameterCall) expressionNode()      {}
func (p *ParameterCall) TokenLiteral() string { return p.Token.Literal }
func (p *ParameterCall) Position() []int      { return p.Token.GetPosition() }
func (p *ParameterCall) String() string {
	var out bytes.Buffer

	out.WriteString(strings.Join(p.Value, "."))

	return out.String()
}
func (p *ParameterCall) GetToken() Token {
	return p.Token
}
func (p *ParameterCall) Type() string {
	t := p.InferredType
	if t != nil {
		return t.Type
	} else {
		return ""
	}
}
func (p *ParameterCall) SetType(ty *Type) {
	p.InferredType = ty
}
func (p *ParameterCall) SetId(id []string) {
	p.ProcessedName = id
}
func (p *ParameterCall) Id() []string {
	return []string{p.ProcessedName[0], strings.Join(p.ProcessedName[1:], "_")}
}
func (p *ParameterCall) IdString() string {
	return strings.Join(p.ProcessedName, "_")
}

func (p *ParameterCall) RawId() []string {
	return p.ProcessedName
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
func (av *AssertVar) String() string {
	return fmt.Sprintf("(%s)", strings.Join(av.Instances, " or "))
}
func (av *AssertVar) Type() string {
	t := av.InferredType
	if t != nil {
		return t.Type
	} else {
		return ""
	}
}
func (av *AssertVar) GetToken() Token {
	return av.Token
}
func (av *AssertVar) SetType(ty *Type) {
	av.InferredType = ty
}

type StructInstance struct {
	Token         Token
	InferredType  *Type
	Complex       bool
	Properties    map[string]*StructProperty
	Order         []string
	Spec          string
	Name          string
	Parent        []string
	ComplexScope  string
	Swaps         []Node
	ProcessedName []string
}

func (si *StructInstance) expressionNode()      {}
func (si *StructInstance) TokenLiteral() string { return si.Token.Literal }
func (si *StructInstance) Position() []int      { return si.Token.GetPosition() }
func (si *StructInstance) String() string {
	var out bytes.Buffer
	for key, value := range si.Properties {
		out.WriteString(fmt.Sprintf("%s_%s_%s:%s", si.Spec, si.Name, key, value.String()))
	}
	return out.String()
}
func (si *StructInstance) GetToken() Token {
	return si.Token
}
func (si *StructInstance) Type() string {
	return string(si.Token.Literal)
}
func (si *StructInstance) SetType(ty *Type) {
	//Skip
}
func (si *StructInstance) Id() []string {
	return []string{si.ProcessedName[0], strings.Join(si.ProcessedName[1:], "_")}
}
func (si *StructInstance) SetId(id []string) {
	si.ProcessedName = id
}
func (si *StructInstance) IdString() string {
	return strings.Join(si.ProcessedName, "_")
}
func (si *StructInstance) RawId() []string {
	return si.ProcessedName
}

type StructProperty struct {
	Token         Token
	InferredType  *Type
	Value         Node
	Spec          string
	Name          string
	ProcessedName []string
}

func (sp *StructProperty) expressionNode()      {}
func (sp *StructProperty) TokenLiteral() string { return sp.Token.Literal }
func (sp *StructProperty) Position() []int      { return sp.Token.GetPosition() }
func (sp *StructProperty) String() string {
	var out bytes.Buffer
	out.WriteString(sp.Value.String())
	return out.String()
}
func (sp *StructProperty) SetId(id []string) {
	sp.ProcessedName = id
}
func (sp *StructProperty) GetToken() Token {
	return sp.Token
}
func (sp *StructProperty) Type() string {
	return string(sp.Value.Type())
}
func (sp *StructProperty) SetType(ty *Type) {
	//skip
}
func (sp *StructProperty) Id() []string {
	return []string{sp.ProcessedName[0], strings.Join(sp.ProcessedName[1:], "_")}
}

func (sp *StructProperty) IdString() string {
	return strings.Join(sp.ProcessedName, "_")
}

func (sp *StructProperty) RawId() []string {
	return sp.ProcessedName
}

type Instance struct {
	Token         Token
	InferredType  *Type
	Value         *Identifier
	Name          string
	Complex       bool //If stock does this stock contain another stock?
	ComplexScope  string
	Processed     *StructInstance
	Swaps         []Node
	ProcessedName []string
	Order         []string
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
func (i *Instance) GetToken() Token {
	return i.Token
}
func (i *Instance) Type() string {
	return i.Token.Literal
}
func (i *Instance) SetType(ty *Type) {
	//Skip
}
func (i *Instance) SetId(id []string) {
	i.ProcessedName = id
}
func (i *Instance) Id() []string {
	return []string{i.ProcessedName[0], strings.Join(i.ProcessedName[1:], "_")}
}

func (i *Instance) IdString() string {
	return strings.Join(i.ProcessedName, "_")
}

func (i *Instance) RawId() []string {
	return i.ProcessedName
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
func (es *ExpressionStatement) GetToken() Token {
	return es.Token
}
func (es *ExpressionStatement) Type() string {
	if es.Expression != nil {
		return es.Expression.Type()
	}

	return ""
}
func (es *ExpressionStatement) SetType(ty *Type) {
	//Skip
}

type IntegerLiteral struct {
	Token         Token
	InferredType  *Type
	Value         int64
	ProcessedName []string
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
func (il *IntegerLiteral) GetToken() Token {
	return il.Token
}
func (il *IntegerLiteral) SetType(ty *Type) {
	il.InferredType = ty
}
func (il *IntegerLiteral) SetId(id []string) {
	il.ProcessedName = id
}
func (il *IntegerLiteral) Id() []string {
	return []string{il.ProcessedName[0], strings.Join(il.ProcessedName[1:], "_")}
}
func (il *IntegerLiteral) IdString() string {
	return strings.Join(il.ProcessedName, "_")
}

func (il *IntegerLiteral) RawId() []string {
	return il.ProcessedName
}

type FloatLiteral struct {
	Token         Token
	InferredType  *Type
	Value         float64
	ProcessedName []string
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
func (fl *FloatLiteral) GetToken() Token {
	return fl.Token
}
func (fl *FloatLiteral) SetType(ty *Type) {
	fl.InferredType = ty
}
func (fl *FloatLiteral) SetId(id []string) {
	fl.ProcessedName = id
}
func (fl *FloatLiteral) Id() []string {
	return []string{fl.ProcessedName[0], strings.Join(fl.ProcessedName[1:], "_")}
}
func (fl *FloatLiteral) IdString() string {
	return strings.Join(fl.ProcessedName, "_")
}

func (fl *FloatLiteral) RawId() []string {
	return fl.ProcessedName
}

type Natural struct {
	Token         Token
	InferredType  *Type
	Value         int64
	ProcessedName []string
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
func (n *Natural) GetToken() Token {
	return n.Token
}
func (n *Natural) SetType(ty *Type) {
	n.InferredType = ty
}
func (n *Natural) SetId(id []string) {
	n.ProcessedName = id
}
func (n *Natural) Id() []string {
	return []string{n.ProcessedName[0], strings.Join(n.ProcessedName[1:], "_")}
}
func (n *Natural) IdString() string {
	return strings.Join(n.ProcessedName, "_")
}

func (n *Natural) RawId() []string {
	return n.ProcessedName
}

type Uncertain struct {
	Token         Token
	InferredType  *Type
	Mean          float64
	Sigma         float64
	ProcessedName []string
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
func (u *Uncertain) GetToken() Token {
	return u.Token
}
func (u *Uncertain) Type() string {
	t := u.InferredType
	if t != nil {
		return t.Type
	} else {
		return "UNCERTAIN"
	}
}
func (u *Uncertain) SetType(ty *Type) {
	u.InferredType = ty
}
func (u *Uncertain) Id() []string {
	return []string{u.ProcessedName[0], strings.Join(u.ProcessedName[1:], "_")}
}
func (u *Uncertain) SetId(id []string) {
	u.ProcessedName = id
}
func (u *Uncertain) IdString() string {
	return strings.Join(u.ProcessedName, "_")
}

func (u *Uncertain) RawId() []string {
	return u.ProcessedName
}

func (u *Uncertain) Position() []int { return u.Token.GetPosition() }

type Unknown struct {
	Token         Token
	InferredType  *Type
	Name          *Identifier
	ProcessedName []string
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
func (u *Unknown) GetToken() Token {
	return u.Token
}
func (u *Unknown) Type() string {
	t := u.InferredType
	if t != nil {
		return t.Type
	} else {
		return "UNKNOWN"
	}
}
func (u *Unknown) SetType(ty *Type) {
	u.InferredType = ty
}
func (u *Unknown) Id() []string {
	return []string{u.ProcessedName[0], strings.Join(u.ProcessedName[1:], "_")}
}
func (u *Unknown) SetId(id []string) {
	u.ProcessedName = id
}

func (u *Unknown) IdString() string {
	return strings.Join(u.ProcessedName, "_")
}

func (u *Unknown) RawId() []string {
	return u.ProcessedName
}

func (u *Unknown) Position() []int { return u.Token.GetPosition() }

type PrefixExpression struct {
	Token         Token
	InferredType  *Type
	Operator      string
	Right         Expression
	ProcessedName []string
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) Position() []int      { return pe.Token.GetPosition() }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())

	return out.String()
}
func (pe *PrefixExpression) GetToken() Token {
	return pe.Token
}
func (pe *PrefixExpression) Type() string {
	t := pe.InferredType
	if t != nil {
		return t.Type
	} else {
		return pe.Right.Type()
	}
}
func (pe *PrefixExpression) SetType(ty *Type) {
	pe.InferredType = ty
}
func (pe *PrefixExpression) Id() []string {
	return []string{pe.ProcessedName[0], strings.Join(pe.ProcessedName[1:], "_")}
}
func (pe *PrefixExpression) SetId(id []string) {
	pe.ProcessedName = id
}
func (pe *PrefixExpression) IdString() string {
	return strings.Join(pe.ProcessedName, "_")
}
func (pe *PrefixExpression) RawId() []string { return pe.ProcessedName }

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
func (ie *InfixExpression) GetToken() Token {
	return ie.Token
}
func (ie *InfixExpression) Type() string {
	t := ie.InferredType
	if t != nil {
		return t.Type
	} else {
		return ie.Right.Type()
	}
}
func (ie *InfixExpression) SetType(ty *Type) {
	ie.InferredType = ty
}

type Boolean struct {
	Token         Token
	InferredType  *Type
	Value         bool
	ProcessedName []string
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
func (b *Boolean) GetToken() Token {
	return b.Token
}
func (b *Boolean) SetType(ty *Type) {
	b.InferredType = ty
}
func (b *Boolean) Id() []string {
	return []string{b.ProcessedName[0], strings.Join(b.ProcessedName[1:], "_")}
}
func (b *Boolean) SetId(id []string) {
	b.ProcessedName = id
}
func (b *Boolean) IdString() string {
	return strings.Join(b.ProcessedName, "_")
}

func (b *Boolean) RawId() []string {
	return b.ProcessedName
}

type This struct {
	Token         Token
	InferredType  *Type
	Value         []string
	ProcessedName []string
}

func (t *This) expressionNode()      {}
func (t *This) operandNode()         {}
func (t *This) TokenLiteral() string { return t.Token.Literal }
func (t *This) Position() []int      { return t.Token.GetPosition() }
func (t *This) String() string       { return strings.Join(t.Value, ".") }
func (t *This) Type() string {
	t2 := t.InferredType
	if t2 != nil {
		return t2.Type
	} else {
		return ""
	}
}
func (t *This) GetToken() Token {
	return t.Token
}
func (t *This) SetType(ty *Type) {
	t.InferredType = ty
}
func (t *This) Id() []string {
	return []string{t.ProcessedName[0], strings.Join(t.ProcessedName[1:], "_")}
}
func (t *This) SetId(id []string) {
	t.ProcessedName = id
}

func (t *This) IdString() string {
	return strings.Join(t.ProcessedName, "_")
}

func (t *This) RawId() []string {
	return t.ProcessedName
}

type Clock struct {
	Token        Token
	InferredType *Type
	Value        string
}

func (c *Clock) expressionNode()      {}
func (c *Clock) TokenLiteral() string { return c.Token.Literal }
func (c *Clock) Position() []int      { return c.Token.GetPosition() }
func (c *Clock) String() string       { return "now" }
func (c *Clock) Type() string {
	t := c.InferredType
	if t != nil {
		return t.Type
	} else {
		return ""
	}
}
func (c *Clock) GetToken() Token {
	return c.Token
}
func (c *Clock) SetType(ty *Type) {
	c.InferredType = ty
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
func (n *Nil) SetType(ty *Type) {
	n.InferredType = ty
}
func (n *Nil) GetToken() Token {
	return n.Token
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
func (bs *BlockStatement) GetToken() Token {
	return bs.Token
}
func (bs *BlockStatement) Type() string { return "" }
func (bs *BlockStatement) SetType(ty *Type) {
	//Skip
}

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
func (pf *ParallelFunctions) GetToken() Token {
	return pf.Token
}
func (pf *ParallelFunctions) SetType(ty *Type) {
	//skip
}

type InitExpression struct {
	Token         Token
	InferredType  *Type
	Expression    Expression
	ProcessedName []string
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
func (ie *InitExpression) GetToken() Token {
	return ie.Token
}
func (ie *InitExpression) Type() string { return "" }
func (ie *InitExpression) SetType(ty *Type) {
	ie.InferredType = ty
}
func (ie *InitExpression) Id() []string { // returns []string{spec, rest_of_the_id}
	return []string{ie.ProcessedName[0], strings.Join(ie.ProcessedName[1:], "_")}
}
func (ie *InitExpression) SetId(id []string) {
	ie.ProcessedName = id
}
func (ie *InitExpression) IdString() string {
	return strings.Join(ie.ProcessedName, "_")
}
func (ie *InitExpression) RawId() []string {
	return ie.ProcessedName
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
func (ie *IfExpression) Position() []int      { return ie.Token.GetPosition() }
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if(")
	out.WriteString(ie.Condition.String())
	out.WriteString("){")
	out.WriteString(ie.Consequence.String())

	if (ie.Elif != &IfExpression{}) && ie.Elif != nil {
		out.WriteString("}else if(")
		out.WriteString(ie.Elif.Condition.String())
		out.WriteString("){")
		out.WriteString(ie.Elif.Consequence.String())
		out.WriteString("}")
	} else if (ie.Alternative != &BlockStatement{}) && ie.Alternative != nil {
		out.WriteString("}else{")
		out.WriteString(ie.Alternative.String())
		out.WriteString("}")
	} else {
		out.WriteString("}")
	}

	return out.String()
}
func (ie *IfExpression) GetToken() Token {
	return ie.Token
}
func (ie *IfExpression) Type() string { return "" }
func (ie *IfExpression) SetType(ty *Type) {
	//skip
}

type FunctionLiteral struct {
	Token         Token
	Parameters    []*Identifier
	Body          *BlockStatement
	ProcessedName []string
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
func (fl *FunctionLiteral) GetToken() Token {
	return fl.Token
}
func (fl *FunctionLiteral) Type() string { return fl.Body.Statements[len(fl.Body.Statements)-1].Type() }
func (fl *FunctionLiteral) SetType(ty *Type) {
	//skip
}
func (fl *FunctionLiteral) Id() []string {
	return []string{fl.ProcessedName[0], strings.Join(fl.ProcessedName[1:], "_")}
}
func (fl *FunctionLiteral) SetId(id []string) {
	fl.ProcessedName = id
}
func (fl *FunctionLiteral) IdString() string {
	return strings.Join(fl.ProcessedName, "_")
}
func (fl *FunctionLiteral) RawId() []string {
	return fl.ProcessedName
}

type BuiltIn struct {
	Token         Token
	Parameters    map[string]Operand
	Function      string
	FromState     string
	ProcessedName []string
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
	out.WriteString(")")

	return out.String()
}
func (b *BuiltIn) GetToken() Token {
	return b.Token
}
func (b *BuiltIn) Type() string { return "BUILTIN" }
func (b *BuiltIn) SetType(ty *Type) {
	//skip
}
func (b *BuiltIn) Id() []string { // returns []string{spec, rest_of_the_id}
	return []string{b.ProcessedName[0], strings.Join(b.ProcessedName[1:], "_")}
}
func (b *BuiltIn) SetId(id []string) {
	b.ProcessedName = id
}
func (b *BuiltIn) IdString() string {
	return strings.Join(b.ProcessedName, "_")
}
func (b *BuiltIn) RawId() []string {
	return b.ProcessedName
}

type StringLiteral struct {
	Token         Token
	InferredType  *Type
	Value         string
	ProcessedName []string
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
func (sl *StringLiteral) GetToken() Token {
	return sl.Token
}
func (sl *StringLiteral) SetType(ty *Type) {
	sl.InferredType = ty
}
func (sl *StringLiteral) Id() []string {
	return []string{sl.ProcessedName[0], strings.Join(sl.ProcessedName[1:], "_")}
}
func (sl *StringLiteral) SetId(id []string) {
	sl.ProcessedName = id
}
func (sl *StringLiteral) IdString() string {
	return strings.Join(sl.ProcessedName, "_")
}
func (sl *StringLiteral) RawId() []string {
	return sl.ProcessedName
}

type IndexExpression struct {
	Token         Token
	InferredType  *Type
	Left          Expression
	Index         Expression
	ProcessedName []string
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
func (ie *IndexExpression) GetToken() Token {
	return ie.Token
}
func (ie *IndexExpression) Type() string {
	return ie.Left.Type()
}
func (ie *IndexExpression) SetType(ty *Type) {
	//skip
}
func (ie *IndexExpression) Id() []string {
	return []string{ie.ProcessedName[0], strings.Join(ie.ProcessedName[1:], "_")}
}
func (ie *IndexExpression) SetId(id []string) {
	ie.ProcessedName = id
}
func (ie *IndexExpression) IdString() string {
	return strings.Join(ie.ProcessedName, "_")
}
func (ie *IndexExpression) RawId() []string {
	return ie.ProcessedName
}

type StockLiteral struct {
	Token         Token
	InferredType  *Type
	Order         []string
	Pairs         map[*Identifier]Expression
	ProcessedName []string
}

func (sl *StockLiteral) expressionNode()      {}
func (sl *StockLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StockLiteral) Position() []int      { return sl.Token.GetPosition() }
func (sl *StockLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, k := range sl.Order {
		key := sl.GetPropertyIdent(k)
		value := sl.Pairs[key]
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
func (sl *StockLiteral) GetToken() Token {
	return sl.Token
}
func (sl *StockLiteral) Type() string { return "STOCK" }
func (sl *StockLiteral) SetType(ty *Type) {
	//skip
}
func (sl *StockLiteral) Id() []string { // returns []string{spec, rest_of_the_id}
	return []string{sl.ProcessedName[0], strings.Join(sl.ProcessedName[1:], "_")}
}
func (sl *StockLiteral) SetId(id []string) {
	sl.ProcessedName = id
}
func (sl *StockLiteral) IdString() string {
	return strings.Join(sl.ProcessedName, "_")
}
func (sl *StockLiteral) RawId() []string {
	return sl.ProcessedName
}
func (sl *StockLiteral) GetPropertyIdent(key string) *Identifier {
	for k := range sl.Pairs {
		if k.Value == key {
			return k
		}
	}
	return nil
}

type FlowLiteral struct {
	Token         Token
	InferredType  *Type
	Order         []string
	Pairs         map[*Identifier]Expression
	ProcessedName []string
}

func (fl *FlowLiteral) expressionNode()      {}
func (fl *FlowLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FlowLiteral) Position() []int      { return fl.Token.GetPosition() }
func (fl *FlowLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, k := range fl.Order {
		key := fl.GetPropertyIdent(k)
		value := fl.Pairs[key]
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
func (fl *FlowLiteral) GetToken() Token {
	return fl.Token
}
func (fl *FlowLiteral) Type() string { return "FLOW" }
func (fl *FlowLiteral) SetType(ty *Type) {
	//skip
}
func (fl *FlowLiteral) Id() []string {
	return []string{fl.ProcessedName[0], strings.Join(fl.ProcessedName[1:], "_")}
}
func (fl *FlowLiteral) SetId(id []string) {
	fl.ProcessedName = id
}
func (fl *FlowLiteral) IdString() string {
	return strings.Join(fl.ProcessedName, "_")
}
func (fl *FlowLiteral) RawId() []string {
	return fl.ProcessedName
}
func (fl *FlowLiteral) GetPropertyIdent(key string) *Identifier {
	for k := range fl.Pairs {
		if k.Value == key {
			return k
		}
	}
	return nil
}

type ComponentLiteral struct {
	Token         Token
	InferredType  *Type
	Order         []string
	Pairs         map[*Identifier]Expression
	ProcessedName []string
}

func (cl *ComponentLiteral) expressionNode()      {}
func (cl *ComponentLiteral) TokenLiteral() string { return cl.Token.Literal }
func (cl *ComponentLiteral) Position() []int      { return cl.Token.GetPosition() }
func (cl *ComponentLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, k := range cl.Order {
		key := cl.GetPropertyIdent(k)
		value := cl.Pairs[key]
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
func (cl *ComponentLiteral) GetToken() Token {
	return cl.Token
}
func (cl *ComponentLiteral) Type() string { return "COMPONENT" }
func (cl *ComponentLiteral) SetType(ty *Type) {
	//skip
}
func (cl *ComponentLiteral) Id() []string {
	return []string{cl.ProcessedName[0], strings.Join(cl.ProcessedName[1:], "_")}
}
func (cl *ComponentLiteral) SetId(id []string) {
	cl.ProcessedName = id
}
func (cl *ComponentLiteral) IdString() string {
	return strings.Join(cl.ProcessedName, "_")
}
func (cl *ComponentLiteral) RawId() []string {
	return cl.ProcessedName
}
func (cl *ComponentLiteral) GetPropertyIdent(key string) *Identifier {
	for k := range cl.Pairs {
		if k.Value == key {
			return k
		}
	}
	return nil
}
