package util

import (
	"fault/ast"
	"os"
	"strings"
)

func Filepath(filepath string) string {
	if host, ok := os.LookupEnv("FAULT_HOST"); ok {
		if strings.Contains(filepath, "~") {
			path := strings.Split(filepath, "~")
			if string(path[1][0]) == "/" {
				filepath = path[1][1:]
			} else {
				filepath = path[1]
			}
			return strings.Join([]string{host, filepath}, "/")
		}
		for strings.Contains(filepath, "..") {
			idx := strings.Index(filepath, "..")
			path := strings.Split(filepath[0:idx], "/")
			if path[len(path)-1] == "" { //Trailing slashes
				path = path[0 : len(path)-1]
			}
			var pathstr string
			if len(path) > 1 {
				pathstr = strings.Join(path[0:len(path)-1], "/")
			} else {
				pathstr = path[0]
			}
			filepath = strings.Join([]string{pathstr, filepath[idx+2:]}, "")
		}

		filepath = strings.Join([]string{host, filepath}, "/")
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
