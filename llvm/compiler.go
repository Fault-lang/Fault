package llvm

import (
	"errors"
	"fault/ast"
	"fault/llvm/name"
	"fault/preprocess"
	"fault/util"
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/metadata"
	irtypes "github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

var DoubleP = &irtypes.PointerType{ElemType: irtypes.Double}
var I1P = &irtypes.PointerType{ElemType: irtypes.I1}

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

	runRound int16

	currentSpec      string
	specs            map[string]*spec
	instances        map[string][]string
	instanceChildren map[string]string
	structPropOrder  map[string][]string

	contextBlock    *ir.Block
	contextFunc     *ir.Func
	contextFuncName string
	contextMetadata *metadata.Attachment

	// Stack of variables that are in scope
	alloc             bool
	allocatedPointers []map[string]*ir.InstAlloca

	// Where a condition should jump when done
	contextCondAfter []*ir.Block

	builtIns        map[string]*ir.Func
	specStructs     map[string]*preprocess.SpecRecord
	specFunctions   map[string]value.Value
	specGlobals     map[string]*ir.Global
	RawAsserts      []*ast.AssertionStatement
	RawAssumes      []*ast.AssumptionStatement
	Asserts         []*ast.AssertionStatement
	Assumes         []*ast.AssumptionStatement
	Uncertains      map[string][]float64
	Unknowns        []string
	Components      map[string]map[string]string
	ComponentStarts map[string]string
}

func NewCompiler() *Compiler {
	c := &Compiler{
		module: ir.NewModule(),
		specs:  make(map[string]*spec),

		alloc:             true,
		allocatedPointers: make([]map[string]*ir.InstAlloca, 0),

		contextCondAfter: make([]*ir.Block, 0),
		instances:        make(map[string][]string),
		instanceChildren: make(map[string]string),
		structPropOrder:  make(map[string][]string),

		runRound:        0,
		builtIns:        make(map[string]*ir.Func),
		specStructs:     make(map[string]*preprocess.SpecRecord),
		specFunctions:   make(map[string]value.Value),
		specGlobals:     make(map[string]*ir.Global),
		Uncertains:      make(map[string][]float64),
		Components:      make(map[string]map[string]string),
		ComponentStarts: make(map[string]string),
	}
	c.addGlobal()
	return c
}

func (c *Compiler) LoadMeta(structs map[string]*preprocess.SpecRecord, uncertains map[string][]float64, unknowns []string) {
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

	switch decl := specfile.Statements[0].(type) {
	case *ast.SpecDeclStatement:
		c.currentSpec = decl.Name.Value
	case *ast.SysDeclStatement:
		c.currentSpec = decl.Name.Value
	default:
		panic(fmt.Sprintf("spec file improperly formatted. Missing spec declaration, got %T", specfile.Statements[0]))
	}

	c.specs[c.currentSpec] = NewCompiledSpec(c.currentSpec)
	for _, fileNode := range specfile.Statements {
		c.compile(fileNode)
	}

	for _, assert := range c.RawAsserts {
		c.compileAssert(assert)
	}
	for _, assert := range c.RawAssumes {
		c.compileAssert(assert)
	}

	return c.Asserts, c.Assumes
}

func (c *Compiler) compile(node ast.Node) {
	switch v := node.(type) {
	case *ast.SpecDeclStatement:
		break
	case *ast.SysDeclStatement:
		break
	case *ast.ImportStatement:
		parent := c.currentSpec
		asserts, assumes := c.processSpec(v.Tree, true) //Move all asserts to the end of the compilation process
		c.Asserts = append(c.Asserts, asserts...)
		c.Assumes = append(c.Assumes, assumes...)
		c.currentSpec = parent
	case *ast.ConstantStatement:
		c.compileConstant(v)
	case *ast.DefStatement:
		c.compileStruct(v)

	case *ast.FunctionLiteral:

	case *ast.InfixExpression:
		c.compileInfix(v)

	case *ast.PrefixExpression:
		c.compilePrefix(v)

	case *ast.AssumptionStatement:
		// Need to do these after the run block so we move them
		c.RawAssumes = append(c.RawAssumes, v)

	case *ast.AssertionStatement:
		c.RawAsserts = append(c.RawAsserts, v)

	case *ast.ForStatement:
		c.contextFuncName = "__run"
		for i := int64(0); i < v.Rounds.Value; i++ {
			c.compileBlock(v.Body)
			c.runRound = c.runRound + 1
		}
		c.contextFuncName = ""

	case *ast.StartStatement:
		for _, p := range v.Pairs {
			branch, err := c.specStructs[c.currentSpec].FetchComponent(p[0])
			if err != nil {
				panic(err)
			}
			id := []string{c.currentSpec, p[0], p[1]}
			c.processFunc(id, branch)
			c.ComponentStarts[p[0]] = p[1]
		}

	default:
		pos := node.Position()
		panic(fmt.Sprintf("node type %T unimplemented line: %d col: %d", v, pos[0], pos[1]))
	}
}

