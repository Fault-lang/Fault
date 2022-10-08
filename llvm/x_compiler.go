package llvm

// import (
// 	"errors"
// 	"fault/ast"
// 	"fault/llvm/name"
// 	"fault/types"
// 	"fault/util"
// 	"fmt"
// 	"runtime/debug"
// 	"strings"

// 	"github.com/llir/llvm/ir"
// 	"github.com/llir/llvm/ir/constant"
// 	"github.com/llir/llvm/ir/enum"
// 	"github.com/llir/llvm/ir/metadata"
// 	irtypes "github.com/llir/llvm/ir/types"
// 	"github.com/llir/llvm/ir/value"
// )

// var DoubleP = &irtypes.PointerType{ElemType: irtypes.Double}
// var I1P = &irtypes.PointerType{ElemType: irtypes.I1}

// var OP_NEGATE = map[string]string{
// 	"==": "!=",
// 	">=": "<",
// 	">":  "<=",
// 	"<=": ">",
// 	"!=": "==",
// 	"<":  ">=",
// 	"&&": "||",
// 	"||": "&&",
// 	//"=": "!=",
// }

// type Compiler struct {
// 	module *ir.Module

// 	specs            map[string]*spec
// 	instances        map[string][]string
// 	instanceChildren map[string]string
// 	currentSpec      *spec

// 	currentSpecName string
// 	currScope       []string
// 	currStruct      string

// 	specStructs     map[string]types.StockFlow
// 	specFunctions   map[string]value.Value
// 	structPropOrder map[string][]string
// 	builtIns        map[string]*ir.Func

// 	contextFuncName string
// 	contextMetadata *metadata.Attachment

// 	alloc    bool
// 	runRound int16

// 	contextBlock *ir.Block
// 	contextFunc  *ir.Func

// 	// Stack of variables that are in scope
// 	allocatedPointers []map[string]*ir.InstAlloca

// 	// Where a condition should jump when done
// 	contextCondAfter []*ir.Block

// 	specGlobals     map[string]*ir.Global
// 	RawAsserts      []*ast.AssertionStatement
// 	RawAssumes      []*ast.AssumptionStatement
// 	Asserts         []*ast.AssertionStatement
// 	Assumes         []*ast.AssumptionStatement
// 	Uncertains      map[string][]float64
// 	Unknowns        []string
// 	Components      map[string]map[string]string
// 	ComponentStarts map[string]string
// }

// func NewCompiler() *Compiler {
// 	c := &Compiler{
// 		module: ir.NewModule(),

// 		specs:            make(map[string]*spec),
// 		instances:        make(map[string][]string),
// 		instanceChildren: make(map[string]string),
// 		specStructs:      make(map[string]types.StockFlow),
// 		specFunctions:    make(map[string]value.Value),
// 		structPropOrder:  make(map[string][]string),
// 		builtIns:         make(map[string]*ir.Func),

// 		contextMetadata: nil,
// 		alloc:           true,

// 		allocatedPointers: make([]map[string]*ir.InstAlloca, 0),
// 		currScope:         []string{""},

// 		contextCondAfter: make([]*ir.Block, 0),

// 		specGlobals: make(map[string]*ir.Global),
// 		runRound:    0,

// 		Uncertains:      make(map[string][]float64),
// 		Components:      make(map[string]map[string]string),
// 		ComponentStarts: make(map[string]string),
// 	}
// 	c.addGlobal()
// 	return c
// }

// func (c *Compiler) LoadMeta(structs map[string]types.StockFlow, uncertains map[string][]float64, unknowns []string) {
// 	c.specStructs = structs
// 	c.Unknowns = unknowns
// 	c.Uncertains = uncertains
// }

// func (c *Compiler) Compile(root ast.Node) (err error) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			// Compile time panics, that are not errors in the compiler
// 			if _, ok := r.(Panic); ok {
// 				err = errors.New(fmt.Sprint(r))
// 				return
// 			}

// 			// Bugs in the compiler
// 			err = fmt.Errorf("%s\n\nInternal compiler stacktrace:\n%s",
// 				fmt.Sprint(r),
// 				string(debug.Stack()),
// 			)
// 		}
// 	}()

// 	c.processSpec(root, false)
// 	return
// }

// func (c *Compiler) processSpec(root ast.Node, isImport bool) ([]*ast.AssertionStatement, []*ast.AssumptionStatement) {
// 	specfile, ok := root.(*ast.Spec)
// 	if !ok {
// 		panic(fmt.Sprintf("spec file improperly formatted. Root node is %T", root))
// 	}

// 	var name string
// 	switch decl := specfile.Statements[0].(type) {
// 	case *ast.SpecDeclStatement:
// 		name = decl.Name.Value
// 	case *ast.SysDeclStatement:
// 		name = decl.Name.Value
// 	default:
// 		panic(fmt.Sprintf("spec file improperly formatted. Missing spec declaration, got %T", specfile.Statements[0]))
// 	}

// 	c.currentSpec = NewCompiledSpec(name)
// 	c.currentSpecName = name
// 	c.specs[c.currentSpecName] = c.currentSpec
// 	for _, fileNode := range specfile.Statements {
// 		c.compile(fileNode)
// 	}

// 	if !isImport {
// 		for _, assert := range c.RawAsserts {
// 			c.compileAssert(assert)
// 		}
// 		for _, assert := range c.RawAssumes {
// 			c.compileAssert(assert)
// 		}
// 	}
// 	return c.Asserts, c.Assumes
// }

// func (c *Compiler) compile(node ast.Node) {
// 	switch v := node.(type) {
// 	case *ast.SpecDeclStatement:
// 		break
// 	case *ast.SysDeclStatement:
// 		break
// 	case *ast.ImportStatement:
// 		parent := c.currentSpecName
// 		parentSp := c.currentSpec
// 		asserts, assumes := c.processSpec(v.Tree, true) //Move all asserts to the end of the compilation process
// 		c.Asserts = append(c.Asserts, asserts...)
// 		c.Assumes = append(c.Assumes, assumes...)
// 		c.currentSpecName = parent
// 		c.currentSpec = parentSp
// 	case *ast.ConstantStatement:
// 		c.compileConstant(v)
// 	case *ast.DefStatement:
// 		c.compileStruct(v)

// 	case *ast.FunctionLiteral:

// 	case *ast.InfixExpression:
// 		c.compileInfix(v)

// 	case *ast.PrefixExpression:
// 		c.compilePrefix(v)

