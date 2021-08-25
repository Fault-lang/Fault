package walker

import "fault/ast"

// Local function scope
// preparsed local struct
// Global
// preparsed global struct

func Preparse(pairs map[ast.Expression]ast.Expression) map[string]ast.Node {
	properties := make(map[string]ast.Node)
	for k, v := range pairs {
		id := k.String()
		switch tree := v.(type) {
		case *ast.FunctionLiteral:
			properties[id] = tree.Body
		default:
			properties[id] = tree
		}
	}
	return properties
}