func (c *Compiler) compileConstant(node *ast.ConstantStatement) {
	value := c.compileValue(node.Value)
	id := []string{c.currentSpec, node.Name.Value}
	c.setConst(id, value)
	c.globalVariable(id, value, node.Position())
}

func (c *Compiler) setConst(rawid []string, val value.Value) {
	c.specs[c.currentSpec].DefineSpecType(rawid, val.Type())
}

func (c *Compiler) compileStruct(def *ast.DefStatement) {
	id := def.Name.Id()
	key := strings.Join(id, "_")
	switch def.Type() {
	case "FLOW":
		c.instances[key] = []string{key}
		c.structPropOrder[key] = def.Value.(*ast.FlowLiteral).Order
	case "STOCK":
		c.instances[key] = []string{key}
		c.structPropOrder[key] = def.Value.(*ast.StockLiteral).Order
	case "GLOBAL":
		c.instances[key] = []string{key}
		//instance, _ := def.Value.(*ast.Instance)
		//c.compileInstance(instance, strings.Join(id, "_"))
	case "COMPONENT":
		c.instances[key] = []string{key}
		c.structPropOrder[key] = def.Value.(*ast.ComponentLiteral).Order
		c.compileComponent(def.Value.(*ast.ComponentLiteral), key)
	}
}

func (c *Compiler) compileValue(node ast.Node) value.Value {
	if node == nil {
		panic("value received by compileValue is nil")
	}
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
	case *ast.PrefixExpression:
		return c.compilePrefix(v)
	case *ast.FunctionLiteral:
		return c.compileFunction(v)
	case *ast.Instance:
		
	case *ast.StructInstance:
		c.compileInstance(v)
	case *ast.ParameterCall:
		return c.compileParameterCall(v)
	case *ast.BlockStatement:
		return c.compileBlock(v)
	case *ast.This:
		return c.compileThis(v)
	default:
		pos := v.Position()
		panic(fmt.Sprintf("unknown value type %T line: %d col: %d", v, pos[0], pos[1]))
	}
	return nil
}

func (c *Compiler) compileInstance(node *ast.StructInstance) {
	if c.runRound > 0 { // Initialize things only once
		return
	}
	if c.contextFuncName == "__run" {
		c.alloc = false
	}
	parentFunction := c.contextFuncName
	c.contextFuncName = node.IdString()

	id := node.Id()
	parent := strings.Join(node.Parent, "_")
	pos := node.Position()
	children := make(map[string]string)

	switch node.Type() {
	case "STOCK":
		children = c.processStruct(node)
	case "FLOW":
		children = c.processStruct(node)
	default:
		panic(fmt.Sprintf("no stock or flow named %s, line: %d, col %d", id, pos[0], pos[1]))
	}
	key := strings.Join(id, "_")
	c.structPropOrder[key] = node.Order
	c.instances[key] = append(c.instances[key], parent)
	c.instances[parent] = append(c.instances[parent], key)
	c.instanceChildren = util.MergeStringMaps(c.instanceChildren, children)

	c.contextFuncName = parentFunction
	if c.contextFuncName == "__run" {
		c.alloc = true
	}
}

