package preprocess

import (
	"fault/ast"
	"fault/util"
	"fmt"
	"strings"
)

type Processor struct {
	Specs       map[string]*SpecRecord
	scope       string
	localIdents map[string][]string
	trail       util.ImportTrail
	structTypes map[string]map[string]string
	Processed   ast.Node
	initialPass bool
	inFunc      bool
}

func NewProcesser() *Processor {
	return &Processor{
		Specs:       make(map[string]*SpecRecord),
		structTypes: make(map[string]map[string]string),
		localIdents: make(map[string][]string),
		initialPass: true,
		inFunc:      false,
	}
}

func (p *Processor) Run(n *ast.Spec) *ast.Spec {
	tree, err := p.walk(n)
	if err != nil {
		panic(err)
	}
	p.initialPass = false

	tree, err = p.walk(tree)
	if err != nil {
		panic(err)
	}
	spec := tree.(*ast.Spec)
	p.Processed = spec
	return spec
}

func (p *Processor) buildIdContext(spec string) []string {
	scopeParts := strings.Split(p.scope, "_")
	if scopeParts[0] == "" {
		return []string{spec}
	}
	return append([]string{spec}, scopeParts...)
}

func (p *Processor) namePairs(pairs map[*ast.Identifier]ast.Expression) (map[*ast.Identifier]ast.Expression, []string) {
	var keys []string
	named := make(map[*ast.Identifier]ast.Expression)
	for k, prs := range pairs {
		keys = append(keys, k.String())
		pn := p.buildIdContext(p.trail.CurrentSpec())
		rawid := append(pn, k.String())
		prs.(ast.Nameable).SetId(rawid)
		k.SetId(rawid)
		named[k] = prs
	}
	return named, keys
}

