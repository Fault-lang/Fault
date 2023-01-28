package smt

func (g *Generator) parseAssert(assert ast.Node) ([]*assrt, []*assrt, string) {
		switch e := assert.(type) {
		case *ast.AssertionStatement:
		case *ast.AssumptionStatement:
		default:
			pos := e.Position()
			panic(fmt.Sprintf("not a valid assert or assumption line: %d, col: %d", pos[0], pos[1]))
		}
	}
	