// 	case *ast.AssumptionStatement:
// 		// Need to do these after the run block so we move them
// 		c.RawAssumes = append(c.RawAssumes, v)

// 	case *ast.AssertionStatement:
// 		c.RawAsserts = append(c.RawAsserts, v)
// 		//c.compileAssertion(v)

// 	case *ast.ForStatement:
// 		c.contextFuncName = "__run"
// 		for i := int64(0); i < v.Rounds.Value; i++ {
// 			c.compileBlock(v.Body)
// 			c.runRound = c.runRound + 1
// 		}
// 		c.contextFuncName = ""

// 	case *ast.StartStatement:
// 		for _, p := range v.Pairs {
// 			id := c.getFullVariableName([]string{p[0], p[1]})
// 			id, _ = c.GetSpec(id)
// 			c.processFunc(id, []string{id[0], p[0]}, 0)
// 			c.ComponentStarts[p[0]] = p[1]
// 		}

// 	default:
// 		pos := node.Position()
// 		panic(fmt.Sprintf("node type %T unimplemented line: %d col: %d", v, pos[0], pos[1]))
// 	}

// 	// InitExpression
// 	// IfExpression
// 	// IndexExpression <-- Is this still used?

// }

// func (c *Compiler) compileComponent(node *ast.ComponentLiteral, cname string) {
// 	for k, p := range node.Pairs {
// 		var pname string
// 		key := k.(*ast.Identifier)
// 		scopeName := []string{cname, key.Value}

// 		oldScope := c.currScope
// 		oldBlock := c.contextBlock

// 		c.contextFuncName = strings.Join(scopeName, "_")
// 		c.currScope = scopeName

// 		switch v := p.(type) {
// 		case *ast.StateLiteral:
// 			params := c.generateParameters([]string{key.Spec, cname}, c.specStructs[key.Spec][cname], []string{key.Spec, cname, key.Value})
// 			c.resetParaState(params)
// 			f := c.module.NewFunc(key.Value, irtypes.Void, params...)
// 			c.contextFunc = f
// 			pname = name.Block()
// 			c.contextBlock = f.NewBlock(pname)
// 			if c.Components[cname] != nil {
// 				c.Components[cname][key.Value] = pname
// 			} else {
// 				c.Components[cname] = map[string]string{key.Value: pname}
// 			}
// 			val := c.compileBlock(v.Body)
// 			c.contextBlock.NewRet(val)
// 			c.contextBlock = oldBlock
// 			c.contextFuncName = "__run"
// 			c.currScope = oldScope
// 			c.contextFunc = nil
// 		case *ast.Instance:
// 			c.compileInstance(v, cname)
// 		default:
// 			val := c.compileValue(v)

// 			if val != nil {
// 				id := []string{cname, key.Value}
// 				id, s := c.GetSpec(id)
// 				if s.GetSpecVar(id) != nil {
// 					name := strings.Join(id, "_")
// 					pointer := s.GetSpecVarPointer(name)
// 					ty := s.GetSpecType(name)
// 					c.contextBlock.NewLoad(ty, pointer)
// 				} else {
// 					s.DefineSpecType(id, val.Type())
// 					s.DefineSpecVar(id, val)
// 					c.allocVariable(id, val, []int{0, 0, 0, 0})
// 				}
// 			}
// 		}
// 	}
// }

// func (c *Compiler) compileIdent(node *ast.Identifier) *ir.InstLoad {
// 	return c.lookupIdent([]string{node.Spec, node.Value}, node.Position())
// }

// func (c *Compiler) compilePrefix(node *ast.PrefixExpression) value.Value {
// 	val := c.compileInfixNode(node.Right)
// 	switch node.Operator {
// 	case "!":
// 		return c.contextBlock.NewXor(val, constant.NewInt(irtypes.I1, 1))
// 	case "-":
// 		return c.contextBlock.NewFNeg(val)
// 	default:
// 		panic(fmt.Sprintf("unrecognized prefix operator %s", node.Operator))
// 	}
// }

// func (c *Compiler) compileInfix(node *ast.InfixExpression) value.Value {
// 	pos := node.Position()
// 	switch node.Operator {
// 	case "=": // Used to store temporary local values
// 		if !c.validOperator(node, true) {
// 			panic(fmt.Sprintf("operator %s cannot be used on variables of type %s and %s", node.Operator, node.Left.Type(), node.Right.Type()))
// 		}

// 		var id []string
// 		var s *spec
// 		r := c.compileValue(node.Right)
// 		if _, ok := node.Right.(*ast.Instance); !ok { // If declaring a new instance don't save
// 			var fvn []string
// 			switch n := node.Left.(type) {
// 			case *ast.Identifier:
// 				fvn = c.getFullVariableName([]string{n.Value})
// 			case *ast.ParameterCall:
// 				fvn = c.getFullVariableName(n.Value)
// 			}
// 			id, s = c.GetSpec(fvn)

// 			if c.isVarSet(id) && c.alloc {
// 				fvns := strings.Join(id, "_")
// 				p := s.GetSpecVarPointer(fvns)
// 				c.contextBlock.NewStore(r, p)
// 				return nil
// 			}
// 			s.DefineSpecVar(id, r)
// 			s.DefineSpecType(id, r.Type())
// 			if c.alloc {
// 				c.allocVariable(id, r, node.Left.Position())
// 			}
// 		}
// 		return nil
// 	case "<-":
// 		if !c.validOperator(node, false) {
// 			panic(fmt.Sprintf("operator %s cannot be used on variables of type %s and %s", node.Operator, node.Left.Type(), node.Right.Type()))
// 		}

// 		r := c.compileValue(node.Right)
// 		n, ok := node.Left.(*ast.ParameterCall)
// 		if !ok {
// 			pos := node.Position()
// 			panic(fmt.Sprintf("cannot use <- or -> operator on a non-stock value col: %d, line: %d", pos[0], pos[1]))
// 		}

// 		pos := n.Position()
// 		var fvn []string
// 		switch n := node.Left.(type) {
// 		case *ast.Identifier:
// 			fvn = c.getFullVariableName([]string{n.Value})
// 		case *ast.ParameterCall:
// 			fvn = c.pcVariableName(n)
// 		}
// 		id, _ := c.GetSpec(fvn)
// 		fvns := strings.Join(id, "_")

// 		if !c.isVarSet(id) {
// 			panic(fmt.Sprintf("cannot send value to variable %s. Variable not defined line: %d, col: %d", fvns, pos[0], pos[1]))
// 		}