func (c *Compiler) compileComponent(node *ast.ComponentLiteral, cname string) {
	for key, p := range node.Pairs {
		var pname string
		scopeName := []string{cname, key.Value}

		oldBlock := c.contextBlock

		c.contextFuncName = strings.Join(scopeName, "_")

		switch v := p.(type) {
		case *ast.StateLiteral:
			id := []string{key.Spec, cname}
			parentID := strings.Join(node.Id(), "_")
			k := strings.Join(id, "_")
			c.structPropOrder[k] = c.structPropOrder[parentID]

			branches := util.Preparse(node.Pairs)
			params := c.generateParameters([]string{key.Spec, cname}, branches)
			c.resetParaState(params)
			f := c.module.NewFunc(key.Value, irtypes.Void, params...)
			c.contextFunc = f
			pname = name.Block()
			c.contextBlock = f.NewBlock(pname)
			if c.Components[cname] != nil {
				c.Components[cname][key.Value] = pname
			} else {
				c.Components[cname] = map[string]string{key.Value: pname}
			}
			val := c.compileBlock(v.Body)
			c.contextBlock.NewRet(val)
			c.contextBlock = oldBlock
			c.contextFuncName = "__run"
			c.contextFunc = nil
		case *ast.Instance:
		
		default:
			val := c.compileValue(v)

			if val != nil {
				id := []string{c.currentSpec, cname, key.Value}
				s := c.specs[id[0]]
				if s.GetSpecVar(id) != nil {
					vname := strings.Join(id, "_")
					pointer := s.GetSpecVarPointer(id)
					ty := s.GetSpecType(vname)
					c.contextBlock.NewLoad(ty, pointer)
				} else {
					s.DefineSpecType(id, val.Type())
					s.DefineSpecVar(id, val)
					c.allocVariable(id, val, []int{0, 0, 0, 0})
				}
			}
		}
	}
}

func (c *Compiler) compileParameterCall(pc *ast.ParameterCall) value.Value {
	var err error
	id := pc.RawId()
	spec := c.specStructs[id[0]]
	ty, _ := spec.GetStructType(id[0:2]) //Removing the key
	st := id[len(id)-2]                  //The struct is the second to last item
	key := id[len(id)-1]

	var branches map[string]ast.Node
	switch ty {
	case "FLOW":
		branches, err = spec.FetchFlow(st)
		if err != nil {
			panic(err)
		}
	case "STOCK":
		branches, err = spec.FetchStock(st)
		if err != nil {
			panic(err)
		}
	default:
		panic(fmt.Sprintf("struct %s not found", id))
	}

	if c.contextFuncName == "__run" &&
		c.isFunction(branches[key]) {
		return c.processFunc(id, branches)
	}

	// Otherwise inline the parameter...
	// if c.currScope[0] == "" {
	// 	c.currScope = []string{pc.Value[0]}
	// }
	parentFunction := c.contextFuncName
	c.contextFuncName = pc.Value[0]

	val := c.compileValue(branches[key])

	// If there's no value, there's nothing to store
	if val != nil || !c.isFunction(branches[key]) {
		s := c.specs[id[0]]
		if s.GetSpecVar(id) != nil {
			vname := strings.Join(id, "_")
			pointer := s.GetSpecVarPointer(id)
			ty := s.GetSpecType(vname)
			c.contextBlock.NewLoad(ty, pointer)
		} else {
			s.DefineSpecType(id, val.Type())
			s.DefineSpecVar(id, val)
			c.allocVariable(id, val, pc.Position())
		}
	}
	// if c.currScope[0] == pc.Value[0] {
	// 	c.currScope = []string{""}
	// }
	c.contextFuncName = parentFunction
	return val
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
			ret = c.compileFunction(exp)
		case *ast.ExpressionStatement:
			ret = c.compileFunction(exp.Expression)
		}
	}
	return ret
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
			id := node.Expressions[i].(*ast.ParameterCall).Id()
			s := c.specs[id[0]]
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

