package util

import (
	"fault/ast"
	"os"
	"strings"
)

func Filepath(filepath string) string {
	if host, ok := os.LookupEnv("FAULT_HOST"); ok {
		hostParts := strings.Split(host, "/")
		for filepath[0:2] == ".." {
			filepath = filepath[3:]
			if len(hostParts) > 0 {
				hostParts = hostParts[0 : len(hostParts)-1]
			}
		}
		filepath = strings.Join(append(hostParts, filepath), "/")
	}
	return filepath
}

func Preparse(pairs map[ast.Expression]ast.Expression) map[string]ast.Node {
	properties := make(map[string]ast.Node)
	for k, v := range pairs {
		id := strings.TrimSpace(k.String())
		switch tree := v.(type) {
		case *ast.FunctionLiteral:
			properties[id] = tree.Body
		default:
			properties[id] = tree
		}
	}
	return properties
}

func Cartesian(list1 []string, list2 []string) [][]string {
	var product [][]string
	for _, a := range list1 {
		for _, b := range list2 {
			product = append(product, []string{a, b})
		}
	}
	return product
}

func CartesianMulti(listOfLists [][]string) [][]string {
	start := Cartesian(listOfLists[0], listOfLists[1])
	for i := 2; i < len(listOfLists); i++ {
		start = product(start, listOfLists[i])
	}
	return start
}

func product(list1 [][]string, list2 []string) [][]string {
	var results [][]string
	for _, l := range list1 {
		for _, l1 := range list2 {
			p := append(l, l1)
			results = append(results, p)
		}
	}
	return results
}

func IsCompare(op string) bool {
	switch op {
	case ">":
		return true
	case "<":
		return true
	case "==":
		return true
	case "!=":
		return true
	case "<=":
		return true
	case ">=":
		return true
	case "&&":
		return true
	case "||":
		return true
	case "!":
		return true
	default:
		return false
	}
}
