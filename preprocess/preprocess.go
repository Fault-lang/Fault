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
	trail       util.ImportTrail
	structTypes map[string]map[string]string
	Processed   ast.Node
}

func NewProcesser() *Processor {
	return &Processor{
		Specs:       make(map[string]*SpecRecord),
		structTypes: make(map[string]map[string]string),
	}
}

func (p *Processor) Run(n *ast.Spec) *ast.Spec {
	tree, err := p.walk(n)
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

func (p *Processor) walk(n ast.Node) (ast.Node, error) {
	var err error
	var pro ast.Node
	switch node := n.(type) {
	case *ast.Spec:
		for i, v := range node.Statements {
			pro, err = p.walk(v)
			node.Statements[i] = pro.(ast.Statement)
		}
		return node, err
	case *ast.SpecDeclStatement:
		p.Specs[node.Name.Value] = NewSpecRecord()
		p.Specs[node.Name.Value].SpecName = node.Name.Value
		p.trail = p.trail.PushSpec(node.Name.Value)
		return node, err
	case *ast.SysDeclStatement:
		p.Specs[node.Name.Value] = NewSpecRecord()
		p.Specs[node.Name.Value].SpecName = node.Name.Value
		p.trail = p.trail.PushSpec(node.Name.Value)
		return node, err
	case *ast.ImportStatement:
		pro, err = p.walk(node.Tree)
		node.Tree = pro.(*ast.Spec)
		_, p.trail = p.trail.PopSpec()
		return node, err
	case *ast.ConstantStatement:
		var spec *SpecRecord
		if p.Specs[p.trail.CurrentSpec()] != nil {
			spec = p.Specs[p.trail.CurrentSpec()]
		}
		spec.AddConstant(node.Name.Value, node.Value)
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

		var spec *SpecRecord
		properties := util.Preparse(node.Pairs)
		if p.Specs[p.trail.CurrentSpec()] != nil {
			spec = p.Specs[p.trail.CurrentSpec()]
		}
		spec.AddStock(p.scope, properties)
		spec.Index("STOCK", p.scope)
		p.Specs[p.trail.CurrentSpec()] = spec
		pn := p.buildIdContext(spec.Id())
		node.ProcessedName = pn

		for k, v := range properties {
			pro, err = p.walk(v)
			if err != nil {
				return node, err
			}
			ident := node.GetPropertyIdent(k)

			node.Pairs[ident] = pro.(ast.Expression)
		}

		p.scope = ""
		return node, err
	case *ast.FlowLiteral:
		if p.structTypes[p.trail.CurrentSpec()] == nil {
			p.structTypes[p.trail.CurrentSpec()] = make(map[string]string)
		}
		p.structTypes[p.trail.CurrentSpec()][p.scope] = "FLOW"

		var spec *SpecRecord
		properties := util.Preparse(node.Pairs)
		if p.Specs[p.trail.CurrentSpec()] != nil {
			spec = p.Specs[p.trail.CurrentSpec()]
		}
		spec.AddFlow(p.scope, properties)
		spec.Index("FLOW", p.scope)
		p.Specs[p.trail.CurrentSpec()] = spec
		pn := p.buildIdContext(spec.Id())
		node.ProcessedName = pn

		for k, v := range properties {
			pro, err = p.walk(v)
			if err != nil {
				return node, err
			}
			ident := node.GetPropertyIdent(k)
			node.Pairs[ident] = pro.(ast.Expression)
		}

		p.scope = ""
		return node, err
	case *ast.ComponentLiteral:
		if p.structTypes[p.trail.CurrentSpec()] == nil {
			p.structTypes[p.trail.CurrentSpec()] = make(map[string]string)
		}
		p.structTypes[p.trail.CurrentSpec()][p.scope] = "COMPONENT"

		var spec *SpecRecord
		properties := util.Preparse(node.Pairs)
		if p.Specs[p.trail.CurrentSpec()] != nil {
			spec = p.Specs[p.trail.CurrentSpec()]
		}
		spec.AddComponent(p.scope, properties)
		spec.Index("COMPONENT", p.scope)
		p.Specs[p.trail.CurrentSpec()] = spec
		node.ProcessedName = p.buildIdContext(spec.Id())

		for k, v := range properties {
			pro, err = p.walk(v)
			if err != nil {
				return node, err
			}
			ident := node.GetPropertyIdent(k)
			node.Pairs[ident] = pro.(ast.Expression)
		}

		p.scope = ""
		return node, err
	case *ast.AssertionStatement:
		return node, err
	case *ast.AssumptionStatement:
		return node, err
	case *ast.ExpressionStatement:
		pro, err = p.walk(node.Expression)
		node.Expression = pro.(ast.Expression)
		return node, err
	case *ast.ForStatement: //TODO
		for i, v := range node.Body.Statements {
			pro, err = p.walk(v)
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
		pro, err = p.walk(node.Body)
		if err != nil {
			return pro, err
		}
		node.Body = pro.(*ast.BlockStatement)
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
		var key string
		importSpec := p.Specs[node.Value.Spec] //Where the struct definition lives

		if p.Specs[p.trail.CurrentSpec()] == nil {
			p.Specs[p.trail.CurrentSpec()] = NewSpecRecord()
		}

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
		switch p.structTypes[node.Value.Spec][node.Value.Value] {
		case "STOCK":
			reference := importSpec.FetchStock(node.Value.Value)
			spec.AddInstance(key, reference, ty)
			properties = spec.FetchStock(key)

			spec.Index("STOCK", key)
			p.Specs[p.trail.CurrentSpec()] = spec

			oldScope := p.scope
			p.scope = key

			pro := &ast.StructInstance{Token: node.Token,
				Spec: p.trail.CurrentSpec(), Name: key,
				Parent:     node.Value.Value,
				Properties: make(map[string]*ast.StructProperty)}

			pn := p.buildIdContext(spec.Id())
			fmt.Println(p.scope)

			var pro2 ast.Node
			for id, v := range properties {
				// Looking for more instances
				if inst, ok := v.(*ast.Instance); ok {
					inst.ComplexScope = key
					pro2, err = p.walk(inst)
				} else {
					pro2, err = p.walk(v)
				}

				if err != nil {
					return node, err
				}

				token := ast.Token{
					Type:     ast.TokenType("STOCK"),
					Literal:  "STOCK",
					Position: node.Position(),
				}
				property := &ast.StructProperty{Token: token, Value: pro2, Spec: p.trail.CurrentSpec(), Name: id}
				property.ProcessedName = append(pn, id)
				pro.Properties[id] = property
			}
			//node.Processed = pro
			//node.ProcessedName = pn
			pro.ProcessedName = pn
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
				Spec: p.trail.CurrentSpec(), Name: key,
				Parent:     node.Value.Value,
				Properties: make(map[string]*ast.StructProperty)}

			var pro2 ast.Node
			pn := p.buildIdContext(spec.Id())
			for id, v := range properties {
				// Looking for more instances
				if inst, ok := v.(*ast.Instance); ok {
					inst.ComplexScope = key
					pro2, err = p.walk(inst)
				} else {
					pro2, err = p.walk(v)
				}

				if err != nil {
					return node, err
				}

				token := ast.Token{
					Type:     ast.TokenType("FLOW"),
					Literal:  "FLOW",
					Position: node.Position(),
				}
				property := &ast.StructProperty{Token: token, Value: pro2, Spec: p.trail.CurrentSpec(), Name: id}
				property.ProcessedName = append(pn, id)
				pro.Properties[id] = property
			}

			//node.Processed = pro
			//node.ProcessedName = pn
			pro.ProcessedName = pn
			p.scope = oldScope
			return pro, err
		default:
			panic(fmt.Sprintf("invalid instance %s", node.Value.Value))
		}

	case *ast.Identifier:
		spec := p.Specs[node.Spec]
		rawid := p.buildIdContext(spec.Id())
		rawid = append(rawid, node.Value)
		node.ProcessedName = rawid
		return node, err

	case *ast.ParameterCall:
		if node.Value[0] == "this" {
			//Convert this
			p.buildIdContext(p.trail.CurrentSpec())
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
		rawid = append(rawid, node.Value...)
		node.ProcessedName = rawid
		return node, err

	default:
		return node, err
	}
}
