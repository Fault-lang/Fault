package preprocess

import (
	"fault/ast"
	"fault/listener"
	"fault/util"
	"fmt"
	"strings"
)

type Processor struct {
	Specs                map[string]*SpecRecord
	scope                string
	localIdents          map[string][]string
	trail                util.ImportTrail
	structTypes          map[string]map[string]string
	Processed            *ast.Spec
	initialPass          bool
	inFunc               bool
	inStruct             string
	inState              string
	inGlobal             bool
	StructsPropertyOrder map[string][]string
	Instances            map[string]*ast.StructInstance
}

func NewProcesser() *Processor {
	return &Processor{
		Specs:                make(map[string]*SpecRecord),
		structTypes:          make(map[string]map[string]string),
		localIdents:          make(map[string][]string),
		initialPass:          true,
		inFunc:               false,
		inGlobal:             false,
		StructsPropertyOrder: make(map[string][]string),
		Instances:            make(map[string]*ast.StructInstance),
	}
}

func Execute(l *listener.FaultListener) *Processor {
	pre := NewProcesser()
	pre.StructsPropertyOrder = l.StructsPropertyOrder
	pre.Run(l.AST)
	return pre
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

func (p *Processor) Partial(spec string, node ast.Node) (ast.Node, error) {
	p.initialPass = true
	p.trail = p.trail.PushSpec(spec)
	return p.walk(node)
}

func (p *Processor) buildIdContext(spec string) []string {
	if p.inState != "" {
		return []string{spec}
	}
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

func (p *Processor) collapsibleIf(n ast.Node) bool {
	//Replace nested ifs with if A && B construction
	// checks that the block has only one statement
	// and that statement is another conditional
	switch b := n.(type) {
	case *ast.BlockStatement:
		if len(b.Statements) != 1 {
			return false
		}
		line, ok := b.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			return false
		}
		if _, ok := line.Expression.(*ast.IfExpression); ok {
			return true
		}
		return false
	case *ast.IfExpression:
		if len(b.Consequence.Statements) != 1 {
			return false
		}

		line, ok := b.Consequence.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			return false
		}
		if _, ok := line.Expression.(*ast.IfExpression); ok {
			return true
		}
		return false
	}
	return false
}

func (p *Processor) collapse(parent *ast.IfExpression, child *ast.IfExpression) *ast.IfExpression {
	node := &ast.IfExpression{}
	// Add the condition of parent to child
	cond := &ast.InfixExpression{Left: parent.Condition, Right: child.Condition, Operator: "&&"}
	node.Condition = cond
	node.Consequence = child.Consequence

	// If elif block add condition of parent
	if child.Elif != nil {
		el := p.collapse(parent, child.Elif)
		node.Elif = p.attachElif(parent, el)
	}

	// If Else, convert to elif block and add condition of parent
	if child.Alternative != nil {
		el := &ast.IfExpression{Condition: parent.Condition, Consequence: child.Alternative}
		node.Elif = p.attachElif(parent, el)
	}
	return node
}