// 		_, s := c.GetSpec(fvn)
// 		pointer := s.GetSpecVarPointer(fvns)
// 		c.contextBlock.NewStore(r, pointer)
// 		return nil
// 	case "+":
// 		if !c.validOperator(node, false) {
// 			panic(fmt.Sprintf("operator %s cannot be used on variables of type %s and %s", node.Operator, node.Left.Type(), node.Right.Type()))
// 		}

// 		l := c.compileInfixNode(node.Left)
// 		r := c.compileInfixNode(node.Right)

// 		return c.contextBlock.NewFAdd(l, r)
// 	case "-":
// 		if !c.validOperator(node, false) {
// 			panic(fmt.Sprintf("operator %s cannot be used on variables of type %s and %s", node.Operator, node.Left.Type(), node.Right.Type()))
// 		}

// 		l := c.compileInfixNode(node.Left)
// 		r := c.compileInfixNode(node.Right)

// 		sub := c.contextBlock.NewFSub(l, r)
// 		return sub
// 	case "*":
// 		if !c.validOperator(node, false) {
// 			panic(fmt.Sprintf("operator %s cannot be used on variables of type %s and %s", node.Operator, node.Left.Type(), node.Right.Type()))
// 		}

// 		l := c.compileInfixNode(node.Left)
// 		r := c.compileInfixNode(node.Right)

// 		mul := c.contextBlock.NewFMul(l, r)
// 		return mul
// 	case "/":
// 		if !c.validOperator(node, false) {
// 			panic(fmt.Sprintf("operator %s cannot be used on variables of type %s and %s", node.Operator, node.Left.Type(), node.Right.Type()))
// 		}

// 		l := c.compileInfixNode(node.Left)
// 		r := c.compileInfixNode(node.Right)

// 		div := c.contextBlock.NewFDiv(l, r)
// 		return div
// 	case "%":
// 		if !c.validOperator(node, false) {
// 			panic(fmt.Sprintf("operator %s cannot be used on variables of type %s and %s", node.Operator, node.Left.Type(), node.Right.Type()))
// 		}

// 		l := c.compileInfixNode(node.Left)
// 		r := c.compileInfixNode(node.Right)

// 		rem := c.contextBlock.NewFRem(l, r)
// 		return rem
// 	case ">":
// 		if !c.validOperator(node, false) {
// 			panic(fmt.Sprintf("operator %s cannot be used on variables of type %s and %s", node.Operator, node.Left.Type(), node.Right.Type()))
// 		}

// 		l := c.compileInfixNode(node.Left)
// 		r := c.compileInfixNode(node.Right)

// 		ogt := c.contextBlock.NewFCmp(enum.FPredOGT, l, r)
// 		return ogt
// 	case ">=":
// 		if !c.validOperator(node, false) {
// 			panic(fmt.Sprintf("operator %s cannot be used on variables of type %s and %s", node.Operator, node.Left.Type(), node.Right.Type()))
// 		}

// 		l := c.compileInfixNode(node.Left)
// 		r := c.compileInfixNode(node.Right)

// 		oge := c.contextBlock.NewFCmp(enum.FPredOGE, l, r)
// 		return oge
// 	case "<":
// 		if !c.validOperator(node, false) {
// 			panic(fmt.Sprintf("operator %s cannot be used on variables of type %s and %s", node.Operator, node.Left.Type(), node.Right.Type()))
// 		}

// 		l := c.compileInfixNode(node.Left)
// 		r := c.compileInfixNode(node.Right)

// 		olt := c.contextBlock.NewFCmp(enum.FPredOLT, l, r)
// 		return olt
// 	case "<=":
// 		if !c.validOperator(node, false) {
// 			panic(fmt.Sprintf("operator %s cannot be used on variables of type %s and %s", node.Operator, node.Left.Type(), node.Right.Type()))
// 		}

// 		l := c.compileInfixNode(node.Left)
// 		r := c.compileInfixNode(node.Right)

// 		ole := c.contextBlock.NewFCmp(enum.FPredOLE, l, r)
// 		return ole
// 	case "==":
// 		if !c.validOperator(node, true) {
// 			panic(fmt.Sprintf("operator %s cannot be used on variables of type %s and %s", node.Operator, node.Left.Type(), node.Right.Type()))
// 		}

// 		l := c.compileInfixNode(node.Left)
// 		r := c.compileInfixNode(node.Right)

// 		if node.Right.Type() == "BOOL" {
// 			return c.contextBlock.NewICmp(enum.IPredEQ, l, r)
// 		} else {
// 			return c.contextBlock.NewFCmp(enum.FPredOEQ, l, r)
// 		}
// 	case "!=":
// 		if !c.validOperator(node, true) {
// 			panic(fmt.Sprintf("operator %s cannot be used on variables of type %s and %s", node.Operator, node.Left.Type(), node.Right.Type()))
// 		}

// 		l := c.compileInfixNode(node.Left)
// 		r := c.compileInfixNode(node.Right)

// 		if node.Right.Type() == "BOOL" {
// 			return c.contextBlock.NewICmp(enum.IPredNE, l, r)
// 		} else {
// 			return c.contextBlock.NewFCmp(enum.FPredONE, l, r)
// 		}
// 	default:
// 		panic(fmt.Sprintf("unknown operator %s. line: %d, col: %d", node.Operator, pos[0], pos[1]))
// 	}
// }

// func (c *Compiler) pcVariableName(v *ast.ParameterCall) []string {
// 	id := []string{c.currentSpecName, c.currStruct}

// 	if c.currScope[0] != "" {
// 		id = append(id, c.currScope...)
// 	}

// 	return append(id, v.Value...)
// }

// func (c *Compiler) compileInfixNode(node ast.Node) value.Value {
// 	switch v := node.(type) {
// 	case *ast.ParameterCall:
// 		var id []string
// 		if c.currStruct != "" {
// 			id = c.pcVariableName(v)
// 		} else {
// 			id = c.getFullVariableName(v.Value)
// 		}
// 		return c.lookupIdent(id, node.Position())
// 	default:
// 		return c.compileValue(node)
// 	}
// }