func (c *Compiler) compileFunction(node ast.Node) value.Value {
	if !c.alloc { //Short circuit this if just initializing
		return nil
	}

	switch v := node.(type) {
	case *ast.FunctionLiteral:
		body := v.Body.Statements
		var ret value.Value
		for i := 0; i < len(body); i++ {
			ret = c.compileFunction(body[i])
		}
		return ret
	case *ast.ExpressionStatement:
		return c.compileFunction(v.Expression)
	case *ast.InfixExpression:
		return c.compileInfix(v)

	case *ast.PrefixExpression:
		c.compilePrefix(v)

	case *ast.IfExpression:
		c.compileIf(v)

	case *ast.Instance:
	
	case *ast.StructInstance:
		c.compileInstance(v)

	case *ast.IndexExpression:

	case *ast.ParameterCall:
		return c.compileParameterCall(v)

	case *ast.BuiltIn:
		//Is this the first time we're seeing this builtin?
		if c.builtIns[v.Function] == nil {
			var param []*ir.Param
			for k := range v.Parameters {
				param = append(param, ir.NewParam(k, irtypes.NewPointer(irtypes.I8)))
			}
			oldBlock := c.contextBlock
			f := c.module.NewFunc(v.Function, irtypes.Void, param...)
			c.contextBlock = f.NewBlock(name.Block())

			c.contextBlock.NewRet(nil)
			c.contextBlock = oldBlock

			c.builtIns[v.Function] = f
		}
		var params []value.Value
		for _, v := range v.Parameters {
			l := uint64(len(v.String()))
			alloc := c.contextBlock.NewAlloca(irtypes.NewArray(l, irtypes.I8))
			c.contextBlock.NewStore(constant.NewCharArrayFromString(v.String()), alloc)
			//load := c.contextBlock.NewLoad(irtypes.I8, alloc)
			cast := c.contextBlock.NewBitCast(alloc, irtypes.I8Ptr)
			params = append(params, cast)
		}

		c.contextBlock.NewCall(c.builtIns[v.Function], params...)

	default:
		pos := node.Position()
		panic(fmt.Sprintf("invalid expression %T in function body. line: %d, col:%d", node, pos[0], pos[1]))
	}
	return nil
}

func (c *Compiler) compilePrefix(node *ast.PrefixExpression) value.Value {
	val := c.compileInfixNode(node.Right)
	switch node.Operator {
	case "!":
		return c.contextBlock.NewXor(val, constant.NewInt(irtypes.I1, 1))
	case "-":
		return c.contextBlock.NewFNeg(val)
	default:
		panic(fmt.Sprintf("unrecognized prefix operator %s", node.Operator))
	}
}