func (p *Processor) walk(n ast.Node) (ast.Node, error) {
	var err error
	var pro ast.Node

	switch node := n.(type) {
	case *ast.Spec:
		for i, v := range node.Statements {
			pro, err = p.walk(v)
			if err != nil {
				return node, err
			}
			node.Statements[i] = pro.(ast.Statement)
		}
		return node, err
	case *ast.SpecDeclStatement:
		if p.initialPass {
			p.Specs[node.Name.Value] = NewSpecRecord()
			p.Specs[node.Name.Value].SpecName = node.Name.Value
			p.trail = p.trail.PushSpec(node.Name.Value)
		}
		return node, err
	case *ast.SysDeclStatement:
		if p.initialPass {
			p.Specs[node.Name.Value] = NewSpecRecord()
			p.Specs[node.Name.Value].SpecName = node.Name.Value
			p.trail = p.trail.PushSpec(node.Name.Value)
		}
		return node, err
	case *ast.ImportStatement:
		pro, err = p.walk(node.Tree)
		if err != nil {
			return node, err
		}
		node.Tree = pro.(*ast.Spec)
		_, p.trail = p.trail.PopSpec()
		return node, err
	case *ast.ConstantStatement:
		if !p.initialPass {
			return node, err
		}

		var spec *SpecRecord
		if p.Specs[p.trail.CurrentSpec()] != nil {
			spec = p.Specs[p.trail.CurrentSpec()]
		}

		// Has this already been defined?
		if spec.FetchConstant(node.Name.Value) != nil {
			panic(fmt.Sprintf("variable %s is a constant and cannot be modified", node.Name.Value))
		}

		pronm, err := p.walk(node.Name)
		if err != nil {
			return node, err
		}
		node.Name = pronm.(*ast.Identifier)
		pro, err = p.walk(node.Value)
		if err != nil {
			return node, err
		}
		node.Value = pro.(ast.Expression)

		spec.AddConstant(node.Name.Value, pro)
		spec.Index(node.Name.Value, "CONSTANT")
		p.Specs[p.trail.CurrentSpec()] = spec

		return node, err
	case *ast.DefStatement:
		p.scope = strings.TrimSpace(node.Name.String())
		pro, err = p.walk(node.Value)
		if err != nil {
			return node, err
		}
		namepro, err := p.walk(node.Name)
		if err != nil {
			return node, err
		}

		node.Name = namepro.(*ast.Identifier)
		node.Value = pro.(ast.Expression)
		return node, err
	case *ast.StockLiteral:
		if p.structTypes[p.trail.CurrentSpec()] == nil {
			p.structTypes[p.trail.CurrentSpec()] = make(map[string]string)
		}
		p.structTypes[p.trail.CurrentSpec()][p.scope] = "STOCK"

		var properties map[string]ast.Node
		var spec *SpecRecord
		var idx []string
		if p.Specs[p.trail.CurrentSpec()] != nil {
			spec = p.Specs[p.trail.CurrentSpec()]
		}

		if p.initialPass {
			node.Pairs, idx = p.namePairs(node.Pairs)
			local := strings.Join([]string{p.trail.CurrentSpec(), p.scope}, "_")
			p.localIdents[local] = idx

			properties = util.Preparse(node.Pairs)
			spec.AddStock(p.scope, properties)
			spec.Index("STOCK", p.scope)
			p.Specs[p.trail.CurrentSpec()] = spec
			pn := p.buildIdContext(spec.Id())
			node.ProcessedName = pn
		} else {
			properties = spec.FetchStock(p.scope)
		}

		for k, v := range properties {
			p.inFunc = true
			pro, err = p.walk(v)
			if err != nil {
				return node, err
			}
			p.inFunc = false

			ident := node.GetPropertyIdent(k)
			ident.SetId(pro.(ast.Nameable).RawId())

			pron, err := p.walk(ident)
			if err != nil {
				return node, err
			}

			node.Pairs[pron.(*ast.Identifier)] = pro.(ast.Expression)
		}
		properties = util.Preparse(node.Pairs)
		spec.UpdateStock(p.scope, properties)

		p.scope = ""
		return node, err
	case *ast.FlowLiteral:
		if p.structTypes[p.trail.CurrentSpec()] == nil {
			p.structTypes[p.trail.CurrentSpec()] = make(map[string]string)
		}
		p.structTypes[p.trail.CurrentSpec()][p.scope] = "FLOW"

		var properties map[string]ast.Node
		var spec *SpecRecord
		var idx []string
		if p.Specs[p.trail.CurrentSpec()] != nil {
			spec = p.Specs[p.trail.CurrentSpec()]
		}

		if p.initialPass {
			node.Pairs, idx = p.namePairs(node.Pairs)
			local := strings.Join([]string{p.trail.CurrentSpec(), p.scope}, "_")
			p.localIdents[local] = idx

			properties = util.Preparse(node.Pairs)
			spec.AddFlow(p.scope, properties)
			spec.Index("FLOW", p.scope)
			p.Specs[p.trail.CurrentSpec()] = spec
			pn := p.buildIdContext(spec.Id())
			node.ProcessedName = pn
		} else {
			properties = spec.FetchFlow(p.scope)
		}

		for k, v := range properties {
			p.inFunc = true
			pro, err = p.walk(v)
			if err != nil {
				return node, err
			}
			p.inFunc = false

			ident := node.GetPropertyIdent(k)
			ident.SetId(pro.(ast.Nameable).RawId())

			pron, err := p.walk(ident)
			if err != nil {
				return node, err
			}

			node.Pairs[pron.(*ast.Identifier)] = pro.(ast.Expression)
		}
		properties = util.Preparse(node.Pairs)
		spec.UpdateFlow(p.scope, properties)

		p.scope = ""
		return node, err
	case *ast.ComponentLiteral:
		if p.structTypes[p.trail.CurrentSpec()] == nil {
			p.structTypes[p.trail.CurrentSpec()] = make(map[string]string)
		}
		p.structTypes[p.trail.CurrentSpec()][p.scope] = "COMPONENT"

		var properties map[string]ast.Node
		var spec *SpecRecord
		var idx []string
		if p.Specs[p.trail.CurrentSpec()] != nil {
			spec = p.Specs[p.trail.CurrentSpec()]
		}

		if p.initialPass {
			node.Pairs, idx = p.namePairs(node.Pairs)
			local := strings.Join([]string{p.trail.CurrentSpec(), p.scope}, "_")
			p.localIdents[local] = idx

			properties = util.Preparse(node.Pairs)
			spec.AddComponent(p.scope, properties)
			spec.Index("COMPONENT", p.scope)
			p.Specs[p.trail.CurrentSpec()] = spec
			pn := p.buildIdContext(spec.Id())
			node.ProcessedName = pn
		} else {
			properties = spec.FetchComponent(p.scope)
		}

		for k, v := range properties {
			p.inFunc = true
			pro, err = p.walk(v)
			if err != nil {
				return node, err
			}
			p.inFunc = false

			ident := node.GetPropertyIdent(k)
			pron, err := p.walk(ident)
			if err != nil {
				return node, err
			}

			node.Pairs[pron.(*ast.Identifier)] = pro.(ast.Expression)
		}
		properties = util.Preparse(node.Pairs)
		spec.UpdateComponent(p.scope, properties)

		p.scope = ""
		return node, err
	case *ast.AssertionStatement:
		pro, err = p.walk(node.Constraints)
		if err != nil {
			return node, err
		}
		node.Constraints = pro.(*ast.InvariantClause)
		return node, err
	case *ast.AssumptionStatement:
		pro, err = p.walk(node.Constraints)
		if err != nil {
			return node, err
		}
		node.Constraints = pro.(*ast.InvariantClause)
		return node, err
	case *ast.InvariantClause:
		l, err := p.walk(node.Left)
		if err != nil {
			return node, err
		}

		r, err := p.walk(node.Right)
		if err != nil {
			return node, err
		}

		node.Left = l.(ast.Expression)
		node.Right = r.(ast.Expression)
		return node, err

	case *ast.ExpressionStatement:
		pro, err = p.walk(node.Expression)
		if err != nil {
			return node, err
		}
		node.Expression = pro.(ast.Expression)
		return node, err
	case *ast.ForStatement:
		for i, v := range node.Body.Statements {
			pro, err = p.walk(v)
			if err != nil {
				return node, err
			}
			node.Body.Statements[i] = pro.(ast.Statement)
		}
		return node, err
	case *ast.StartStatement: //TODO
		return node, err
	case *ast.StateLiteral:
		pro, err = p.walk(node.Body)
		if err != nil {
			return pro, err
		}
		node.Body = pro.(*ast.BlockStatement)
		return node, err
	case *ast.FunctionLiteral:
		p.inFunc = true
		pro, err = p.walk(node.Body)
		if err != nil {
			return pro, err
		}
		node.Body = pro.(*ast.BlockStatement)
		p.inFunc = false
		return node, err
	case *ast.BlockStatement:
		var statements []ast.Statement
		for _, v := range node.Statements {
			pro, err = p.walk(v)
			if err != nil {
				return node, err
			}
			statements = append(statements, pro.(ast.Statement))
		}
		node.Statements = statements
		return node, err
	case *ast.InfixExpression:
		l, err := p.walk(node.Left)
		if err != nil {
			return node, err
		}

		r, err := p.walk(node.Right)
		if err != nil {
			return node, err
		}

		node.Left = l.(ast.Expression)
		node.Right = r.(ast.Expression)
		return node, err

	case *ast.IndexExpression:

		l, err := p.walk(node.Left)
		if err != nil {
			return node, err
		}
		node.Left = l.(ast.Expression)
		node.ProcessedName = l.(ast.Nameable).RawId()
		return node, err

	case *ast.PrefixExpression:
		r, err := p.walk(node.Right)
		if err != nil {
			return node, err
		}

		node.Right = r.(ast.Expression)

		return node, err

	case *ast.IfExpression:
		cond, err := p.walk(node.Condition)
		if err != nil {
			return node, err
		}

		conseq, err := p.walk(node.Consequence)
		if err != nil {
			return node, err
		}

		var alt ast.Node
		if node.Alternative != nil {
			alt, err = p.walk(node.Alternative)
			if err != nil {
				return node, err
			}
		}

		var elif ast.Node
		if node.Elif != nil {
			elif, err = p.walk(node.Elif)
			if err != nil {
				return node, err
			}
		}

		node.Condition = cond.(ast.Expression)
		node.Consequence = conseq.(*ast.BlockStatement)
		if node.Alternative != nil {
			node.Alternative = alt.(*ast.BlockStatement)
		}
		if node.Elif != nil {
			node.Elif = elif.(*ast.IfExpression)
		}
		return node, err

	case *ast.Instance:
		if p.Specs[p.trail.CurrentSpec()] == nil {
			p.Specs[p.trail.CurrentSpec()] = NewSpecRecord()
		}

		if p.initialPass {
			pro, err = p.walk(node.Value)
			if err != nil {
				return node, err
			}
			pn := p.buildIdContext(p.trail.CurrentSpec())
			node.Value = pro.(*ast.Identifier)
			node.ProcessedName = append(pn, node.Name)
			return node, err
		}

		var key string
		importSpec := p.Specs[node.Value.Spec] //Where the struct definition lives

		spec := p.Specs[p.trail.CurrentSpec()] //Where the instance is being declared

		if node.ComplexScope != "" {
			key = strings.Join([]string{node.ComplexScope, node.Name}, "_")
		} else if p.scope == "" { //For example if it's initialized in the run block
			key = node.Name
		} else {
			key = strings.Join([]string{p.scope, node.Name}, "_")
		}

		ty := p.structTypes[node.Value.Spec][node.Value.Value]
		var properties map[string]ast.Node
		switch ty {
		case "STOCK":
			reference := importSpec.FetchStock(node.Value.Value)
			spec.AddInstance(key, reference, ty)
			properties = spec.FetchStock(key)
			spec.Index("STOCK", key)
			p.Specs[p.trail.CurrentSpec()] = spec

			oldScope := p.scope
			p.scope = key

			pro := &ast.StructInstance{Token: node.Token,
				Spec: p.trail.CurrentSpec(), Name: node.Name,
				Parent:       []string{node.Value.Spec, node.Value.Value},
				Properties:   make(map[string]*ast.StructProperty),
				ComplexScope: node.ComplexScope}

			pro.Token.Literal = "STOCK"

			pn := p.buildIdContext(spec.Id())

			var pro2 ast.Node
			properties2 := make(map[string]ast.Node)
			var order []string
			for id, v := range properties {
				name := append(pn, id)
				order = append(order, id)
				// Looking for more instances
				switch inst := v.(type) {
				case *ast.StructInstance:
					inst.ComplexScope = key
					pro2, err = p.walk(inst)
				case *ast.Instance:
					inst.ComplexScope = key
					pro2, err = p.walk(inst)
				default:
					pro2, err = p.walk(v)
				}

				if err != nil {
					return node, err
				}

				pro2.(ast.Nameable).SetId(name)

				token := ast.Token{
					Type:     ast.TokenType("STOCK"),
					Literal:  "STOCK",
					Position: node.Position(),
				}
				properties2[id] = pro2
				property := &ast.StructProperty{Token: token, Value: pro2, Spec: p.trail.CurrentSpec(), Name: id}
				property.ProcessedName = name
				pro.Properties[id] = property
			}
			spec.UpdateStock(key, properties2)
			pro.ProcessedName = pn
			pro.Order = util.StableSortKeys(order)
			p.scope = oldScope
			return pro, err
		case "FLOW":
			reference := importSpec.FetchFlow(node.Value.Value)
			spec.AddInstance(key, reference, ty)
			properties = spec.FetchFlow(key)

			spec.Index("FLOW", key)
			p.Specs[p.trail.CurrentSpec()] = spec

			oldScope := p.scope
			p.scope = key
			pro := &ast.StructInstance{Token: node.Token,
				Spec: p.trail.CurrentSpec(), Name: node.Name,
				Parent:       []string{node.Value.Spec, node.Value.Value},
				Properties:   make(map[string]*ast.StructProperty),
				ComplexScope: node.ComplexScope}

			pro.Token.Literal = "FLOW"

			var pro2 ast.Node
			properties2 := make(map[string]ast.Node)
			pn := p.buildIdContext(spec.Id())
			var order []string
			for id, v := range properties {
				order = append(order, id)
				name := append(pn, id)
				// Looking for more instances
				switch inst := v.(type) {
				case *ast.StructInstance:
					inst.ComplexScope = key
					pro2, err = p.walk(inst)
				case *ast.Instance:
					inst.ComplexScope = key
					pro2, err = p.walk(inst)
				default:
					pro2, err = p.walk(v)
				}

				if err != nil {
					return node, err
				}

				pro2.(ast.Nameable).SetId(name)

				token := ast.Token{
					Type:     ast.TokenType("FLOW"),
					Literal:  "FLOW",
					Position: node.Position(),
				}
				properties2[id] = pro2
				property := &ast.StructProperty{Token: token, Value: pro2, Spec: p.trail.CurrentSpec(), Name: id}
				property.ProcessedName = name
				pro.Properties[id] = property
			}
			spec.UpdateFlow(key, properties2)
			pro.ProcessedName = pn
			pro.Order = util.StableSortKeys(order)
			p.scope = oldScope
			return pro, err
		default:
			panic(fmt.Sprintf("can't find an instance named %s", node.Value.Value))
		}
	case *ast.StructInstance:
		if p.Specs[p.trail.CurrentSpec()] == nil {
			p.Specs[p.trail.CurrentSpec()] = NewSpecRecord()
		}

		var key string
		importSpec := p.Specs[node.Spec] //Where the struct definition lives

		spec := p.Specs[p.trail.CurrentSpec()] //Where the instance is being declared

		if node.ComplexScope != "" {
			key = strings.Join([]string{node.ComplexScope, node.Name}, "_")
		} else if p.scope == "" { //For example if it's initialized in the run block
			key = node.Name
		} else {
			key = strings.Join([]string{p.scope, node.Name}, "_")
		}

		parent := node.Parent
		pname := strings.Join(parent[1:], "_")
		ty := p.structTypes[parent[0]][pname]
		var properties map[string]ast.Node
		switch ty {
		case "STOCK":
			reference := importSpec.FetchStock(pname)
			spec.AddInstance(key, reference, ty)
			properties = spec.FetchStock(key)

			spec.Index("STOCK", key)
			p.Specs[p.trail.CurrentSpec()] = spec

			oldScope := p.scope
			p.scope = key

			pn := p.buildIdContext(spec.Id())

			var pro2 ast.Node
			properties2 := make(map[string]ast.Node)
			for id, v := range properties {
				name := append(pn, id)
				// Looking for more instances
				switch inst := v.(type) {
				case *ast.StructInstance:
					inst.ComplexScope = key
					pro2, err = p.walk(inst)
				case *ast.Instance:
					inst.ComplexScope = key
					pro2, err = p.walk(inst)
				default:
					pro2, err = p.walk(v)
				}

				if err != nil {
					return node, err
				}

				pro2.(ast.Nameable).SetId(name)

				token := ast.Token{
					Type:     ast.TokenType("STOCK"),
					Literal:  "STOCK",
					Position: node.Position(),
				}
				properties2[id] = pro2
				property := &ast.StructProperty{Token: token, Value: pro2, Spec: p.trail.CurrentSpec(), Name: id}
				property.ProcessedName = name
				node.Properties[id] = property
			}
			spec.UpdateStock(key, properties2)
			node.ProcessedName = pn
			p.scope = oldScope
			return node, err
		case "FLOW":
			reference := importSpec.FetchFlow(pname)
			spec.AddInstance(key, reference, ty)
			properties = spec.FetchFlow(key)

			spec.Index("FLOW", key)
			p.Specs[p.trail.CurrentSpec()] = spec

			oldScope := p.scope
			p.scope = key

			var pro2 ast.Node
			properties2 := make(map[string]ast.Node)
			pn := p.buildIdContext(spec.Id())
			for id, v := range properties {
				name := append(pn, id)
				// Looking for more instances
				switch inst := v.(type) {
				case *ast.StructInstance:
					inst.ComplexScope = key
					pro2, err = p.walk(inst)
				case *ast.Instance:
					inst.ComplexScope = key
					pro2, err = p.walk(inst)
				default:
					pro2, err = p.walk(v)
				}

				if err != nil {
					return node, err
				}

				pro2.(ast.Nameable).SetId(name)

				token := ast.Token{
					Type:     ast.TokenType("FLOW"),
					Literal:  "FLOW",
					Position: node.Position(),
				}
				properties2[id] = pro2
				property := &ast.StructProperty{Token: token, Value: pro2, Spec: p.trail.CurrentSpec(), Name: id}
				property.ProcessedName = name

				node.Properties[id] = property
			}
			spec.UpdateFlow(key, properties2)
			node.ProcessedName = pn
			p.scope = oldScope
			return node, err
		default:
			panic(fmt.Sprintf("can't find an struct instance named %s", node.Parent))
		}
	case *ast.Identifier:
		if !p.initialPass {
			return node, err
		}

		spec := p.Specs[node.Spec]
		rawid := p.buildIdContext(spec.Id())

		if p.inFunc {
			// If this variable is not defined in the struct properties
			// (it's scope) then it's a global variable and should be
			// named that way
			local := strings.Join(rawid, "_")
			if util.InStringSlice(p.localIdents[local], node.Value) {
				rawid = append(rawid, node.Value)
			} else {
				rawid = append(rawid[0:1], node.Value) // spec_struct_var -> spec_var
			}
		} else {
			rawid = append(rawid, node.Value)
		}

		node.ProcessedName = rawid
		return node, err

	case *ast.Unknown:
		if !p.initialPass {
			return node, err
		}

		spec := p.Specs[node.Name.Spec]
		rawid := p.buildIdContext(spec.Id())

		rawid = append(rawid, node.Name.Value)

		node.ProcessedName = rawid
		return node, err

	case *ast.ParameterCall:
		if p.initialPass {
			return node, err
		}
		if node.Value[0] == "this" {
			//Convert this
			rawid := p.buildIdContext(p.trail.CurrentSpec())
			rawid = append(rawid, node.Value[1:]...)
			node2 := &ast.This{
				Token:         node.Token,
				Value:         append([]string{rawid[len(rawid)-1]}, node.Value[1:]...),
				ProcessedName: rawid,
			}
			return node2, err
		}
		spec := p.Specs[node.Spec]
		rawid := p.buildIdContext(spec.Id())

		if rawid[len(rawid)-1] == node.Value[0] {
			//Happens when it's being called from the run block
			rawid = append(rawid, node.Value[1])
		} else {
			rawid = append(rawid, node.Value...)
		}

		node.ProcessedName = rawid

		// If the call is to a function we need to make
		// sure the names of variables inside the function
		// reflect the namespace scope (eg calling from
		// instances created in the runblock)

		ty, _ := spec.GetStructType(rawid)
		branch := spec.FetchVar(rawid, ty)
		if fn, ok := branch.(*ast.FunctionLiteral); ok {
			fn2, err := p.walk(fn)
			if err != nil {
				return node, err
			}
			proFn := fn2.(*ast.FunctionLiteral)
			spec.UpdateVar(rawid, ty, proFn)
		}

		return node, err

	case *ast.ParallelFunctions:
		if p.initialPass {
			return node, err
		}

		for i, v := range node.Expressions {
			// Not sure we ever want this to be anything
			// other than a call actually :/
			scope := v.(*ast.ParameterCall).Value
			oldScope := p.scope
			p.scope = strings.Join(scope[0:len(scope)-1], "_")

			n, err := p.walk(v)
			if err != nil {
				return node, err
			}
			node.Expressions[i] = n.(ast.Expression)
			p.scope = oldScope
		}
		return node, err

	default:
		return node, err
	}
}