// func (c *Compiler) compileParallel(node *ast.ParallelFunctions) {
// 	if c.contextFuncName != "__run" {
// 		pos := node.Position()
// 		panic(fmt.Sprintf("cannot use parallel operator outside of the run block. line: %d, col: %d", pos[0], pos[1]))
// 	}
// 	gname := name.ParallelGroup(node.String())
// 	for i := 0; i < len(node.Expressions); i++ {
// 		l := c.compileValue(node.Expressions[i])
// 		md := &metadata.Attachment{
// 			Name: gname,
// 			Node: &metadata.DIBasicType{
// 				MetadataID: -1,
// 				Tag:        enum.DwarfTagStringType,
// 			}}
// 		switch exp := l.(type) {
// 		case *ir.Func:
// 			id := node.Expressions[i].(*ast.ParameterCall).Value
// 			id, s := c.GetSpec(id)
// 			params := s.GetParams(id)
// 			l_func := c.contextBlock.NewCall(exp, params...)
// 			l_func.Metadata = append(l_func.Metadata, md)
// 		case *ir.InstFAdd:
// 			exp.Metadata = append(exp.Metadata, md)
// 		case *ir.InstFSub:
// 			exp.Metadata = append(exp.Metadata, md)
// 		case *ir.InstFMul:
// 			exp.Metadata = append(exp.Metadata, md)
// 		case *ir.InstFDiv:
// 			exp.Metadata = append(exp.Metadata, md)
// 		case *ir.InstFRem:
// 			exp.Metadata = append(exp.Metadata, md)
// 		case *ir.InstFCmp:
// 			exp.Metadata = append(exp.Metadata, md)
// 		}
// 	}
// 	c.contextMetadata = nil
// }

// func (c *Compiler) compileValue(node ast.Node) value.Value {
// 	if node == nil {
// 		panic("value received by compileValue is nil")
// 	}
// 	switch v := node.(type) {
// 	case *ast.IntegerLiteral:
// 		return constant.NewFloat(irtypes.Double, float64(v.Value))
// 	case *ast.FloatLiteral:
// 		return constant.NewFloat(irtypes.Double, v.Value)
// 	case *ast.StringLiteral:
// 		return constant.NewCharArrayFromString(v.Value)
// 	case *ast.Boolean:
// 		return constant.NewBool(v.Value)
// 	case *ast.Natural:
// 		return constant.NewFloat(irtypes.Double, float64(v.Value))
// 	case *ast.Uncertain: //Set to dummy value for LLVM IR, catch during SMT generation
// 		return constant.NewFloat(irtypes.Double, float64(0.000000000009))
// 	case *ast.Unknown:
// 		return constant.NewFloat(irtypes.Double, float64(0.000000000009))
// 	case *ast.Nil:
// 		return constant.NewNull(&irtypes.PointerType{})
// 	case *ast.Identifier:
// 		return c.compileIdent(v)
// 	case *ast.InfixExpression:
// 		return c.compileInfix(v)
// 	case *ast.PrefixExpression:
// 		return c.compilePrefix(v)
// 	case *ast.FunctionLiteral:
// 		return c.compileFunction(v)
// 	case *ast.Instance:
// 		c.compileInstance(v, v.Name)
// 	case *ast.ParameterCall:
// 		return c.compileParameterCall(v)
// 	case *ast.BlockStatement:
// 		return c.compileBlock(v)
// 	default:
// 		pos := v.Position()
// 		panic(fmt.Sprintf("unknown value type %T line: %d col: %d", v, pos[0], pos[1]))
// 	}
// 	return nil
// }

// func (c *Compiler) compileConstant(node *ast.ConstantStatement) {
// 	value := c.compileValue(node.Value)
// 	id := c.getFullVariableName([]string{node.Name.Value})
// 	id, _ = c.GetSpec(id)
// 	c.setConst(id, value)
// 	c.globalVariable(id, value, node.Position())
// }

// func (c *Compiler) compileBlock(node *ast.BlockStatement) value.Value {
// 	if !c.alloc {
// 		return nil
// 	}
// 	body := node.Statements
// 	var ret value.Value
// 	for i := 0; i < len(body); i++ {
// 		switch exp := body[i].(type) {
// 		case *ast.ParallelFunctions:
// 			c.compileParallel(exp)
// 		case ast.Expression:
// 			ret = c.compileFunctionBody(exp)
// 		case *ast.ExpressionStatement:
// 			ret = c.compileFunctionBody(exp.Expression)
// 		}
// 	}
// 	return ret
// }

// func (c *Compiler) GetIR() string {
// 	return c.module.String()
// }

// func (c *Compiler) addGlobal() {
// 	global := NewCompiledSpec("__global")

// 	c.specs["__global"] = global

// 	// run block
// 	c.contextFunc = c.module.NewFunc("__run", irtypes.Void)
// 	mainBlock := c.contextFunc.NewBlock(name.Block())
// 	mainBlock.NewRet(nil)
// 	c.contextBlock = mainBlock
// }

// func (c *Compiler) compileStruct(def *ast.DefStatement) {
// 	namespace := []string{def.Name.Spec, def.Name.Value}
// 	key := strings.Join(namespace, "_")
// 	switch def.Type() {
// 	case "FLOW":
// 		c.instances[key] = namespace
// 		c.structPropOrder[key] = def.Value.(*ast.FlowLiteral).Order
// 	case "STOCK":
// 		c.instances[key] = namespace
// 		c.structPropOrder[key] = def.Value.(*ast.StockLiteral).Order
// 	case "GLOBAL":
// 		instance, _ := def.Value.(*ast.Instance)
// 		importSpace := []string{instance.Value.Spec, instance.Value.Value}
// 		c.instances[key] = importSpace
// 		c.compileInstance(def.Value.(*ast.Instance), strings.Join(importSpace, "_"))
// 	case "COMPONENT":
// 		c.instances[key] = namespace
// 		c.structPropOrder[key] = def.Value.(*ast.ComponentLiteral).Order
// 		c.compileComponent(def.Value.(*ast.ComponentLiteral), def.Name.Value)
// 	}
// }

// func (c *Compiler) compileAssert(assert ast.Node) {
// 	var l, r ast.Expression
// 	switch a := assert.(type) {
// 	case *ast.AssertionStatement:
// 		if a.TemporalFilter == "" { //If there is a temporal filter this is negated instead
// 			l = negate(a.Constraints.Left)
// 			r = negate(a.Constraints.Right)
// 			a.Constraints.Operator = OP_NEGATE[a.Constraints.Operator]
// 		} else {
// 			l = a.Constraints.Left
// 			r = a.Constraints.Right
// 			a.TemporalFilter, a.TemporalN = negateTemporal(a.TemporalFilter, a.TemporalN)
// 			if a.TemporalN < 0 {
// 				pos := a.Position()
// 				panic(fmt.Sprintf("temporal logic not value, filter searching for fewer than 0 states: line %d col %d", pos[0], pos[1]))
// 			}
// 		}
// 		a.Constraints.Left = c.convertAssertVariables(l)
// 		a.Constraints.Right = c.convertAssertVariables(r)
// 		c.Asserts = append(c.Asserts, a)
// 	case *ast.AssumptionStatement:
// 		a.Constraints.Left = c.convertAssertVariables(a.Constraints.Left)
// 		a.Constraints.Right = c.convertAssertVariables(a.Constraints.Right)
// 		c.Assumes = append(c.Assumes, a)
// 	default:
// 		panic("statement must be an assert or an assumption.")
// 	}
// }

