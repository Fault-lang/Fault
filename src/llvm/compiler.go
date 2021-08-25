package llvm

import (
	"errors"
	"fault/ast"
	"fault/llvm/name"
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

type Compiler struct {
	module *ir.Module

	specs       map[string]*spec
	instances   map[string]string
	currentSpec *spec

	currentSpecName string
	currScope       string

	specStructs   map[string]map[string]ast.Node
	specFunctions map[string]value.Value

	contextFuncName string

	runBlock     *ir.Func
	runRound     int16
	funcStatePos map[string]map[string]interface{}

	// Stack of return values pointers, is used both used if a function returns more
	// than one value (arg pointers), and single stack based returns
	contextFuncRetVals [][]value.Value

	contextBlock *ir.Block

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

func NewCompiler(structs map[string]map[string]ast.Node) *Compiler {
	c := &Compiler{
		module: ir.NewModule(),

		specs:         make(map[string]*spec),
		instances:     make(map[string]string),
		specStructs:   structs,
		specFunctions: make(map[string]value.Value),

		contextFuncRetVals: make([][]value.Value, 0),

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
	fmt.Println(c.GetIR())
	return
}

func (c *Compiler) compileIdent(node *ast.Identifier) *ir.InstLoad {
	id := c.getFullVariableName([]string{node.Value})
	id, s := c.GetSpec(id)
	var name string

	local := s.GetSpecVar(id)
	if local != nil {
		name = c.getVariableStateName(id)
		pointer := s.GetSpecVarPointer(name)
		return c.contextBlock.NewLoad(irtypes.Double, pointer)
	} else {

		// Might be a spec global constant
		g := id[len(id)-1]
		global := s.GetSpecVar([]string{id[0], g})

		if global == nil {
			pos := node.Position()
			panic(fmt.Sprintf("variable %s not defined line: %d col: %d", strings.Join(id, "_"), pos[0], pos[1]))
		}
		name = strings.Join([]string{id[0], g}, "_")

	}
	pointer := c.specGlobals[name]
	return c.contextBlock.NewLoad(irtypes.Double, pointer)
}

func (c *Compiler) compileInfix(node *ast.InfixExpression) value.Value {
	pos := node.Position()

	switch node.Operator {
	case "=": // Used to store temporary local values
		r := c.compileValue(node.Right)
		if _, ok := node.Right.(*ast.Instance); !ok { // If declaring a new instance don't save
			fvn := c.getFullVariableName([]string{node.Left.(*ast.Identifier).Value})
			//fvns := c.getVariableStateName(fvn)
			if c.isVarSet(fvn) {
				p := c.fetchAllocation(fvn)
				c.contextBlock.NewStore(r, p)
				return nil
			}
			id, s := c.GetSpec(fvn)
			s.DefineSpecVar(id, r)
			c.allocVariable(id, r, node.Left.Position())
		}
		return nil
	case "<-":
		r := c.compileValue(node.Right)
		id := node.Left.(*ast.ParameterCall).Value
		pos := node.Left.(*ast.ParameterCall).Position()
		fvn := c.getFullVariableName(id)
		fvns := c.getVariableStateName(fvn)

		if !c.isVarSet(fvn) {
			panic(fmt.Sprintf("cannot send value to variable %s. Variable not defined line: %d, col: %d", fvns, pos[0], pos[1]))
		}

		id, s := c.GetSpec(fvn)
		s.DefineSpecVar(id, r)
		c.allocVariable(fvn, r, pos)
		return nil
	case "+":
		l := c.compileValue(node.Left)
		r := c.compileValue(node.Right)

		return c.contextBlock.NewFAdd(l, r)
	case "-":
		l := c.compileValue(node.Left)
		r := c.compileValue(node.Right)

		return c.contextBlock.NewFSub(l, r)
	case "*":
		l := c.compileValue(node.Left)
		r := c.compileValue(node.Right)

		return c.contextBlock.NewFMul(l, r)
	case "/":
		l := c.compileValue(node.Left)
		r := c.compileValue(node.Right)

		return c.contextBlock.NewFDiv(l, r)
	case "%":
		l := c.compileValue(node.Left)
		r := c.compileValue(node.Right)

		return c.contextBlock.NewFRem(l, r)
	case ">":
		l := c.compileValue(node.Left)
		r := c.compileValue(node.Right)

		return c.contextBlock.NewFCmp(enum.FPredOGT, l, r)
	case ">=":
		l := c.compileValue(node.Left)
		r := c.compileValue(node.Right)

		return c.contextBlock.NewFCmp(enum.FPredOGE, l, r)
	case "<":
		l := c.compileValue(node.Left)
		r := c.compileValue(node.Right)

		return c.contextBlock.NewFCmp(enum.FPredOLT, l, r)
	case "<=":
		l := c.compileValue(node.Left)
		r := c.compileValue(node.Right)

		return c.contextBlock.NewFCmp(enum.FPredOLE, l, r)
	case "==":
		l := c.compileValue(node.Left)
		r := c.compileValue(node.Right)

		return c.contextBlock.NewFCmp(enum.FPredOEQ, l, r)
	case "!=":
		l := c.compileValue(node.Left)
		r := c.compileValue(node.Right)

		return c.contextBlock.NewFCmp(enum.FPredONE, l, r)
	case "|":
		l := c.compileValue(node.Left)
		r := c.compileValue(node.Right)

		if c.contextFuncName != "__run" {
			panic(fmt.Sprintf("cannot use parallel operator outside of the run block. line: %d, col: %d", pos[0], pos[1]))
		}
		switch node.Left.(type) {
		case *ast.Identifier:
			block := c.runBlock.NewBlock(name.Block())
			l_func := block.NewCall(l)

			// Kind of a BS approach but fine for right now
			md := &metadata.Attachment{
				Name: "parallel",
				Node: nil,
			}
			l_func.Metadata = append(l_func.Metadata, md)
		}

		switch node.Right.(type) {
		case *ast.Identifier:
			block := c.runBlock.NewBlock(name.Block())
			r_func := block.NewCall(r)
			md := &metadata.Attachment{
				Name: "parallel",
				Node: nil,
			}
			r_func.Metadata = append(r_func.Metadata, md)
		}
		return nil
	default:
		panic(fmt.Sprintf("unknown operator %s. line: %d, col: %d", node.Operator, pos[0], pos[1]))
	}
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
		c.compileInstance(v.Value.Value, v.Name, v.Position())
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
	body := node.Statements
	var ret value.Value
	for i := 0; i < len(body); i++ {
		switch exp := body[i].(type) {
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
	c.runBlock = c.module.NewFunc("__run", irtypes.Void)
	mainBlock := c.runBlock.NewBlock(name.Block())
	mainBlock.NewRet(nil)
	c.contextBlock = mainBlock
}

func (c *Compiler) compile(node ast.Node) {
	switch v := node.(type) {
	case *ast.SpecDeclStatement:
		break
	case *ast.ImportStatement:

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

func (c *Compiler) compileInstance(structName string, instName string, pos []int) {
	if c.runRound > 0 { // Initialize things only once
		return
	}
	if c.currScope == "" { //For example we're in the run block
		c.currScope = instName
	}
	parentFunction := c.contextFuncName
	c.contextFuncName = instName
	if c.specStructs[structName] == nil {
		panic(fmt.Sprintf("no stock or flow named %s, line: %d, col %d", structName, pos[0], pos[1]))
	}
	keys := c.generateOrder(c.specStructs[structName])
	for _, k := range keys {
		id := c.getFullVariableName([]string{instName, k})
		switch pv := c.specStructs[structName][k].(type) {
		case *ast.Instance:
			c.compileInstance(pv.Value.Value, k, pos) // Copy instance data over
		case *ast.FunctionLiteral:
			c.compileFunction(pv)
		case *ast.InfixExpression:
			c.compileInfix(pv)
		case *ast.BlockStatement:
			c.compileBlock(pv)
		default:
			val := c.compileValue(c.specStructs[structName][k])
			id, s := c.GetSpec(id)
			s.DefineSpecVar(id, val)
			c.allocVariable(id, val, pos)
		}
	}
	c.instances[instName] = structName
	c.contextFuncName = parentFunction
	if c.currScope == instName { //Again initiating in the run block
		c.currScope = ""
	}
}

func (c *Compiler) compileParameterCall(pc *ast.ParameterCall) value.Value {
	if c.currScope == "" {
		c.currScope = pc.Value[0]
	}
	parentFunction := c.contextFuncName
	c.contextFuncName = pc.Value[0]
	structName := c.instances[pc.Value[0]]
	id := c.getFullVariableName(pc.Value)
	val := c.compileValue(c.specStructs[structName][pc.Value[1]])
	id, s := c.GetSpec(id)
	if val != nil || !c.isFunction(c.specStructs[structName][pc.Value[1]]) {
		s.DefineSpecVar(id, val)
		c.allocVariable(id, val, pc.Position())
	}
	if c.currScope == pc.Value[0] {
		c.currScope = ""
	}
	c.contextFuncName = parentFunction
	return val
}

func (c *Compiler) compileFunction(node *ast.FunctionLiteral) value.Value {
	fn := c.module.NewFunc(c.contextFuncName, irtypes.Double) //Change this to match type
	c.contextBlock = fn.NewBlock(name.Block())
	c.pushVariablesStack()
	c.pushAllocations()

	body := node.Body.Statements
	var retValue value.Value
	for i := 0; i < len(body); i++ {
		exp := body[i].(*ast.ExpressionStatement).Expression
		init, ok := exp.(*ast.InitExpression)
		if ok {
			return c.compileValue(init.Expression)
		}
	}
	retValue = constant.NewFloat(irtypes.Double, 0.0)
	c.contextBlock.NewRet(retValue)
	c.specFunctions[c.contextFuncName] = fn
	c.contextBlock = c.runBlock.NewBlock(name.Block())
	c.contextBlock.NewRet(retValue)
	c.popVariablesStack()
	c.popAllocations()
	return retValue
}

func (c *Compiler) compileFunctionBody(node ast.Expression) value.Value {
	switch v := node.(type) {
	case *ast.InfixExpression:
		return c.compileInfix(v)

	case *ast.PrefixExpression:

	case *ast.IfExpression:

	case *ast.Instance:
		c.compileInstance(v.Name, v.Value.Value, v.Position())

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
		fmt.Printf("%T\n", node)
		return false
	}
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