func (c *Compiler) compileInfix(node *ast.InfixExpression) value.Value {
	var s *spec
	var id []string
	pos := node.Position()
	switch node.Operator {
	case "=": // Used to store temporary local values
		if !c.validOperator(node, true) {
			panic(fmt.Sprintf("operator %s cannot be used on variables of type %s and %s", node.Operator, node.Left.Type(), node.Right.Type()))
		}
		r := c.compileValue(node.Right)
		if _, ok := node.Right.(*ast.Instance); !ok { // If declaring a new instance don't save
			switch n := node.Left.(type) {
			case *ast.Identifier:
				id = n.RawId()
			case *ast.ParameterCall:
				id = n.RawId()
			}

			s = c.specs[id[0]]

			if c.isVarSet(id) && c.alloc {
				p := s.GetSpecVarPointer(id)
				c.contextBlock.NewStore(r, p)
				return nil
			}

			if c.isConstant(id) {
				panic(fmt.Sprintf("variable %s is a constant and cannot be modified", id[len(id)-1]))
			}

			s.DefineSpecVar(id, r)
			s.DefineSpecType(id, r.Type())
			if c.alloc {
				c.allocVariable(id, r, node.Left.Position())
			}
		}
		return nil
	case "<-":
		if !c.validOperator(node, false) {
			panic(fmt.Sprintf("operator %s cannot be used on variables of type %s and %s", node.Operator, node.Left.Type(), node.Right.Type()))
		}

		r := c.compileValue(node.Right)
		n, ok := node.Left.(*ast.ParameterCall)
		if !ok {
			pos := node.Position()
			panic(fmt.Sprintf("cannot use <- or -> operator on a non-stock value col: %d, line: %d", pos[0], pos[1]))
		}

		pos := n.Position()
		id = n.RawId()

		s = c.specs[id[0]]

		if !c.isVarSet(id) {
			panic(fmt.Sprintf("cannot send value to variable %s. Variable not defined line: %d, col: %d", strings.Join(id, "_"), pos[0], pos[1]))
		}

		if c.isConstant(id) {
			panic(fmt.Sprintf("variable %s is a constant and cannot be modified", id[len(id)-1]))
		}

		pointer := s.GetSpecVarPointer(id)
		c.contextBlock.NewStore(r, pointer)
		return nil
	case "+":
		if !c.validOperator(node, false) {
			panic(fmt.Sprintf("operator %s cannot be used on variables of type %s and %s", node.Operator, node.Left.Type(), node.Right.Type()))
		}

		l := c.compileInfixNode(node.Left)
		r := c.compileInfixNode(node.Right)

		return c.contextBlock.NewFAdd(l, r)
	case "-":
		if !c.validOperator(node, false) {
			panic(fmt.Sprintf("operator %s cannot be used on variables of type %s and %s", node.Operator, node.Left.Type(), node.Right.Type()))
		}

		l := c.compileInfixNode(node.Left)
		r := c.compileInfixNode(node.Right)

		sub := c.contextBlock.NewFSub(l, r)
		return sub
	case "*":
		if !c.validOperator(node, false) {
			panic(fmt.Sprintf("operator %s cannot be used on variables of type %s and %s", node.Operator, node.Left.Type(), node.Right.Type()))
		}

		l := c.compileInfixNode(node.Left)
		r := c.compileInfixNode(node.Right)

		mul := c.contextBlock.NewFMul(l, r)
		return mul
	case "/":
		if !c.validOperator(node, false) {
			panic(fmt.Sprintf("operator %s cannot be used on variables of type %s and %s", node.Operator, node.Left.Type(), node.Right.Type()))
		}

		l := c.compileInfixNode(node.Left)
		r := c.compileInfixNode(node.Right)

		div := c.contextBlock.NewFDiv(l, r)
		return div
	case "%":
		if !c.validOperator(node, false) {
			panic(fmt.Sprintf("operator %s cannot be used on variables of type %s and %s", node.Operator, node.Left.Type(), node.Right.Type()))
		}

		l := c.compileInfixNode(node.Left)
		r := c.compileInfixNode(node.Right)

		rem := c.contextBlock.NewFRem(l, r)
		return rem
	case ">":
		if !c.validOperator(node, false) {
			panic(fmt.Sprintf("operator %s cannot be used on variables of type %s and %s", node.Operator, node.Left.Type(), node.Right.Type()))
		}

		l := c.compileInfixNode(node.Left)
		r := c.compileInfixNode(node.Right)

		ogt := c.contextBlock.NewFCmp(enum.FPredOGT, l, r)
		return ogt
	case ">=":
		if !c.validOperator(node, false) {
			panic(fmt.Sprintf("operator %s cannot be used on variables of type %s and %s", node.Operator, node.Left.Type(), node.Right.Type()))
		}

		l := c.compileInfixNode(node.Left)
		r := c.compileInfixNode(node.Right)

		oge := c.contextBlock.NewFCmp(enum.FPredOGE, l, r)
		return oge
	case "<":
		if !c.validOperator(node, false) {
			panic(fmt.Sprintf("operator %s cannot be used on variables of type %s and %s", node.Operator, node.Left.Type(), node.Right.Type()))
		}

		l := c.compileInfixNode(node.Left)
		r := c.compileInfixNode(node.Right)

		olt := c.contextBlock.NewFCmp(enum.FPredOLT, l, r)
		return olt
	case "<=":
		if !c.validOperator(node, false) {
			panic(fmt.Sprintf("operator %s cannot be used on variables of type %s and %s", node.Operator, node.Left.Type(), node.Right.Type()))
		}

		l := c.compileInfixNode(node.Left)
		r := c.compileInfixNode(node.Right)

		ole := c.contextBlock.NewFCmp(enum.FPredOLE, l, r)
		return ole
	case "==":
		if !c.validOperator(node, true) {
			panic(fmt.Sprintf("operator %s cannot be used on variables of type %s and %s", node.Operator, node.Left.Type(), node.Right.Type()))
		}

		l := c.compileInfixNode(node.Left)
		r := c.compileInfixNode(node.Right)

		if node.Right.Type() == "BOOL" {
			return c.contextBlock.NewICmp(enum.IPredEQ, l, r)
		} else {
			return c.contextBlock.NewFCmp(enum.FPredOEQ, l, r)
		}
	case "!=":
		if !c.validOperator(node, true) {
			panic(fmt.Sprintf("operator %s cannot be used on variables of type %s and %s", node.Operator, node.Left.Type(), node.Right.Type()))
		}

		l := c.compileInfixNode(node.Left)
		r := c.compileInfixNode(node.Right)

		if node.Right.Type() == "BOOL" {
			return c.contextBlock.NewICmp(enum.IPredNE, l, r)
		} else {
			return c.contextBlock.NewFCmp(enum.FPredONE, l, r)
		}
	default:
		panic(fmt.Sprintf("unknown operator %s. line: %d, col: %d", node.Operator, pos[0], pos[1]))
	}
}