// func (c *Compiler) generateOrder(structName []string, pairs map[string]ast.Node) []string {
// 	key := strings.Join(structName, "_")
// 	if c.structPropOrder[key] != nil {
// 		return c.structPropOrder[key]
// 	}
// 	panic(fmt.Sprintf("no property order found for struct %s ", structName))
// }

// func (c *Compiler) compileInstance(base *ast.Instance, instName string) {
// 	if c.runRound > 0 { // Initialize things only once
// 		return
// 	}
// 	if c.contextFuncName == "__run" {
// 		c.currScope = []string{instName}
// 		c.alloc = false
// 	}

// 	oldStruct := c.currStruct
// 	c.currStruct = base.Name

// 	pos := base.Position()
// 	structName := []string{base.Value.Spec, base.Value.Value}
// 	parentFunction := c.contextFuncName
// 	c.contextFuncName = instName
// 	if c.specStructs[structName[0]][structName[1]] == nil {
// 		panic(fmt.Sprintf("no stock or flow named %s, line: %d, col %d", structName, pos[0], pos[1]))
// 	}

// 	children := c.processStruct(structName, instName, pos)
// 	c.instanceChildren = util.MergeStringMaps(c.instanceChildren, children)

// 	key := strings.Join([]string{c.currentSpecName, instName}, "_")
// 	c.instances[key] = structName
// 	c.contextFuncName = parentFunction
// 	if c.contextFuncName == "__run" {
// 		c.currScope = []string{""}
// 		c.alloc = true
// 	}
// 	c.currStruct = oldStruct
// }

// func (c *Compiler) processFunc(id []string, structName []string, round int /*pos []int*/) value.Value {
// 	key := id[2]
// 	fname := strings.Join(id, "_")

// 	if round == 0 { //initialize
// 		params := c.generateParameters(structName, c.specStructs[structName[0]][structName[1]], id)
// 		c.resetParaState(params)
// 		f := c.module.NewFunc(fname, irtypes.Void, params...)
// 		c.contextFunc = f

// 		oldScope := c.currScope
// 		oldBlock := c.contextBlock

// 		c.contextFuncName = fname
// 		c.currScope = []string{id[1]} // NOT necessarily the same as structName
// 		c.contextBlock = f.NewBlock(name.Block())

// 		// There's obviously something very wrong here
// 		// this is kind of bad patch for right now until
// 		// I get through some functions of the compiler
// 		var branch ast.Node
// 		if c.specStructs[structName[0]][structName[1]][key] != nil {
// 			branch = c.specStructs[structName[0]][structName[1]][key]
// 		} else {
// 			branch = c.specStructs[id[0]][id[1]][key]
// 		}
// 		val := c.compileValue(branch)
// 		c.contextBlock.NewRet(val)

// 		c.contextBlock = oldBlock
// 		c.contextFuncName = "__run"
// 		c.currScope = oldScope
// 		c.specFunctions[fname] = f
// 		c.contextFunc = nil
// 		c.resetParaState(params)
// 	}

// 	return c.specFunctions[fname]
// }

// func (c *Compiler) processStruct(structName []string, instName string, pos []int) map[string]string {
// 	baseSpec := c.currentSpecName
// 	children := make(map[string]string)
// 	keys := c.generateOrder(structName, c.specStructs[structName[0]][structName[1]])
// 	for _, k := range keys {
// 		var isUncertain []float64
// 		var isUnknown bool
// 		id := strings.Split(instName, "_")            // Slightly hacky solution to nestled instances
// 		if util.InStringSlice(c.ListSpecs(), id[0]) { // <-- This is an edge case when importing structs
// 			id = append([]string{baseSpec}, id[1:]...) //baseSpec will be different from structName[0] if the struct is imported

// 		} else {
// 			id = append([]string{baseSpec}, id...) //baseSpec will be different from structName[0] if the struct is imported
// 		}
// 		id = append(id, k)

// 		switch pv := c.specStructs[structName[0]][structName[1]][k].(type) {
// 		case *ast.Instance:
// 			switch len(id) {
// 			case 3:
// 				// Slightly hacky solution to nestled instances
// 				c.compileInstance(pv, strings.Join(id[1:], "_"))
// 			default:
// 				c.compileInstance(pv, k) // Copy instance data over
// 			}
// 		case *ast.FunctionLiteral:
// 			c.compileFunction(pv)
// 		case *ast.InfixExpression:
// 			c.compileInfix(pv)
// 		case *ast.BlockStatement:
// 			c.compileBlock(pv)
// 		default:
// 			_, ok := pv.(*ast.Uncertain)
// 			if ok {
// 				isUnknown = true
// 			}

// 			uncertain, ok2 := pv.(*ast.Uncertain)
// 			if ok2 {
// 				isUncertain = []float64{uncertain.Mean, uncertain.Sigma}
// 			}
// 			val := c.compileValue(c.specStructs[structName[0]][structName[1]][k])
// 			id, s := c.GetSpec(id)
// 			s.DefineSpecVar(id, val)
// 			s.DefineSpecType(id, val.Type())
// 			c.allocVariable(id, val, pos)
// 			s.vars.ResetState(id)
// 			name := strings.Join(id, "_")
// 			ty := s.GetPointerType(name)
// 			p := ir.NewParam(name, ty)
// 			s.AddParam(id, p)

// 		}
// 		//Track properties of instances so that we can write
// 		// asserts on the struct and honor them for all instances
// 		id, _ = c.GetSpec(id)
// 		fvn := strings.Join(id, "_")
// 		if isUnknown {
// 			c.Unknowns = append(c.Unknowns, fvn)
// 		}
// 		if isUncertain != nil {
// 			c.Uncertains[fvn] = isUncertain
// 		}
// 		children[fvn] = structName[1]
// 	}
// 	return children
// }

