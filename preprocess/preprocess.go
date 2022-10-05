package preprocess

import (
	"fault/ast"
	"fault/types"
	"fault/util"
	"fmt"
	"strings"
)

type Processor struct {
	Specs       map[string]*SpecRecord
	scope       string
	trail       types.ImportTrail
	structTypes map[string]map[string]string
}

func NewProcesser() *Processor {
	return &Processor{
		Specs:       make(map[string]*SpecRecord),
		structTypes: make(map[string]map[string]string),
	}
}

func (p *Processor) Run(n *ast.Spec) error {
	return p.walk(n)
}

func (p *Processor) walk(n ast.Node) error {
	var err error
	switch node := n.(type) {
	case *ast.Spec:
		for _, v := range node.Statements {
			err = p.walk(v)
		}
		return err
	case *ast.SpecDeclStatement:
		p.Specs[node.Name.Value] = NewSpecRecord()
		p.trail = p.trail.PushSpec(node.Name.Value)
		return nil
	case *ast.SysDeclStatement:
		p.Specs[node.Name.Value] = NewSpecRecord()
		p.trail = p.trail.PushSpec(node.Name.Value)
		return nil
	case *ast.ImportStatement:
		err = p.walk(node.Tree)
		return err
	case *ast.ConstantStatement:
		var spec *SpecRecord
		if p.Specs[p.trail.CurrentSpec()] != nil {
			spec = p.Specs[p.trail.CurrentSpec()]
		}
		spec.AddConstant(node.Name.Value, node.Value)
		p.Specs[p.trail.CurrentSpec()] = spec
		return err
	case *ast.DefStatement:
		p.scope = strings.TrimSpace(node.Name.String())
		err = p.walk(node.Value)
		return err
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
		p.Specs[p.trail.CurrentSpec()] = spec

		for _, v := range properties {
			err = p.walk(v)
			if err != nil {
				return err
			}
		}

		p.scope = ""
		return err
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
		p.Specs[p.trail.CurrentSpec()] = spec

		for _, v := range properties {
			err = p.walk(v)
			if err != nil {
				return err
			}
		}

		p.scope = ""
		return err
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
		p.Specs[p.trail.CurrentSpec()] = spec

		for _, v := range properties {
			err = p.walk(v)
			if err != nil {
				return err
			}
		}

		p.scope = ""
		return err
	case *ast.AssertionStatement:
		return err
	case *ast.AssumptionStatement:
		return err
	case *ast.ForStatement:
		return err
	case *ast.StartStatement:
		return err
	case *ast.Instance:
		var key string
		importSpec := p.Specs[node.Value.Spec] //Where the struct definition lives

		if p.Specs[p.trail.CurrentSpec()] == nil {
			p.Specs[p.trail.CurrentSpec()] = NewSpecRecord()
		}

		spec := p.Specs[p.trail.CurrentSpec()] //Where the instance is being declared

		if node.ComplexScope != "" {
			key = strings.Join([]string{node.ComplexScope, node.Name}, "_")
		} else {
			key = strings.Join([]string{p.scope, node.Name}, "_")
		}

		var properties map[string]ast.Node
		switch p.structTypes[node.Value.Spec][node.Value.Value] {
		case "STOCK":
			properties = importSpec.FetchStock(node.Value.Value)
			spec.AddStock(key, properties)

			oldScope := p.scope
			p.scope = key
			for _, v := range properties {
				// Looking for more instances
				if inst, ok := v.(*ast.Instance); ok {
					inst.ComplexScope = key
					err = p.walk(inst)
				} else {
					err = p.walk(v)
				}

				if err != nil {
					return err
				}
			}
			p.scope = oldScope
		case "FLOW":
			properties = importSpec.FetchFlow(node.Value.Value)
			spec.AddFlow(key, properties)

			oldScope := p.scope
			p.scope = key
			for _, v := range properties {
				if inst, ok := v.(*ast.Instance); ok {
					inst.ComplexScope = key
					err = p.walk(inst)
				} else {
					err = p.walk(v)
				}

				if err != nil {
					return err
				}
			}
			p.scope = oldScope
		default:
			panic(fmt.Sprintf("invalid instance %s", node.Value.Value))
		}

		return err
	default:
		return err
	}
}
