package llvm

import (
	"errors"
	"fault/ast"
	"fault/llvm/name"
	"fault/types"
	"fault/util"
	"fmt"
	"runtime/debug"
	"sort"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/metadata"
	irtypes "github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// Cribs a bit from https://github.com/zegl/tre
// Will likely remove most of that influence over
// time. For now Tre is copyright (c) 2018
// Gustav Westling <gustav@westling.xyz>

var DoubleP = &irtypes.PointerType{ElemType: irtypes.Double}

var OP_NEGATE = map[string]string{
	"==": "!=",
	">=": "<",
	">":  "<=",
	"<=": ">",
	"!=": "==",
	"<":  ">=",
	"&&": "||",
	"||": "&&",
	//"=": "!=",
}

type Compiler struct {
	module *ir.Module

	specs            map[string]*spec
	instances        map[string]string
	instanceChildren map[string]string
	currentSpec      *spec

	currentSpecName string
	currScope       string

	specStructs   map[string]types.StockFlow
	specFunctions map[string]value.Value

	contextFuncName string
	contextMetadata *metadata.Attachment

	alloc    bool
	runRound int16

	contextBlock *ir.Block
	contextFunc  *ir.Func

	// Stack of variables that are in scope
	allocatedPointers []map[string]*ir.InstAlloca

	// Where a condition should jump when done
	contextCondAfter []*ir.Block

	specGlobals map[string]*ir.Global
	Asserts     []*ast.AssertionStatement
	Assumes     []*ast.AssumptionStatement
	Uncertains  map[string][]float64
	Unknowns    []string
}

func NewCompiler() *Compiler {
	c := &Compiler{
		module: ir.NewModule(),

		specs:            make(map[string]*spec),
		instances:        make(map[string]string),
		instanceChildren: make(map[string]string),
		specStructs:      make(map[string]types.StockFlow),
		specFunctions:    make(map[string]value.Value),

		contextMetadata: nil,
		alloc:           true,

		allocatedPointers: make([]map[string]*ir.InstAlloca, 0),

		contextCondAfter: make([]*ir.Block, 0),

		specGlobals: make(map[string]*ir.Global),
		runRound:    0,

		Uncertains: make(map[string][]float64),
	}
	c.addGlobal()
	return c
}

func (c *Compiler) LoadMeta(structs map[string]types.StockFlow, uncertains map[string][]float64, unknowns []string) {
	c.specStructs = structs
	c.Unknowns = unknowns
	c.Uncertains = uncertains
}

func (c *Compiler) Compile(root ast.Node) (err error) {
	defer func() {
		if r := recover(); r != nil {
			// Compile time panics, that are not errors in the compiler
			if _, ok := r.(Panic); ok {
				err = errors.New(fmt.Sprint(r))
				return
			}

			// Bugs in the compiler
			err = fmt.Errorf("%s\n\nInternal compiler stacktrace:\n%s",
				fmt.Sprint(r),
				string(debug.Stack()),
			)
		}
	}()
	c.processSpec(root, false)
	return
}

func (c *Compiler) processSpec(root ast.Node, isImport bool) ([]*ast.AssertionStatement, []*ast.AssumptionStatement) {
	specfile, ok := root.(*ast.Spec)
	if !ok {
		panic(fmt.Sprintf("spec file improperly formatted. Root node is %T", root))
	}
	decl, ok := specfile.Statements[0].(*ast.SpecDeclStatement)
	if !ok {
		panic(fmt.Sprintf("spec file improperly formatted. Missing spec declaration, got %T", decl))
	}

	name := decl.Name.Value
	c.currentSpec = NewCompiledSpec(name)
	c.currentSpecName = name
	c.specs[c.currentSpecName] = c.currentSpec
	for _, fileNode := range specfile.Statements {
		c.compile(fileNode)
	}

	if !isImport {
		for _, assert := range c.Asserts {
			c.Asserts = c.Asserts[1:] //Pop
			c.compileAssert(assert)
		}
		for _, assert := range c.Assumes {
			c.Assumes = c.Assumes[1:] //Pop
			c.compileAssert(assert)
		}
	}
	return c.Asserts, c.Assumes
}

func (c *Compiler) compile(node ast.Node) {
	switch v := node.(type) {
	case *ast.SpecDeclStatement:
		break
	case *ast.ImportStatement:
		parent := c.currentSpecName
		parentSp := c.currentSpec
		asserts, assumes := c.processSpec(v.Tree, true) //Move all asserts to the end of the compilation process
		c.Asserts = append(c.Asserts, asserts...)
		c.Assumes = append(c.Assumes, assumes...)
		c.currentSpecName = parent
		c.currentSpec = parentSp
	case *ast.ConstantStatement:
		c.compileConstant(v)
	case *ast.DefStatement:
		c.compileStruct(v)

	case *ast.FunctionLiteral:

	case *ast.InfixExpression:
		c.compileInfix(v)

	case *ast.PrefixExpression:

	case *ast.AssumptionStatement:
		// Need to do these after the run block so we move them
		c.Assumes = append(c.Assumes, v)

	case *ast.AssertionStatement:
		c.Asserts = append(c.Asserts, v)
		//c.compileAssertion(v)

	case *ast.ForStatement:
		c.contextFuncName = "__run"
		for i := int64(0); i < v.Rounds.Value; i++ {
			c.compileBlock(v.Body)
			c.runRound = c.runRound + 1
		}
		c.contextFuncName = ""
	default:
		pos := node.Position()
		panic(fmt.Sprintf("node type %T unimplemented line: %d col: %d", v, pos[0], pos[1]))
	}

	// InitExpression
	// IfExpression
	// IndexExpression <-- Is this still used?

}

func (c *Compiler) compileIdent(node *ast.Identifier) *ir.InstLoad {
	return c.lookupIdent([]string{node.Value}, node.Position())
}

func (c *Compiler) compileInfix(node *ast.InfixExpression) value.Value {
	pos := node.Position()
	switch node.Operator {
	case "=": // Used to store temporary local values
		r := c.compileValue(node.Right)
		if _, ok := node.Right.(*ast.Instance); !ok { // If declaring a new instance don't save
			var fvn []string
			switch n := node.Left.(type) {
			case *ast.Identifier:
				fvn = c.getFullVariableName([]string{n.Value})
			case *ast.ParameterCall:
				fvn = c.getFullVariableName(n.Value)
			}

			if c.isVarSet(fvn) && c.alloc {
				_, s := c.GetSpec(fvn)
				fvns := c.getVariableName(fvn)
				p := s.GetSpecVarPointer(fvns)
				c.contextBlock.NewStore(r, p)
				return nil
			}
			id, s := c.GetSpec(fvn)
			s.DefineSpecVar(id, r)
			if c.alloc {
				c.allocVariable(id, r, node.Left.Position())
			}
		}
		return nil
	case "<-":
		r := c.compileValue(node.Right)
		n, ok := node.Left.(*ast.ParameterCall)
		if !ok {
			pos := node.Position()
			panic(fmt.Sprintf("cannot use <- or -> operator on a non-stock value col: %d, line: %d", pos[0], pos[1]))
		}

		id := n.Value
		pos := n.Position()
		fvn := c.getFullVariableName(id)
		//fvns := c.getVariableStateName(fvn)
		fvns := c.getVariableName(fvn)

		if !c.isVarSet(fvn) {
			panic(fmt.Sprintf("cannot send value to variable %s. Variable not defined line: %d, col: %d", fvns, pos[0], pos[1]))
		}

		_, s := c.GetSpec(fvn)
		pointer := s.GetSpecVarPointer(fvns)
		c.contextBlock.NewStore(r, pointer)
		return nil
	case "+":
		l := c.compilerInfixNode(node.Left)
		r := c.compilerInfixNode(node.Right)

		add := c.contextBlock.NewFAdd(l, r)
		return add
	case "-":
		l := c.compilerInfixNode(node.Left)
		r := c.compilerInfixNode(node.Right)

		sub := c.contextBlock.NewFSub(l, r)
		return sub
	case "*":
		l := c.compilerInfixNode(node.Left)
		r := c.compilerInfixNode(node.Right)

		mul := c.contextBlock.NewFMul(l, r)
		return mul
	case "/":
		l := c.compilerInfixNode(node.Left)
		r := c.compilerInfixNode(node.Right)

		div := c.contextBlock.NewFDiv(l, r)
		return div
	case "%":
		l := c.compilerInfixNode(node.Left)
		r := c.compilerInfixNode(node.Right)

		rem := c.contextBlock.NewFRem(l, r)
		return rem
	case ">":
		l := c.compilerInfixNode(node.Left)
		r := c.compilerInfixNode(node.Right)
		ogt := c.contextBlock.NewFCmp(enum.FPredOGT, l, r)
		return ogt
	case ">=":
		l := c.compilerInfixNode(node.Left)
		r := c.compilerInfixNode(node.Right)

		oge := c.contextBlock.NewFCmp(enum.FPredOGE, l, r)
		return oge
	case "<":
		l := c.compilerInfixNode(node.Left)
		r := c.compilerInfixNode(node.Right)

		olt := c.contextBlock.NewFCmp(enum.FPredOLT, l, r)
		return olt
	case "<=":
		l := c.compilerInfixNode(node.Left)
		r := c.compilerInfixNode(node.Right)

		ole := c.contextBlock.NewFCmp(enum.FPredOLE, l, r)
		return ole
	case "==":
		l := c.compilerInfixNode(node.Left)
		r := c.compilerInfixNode(node.Right)

		oeq := c.contextBlock.NewFCmp(enum.FPredOEQ, l, r)
		return oeq
	case "!=":
		l := c.compilerInfixNode(node.Left)
		r := c.compilerInfixNode(node.Right)

		one := c.contextBlock.NewFCmp(enum.FPredONE, l, r)
		return one
	default:
		panic(fmt.Sprintf("unknown operator %s. line: %d, col: %d", node.Operator, pos[0], pos[1]))
	}
}

func (c *Compiler) compilerInfixNode(node ast.Node) value.Value {
	switch v := node.(type) {
	case *ast.ParameterCall:
		return c.lookupIdent(v.Value, node.Position())
	default:
		return c.compileValue(node)
	}
}

func (c *Compiler) compileParallel(node *ast.ParallelFunctions) {
	if c.contextFuncName != "__run" {
		pos := node.Position()
		panic(fmt.Sprintf("cannot use parallel operator outside of the run block. line: %d, col: %d", pos[0], pos[1]))
	}
	gname := name.ParallelGroup(node.String())
	for i := 0; i < len(node.Expressions); i++ {
		l := c.compileValue(node.Expressions[i])
		md := &metadata.Attachment{
			Name: gname,
			Node: &metadata.DIBasicType{
				MetadataID: -1,
				Tag:        enum.DwarfTagStringType,
			}}
		switch exp := l.(type) {
		case *ir.Func:
			id := node.Expressions[i].(*ast.ParameterCall).Value
			id, s := c.GetSpec(id)
			params := s.GetParams(id)
			l_func := c.contextBlock.NewCall(exp, params...)
			l_func.Metadata = append(l_func.Metadata, md)
		case *ir.InstFAdd:
			exp.Metadata = append(exp.Metadata, md)
		case *ir.InstFSub:
			exp.Metadata = append(exp.Metadata, md)
		case *ir.InstFMul:
			exp.Metadata = append(exp.Metadata, md)
		case *ir.InstFDiv:
			exp.Metadata = append(exp.Metadata, md)
		case *ir.InstFRem:
			exp.Metadata = append(exp.Metadata, md)
		case *ir.InstFCmp:
			exp.Metadata = append(exp.Metadata, md)
		}
	}
	c.contextMetadata = nil
}

func (c *Compiler) compileValue(node ast.Node) value.Value {
	switch v := node.(type) {
	case *ast.IntegerLiteral:
		return constant.NewFloat(irtypes.Double, float64(v.Value))
	case *ast.FloatLiteral:
		return constant.NewFloat(irtypes.Double, v.Value)
	case *ast.StringLiteral:
		return constant.NewCharArrayFromString(v.Value)
	case *ast.Boolean:
		return constant.NewBool(v.Value)
	case *ast.Natural:
		return constant.NewFloat(irtypes.Double, float64(v.Value))
	case *ast.Uncertain: //Set to dummy value for LLVM IR, catch during SMT generation
		return constant.NewFloat(irtypes.Double, float64(0.000000000009))
	case *ast.Unknown:
		return constant.NewFloat(irtypes.Double, float64(0.000000000009))
	case *ast.Nil:
		return constant.NewNull(&irtypes.PointerType{})
	case *ast.Identifier:
		return c.compileIdent(v)
	case *ast.InfixExpression:
		return c.compileInfix(v)
	case *ast.FunctionLiteral:
		return c.compileFunction(v)
	case *ast.Instance:
		c.compileInstance(v, v.Name)
	case *ast.ParameterCall:
		return c.compileParameterCall(v)
	case *ast.BlockStatement:
		return c.compileBlock(v)
	default:
		pos := v.Position()
		panic(fmt.Sprintf("unknown value type %T line: %d col: %d", v, pos[0], pos[1]))
	}
	return nil
}

func (c *Compiler) compileConstant(node *ast.ConstantStatement) {
	value := c.compileValue(node.Value)
	id := c.getFullVariableName([]string{node.Name.Value})
	id, _ = c.GetSpec(id)
	c.setConst(id, value)
	c.globalVariable(id, value, node.Position())
}

func (c *Compiler) compileBlock(node *ast.BlockStatement) value.Value {
	if !c.alloc {
		return nil
	}
	body := node.Statements
	var ret value.Value
	for i := 0; i < len(body); i++ {
		switch exp := body[i].(type) {
		case *ast.ParallelFunctions:
			c.compileParallel(exp)
		case ast.Expression:
			ret = c.compileFunctionBody(exp)
		case *ast.ExpressionStatement:
			ret = c.compileFunctionBody(exp.Expression)
		}
	}
	return ret
}

func (c *Compiler) GetIR() string {
	return c.module.String()
}

func (c *Compiler) addGlobal() {
	global := NewCompiledSpec("__global")

	c.specs["__global"] = global

	// run block
	c.contextFunc = c.module.NewFunc("__run", irtypes.Void)
	mainBlock := c.contextFunc.NewBlock(name.Block())
	mainBlock.NewRet(nil)
	c.contextBlock = mainBlock
}

func (c *Compiler) compileStruct(def *ast.DefStatement) {
	//Not implemented, using preparse from type checker
}

func (c *Compiler) compileAssert(assert ast.Node) {
	var l, r ast.Expression
	switch a := assert.(type) {
	case *ast.AssertionStatement:
		if a.TemporalFilter == "" { //If there is a temporal filter this is negated instead
			l = negate(a.Constraints.Left)
			r = negate(a.Constraints.Right)
			a.Constraints.Operator = OP_NEGATE[a.Constraints.Operator]
		} else {
			l = a.Constraints.Left
			r = a.Constraints.Right
			a.TemporalFilter, a.TemporalN = negateTemporal(a.TemporalFilter, a.TemporalN)
			if a.TemporalN < 0 {
				pos := a.Position()
				panic(fmt.Sprintf("temporal logic not value, filter searching for fewer than 0 states: line %d col %d", pos[0], pos[1]))
			}
		}
		a.Constraints.Left = c.convertAssertVariables(l)
		a.Constraints.Right = c.convertAssertVariables(r)
		c.Asserts = append(c.Asserts, a)
	case *ast.AssumptionStatement:
		a.Constraints.Left = c.convertAssertVariables(a.Constraints.Left)
		a.Constraints.Right = c.convertAssertVariables(a.Constraints.Right)
		c.Assumes = append(c.Assumes, a)
	default:
		panic("statement must be an assert or an assumption.")
	}
}

func (c *Compiler) getInstances(ex ast.Expression) map[string][]string {
	vars := make(map[string][]string)
	switch e := ex.(type) {
	case *ast.InfixExpression:
		left := c.getInstances(e.Left)
		for k, v := range left {
			vars[k] = util.MergeStrSlices(vars[k], v)
		}

		right := c.getInstances(e.Right)
		for k, v := range right {
			vars[k] = util.MergeStrSlices(vars[k], v)
		}
		return vars
	case *ast.Identifier:
		return nil
	case *ast.ParameterCall:
		id := e.Value
		for k, v := range c.instances {
			if v == id[0] {
				vars[v] = util.MergeStrSlices(vars[v], []string{k})
			}
		}
		return vars

	case *ast.PrefixExpression:
		right := c.getInstances(e.Right)
		for k, v := range right {
			vars[k] = util.MergeStrSlices(vars[k], v)
		}
		return vars
	case *ast.IndexExpression:
		left := c.getInstances(e.Left)
		for k, v := range left {
			vars[k] = util.MergeStrSlices(vars[k], v)
		}
		return vars
	default:
		return nil
	}
}

func (c *Compiler) generateOrder(pairs map[string]ast.Node) []string {
	keys := []string{}
	for k := range pairs {
		if k != "___base" {
			keys = append(keys, k)
		}
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

func (c *Compiler) compileInstance(base *ast.Instance, instName string) {
	if c.runRound > 0 { // Initialize things only once
		return
	}
	if c.contextFuncName == "__run" {
		c.currScope = instName
		c.alloc = false
	}

	pos := base.Position()
	structName := base.Value.Value
	parentFunction := c.contextFuncName
	c.contextFuncName = instName
	if c.specStructs[base.Value.Spec][structName] == nil {
		panic(fmt.Sprintf("no stock or flow named %s, line: %d, col %d", structName, pos[0], pos[1]))
	}

	keys := c.generateOrder(c.specStructs[base.Value.Spec][structName])
	for _, k := range keys {
		var isUncertain []float64
		var isUnknown bool

		id := c.getFullVariableName([]string{instName, k})
		switch pv := c.specStructs[base.Value.Spec][structName][k].(type) {
		case *ast.Instance:
			switch len(id) {
			case 3:
				// Slightly hacky solution to nestled instances
				c.compileInstance(pv, strings.Join(id[1:], "_"))
			default:
				c.compileInstance(pv, k) // Copy instance data over
			}
		case *ast.FunctionLiteral:
			c.compileFunction(pv)
		case *ast.InfixExpression:
			c.compileInfix(pv)
		case *ast.BlockStatement:
			c.compileBlock(pv)
		default:
			_, ok := pv.(*ast.Uncertain)
			if ok {
				isUnknown = true
			}

			uncertain, ok2 := pv.(*ast.Uncertain)
			if ok2 {
				isUncertain = []float64{uncertain.Mean, uncertain.Sigma}
			}
			val := c.compileValue(c.specStructs[base.Value.Spec][structName][k])
			id, s := c.GetSpec(id)
			s.DefineSpecVar(id, val)
			c.allocVariable(id, val, pos)
			s.vars.ResetState(id)
			name := c.getVariableName(id)
			p := ir.NewParam(name, DoubleP)
			s.AddParam(id, p)

		}
		//Track properties of instances so that we can write
		// asserts on the struct and honor them for all instances
		id, _ = c.GetSpec(id)
		fvn := strings.Join(id, "_")
		if isUnknown {
			c.Unknowns = append(c.Unknowns, fvn)
		}
		if isUncertain != nil {
			c.Uncertains[fvn] = isUncertain
		}
		c.instanceChildren[strings.Join(id, "_")] = structName
	}
	c.instances[instName] = structName
	c.contextFuncName = parentFunction
	if c.contextFuncName == "__run" {
		c.currScope = ""
		c.alloc = true
	}
}

func (c *Compiler) compileIf(n *ast.IfExpression) {
	cond := c.compileValue(n.Condition)

	afterBlock := c.contextBlock.Parent.NewBlock(name.Block() + "-after")
	trueBlock := c.contextBlock.Parent.NewBlock(name.Block() + "-true")
	falseBlock := afterBlock

	c.contextCondAfter = append(c.contextCondAfter, afterBlock)

	if n.Alternative != nil {
		falseBlock = c.contextBlock.Parent.NewBlock(name.Block() + "-false")
	}

	c.contextBlock.NewCondBr(cond, trueBlock, falseBlock)

	c.contextBlock = trueBlock
	c.compileBlock(n.Consequence)

	// Jump to after-block if no terminator has been set (such as a return statement)
	if trueBlock.Term == nil {
		trueBlock.NewBr(afterBlock)
	}

	if n.Alternative != nil {
		c.contextBlock = falseBlock
		c.compileBlock(n.Alternative)

		// Jump to after-block if no terminator has been set (such as a return statement)
		if falseBlock.Term == nil {
			falseBlock.NewBr(afterBlock)
		}
	}

	c.contextBlock = afterBlock

	// pop after block stack
	c.contextCondAfter = c.contextCondAfter[0 : len(c.contextCondAfter)-1]

	// set after block to jump to the after block
	if len(c.contextCondAfter) > 0 {
		afterBlock.NewBr(c.contextCondAfter[len(c.contextCondAfter)-1])
	}
}

func (c *Compiler) compileParameterCall(pc *ast.ParameterCall) value.Value {
	id := c.getFullVariableName(pc.Value)
	id, s := c.GetSpec(id)
	structName := c.instances[pc.Value[0]]
	// If we're in the run block and the parameter is defined as a function
	// define it as a function and call it from run block
	if c.contextFuncName == "__run" &&
		c.isFunction(c.specStructs[id[0]][structName][pc.Value[1]]) {
		//IR Function + Call
		fname := strings.Join(id, "_")

		if c.runRound == 0 {
			params := c.generateParameters(c.specStructs[id[0]][structName], id)
			c.resetParaState(params)
			f := c.module.NewFunc(fname, irtypes.Void, params...)
			c.contextFunc = f

			oldScope := c.currScope
			oldBlock := c.contextBlock

			c.contextFuncName = fname
			c.currScope = pc.Value[0]
			c.contextBlock = f.NewBlock(name.Block())

			val := c.compileValue(c.specStructs[id[0]][structName][pc.Value[1]])
			c.contextBlock.NewRet(val)

			c.contextBlock = oldBlock
			c.contextFuncName = "__run"
			c.currScope = oldScope
			c.specFunctions[fname] = f
			c.contextFunc = nil
			c.resetParaState(params)
		}

		return c.specFunctions[fname]
	}

	// Otherwise inline the parameter...
	if c.currScope == "" {
		c.currScope = pc.Value[0]
	}
	parentFunction := c.contextFuncName
	c.contextFuncName = pc.Value[0]

	val := c.compileValue(c.specStructs[id[0]][structName][pc.Value[1]])

	// If there's no value, there's nothing to store
	if val != nil || !c.isFunction(c.specStructs[id[0]][structName][pc.Value[1]]) {
		if s.GetSpecVar(id) != nil {
			//name := c.getVariableStateName(id)
			name := c.getVariableName(id)
			pointer := s.GetSpecVarPointer(name)
			c.contextBlock.NewLoad(irtypes.Double, pointer)
		} else {
			s.DefineSpecVar(id, val)
			c.allocVariable(id, val, pc.Position())
		}
	}
	if c.currScope == pc.Value[0] {
		c.currScope = ""
	}
	c.contextFuncName = parentFunction
	return val
}

func (c *Compiler) compileFunction(node *ast.FunctionLiteral) value.Value {
	body := node.Body.Statements
	var retValue value.Value
	for i := 0; i < len(body); i++ {
		exp := body[i].(*ast.ExpressionStatement).Expression
		init, ok := exp.(*ast.InitExpression)
		if ok {
			return c.compileValue(init.Expression)
		}
	}
	return retValue
}

func (c *Compiler) compileFunctionBody(node ast.Expression) value.Value {
	if !c.alloc { //Short circuit this if just initializing
		return nil
	}
	switch v := node.(type) {
	case *ast.InfixExpression:
		return c.compileInfix(v)

	case *ast.PrefixExpression:

	case *ast.IfExpression:
		c.compileIf(v)

	case *ast.Instance:
		orign := v.Name
		origv := v.Value.Value
		v.Value.Value = orign
		v.Name = origv
		c.compileInstance(v, v.Name)

	case *ast.IndexExpression:

	case *ast.ParameterCall:
		return c.compileParameterCall(v)

	default:
		pos := node.Position()
		panic(fmt.Sprintf("invalid expression %T in function body. line: %d, col:%d", node, pos[0], pos[1]))
	}
	return nil
}

func (c *Compiler) GetSpec(id []string) ([]string, *spec) {
	// Returns full namespace of variable and the spec it belongs to
	// assumes current spec if none specified
	if c.specs[id[0]] == nil {
		id = append([]string{c.currentSpecName}, id...)
	}
	return id, c.specs[id[0]]
}

func (c *Compiler) ListSpecs() []string {
	// Lists all specs the compiler knows about
	var specs []string
	for k := range c.specs {
		specs = append(specs, k)
	}
	return specs
}

func (c *Compiler) ListSpecsAndVars() map[string][]string {
	// Lists all specs and their variables
	specs := make(map[string][]string)
	for k, v := range c.specs {
		specs[k] = v.vars.List()
	}
	return specs
}

func (c *Compiler) setConst(id []string, val value.Value) {
	if c.isVarSet(id) {
		fid := strings.Join(id, "_")
		panic(fmt.Sprintf("variable %s is a constant and cannot be reassigned", fid))
	}
	c.specs[c.currentSpecName].DefineSpecVar(id, val)
}

func (c *Compiler) isVarSet(id []string) bool {
	id, s := c.GetSpec(id)
	return s.GetSpecVar(id) != nil
}

func (c *Compiler) isVarSetAssert(id []string) bool {
	//If this is for an assert the var might reference
	//a rule on the struct level
	if c.isVarSet(id) || c.isInstance(id) {
		return true
	}
	return false
}

func (c *Compiler) isInstance(id []string) bool {
	for _, v := range c.instances {
		if v == id[0] {
			return true
		}
	}
	return false
}

func (c *Compiler) isFunction(node ast.Node) bool {
	switch node.(type) {
	case *ast.FunctionLiteral:
		return true
	case *ast.BlockStatement:
		return true
	default:
		return false
	}
}

func (c *Compiler) generateParameters(data map[string]ast.Node, id []string) []*ir.Param {
	var p []*ir.Param
	keys := c.generateOrder(data)
	for _, k := range keys {
		switch n := data[k].(type) {
		case *ast.Instance:
			var ip []*ir.Param
			if n.Complex {
				ip = c.generateParameters(c.specStructs[n.Value.Spec][n.Value.Value], append(id, k))
			} else {
				ip = c.generateParameters(c.specStructs[n.Value.Spec][n.Value.Value], []string{id[0], id[1], k})
			}
			p = append(p, ip...)
		default:
			if !c.isFunction(n) {
				pid := append(id, k)
				name := c.getVariableName(c.getFullVariableName(pid))
				p = append(p, ir.NewParam(name, DoubleP))
			}
		}
	}
	return p
}

func (c *Compiler) resetParaState(p []*ir.Param) {
	for i := 0; i < len(p); i++ {
		name := p[i].LocalName
		id := strings.Split(name, "_")
		id, s := c.GetSpec(id)
		s.vars.ResetState(id)
	}
}

func (c *Compiler) convertAssertVariables(ex ast.Expression) ast.Expression {
	switch e := ex.(type) {
	case *ast.InfixExpression:

		e.Left = c.convertAssertVariables(e.Left)
		e.Right = c.convertAssertVariables(e.Right)
		return e
	case *ast.Identifier:
		id := strings.Split(e.Value, "_")
		pos := e.Position()
		fvn := c.getFullVariableName(id)
		fvns := c.getVariableName(fvn)

		if !c.isVarSetAssert(fvn) {
			panic(fmt.Sprintf("cannot send value to variable %s. Variable not defined line: %d, col: %d", fvns, pos[0], pos[1]))
		}

		instas := c.fetchInstances(fvn)
		if len(instas) == 0 {
			instas = []string{fvns}
		}
		return &ast.AssertVar{
			Token:        e.Token,
			InferredType: e.InferredType,
			Instances:    instas,
		}
	case *ast.ParameterCall:
		id := e.Value
		pos := e.Position()
		fvn := c.getFullVariableName(id)
		fvns := c.getVariableName(fvn)

		if !c.isVarSetAssert(fvn) {
			panic(fmt.Sprintf("cannot send value to variable %s. Variable not defined line: %d, col: %d", fvns, pos[0], pos[1]))
		}

		instas := c.fetchInstances(fvn)
		if len(instas) == 0 {
			instas = []string{fvns}
		}
		return &ast.AssertVar{
			Token:        e.Token,
			InferredType: e.InferredType,
			Instances:    instas,
		}

	case *ast.AssertVar:
		return e
	case *ast.IntegerLiteral:
		return e
	case *ast.FloatLiteral:
		return e
	case *ast.Boolean:
		return e
	case *ast.StringLiteral:
		return e
	case *ast.Natural:
		return e
	case *ast.Uncertain:
		return e
	case *ast.Unknown:
		return e
	case *ast.PrefixExpression:
		e.Right = c.convertAssertVariables(e.Right)
		return e
	case *ast.Nil:
		return e
	case *ast.IndexExpression:
		e.Left = c.convertAssertVariables(e.Left)
		return e
	default:
		pos := e.Position()
		panic(fmt.Sprintf("illegal node %T in assert or assume line: %d, col: %d", e, pos[0], pos[1]))
	}
}

func (c *Compiler) lookupIdent(ident []string, pos []int) *ir.InstLoad {
	id := c.getFullVariableName(ident)
	id, s := c.GetSpec(id)
	var name string

	local := s.GetSpecVar(id)
	if local != nil {
		//name = c.getVariableStateName(id)
		name = c.getVariableName(id)
		pointer := s.GetSpecVarPointer(name)
		load := c.contextBlock.NewLoad(irtypes.Double, pointer)
		return load
	} else {

		// Might be a spec global constant
		g := id[len(id)-1]
		global := s.GetSpecVar([]string{id[0], g})
		if global == nil {
			panic(fmt.Sprintf("variable %s not defined line: %d col: %d", strings.Join(id, "_"), pos[0], pos[1]))
		}
		name = strings.Join([]string{id[0], g}, "_")

	}
	pointer := c.specGlobals[name]
	load := c.contextBlock.NewLoad(irtypes.Double, pointer)
	return load
}

func (c *Compiler) fetchInstances(ident []string) []string {
	// this and convertAssertVariables need a rethink. Seems brittle
	// with lots of edge cases
	var insta []string
	id, _ := c.GetSpec(ident)
	for k, v := range c.instanceChildren {
		if v == id[1] {
			id2 := strings.Split(k, "_")
			if id2[len(id2)-1] == id[len(id)-1] { //Same parameter of a different instance
				insta = append(insta, k)
			}
		}
	}
	return insta
}

func negate(e ast.Expression) ast.Expression {
	//Negate the expression so that the solver attempts to disprove it
	switch n := e.(type) {
	case *ast.InfixExpression:
		op, ok := OP_NEGATE[n.Operator]
		if ok {
			//pos := n.Position()
			//panic(fmt.Sprintf("operator %s not valid from an assertion. line: %d, col: %d", n.Operator, pos[0], pos[1]))
			n.Operator = op
		}
		n.Left = negate(n.Left)
		n.Right = negate(n.Right)

		node := evaluate(n) // If Int/Float, evaluate and return the value
		return node
	case *ast.Boolean:
		if n.Value {
			n.Value = false
		} else {
			n.Value = true
		}
		return n
	case *ast.PrefixExpression:
		return negate(n.Right)
	}
	return e
}

func negateTemporal(op string, n int) (string, int) {
	var op2 string
	var n2 int
	switch op {
	case "nmt":
		op2 = "nft"
		n2 = n + 1
	case "nft":
		op2 = "nmt"
		n2 = n - 1
	}
	return op2, n2
}

func evaluate(n *ast.InfixExpression) ast.Expression {
	if util.IsCompare(n.Operator) {
		return n
	}
	f1, ok1 := n.Left.(*ast.FloatLiteral)
	i1, ok2 := n.Left.(*ast.IntegerLiteral)

	if !ok1 && !ok2 {
		return n
	}

	f2, ok1 := n.Right.(*ast.FloatLiteral)
	i2, ok2 := n.Right.(*ast.IntegerLiteral)

	if !ok1 && !ok2 {
		return n
	}

	if f1 != nil {
		if f2 != nil {
			v := evalFloat(f1.Value, f2.Value, n.Operator)
			return &ast.FloatLiteral{
				Token: n.Token,
				Value: v,
			}
		} else {
			v := evalFloat(f1.Value, float64(i2.Value), n.Operator)
			return &ast.FloatLiteral{
				Token: n.Token,
				Value: v,
			}
		}
	} else {
		if f2 != nil {
			v := evalFloat(float64(i1.Value), f2.Value, n.Operator)
			return &ast.FloatLiteral{
				Token: n.Token,
				Value: v,
			}
		} else {
			if n.Operator == "/" {
				//Return a float in the case of division
				v := evalFloat(float64(i1.Value), float64(i2.Value), n.Operator)
				return &ast.FloatLiteral{
					Token: n.Token,
					Value: v,
				}
			}
			v := evalInt(i1.Value, i2.Value, n.Operator)
			return &ast.IntegerLiteral{
				Token: n.Token,
				Value: v,
			}
		}
	}
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

type Panic string