// func (c *Compiler) compileIf(n *ast.IfExpression) {
// 	cond := c.compileConditional(n.Condition)

// 	afterBlock := c.contextBlock.Parent.NewBlock(name.Block() + "-after")
// 	trueBlock := c.contextBlock.Parent.NewBlock(name.Block() + "-true")
// 	falseBlock := afterBlock

// 	c.contextCondAfter = append(c.contextCondAfter, afterBlock)

// 	if n.Alternative != nil {
// 		falseBlock = c.contextBlock.Parent.NewBlock(name.Block() + "-false")
// 	}

// 	c.contextBlock.NewCondBr(cond, trueBlock, falseBlock)

// 	c.contextBlock = trueBlock
// 	c.compileBlock(n.Consequence)

// 	// Jump to after-block if no terminator has been set (such as a return statement)
// 	if trueBlock.Term == nil {
// 		trueBlock.NewBr(afterBlock)
// 	}

// 	if n.Alternative != nil {
// 		c.contextBlock = falseBlock
// 		c.compileBlock(n.Alternative)

// 		// Jump to after-block if no terminator has been set (such as a return statement)
// 		if falseBlock.Term == nil {
// 			falseBlock.NewBr(afterBlock)
// 		}
// 	}

// 	c.contextBlock = afterBlock

// 	// pop after block stack
// 	c.contextCondAfter = c.contextCondAfter[0 : len(c.contextCondAfter)-1]

// 	// set after block to jump to the after block
// 	if len(c.contextCondAfter) > 0 {
// 		afterBlock.NewBr(c.contextCondAfter[len(c.contextCondAfter)-1])
// 	} else {
// 		afterBlock.NewRet(nil)
// 	}
// }

// func (c *Compiler) compileConditional(n ast.Node) value.Value {
// 	// Reformat the conditional clause to accept
// 	// things like if a {} or if !a {} and replace them
// 	// with a == true or a == false
// 	switch conditional := n.(type) {
// 	case *ast.InfixExpression:
// 		return c.compileValue(conditional)
// 	case *ast.PrefixExpression:
// 		return c.compilePrefix(conditional)
// 	case *ast.Identifier:
// 		right := &ast.Boolean{
// 			Token:        conditional.Token,
// 			InferredType: conditional.InferredType,
// 			Value:        true,
// 		}
// 		n := &ast.InfixExpression{Token: conditional.Token,
// 			InferredType: conditional.InferredType,
// 			Left:         conditional,
// 			Operator:     "==",
// 			Right:        right}
// 		return c.compileValue(n)
// 	case *ast.ParameterCall:
// 		right := &ast.Boolean{
// 			Token:        conditional.Token,
// 			InferredType: conditional.InferredType,
// 			Value:        true,
// 		}
// 		n := &ast.InfixExpression{Token: conditional.Token,
// 			InferredType: conditional.InferredType,
// 			Left:         conditional,
// 			Operator:     "==",
// 			Right:        right}
// 		return c.compileValue(n)
// 	}
// 	return c.compileValue(n)
// }

// func (c *Compiler) compileParameterCall(pc *ast.ParameterCall) value.Value {
// 	id := c.getFullVariableName(pc.Value)
// 	id, s := c.GetSpec(id)
// 	key := strings.Join([]string{id[0], pc.Value[0]}, "_")
// 	structName := c.instances[key]
// 	// If we're in the run block and the parameter is defined as a function
// 	// define it as a function and call it from run block
// 	if c.contextFuncName == "__run" &&
// 		c.isFunction(c.specStructs[structName[0]][structName[1]][pc.Value[1]]) {
// 		return c.processFunc(id, structName, int(c.runRound))
// 	}

// 	// Otherwise inline the parameter...
// 	if c.currScope[0] == "" {
// 		c.currScope = []string{pc.Value[0]}
// 	}
// 	parentFunction := c.contextFuncName
// 	c.contextFuncName = pc.Value[0]

// 	val := c.compileValue(c.specStructs[id[0]][structName[1]][pc.Value[1]])

// 	// If there's no value, there's nothing to store
// 	if val != nil || !c.isFunction(c.specStructs[id[0]][structName[1]][pc.Value[1]]) {
// 		if s.GetSpecVar(id) != nil {
// 			name := strings.Join(id, "_")
// 			pointer := s.GetSpecVarPointer(name)
// 			ty := s.GetSpecType(name)
// 			c.contextBlock.NewLoad(ty, pointer)
// 		} else {
// 			s.DefineSpecType(id, val.Type())
// 			s.DefineSpecVar(id, val)
// 			c.allocVariable(id, val, pc.Position())
// 		}
// 	}
// 	if c.currScope[0] == pc.Value[0] {
// 		c.currScope = []string{""}
// 	}
// 	c.contextFuncName = parentFunction
// 	return val
// }

// func (c *Compiler) compileFunction(node *ast.FunctionLiteral) value.Value {
// 	body := node.Body.Statements
// 	var retValue value.Value
// 	for i := 0; i < len(body); i++ {
// 		exp := body[i].(*ast.ExpressionStatement).Expression
// 		init, ok := exp.(*ast.InitExpression)
// 		if ok {
// 			return c.compileValue(init.Expression)
// 		}
// 	}
// 	return retValue
// }

// func (c *Compiler) compileFunctionBody(node ast.Expression) value.Value {
// 	if !c.alloc { //Short circuit this if just initializing
// 		return nil
// 	}
// 	switch v := node.(type) {
// 	case *ast.InfixExpression:
// 		return c.compileInfix(v)

// 	case *ast.PrefixExpression:
// 		c.compilePrefix(v)

// 	case *ast.IfExpression:
// 		c.compileIf(v)

// 	case *ast.Instance:
// 		orign := v.Name
// 		origv := v.Value.Value
// 		v.Value.Value = orign
// 		v.Name = origv
// 		c.compileInstance(v, v.Name)

// 	case *ast.IndexExpression:

// 	case *ast.ParameterCall:
// 		return c.compileParameterCall(v)

// 	case *ast.BuiltIn:
// 		//Is this the first time we're seeing this builtin?
// 		if c.builtIns[v.Function] == nil {
// 			var param []*ir.Param
// 			for k := range v.Parameters {
// 				param = append(param, ir.NewParam(k, irtypes.NewPointer(irtypes.I8)))
// 			}
// 			oldBlock := c.contextBlock
// 			f := c.module.NewFunc(v.Function, irtypes.Void, param...)
// 			c.contextBlock = f.NewBlock(name.Block())

