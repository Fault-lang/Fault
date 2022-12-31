package visualize

import (
	"fault/ast"
	"fmt"
	"strings"
)

type Visual struct {
	tree        ast.Node
	stateSet    map[string]bool
	VisualState map[string][]string
	systemState []string
}

func NewVisual(tree ast.Node) *Visual {
	return &Visual{
		tree:        tree,
		stateSet:    make(map[string]bool),
		VisualState: make(map[string][]string),
		systemState: []string{"flowchart TD"},
	}
}

func (v *Visual) Build() {
	err := v.walk(v.tree)
	if err != nil {
		panic(err)
	}
}

func (vis *Visual) walk(n ast.Node) error {
	var err error

	switch node := n.(type) {
	case *ast.Spec:
		for _, v := range node.Statements {
			err = vis.walk(v)
			if err != nil {
				return err
			}
		}
		return err
	case *ast.DefStatement:
		if node.TokenLiteral() != "COMPONENT" &&
			node.TokenLiteral() != "GLOBAL" {
			return err
		}

		err = vis.walk(node.Value)
		if err != nil {
			return err
		}

		return err
	case *ast.ComponentLiteral:
		for _, v := range node.Pairs {
			err = vis.walk(v)
			if err != nil {
				return err
			}
		}
		return err
	case *ast.ExpressionStatement:
		err = vis.walk(node.Expression)
		if err != nil {
			return err
		}
		return err
	case *ast.ForStatement:
		for _, v := range node.Body.Statements {
			err = vis.walk(v)
			if err != nil {
				return err
			}
		}
		return err
	case *ast.StartStatement:
		return err
	case *ast.FunctionLiteral:
		err = vis.walk(node.Body)
		if err != nil {
			return err
		}
		return err
	case *ast.BlockStatement:
		for _, v := range node.Statements {
			err = vis.walk(v)
			if err != nil {
				return err
			}
		}
		return err
	case *ast.InfixExpression:
		err = vis.walk(node.Left)
		if err != nil {
			return err
		}

		err = vis.walk(node.Right)
		if err != nil {
			return err
		}

		return err

	case *ast.IndexExpression:
		err = vis.walk(node.Left)
		if err != nil {
			return err
		}
		return err

	case *ast.PrefixExpression:
		err := vis.walk(node.Right)
		if err != nil {
			return err
		}
		return err

	case *ast.IfExpression:
		err = vis.walk(node.Consequence)
		if err != nil {
			return err
		}

		if node.Elif != nil {
			err = vis.walk(node.Elif)
			if err != nil {
				return err
			}
		}
		if node.Alternative != nil {
			err = vis.walk(node.Alternative)
			if err != nil {
				return err
			}
		}
		return err

	case *ast.StructInstance:
		ty := node.Type()
		switch ty {
		case "STOCK":
			for _, v := range node.Properties {
				switch inst := v.Value.(type) {
				case *ast.StructInstance:
					to, err := vis.getShape(inst)
					if err != nil {
						return err
					}
					nid := node.IdString()
					stock := fmt.Sprintf("\t%s[%s]-->%s", nid, nid, to)
					vis.systemState = append(vis.systemState, stock)
				}
			}
			return err
		case "FLOW":
			for _, v := range node.Properties {
				switch inst := v.Value.(type) {
				case *ast.StructInstance:
					to, err := vis.getShape(inst)
					if err != nil {
						return err
					}
					nid := node.IdString()
					flow := fmt.Sprintf("\t%s{{%s}}-->%s", nid, nid, to)
					vis.systemState = append(vis.systemState, flow)
				}
			}
			return err
		default:
			return fmt.Errorf("can't find a struct instance named %s", node.Parent)
		}

	case *ast.BuiltIn:
		if node.Function == "advance" {
			temp := node.RawId()
			froms := temp[1:3]
			tos := node.Parameters["toState"].(ast.Nameable).Id()
			vis.addLine(froms[0], fmt.Sprintf("\t%s --> %s", strings.Join(froms, "_"), tos[1]))
		}

		return err
	default:
		return err
	}
}

func (v *Visual) getShape(inst *ast.StructInstance) (string, error) {
	switch inst.Type() {
	case "STOCK":
		return fmt.Sprintf("%s[%s]", inst.IdString(), inst.IdString()), nil
	case "FLOW":
		return fmt.Sprintf("%s{{%s}}", inst.IdString(), inst.IdString()), nil
	default:
		return "", fmt.Errorf("invalid type in getShape")
	}
}

func (v *Visual) addLine(k string, s string) {
	if s == "}" {
		v.VisualState[k] = append(v.VisualState[k], s)
		return
	}

	if _, ok := v.stateSet[s]; !ok {
		v.VisualState[k] = append(v.VisualState[k], s)
		v.stateSet[s] = true
	}
}

func (v *Visual) Render() string {
	var s string
	if len(v.VisualState) > 0 {
		s = "stateDiagram"
		for k, v := range v.VisualState {
			block := strings.Join(v, "\n")
			s = fmt.Sprintf("%s\nstate %s {\n%s\n}", s, k, block)
		}
	}

	if len(v.systemState) > 1 && s != "" {
		sys := strings.Join(v.systemState, "\n")
		return strings.Join([]string{s, sys}, "\n\n")
	} else if len(v.systemState) > 1 {
		return strings.Join(v.systemState, "\n")
	} else {
		return s
	}
}