func (c *Compiler) compileInfixNode(node ast.Node) value.Value {
	switch v := node.(type) {
	case *ast.ParameterCall:
		id := v.Id()
		return c.lookupIdent(id, node.Position())
	default:
		return c.compileValue(node)
	}
}

func (c *Compiler) compileIdent(node *ast.Identifier) *ir.InstLoad {
	return c.lookupIdent(node.Id(), node.Position())
}

func (c *Compiler) compileThis(node *ast.This) *ir.InstLoad {
	return c.lookupIdent(node.Id(), node.Position())
}

func (c *Compiler) compileIf(n *ast.IfExpression) {
	cond := c.compileConditional(n.Condition)

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
	} else {
		afterBlock.NewRet(nil)
	}
}

func (c *Compiler) compileConditional(n ast.Node) value.Value {
	// Reformat the conditional clause to accept
	// things like if a {} or if !a {} and replace them
	// with a == true or a == false
	switch conditional := n.(type) {
	case *ast.InfixExpression:
		return c.compileValue(conditional)
	case *ast.PrefixExpression:
		return c.compilePrefix(conditional)
	case *ast.Identifier:
		right := &ast.Boolean{
			Token:        conditional.Token,
			InferredType: conditional.InferredType,
			Value:        true,
		}
		n := &ast.InfixExpression{Token: conditional.Token,
			InferredType: conditional.InferredType,
			Left:         conditional,
			Operator:     "==",
			Right:        right}
		return c.compileValue(n)
	case *ast.ParameterCall:
		right := &ast.Boolean{
			Token:        conditional.Token,
			InferredType: conditional.InferredType,
			Value:        true,
		}
		n := &ast.InfixExpression{Token: conditional.Token,
			InferredType: conditional.InferredType,
			Left:         conditional,
			Operator:     "==",
			Right:        right}
		return c.compileValue(n)
	}
	return c.compileValue(n)
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

func (c *Compiler) convertAssertVariables(ex ast.Expression) ast.Expression {
	switch e := ex.(type) {
	case *ast.InfixExpression:

		e.Left = c.convertAssertVariables(e.Left)
		e.Right = c.convertAssertVariables(e.Right)
		return e
	case *ast.Identifier:
		id := e.RawId()
		pos := e.Position()
		vname := strings.Join(id, "_")

		if !c.isVarSetAssert(id) {
			panic(fmt.Sprintf("cannot send value to variable %s. Variable not defined line: %d, col: %d", vname, pos[0], pos[1]))
		}

		instas := c.fetchInstances(id)
		if len(instas) == 0 {
			instas = []string{vname}
		}
		return &ast.AssertVar{
			Token:        e.Token,
			InferredType: e.InferredType,
			Instances:    instas,
		}
	case *ast.ParameterCall:
		id := e.RawId()
		pos := e.Position()
		vname := strings.Join(id, "_")

		if !c.isVarSetAssert(id) {
			panic(fmt.Sprintf("cannot send value to variable %s. Variable not defined line: %d, col: %d", vname, pos[0], pos[1]))
		}

		instas := c.fetchInstances(id)
		if len(instas) == 0 {
			instas = []string{vname}
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

func (c *Compiler) lookupIdent(id []string, pos []int) *ir.InstLoad {
	s := c.specs[id[0]]
	vname := strings.Join(id, "_")
	local := s.GetSpecVar(id)
	if local != nil {
		pointer := s.GetSpecVarPointer(id)
		ty := s.GetSpecType(vname)
		load := c.contextBlock.NewLoad(ty, pointer)
		return load
	}

	pointer := c.specGlobals[vname]
	if pointer != nil {
		pointer := c.specGlobals[vname]
		ty := s.GetSpecType(vname)
		load := c.contextBlock.NewLoad(ty, pointer)
		return load
	}
	return nil
}

func (c *Compiler) processFunc(rawId []string, branch map[string]ast.Node) value.Value {
	fname := strings.Join(rawId, "_")

	if c.runRound == 0 { //initialize
		params := c.generateParameters(rawId, branch)
		f := c.module.NewFunc(fname, irtypes.Void, params...)
		c.contextFunc = f

		//oldScope := c.currScope
		oldBlock := c.contextBlock

		c.contextFuncName = fname
		//c.currScope = []string{id[1]} // NOT necessarily the same as structName
		c.contextBlock = f.NewBlock(name.Block())

		val := c.compileValue(branch[rawId[len(rawId)-1]])
		c.contextBlock.NewRet(val)

		c.contextBlock = oldBlock
		c.contextFuncName = "__run"
		//c.currScope = oldScope
		c.specFunctions[fname] = f
		c.contextFunc = nil
		c.resetParaState(params)
	}

	return c.specFunctions[fname]
}

func (c *Compiler) processStruct(node *ast.StructInstance) map[string]string {
	keys := node.Order
	tree := node.Properties
	pos := node.Position()
	parentId := node.Id()
	var s *spec
	children := make(map[string]string)
	var params []value.Value
	var funcs [][]string

	for _, k := range keys {
		var isUncertain []float64
		var isUnknown bool
		var id []string

		switch pv := tree[k].Value.(type) {
		case *ast.Instance:
		case *ast.StructInstance:
			c.compileInstance(pv)
			id = pv.Id()
			sInner := c.specs[id[0]]
			pInner := sInner.GetParams(id)
			params = append(params, pInner...)
		case *ast.FunctionLiteral:
			c.compileFunction(pv)
			id = pv.Id()
			funcs = append(funcs, id)
		default:
			if n, ok := pv.(*ast.Unknown); ok {
				isUnknown = true
				id = n.Id()
			} else if uncertain, ok2 := pv.(*ast.Uncertain); ok2 {
				isUncertain = []float64{uncertain.Mean, uncertain.Sigma}
				id = n.Id()
			} else if n, ok := pv.(*ast.IntegerLiteral); ok {
				id = n.Id()
			} else if n, ok := pv.(*ast.FloatLiteral); ok {
				id = n.Id()
			} else if n, ok := pv.(*ast.Boolean); ok {
				id = n.Id()
			} else if n, ok := pv.(*ast.StringLiteral); ok {
				id = n.Id()
			} else if n, ok := pv.(*ast.Identifier); ok {
				id = n.Id()
			} else if n, ok := pv.(*ast.ParameterCall); ok {
				id = n.Id()
			}

			val := c.compileValue(pv)
			s = c.specs[id[0]]
			s.DefineSpecVar(id, val)
			s.DefineSpecType(id, val.Type())
			c.allocVariable(id, val, pos)
			vname := strings.Join(id, "_")
			s.vars.ResetState(vname)
			ty := s.GetPointerType(vname)
			p := ir.NewParam(vname, ty)
			params = append(params, p)
			s.AddParam(parentId, p)

		}
		//Track properties of instances so that we can write
		// asserts on the struct and honor them for all instances
		vname := strings.Join(id, "_")
		if isUnknown {
			c.Unknowns = append(c.Unknowns, vname)
		}
		if isUncertain != nil {
			c.Uncertains[vname] = isUncertain
		}
		children[vname] = node.Parent[1]
	}

	//Add the params for all the functions
	s = c.specs[parentId[0]]
	for _, f := range funcs {
		if len(params) > 0 {
			s.AddParams(f, params)
		}
	}
	return children
}

func (c *Compiler) generateParameters(id []string, data map[string]ast.Node) []*ir.Param {
	var p []*ir.Param
	var s *spec

	var keys []string
	keys = c.fetchOrder(id)
	if len(keys) == 0 { //If no order if found (ie components) fall back to alphabetically
		keys = c.generateOrder(data)
	}

	sr := c.specStructs[id[0]]
	for _, k := range keys {
		switch n := data[k].(type) {
		case *ast.StructInstance:
			var ip []*ir.Param
			child := n.Id()
			strInst, err := sr.Fetch(child[1], n.Type())
			if err != nil {
				panic(err)
			}

			if n.Complex {
				ip = c.generateParameters(child, strInst)
			} else {
				ip = c.generateParameters(n.Id(), strInst)
			}
			p = append(p, ip...)
		case *ast.StructProperty:
			if _, ok := n.Value.(*ast.FunctionLiteral); !ok {
				rawid := n.Value.(ast.Nameable).RawId()
				s = c.specs[rawid[0]]
				vname := strings.Join(rawid, "_")
				ty := s.GetPointerType(vname)
				p = append(p, ir.NewParam(vname, ty))
			}
		case *ast.FunctionLiteral:
			// Skip, do nothing

		default:
			rawid := n.(ast.Nameable).RawId()
			s = c.specs[rawid[0]]
			vname := strings.Join(rawid, "_")
			ty := s.GetPointerType(vname)
			p = append(p, ir.NewParam(vname, ty))

		}
	}
	return p
}

func (c *Compiler) fetchOrder(id []string) []string {
	key := strings.Join(id, "_")
	return c.structPropOrder[key]
}

func (c *Compiler) generateOrder(pairs map[string]ast.Node) []string {
	keys := []string{}
	for k := range pairs {
		if k != "___base" {
			keys = append(keys, k)
		}
	}
	return util.StableSortKeys(keys)
}

func (c *Compiler) resetParaState(p []*ir.Param) {
	for i := 0; i < len(p); i++ {
		id := p[i].LocalName
		parts := strings.Split(id, "_")
		s := c.specs[parts[0]]
		s.vars.ResetState(id)
	}
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

func (c *Compiler) isConstant(rawid []string) bool {
	spec := c.specStructs[rawid[0]]
	_, err := spec.FetchConstant(rawid[1])
	if err == nil {
		return true
	}

	return false
}

func (c *Compiler) isVarSet(rawid []string) bool {
	var err error
	s := c.specStructs[rawid[0]]
	if len(rawid) > 2 {
		return c.isStrVarSet(rawid)
	}

	_, err = s.FetchStock(rawid[1])
	if err == nil {
		return true
	}

	_, err = s.FetchFlow(rawid[1])
	if err == nil {
		return true
	}

	_, err = s.FetchConstant(rawid[1])
	if err == nil {
		return true
	}

	_, err = s.FetchComponent(rawid[1])
	if err == nil {
		return true
	}

	return false
}

func (c *Compiler) isStrVarSet(rawid []string) bool {
	var err error
	s := c.specStructs[rawid[0]]
	ty, structId := s.GetStructType(rawid)
	name := strings.Join(structId[1:], "_")

	var st map[string]ast.Node
	switch ty {
	case "STOCK":
		st, err = s.FetchStock(name)
	case "FLOW":
		st, err = s.FetchFlow(name)
	case "COMPONENT":
		st, err = s.FetchComponent(name)
	default:
		return false
	}

	if err != nil {
		panic(err)
	}

	if st[rawid[len(rawid)-1]] != nil {
		return true
	}
	return false
}

func (c *Compiler) isVarSetAssert(rawid []string) bool {
	//If this is for an assert the var might reference
	//a rule on the struct level
	if c.isVarSet(rawid) || c.isInstance(rawid[0:len(rawid)-1]) {
		return true
	}
	return false
}

func (c *Compiler) isInstance(id []string) bool {
	key := strings.Join(id, "_")
	for k, _ := range c.instances {
		if k == key {
			return true
		}
	}
	return false
}

func (c *Compiler) fetchInstances(id []string) []string {
	// this and convertAssertVariables need a rethink. Seems brittle
	// with lots of edge cases
	var insta []string
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

func (c *Compiler) validOperator(node *ast.InfixExpression, boolsAllowed bool) bool {
	if !boolsAllowed && (node.Left.Type() == "BOOL" || node.Right.Type() == "BOOL") {
		return false
	}
	return true
}

func negate(e ast.Expression) ast.Expression {
	//Negate the expression so that the solver attempts to disprove it
	switch n := e.(type) {
	case *ast.InfixExpression:
		op, ok := OP_NEGATE[n.Operator]
		if ok {
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

func (c *Compiler) GetIR() string {
	return c.module.String()
}

func (c *Compiler) addGlobal() {
	// run block
	c.contextFunc = c.module.NewFunc("__run", irtypes.Void)
	mainBlock := c.contextFunc.NewBlock(name.Block())
	mainBlock.NewRet(nil)
	c.contextBlock = mainBlock
}

type Panic string