func (p *Processor) attachElif(n *ast.IfExpression, el *ast.IfExpression) *ast.IfExpression {
	if n.Elif != nil && n.Elif.Condition.String() == el.Condition.String() {
		return n
	} else if n.Elif != nil {
		return p.attachElif(n.Elif, el)
	}
	n.Elif = el
	return n
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
		} else {
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
		_, err := spec.FetchConstant(node.Name.Value)
		if err == nil {
			return node, fmt.Errorf("variable %s is a constant and cannot be modified", node.Name.Value)
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
		if node.TokenLiteral() != "GLOBAL" {
			p.scope = strings.TrimSpace(node.Name.String())
		} else {
			p.inGlobal = true
		}

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

		if node.TokenLiteral() == "GLOBAL" {
			p.inGlobal = false
		} else {
			p.scope = ""
		}

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
			properties, err = spec.FetchStock(p.scope)
			if err != nil {
				return node, err
			}
		}

		for _, k := range node.Order {
			v := properties[k]
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
			properties, err = spec.FetchFlow(p.scope)
			if err != nil {
				return node, err
			}
		}

		for _, k := range node.Order {
			v := properties[k]
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
			properties, err = spec.FetchComponent(p.scope)
			if err != nil {
				return node, err
			}
		}

		for _, k := range node.Order {
			v := properties[k]
			p.inFunc = true
			p.inState = k
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
			p.inState = ""
			node.Pairs[pron.(*ast.Identifier)] = pro.(ast.Expression)
		}
		properties = util.Preparse(node.Pairs)
		spec.UpdateComponent(p.scope, properties)

		p.scope = ""
		return node, err
	case *ast.AssertionStatement:
		pro, err = p.walk(node.Constraint)
		if err != nil {
			return node, err
		}
		node.Constraint = pro.(*ast.InvariantClause)
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
		for i, v := range node.Inits.Statements {
			pro, err = p.walk(v)
			if err != nil {
				return node, err
			}
			node.Inits.Statements[i] = pro.(ast.Statement)
		}
		for i, v := range node.Body.Statements {
			pro, err = p.walk(v)
			if err != nil {
				return node, err
			}
			node.Body.Statements[i] = pro.(ast.Statement)
		}
		return node, err
	case *ast.StartStatement:
		return node, err
	case *ast.FunctionLiteral:
		oldStruct := p.inStruct
		rawid := node.RawId()
		p.inStruct = rawid[1]

		p.inFunc = true
		pro, err = p.walk(node.Body)
		if err != nil {
			return pro, err
		}
		node.Body = pro.(*ast.BlockStatement)
		p.inFunc = false
		p.inStruct = oldStruct
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
		if node.TokenLiteral() == "SWAP" {
			return p.walkSwap(node)
		}

		l, err := p.walk(node.Left)
		if err != nil {
			if node.Token.Type == "ASSIGN" {
				return node, fmt.Errorf("illegal assignment %s", node.String())
			}
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
		rawid := l.(ast.Nameable).RawId()
		rawid = append(rawid, node.Index.String())
		node.ProcessedName = rawid
		return node, err

	case *ast.PrefixExpression:
		r, err := p.walk(node.Right)
		if err != nil {
			return node, err
		}

		node.Right = r.(ast.Expression)

		return node, err

	case *ast.IfExpression:
		pro := &ast.IfExpression{}
		cond, err := p.walk(node.Condition)
		if err != nil {
			return node, err
		}

		conseq, err := p.walk(node.Consequence)
		if err != nil {
			return node, err
		}

		if p.collapsibleIf(conseq) {
			pro = p.collapse(node, conseq.(*ast.BlockStatement).Statements[0].(*ast.ExpressionStatement).Expression.(*ast.IfExpression))
		} else {
			pro.Condition = cond.(ast.Expression)
			pro.Consequence = conseq.(*ast.BlockStatement)
		}

		var elif ast.Node
		if node.Elif != nil {
			elif, err = p.walk(node.Elif) // Since Elif is an IfExpression we already check for collapsibility
			if err != nil {
				return node, err
			}
			pro.Elif = elif.(*ast.IfExpression)
		}

		var alt ast.Node
		if node.Alternative != nil {
			alt, err = p.walk(node.Alternative)
			if err != nil {
				return node, err
			}

			if p.collapsibleIf(alt) {
				el := alt.(*ast.BlockStatement).Statements[0].(*ast.ExpressionStatement).Expression.(*ast.IfExpression)
				if node.Elif == nil {
					pro.Elif = el
					pro.Alternative = nil
				} else {
					pro.Elif = p.attachElif(node.Elif, el)
					pro.Alternative = nil
				}
			} else {
				pro.Alternative = alt.(*ast.BlockStatement)
			}
		}
		return pro, err

	case *ast.Instance:
		if p.Specs[p.trail.CurrentSpec()] == nil {
			p.Specs[p.trail.CurrentSpec()] = NewSpecRecord()
		}

		order := node.Order

		if p.initialPass {
			pro, err = p.walk(node.Value)
			if err != nil {
				return node, err
			}
			pn := p.buildIdContext(p.trail.CurrentSpec())
			pn = append(pn, node.Name)
			node.Value = pro.(*ast.Identifier)
			node.ProcessedName = pn

			if !node.Complex && p.inGlobal {
				local := strings.Join(pn[0:2], "_")
				p.localIdents[local] = order
			}

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

		var swaps []ast.Node
		oldScope := p.scope
		p.scope = "" //Just so we can do swaps
		for _, s := range node.Swaps {

			sw, err := p.walk(s)
			if err != nil {
				return node, err
			}
			swaps = append(swaps, sw)
		}
		p.scope = key

		ty := p.structTypes[node.Value.Spec][node.Value.Value]
		var properties map[string]ast.Node
		switch ty {
		case "STOCK":
			reference, err := importSpec.FetchStock(node.Value.Value)
			if err != nil {
				return node, err
			}

			spec.AddInstance(key, reference, ty)
			properties, err = spec.FetchStock(key)
			if err != nil {
				return node, err
			}

			spec.Index("STOCK", key)
			p.Specs[p.trail.CurrentSpec()] = spec

			if len(order) == 0 { //Sometimes happens if Instance node is referenced before struct is def
				strkey := strings.Join([]string{node.Value.Spec, node.Value.Value}, "_")
				order = p.StructsPropertyOrder[strkey]
			}

			pro := &ast.StructInstance{Token: node.Token,
				Spec: p.trail.CurrentSpec(), Name: node.Name,
				Parent:       []string{node.Value.Spec, node.Value.Value},
				Properties:   make(map[string]*ast.StructProperty),
				ComplexScope: node.ComplexScope,
				Swaps:        swaps,
				Order:        order}

			pro.Token.Literal = "STOCK"

			pn := p.buildIdContext(spec.Id())

			var pro2 ast.Node
			properties2 := make(map[string]ast.Node)

			for _, id := range order {
				v := properties[id]
				name := append(pn, id)
				// Looking for more instances
				switch inst := v.(type) {
				case *ast.StructInstance:
					inst.ComplexScope = key
					pro2, err = p.walk(inst)
					pro2.(ast.Nameable).SetId(name)
					p.Instances[pro2.(ast.Nameable).IdString()] = pro2.(*ast.StructInstance)
				case *ast.Instance:
					inst.ComplexScope = key
					pro2, err = p.walk(inst)
					pro2.(ast.Nameable).SetId(name)
					p.Instances[pro2.(ast.Nameable).IdString()] = pro2.(*ast.StructInstance)
				default:
					pro2, err = p.walk(v)
					pro2.(ast.Nameable).SetId(name)
				}

				if err != nil {
					return node, err
				}

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
			p.Instances[node.IdString()] = pro

			p.scope = oldScope
			return pro, err
		case "FLOW":
			reference, err := importSpec.FetchFlow(node.Value.Value)
			if err != nil {
				return node, err
			}

			spec.AddInstance(key, reference, ty)
			properties, err = spec.FetchFlow(key)
			if err != nil {
				return node, err
			}

			spec.Index("FLOW", key)
			p.Specs[p.trail.CurrentSpec()] = spec

			if len(order) == 0 {
				strkey := strings.Join([]string{node.Value.Spec, node.Value.Value}, "_")
				order = p.StructsPropertyOrder[strkey]
			}

			pro := &ast.StructInstance{Token: node.Token,
				Spec: p.trail.CurrentSpec(), Name: node.Name,
				Parent:       []string{node.Value.Spec, node.Value.Value},
				Properties:   make(map[string]*ast.StructProperty),
				Swaps:        swaps,
				ComplexScope: node.ComplexScope,
				Order:        order}

			pro.Token.Literal = "FLOW"

			var pro2 ast.Node
			properties2 := make(map[string]ast.Node)
			pn := p.buildIdContext(spec.Id())

			for _, id := range order {
				v := properties[id]
				name := append(pn, id)
				// Looking for more instances
				switch inst := v.(type) {
				case *ast.StructInstance:
					inst.ComplexScope = key
					pro2, err = p.walk(inst)
					pro2.(ast.Nameable).SetId(name)
					p.Instances[pro2.(ast.Nameable).IdString()] = pro2.(*ast.StructInstance)
				case *ast.Instance:
					inst.ComplexScope = key
					pro2, err = p.walk(inst)
					pro2.(ast.Nameable).SetId(name)
					p.Instances[pro2.(ast.Nameable).IdString()] = pro2.(*ast.StructInstance)

				default:
					pro2, err = p.walk(v)
					pro2.(ast.Nameable).SetId(name)
				}

				if err != nil {
					return node, err
				}

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
			p.Instances[node.IdString()] = pro

			p.scope = oldScope
			return pro, err
		default:
			return node, fmt.Errorf("can't find an instance named %s", node.Value.Value)
		}
	case *ast.StructInstance:
		if p.Specs[p.trail.CurrentSpec()] == nil {
			p.Specs[p.trail.CurrentSpec()] = NewSpecRecord()
		}

		order := node.Order

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
			reference, err := importSpec.FetchStock(pname)
			if err != nil {
				return node, err
			}
			spec.AddInstance(key, reference, ty)

			if len(node.Properties) > 0 {
				properties = util.ExtractBranches(node.Properties)
			} else {
				properties, err = spec.FetchStock(key)
				if err != nil {
					return node, err
				}

			}

			spec.Index("STOCK", key)
			p.Specs[p.trail.CurrentSpec()] = spec

			oldScope := p.scope
			p.scope = key

			pn := p.buildIdContext(spec.Id())

			var pro2 ast.Node
			properties2 := make(map[string]ast.Node)
			for _, id := range order {
				v := properties[id]
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
			reference, err := importSpec.FetchFlow(pname)
			if err != nil {
				return node, err
			}

			spec.AddInstance(key, reference, ty)
			if len(node.Properties) > 0 {
				properties = util.ExtractBranches(node.Properties)
			} else {
				properties, err = spec.FetchFlow(key)
				if err != nil {
					return node, err
				}
			}

			spec.Index("FLOW", key)
			p.Specs[p.trail.CurrentSpec()] = spec

			oldScope := p.scope
			p.scope = key

			var pro2 ast.Node
			properties2 := make(map[string]ast.Node)
			pn := p.buildIdContext(spec.Id())
			for _, id := range order {
				v := properties[id]
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
			return node, fmt.Errorf("can't find a struct instance named %s", node.Parent)
		}
	case *ast.Identifier:
		var spec *SpecRecord

		// Check to see if this is a constant from
		// an import
		im := p.Specs[node.Spec]
		_, check := im.FetchConstant(node.Value)
		if check == nil {
			spec = im
		} else {
			spec = p.Specs[p.trail.CurrentSpec()]
		}
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

		spec := p.Specs[p.trail.CurrentSpec()]
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
			rawid := append([]string{p.trail.CurrentSpec(), p.scope}, node.Value[1:]...)
			node2 := &ast.This{
				Token:         node.Token,
				Value:         rawid[len(rawid)-2:],
				ProcessedName: rawid,
			}
			return node2, err
		}

		var rawid []string
		var spec *SpecRecord
		spec = p.Specs[p.trail.CurrentSpec()]
		rawid = p.buildIdContext(p.trail.CurrentSpec())

		rawid = append(rawid, node.Value...)

		node.ProcessedName = rawid

		// If the call is to a function we need to make
		// sure the names of variables inside the function
		// reflect the namespace scope (eg calling from
		// instances created in the runblock)

		ty, _ := spec.GetStructType(rawid)

		if ty == "NIL" && p.inStruct != "" { // We might be in a function
			rawid2 := append([]string{rawid[0]}, p.inStruct)
			rawid = append(rawid2, rawid[1:]...) //In which case we can find the node by inserting the struct name
			ty, _ = spec.GetStructType(rawid)

			node.ProcessedName = rawid
		}

		branch, err := spec.FetchVar(rawid, ty)
		if err != nil {
			return node, err
		}

		// State charts tend to create endless loops by design
		// short-curcuit if we've already processed this node
		brName := branch.(ast.Nameable).RawId()
		if brName[0] == rawid[0] {
			return node, err
		}

		if fn, ok := branch.(*ast.FunctionLiteral); ok {
			var oldScope string
			var oldState string
			if ty == "COMPONENT" {
				oldScope = p.scope
				oldState = p.inState
				p.scope = rawid[1]
				p.inState = rawid[2]
			}

			fn2, err := p.walk(fn)
			if err != nil {
				return node, err
			}
			proFn := fn2.(*ast.FunctionLiteral)
			spec.UpdateVar(rawid, ty, proFn)

			if ty == "COMPONENT" {
				p.scope = oldScope
				p.inState = oldState
			}
		}

		return node, err

	case *ast.ParallelFunctions:
		if p.initialPass {
			return node, err
		}

		for i, v := range node.Expressions {
			// Not sure we ever want this to be anything
			// other than a call actually :/

			n, err := p.walk(v)
			if err != nil {
				return node, err
			}
			node.Expressions[i] = n.(ast.Expression)
		}
		return node, err

	case *ast.BuiltIn:
		if p.initialPass {
			return node, err
		}

		spec := p.Specs[p.trail.CurrentSpec()]
		rawid := []string{spec.Id()}
		rawid = append(rawid, p.scope, p.inState, node.Function)
		node.FromState = p.inState
		node.ProcessedName = rawid

		if node.Function == "advance" {
			pro, err := p.walk(node.Parameters["toState"])
			if err != nil {
				return node, err
			}
			node.Parameters["toState"] = pro.(ast.Operand)
		}

		return node, err
	default:
		return node, err
	}
}

func (p *Processor) walkSwap(node *ast.InfixExpression) (ast.Node, error) {
	left, err := p.swapNode(node.Left)
	if err != nil {
		return node, err
	}

	right, err := p.swapNode(node.Right)
	if err != nil {
		return node, err
	}

	node.Left = left.(ast.Expression)
	node.Right = right.(ast.Expression)
	return node, err
}

func (p *Processor) swapNode(node ast.Node) (ast.Node, error) {
	var err error
	if n, ok := node.(*ast.ParameterCall); ok {
		if p.initialPass {
			return n, err
		}
		if n.Value[0] == "this" {
			return nil, fmt.Errorf("incorrect left side value %s", n.Value[0])
		}

		var rawid []string
		rawid = p.buildIdContext(p.trail.CurrentSpec())

		rawid = append(rawid, n.Value...)

		n.ProcessedName = rawid
		return n, err
	}
	return p.walk(node)
}

func alreadyNamed(n1 []string, n2 []string) bool {
	if len(n1) != len(n2) {
		return false
	}
	for i, v := range n1 {
		if v != n2[i] {
			return false
		}
	}
	return true
}