// 			c.contextBlock.NewRet(nil)
// 			c.contextBlock = oldBlock

// 			c.builtIns[v.Function] = f
// 		}
// 		var params []value.Value
// 		for _, v := range v.Parameters {
// 			l := uint64(len(v.String()))
// 			alloc := c.contextBlock.NewAlloca(irtypes.NewArray(l, irtypes.I8))
// 			c.contextBlock.NewStore(constant.NewCharArrayFromString(v.String()), alloc)
// 			//load := c.contextBlock.NewLoad(irtypes.I8, alloc)
// 			cast := c.contextBlock.NewBitCast(alloc, irtypes.I8Ptr)
// 			params = append(params, cast)
// 		}

// 		c.contextBlock.NewCall(c.builtIns[v.Function], params...)

// 	default:
// 		pos := node.Position()
// 		panic(fmt.Sprintf("invalid expression %T in function body. line: %d, col:%d", node, pos[0], pos[1]))
// 	}
// 	return nil
// }

// func (c *Compiler) GetSpec(id []string) ([]string, *spec) {
// 	// Returns full namespace of variable and the spec it belongs to
// 	// assumes current spec if none specified
// 	if c.specs[id[0]] == nil {
// 		id = append([]string{c.currentSpecName}, id...)
// 	}
// 	return id, c.currentSpec
// }

// func (c *Compiler) ListSpecs() []string {
// 	// Lists all specs the compiler knows about
// 	var specs []string
// 	for k := range c.specs {
// 		specs = append(specs, k)
// 	}
// 	return specs
// }

// func (c *Compiler) ListSpecsAndVars() map[string][]string {
// 	// Lists all specs and their variables
// 	specs := make(map[string][]string)
// 	for k, v := range c.specs {
// 		specs[k] = v.vars.List()
// 	}
// 	return specs
// }

// func (c *Compiler) setConst(id []string, val value.Value) {
// 	if c.isVarSet(id) {
// 		fid := strings.Join(id, "_")
// 		panic(fmt.Sprintf("variable %s is a constant and cannot be reassigned", fid))
// 	}
// 	c.specs[c.currentSpecName].DefineSpecVar(id, val)
// 	c.specs[c.currentSpecName].DefineSpecType(id, val.Type())
// }

// func (c *Compiler) isVarSet(id []string) bool {
// 	id, s := c.GetSpec(id)
// 	return s.GetSpecVar(id) != nil
// }

// func (c *Compiler) isVarSetAssert(id []string) bool {
// 	//If this is for an assert the var might reference
// 	//a rule on the struct level
// 	if c.isVarSet(id) || c.isInstance(id) {
// 		return true
// 	}
// 	return false
// }

// func (c *Compiler) isInstance(id []string) bool {
// 	for _, v := range c.instances {
// 		if v[1] == id[0] {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (c *Compiler) isFunction(node ast.Node) bool {
// 	switch node.(type) {
// 	case *ast.FunctionLiteral:
// 		return true
// 	case *ast.BlockStatement:
// 		return true
// 	default:
// 		return false
// 	}
// }

// func (c *Compiler) validOperator(node *ast.InfixExpression, boolsAllowed bool) bool {
// 	if !boolsAllowed && (node.Left.Type() == "BOOL" || node.Right.Type() == "BOOL") {
// 		return false
// 	}
// 	return true
// }

// func (c *Compiler) generateParameters(structName []string, data map[string]ast.Node, id []string) []*ir.Param {
// 	var p []*ir.Param
// 	var s *spec
// 	keys := c.generateOrder(structName, data)
// 	for _, k := range keys {
// 		switch n := data[k].(type) {
// 		case *ast.Instance:
// 			var ip []*ir.Param
// 			if n.Complex {
// 				ip = c.generateParameters([]string{n.Value.Spec, n.Value.Value}, c.specStructs[n.Value.Spec][n.Value.Value], append(id, k))
// 			} else {
// 				ip = c.generateParameters([]string{n.Value.Spec, n.Value.Value}, c.specStructs[n.Value.Spec][n.Value.Value], []string{id[0], id[1], k})
// 			}
// 			p = append(p, ip...)
// 		default:
// 			if !c.isFunction(n) {
// 				pid := append(id, k)
// 				pid, s = c.GetSpec(pid)
// 				name := strings.Join(pid, "_")
// 				ty := s.GetPointerType(name)
// 				p = append(p, ir.NewParam(name, ty))
// 			}
// 		}
// 	}
// 	return p
// }

// func (c *Compiler) resetParaState(p []*ir.Param) {
// 	for i := 0; i < len(p); i++ {
// 		name := p[i].LocalName
// 		id := strings.Split(name, "_")
// 		id, s := c.GetSpec(id)
// 		s.vars.ResetState(id)
// 	}
// }

// func (c *Compiler) convertAssertVariables(ex ast.Expression) ast.Expression {
// 	switch e := ex.(type) {
// 	case *ast.InfixExpression:

// 		e.Left = c.convertAssertVariables(e.Left)
// 		e.Right = c.convertAssertVariables(e.Right)
// 		return e
// 	case *ast.Identifier:
// 		id := strings.Split(e.Value, "_")
// 		pos := e.Position()
// 		fvn := c.getFullVariableName(id)
// 		id, _ = c.GetSpec(fvn)
// 		fvns := strings.Join(id, "_")

// 		if !c.isVarSetAssert(fvn) {
// 			panic(fmt.Sprintf("cannot send value to variable %s. Variable not defined line: %d, col: %d", fvns, pos[0], pos[1]))
// 		}

// 		instas := c.fetchInstances(fvn)
// 		if len(instas) == 0 {
// 			instas = []string{fvns}
// 		}
// 		return &ast.AssertVar{
// 			Token:        e.Token,
// 			InferredType: e.InferredType,
// 			Instances:    instas,
// 		}
// 	case *ast.ParameterCall:
// 		id := e.Value
// 		pos := e.Position()
// 		fvn := c.getFullVariableName(id)
// 		id, _ = c.GetSpec(id)
// 		fvns := strings.Join(id, "_")

// 		if !c.isVarSetAssert(fvn) {
// 			panic(fmt.Sprintf("cannot send value to variable %s. Variable not defined line: %d, col: %d", fvns, pos[0], pos[1]))
// 		}

// 		instas := c.fetchInstances(fvn)
// 		if len(instas) == 0 {
// 			instas = []string{fvns}
// 		}
// 		return &ast.AssertVar{
// 			Token:        e.Token,
// 			InferredType: e.InferredType,
// 			Instances:    instas,
// 		}

