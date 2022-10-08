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

	currentSpec     string
	specs           map[string]*spec
	structPropOrder map[string][]string

	contextBlock    *ir.Block
	contextFunc     *ir.Func
	contextFuncName string

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

		allocatedPointers: make([]map[string]*ir.InstAlloca, 0),

		contextCondAfter: make([]*ir.Block, 0),
		structPropOrder:  make(map[string][]string),

		runRound:        0,
		builtIns:        make(map[string]*ir.Func),
		specStructs:     make(map[string]*preprocess.SpecRecord),
		specFunctions:   make(map[string]value.Value),
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
			id := []string{c.currentSpec, p[0], p[1]}
			c.processFunc(id, []string{id[0], p[0]}, 0)
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

func (c *Compiler) setConst(id []string, val value.Value) {
	if c.isVarSet(id) {
		fid := strings.Join(id, "_")
		panic(fmt.Sprintf("variable %s is a constant and cannot be reassigned", fid))
	}
	c.specs[c.currentSpec].DefineSpecVar(id, val)
	c.specs[c.currentSpec].DefineSpecType(id, val.Type())
}

func (c *Compiler) compileStruct(def *ast.DefStatement) {
	id := def.Name.ProcessedName
	key := strings.Join(id, "_")
	switch def.Type() {
	case "FLOW":
		c.structPropOrder[key] = def.Value.(*ast.FlowLiteral).Order
	case "STOCK":
		c.structPropOrder[key] = def.Value.(*ast.StockLiteral).Order
	case "GLOBAL":
		instance, _ := def.Value.(*ast.Instance)
		c.compileInstance(instance, strings.Join(id, "_"))
	case "COMPONENT":
		c.structPropOrder[key] = def.Value.(*ast.ComponentLiteral).Order
		c.compileComponent(def.Value.(*ast.ComponentLiteral), id)
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

func (c *Compiler) compileComponent(node *ast.ComponentLiteral, cname string) {
	for key, p := range node.Pairs {
		var pname string
		scopeName := []string{cname, key.Value}

		//oldScope := c.currScope
		oldBlock := c.contextBlock

		c.contextFuncName = strings.Join(scopeName, "_")
		//c.currScope = scopeName

		switch v := p.(type) {
		case *ast.StateLiteral:
			params := c.generateParameters([]string{key.Spec, cname}, c.specStructs[key.Spec][cname], []string{key.Spec, cname, key.Value})
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
			c.currScope = oldScope
			c.contextFunc = nil
		case *ast.Instance:
			c.compileInstance(v, cname)
		default:
			val := c.compileValue(v)

			if val != nil {
				id := []string{cname, key.Value}
				id, s := c.GetSpec(id)
				if s.GetSpecVar(id) != nil {
					name := strings.Join(id, "_")
					pointer := s.GetSpecVarPointer(name)
					ty := s.GetSpecType(name)
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

func (c *Compiler) compileFunctionBody(node ast.Expression) value.Value {
	if !c.alloc { //Short circuit this if just initializing
		return nil
	}
	switch v := node.(type) {
	case *ast.InfixExpression:
		return c.compileInfix(v)

	case *ast.PrefixExpression:
		c.compilePrefix(v)

	case *ast.IfExpression:
		c.compileIf(v)

	case *ast.Instance:
		// orign := v.Value.Spec
		// origv := v.Value.Value
		// v.Value.Value = orign
		// v.Name = origv
		c.compileInstance(v, v.Name)

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
				id = n.ProcessedName
			case *ast.ParameterCall:
				id = n.ProcessedName
			}
			if c.isVarSet(id) && c.alloc {
				p := s.GetSpecVarPointer(id)
				c.contextBlock.NewStore(r, p)
				return nil
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
		switch n := node.Left.(type) {
		case *ast.Identifier:
			id = n.ProcessedName
		case *ast.ParameterCall:
			id = n.ProcessedName
		}

		if !c.isVarSet(id) {
			panic(fmt.Sprintf("cannot send value to variable %s. Variable not defined line: %d, col: %d", strings.Join(id, "_"), pos[0], pos[1]))
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
		id := v.ProcessedName
		return c.lookupIdent(id, node.Position())
	default:
		return c.compileValue(node)
	}
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

func (c *Compiler) lookupIdent(id []string, pos []int) *ir.InstLoad {
	s := c.specs[id[0]]
	name := strings.Join(id, "_")

	local := s.GetSpecVar(id)
	if local != nil {
		pointer := s.GetSpecVarPointer(id)
		ty := s.GetSpecType(name)
		load := c.contextBlock.NewLoad(ty, pointer)
		return load
	}

	pointer := c.specGlobals[name]
	if pointer != nil {
		pointer := c.specGlobals[name]
		ty := s.GetSpecType(name)
		load := c.contextBlock.NewLoad(ty, pointer)
		return load
	}
	return nil
}

func (c *Compiler) processFunc(id []string, structName []string, round int /*pos []int*/) value.Value {
	fname := strings.Join(id, "_")

	if round == 0 { //initialize
		spec := c.specStructs[id[0]]
		ty := spec.GetStructType(id)
		branch := spec.Fetch(id[1], ty)
		params := c.generateParameters(id, branch)
		c.resetParaState(params)
		f := c.module.NewFunc(fname, irtypes.Void, params...)
		c.contextFunc = f

		//oldScope := c.currScope
		oldBlock := c.contextBlock

		c.contextFuncName = fname
		//c.currScope = []string{id[1]} // NOT necessarily the same as structName
		c.contextBlock = f.NewBlock(name.Block())

		val := c.compileValue(branch[id[2]])
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

func (c *Compiler) generateParameters(id []string, data map[string]ast.Node) []*ir.Param {
	var p []*ir.Param
	var s *spec
	keys := c.fetchOrder(id)
	sr := c.specStructs[id[0]]
	for _, k := range keys {
		switch n := data[k].(type) {
		case *ast.Instance:
			var ip []*ir.Param
			child := n.ProcessedName
			if n.Complex {
				ip = c.generateParameters(child, sr.Fetch(child[1], n.InferredType.Type))
			} else {
				ip = c.generateParameters(n.ProcessedName, sr.Fetch(child[1], n.InferredType.Type))
			}
			p = append(p, ip...)
		default:
			if n2, ok := n.(*ast.StructProperty); ok {
				child := n2.ProcessedName
				name := strings.Join(child, "_")
				if n2.Value.Type() != "FUNCTION" {
					ty := s.GetPointerType(name)
					p = append(p, ir.NewParam(name, ty))
				}
			}
		}
	}
	return p
}

func (c *Compiler) fetchOrder(id []string) []string {
	key := strings.Join(id, "_")
	if c.structPropOrder[key] != nil {
		return c.structPropOrder[key]
	}
	panic(fmt.Sprintf("no property order found for struct %s ", key))
}

func (c *Compiler) resetParaState(p []*ir.Param) {
	for i := 0; i < len(p); i++ {
		id := p[i].LocalName
		parts := strings.Split(id, "_")
		s := c.specs[parts[0]]
		s.vars.ResetState(id)
	}
}

func (c *Compiler) isVarSet(id []string) bool {
	s := c.specStructs[id[0]]
	if s.FetchStock(id[1]) != nil {
		return true
	}
	if s.FetchFlow(id[1]) != nil {
		return true
	}
	if s.FetchConstant(id[1]) != nil {
		return true
	}
	if s.FetchComponent(id[1]) != nil {
		return true
	}
	return false
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

func (c *Compiler) GetIR() string {
	return c.module.String()
}

func (c *Compiler) addGlobal() {
	//global := NewCompiledSpec("__global")

	//c.specs["__global"] = global

	// run block
	c.contextFunc = c.module.NewFunc("__run", irtypes.Void)
	mainBlock := c.contextFunc.NewBlock(name.Block())
	mainBlock.NewRet(nil)
	c.contextBlock = mainBlock
}

type Panic string
