package llvm

import (
	"errors"
	"fault/ast"
	"fault/llvm/name"
	"fault/types"
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

type Compiler struct {
	module *ir.Module

	specs       map[string]*spec
	instances   map[string]string
	currentSpec *spec

	currentSpecName string
	currScope       string

	specStructs   map[string]types.StockFlow
	specFunctions map[string]value.Value

	contextFuncName string
	contextMetadata *metadata.Attachment

	alloc        bool
	runRound     int16
	funcStatePos map[string]map[string]interface{}

	// Stack of return values pointers, is used both used if a function returns more
	// than one value (arg pointers), and single stack based returns
	contextFuncRetVals [][]value.Value

	contextBlock *ir.Block
	contextFunc  *ir.Func

	// Stack of variables that are in scope
	contextBlockVariables []map[string]value.Value
	allocatedPointers     []map[string]*ir.InstAlloca

	// What a break or continue should resolve to
	contextLoopBreak    []*ir.Block
	contextLoopContinue []*ir.Block

	// Where a condition should jump when done
	contextCondAfter []*ir.Block

	// What type the current assign operation is assigning to.
	// Is used when evaluating what type an integer constant should be.
	contextAssignDest []value.Value

	// Stack of Alloc instructions
	// Is used to decide if values should be stack or heap allocated
	//contextAlloc []*parser.AllocNode

	specGlobals map[string]*ir.Global
}

func NewCompiler(structs map[string]types.StockFlow) *Compiler {
	c := &Compiler{
		module: ir.NewModule(),

		specs:         make(map[string]*spec),
		instances:     make(map[string]string),
		specStructs:   structs,
		specFunctions: make(map[string]value.Value),

		contextFuncRetVals: make([][]value.Value, 0),
		contextMetadata:    nil,
		alloc:              true,

		funcStatePos: make(map[string]map[string]interface{}),

		contextBlockVariables: make([]map[string]value.Value, 0),
		allocatedPointers:     make([]map[string]*ir.InstAlloca, 0),

		contextLoopBreak:    make([]*ir.Block, 0),
		contextLoopContinue: make([]*ir.Block, 0),
		contextCondAfter:    make([]*ir.Block, 0),

		contextAssignDest: make([]value.Value, 0),

		specGlobals: make(map[string]*ir.Global),
		runRound:    0,
	}
	c.addGlobal()
	c.pushVariablesStack()
	c.pushAllocations()
	return c
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
	return
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
		id := node.Left.(*ast.ParameterCall).Value
		pos := node.Left.(*ast.ParameterCall).Position()
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
	case *ast.Uncertain:
		return constant.NewStruct(&irtypes.StructType{},
			constant.NewFloat(irtypes.Double, v.Mean),
			constant.NewFloat(irtypes.Double, v.Sigma))
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

func (c *Compiler) compile(node ast.Node) {
	switch v := node.(type) {
	case *ast.SpecDeclStatement:
		break
	case *ast.ImportStatement:
		parent := c.currentSpecName
		parentSp := c.currentSpec
		err := c.Compile(v.Tree)
		if err != nil {
			panic(err)
		}
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

	case *ast.AssertionStatement:

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
	// InstanceExpression
	// IndexExpression <-- Is this still used?

}

func (c *Compiler) compileStruct(def *ast.DefStatement) {
	//Not implemented, using preparse from type checker
}

func (c *Compiler) generateOrder(pairs map[string]ast.Node) []string {
	keys := []string{}
	for k, _ := range pairs {
		keys = append(keys, k)
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
		id := c.getFullVariableName([]string{instName, k})
		switch pv := c.specStructs[base.Value.Spec][structName][k].(type) {
		case *ast.Instance:
			c.compileInstance(pv, k) // Copy instance data over
		case *ast.FunctionLiteral:
			c.compileFunction(pv)
		case *ast.InfixExpression:
			c.compileInfix(pv)
		case *ast.BlockStatement:
			c.compileBlock(pv)
		default:
			val := c.compileValue(c.specStructs[base.Value.Spec][structName][k])
			id, s := c.GetSpec(id)
			s.DefineSpecVar(id, val)
			c.allocVariable(id, val, pos)
			s.vars.ResetState(id)
			//name := c.getVariableStateName(id)
			name := c.getVariableName(id)
			p := ir.NewParam(name, DoubleP)
			s.AddParam(id, p)
		}
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

// func (c *Compiler) compileBranch(b *ast.BlockStatement) *ir.Block {
// 	c.contextBlock = c.contextFunc.NewBlock(name.Block())
// 	val := c.compileBlock(b)
// 	c.contextBlock.NewRet(val)
// 	child := c.contextBlock
// 	//c.contextBlock = c.contextFunc.NewBlock(name.Block())
// 	return child
//}

func (c *Compiler) compileParameterCall(pc *ast.ParameterCall) value.Value {
	structName := c.instances[pc.Value[0]]
	id := c.getFullVariableName(pc.Value)
	id, s := c.GetSpec(id)
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
	for k, _ := range c.specs {
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
			ip := c.generateParameters(c.specStructs[n.Value.Spec][n.Value.Value], []string{id[0], id[1], k})
			p = append(p, ip...)
		default:
			if !c.isFunction(n) {
				pid := append(id, k)
				//name := c.getVariableStateName(c.getFullVariableName(pid))
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
		id[len(id)-1] = id[len(id)-1][:len(id[len(id)-1])-1]
		s.vars.ResetState(id)
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

func (c *Compiler) pushVariablesStack() {
	c.contextBlockVariables = append(c.contextBlockVariables, make(map[string]value.Value))
}

func (c *Compiler) popVariablesStack() {
	c.contextBlockVariables = c.contextBlockVariables[0 : len(c.contextBlockVariables)-1]
}

func (c *Compiler) pushAllocations() {
	c.allocatedPointers = append(c.allocatedPointers, make(map[string]*ir.InstAlloca))
}

func (c *Compiler) popAllocations() {
	c.allocatedPointers = c.allocatedPointers[0 : len(c.allocatedPointers)-1]
}

type Panic string

func compilePanic(message string) {
	panic(Panic(fmt.Sprintf("compile panic: %s\n", message)))
}