// 	case *ast.AssertVar:
// 		return e
// 	case *ast.IntegerLiteral:
// 		return e
// 	case *ast.FloatLiteral:
// 		return e
// 	case *ast.Boolean:
// 		return e
// 	case *ast.StringLiteral:
// 		return e
// 	case *ast.Natural:
// 		return e
// 	case *ast.Uncertain:
// 		return e
// 	case *ast.Unknown:
// 		return e
// 	case *ast.PrefixExpression:
// 		e.Right = c.convertAssertVariables(e.Right)
// 		return e
// 	case *ast.Nil:
// 		return e
// 	case *ast.IndexExpression:
// 		e.Left = c.convertAssertVariables(e.Left)
// 		return e
// 	default:
// 		pos := e.Position()
// 		panic(fmt.Sprintf("illegal node %T in assert or assume line: %d, col: %d", e, pos[0], pos[1]))
// 	}
// }

// func (c *Compiler) lookupIdent(ident []string, pos []int) *ir.InstLoad {
// 	id, s := c.GetSpec(ident)
// 	name := strings.Join(id, "_")
// 	pointer := c.specGlobals[name]
// 	if pointer != nil {
// 		pointer := c.specGlobals[name]
// 		ty := s.GetSpecType(name)
// 		load := c.contextBlock.NewLoad(ty, pointer)
// 		return load
// 	}

// 	local := s.GetSpecVar(id)
// 	if local != nil {
// 		pointer := s.GetSpecVarPointer(name)
// 		ty := s.GetSpecType(name)
// 		load := c.contextBlock.NewLoad(ty, pointer)
// 		return load
// 	}

// 	// Might be a spec global constant
// 	g := id[len(id)-1]
// 	global := s.GetSpecVar([]string{id[0], g})
// 	if global != nil {
// 		name = strings.Join([]string{id[0], g}, "_")
// 	}
// 	pointer = c.specGlobals[name]
// 	ty := s.GetSpecType(name)
// 	load := c.contextBlock.NewLoad(ty, pointer)
// 	return load
// }

// func (c *Compiler) fetchInstances(ident []string) []string {
// 	// this and convertAssertVariables need a rethink. Seems brittle
// 	// with lots of edge cases
// 	var insta []string
// 	id, _ := c.GetSpec(ident)
// 	for k, v := range c.instanceChildren {
// 		if v == id[1] {
// 			id2 := strings.Split(k, "_")
// 			if id2[len(id2)-1] == id[len(id)-1] { //Same parameter of a different instance
// 				insta = append(insta, k)
// 			}
// 		}
// 	}
// 	return insta
// }

// func negate(e ast.Expression) ast.Expression {
// 	//Negate the expression so that the solver attempts to disprove it
// 	switch n := e.(type) {
// 	case *ast.InfixExpression:
// 		op, ok := OP_NEGATE[n.Operator]
// 		if ok {
// 			//pos := n.Position()
// 			//panic(fmt.Sprintf("operator %s not valid from an assertion. line: %d, col: %d", n.Operator, pos[0], pos[1]))
// 			n.Operator = op
// 		}
// 		n.Left = negate(n.Left)
// 		n.Right = negate(n.Right)

// 		node := evaluate(n) // If Int/Float, evaluate and return the value
// 		return node
// 	case *ast.Boolean:
// 		if n.Value {
// 			n.Value = false
// 		} else {
// 			n.Value = true
// 		}
// 		return n
// 	case *ast.PrefixExpression:
// 		return negate(n.Right)
// 	}
// 	return e
// }

// func negateTemporal(op string, n int) (string, int) {
// 	var op2 string
// 	var n2 int
// 	switch op {
// 	case "nmt":
// 		op2 = "nft"
// 		n2 = n + 1
// 	case "nft":
// 		op2 = "nmt"
// 		n2 = n - 1
// 	}
// 	return op2, n2
// }

// func evaluate(n *ast.InfixExpression) ast.Expression {
// 	if util.IsCompare(n.Operator) {
// 		return n
// 	}
// 	f1, ok1 := n.Left.(*ast.FloatLiteral)
// 	i1, ok2 := n.Left.(*ast.IntegerLiteral)

// 	if !ok1 && !ok2 {
// 		return n
// 	}

// 	f2, ok1 := n.Right.(*ast.FloatLiteral)
// 	i2, ok2 := n.Right.(*ast.IntegerLiteral)

// 	if !ok1 && !ok2 {
// 		return n
// 	}

// 	if f1 != nil {
// 		if f2 != nil {
// 			v := evalFloat(f1.Value, f2.Value, n.Operator)
// 			return &ast.FloatLiteral{
// 				Token: n.Token,
// 				Value: v,
// 			}
// 		} else {
// 			v := evalFloat(f1.Value, float64(i2.Value), n.Operator)
// 			return &ast.FloatLiteral{
// 				Token: n.Token,
// 				Value: v,
// 			}
// 		}
// 	} else {
// 		if f2 != nil {
// 			v := evalFloat(float64(i1.Value), f2.Value, n.Operator)
// 			return &ast.FloatLiteral{
// 				Token: n.Token,
// 				Value: v,
// 			}
// 		} else {
// 			if n.Operator == "/" {
// 				//Return a float in the case of division
// 				v := evalFloat(float64(i1.Value), float64(i2.Value), n.Operator)
// 				return &ast.FloatLiteral{
// 					Token: n.Token,
// 					Value: v,
// 				}
// 			}
// 			v := evalInt(i1.Value, i2.Value, n.Operator)
// 			return &ast.IntegerLiteral{
// 				Token: n.Token,
// 				Value: v,
// 			}
// 		}
// 	}
// }

// func evalFloat(f1 float64, f2 float64, op string) float64 {
// 	switch op {
// 	case "+":
// 		return f1 + f2
// 	case "-":
// 		return f1 - f2
// 	case "*":
// 		return f1 * f2
// 	case "/":
// 		return f1 / f2
// 	default:
// 		panic(fmt.Sprintf("unsupported operator %s", op))
// 	}
// }

// func evalInt(i1 int64, i2 int64, op string) int64 {
// 	switch op {
// 	case "+":
// 		return i1 + i2
// 	case "-":
// 		return i1 - i2
// 	case "*":
// 		return i1 * i2
// 	default:
// 		panic(fmt.Sprintf("unsupported operator %s", op))
// 	}
// }

// type Panic string
