package util

import (
	"fault/ast"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func GenerateToken(token string, literal string, start antlr.Token, stop antlr.Token) ast.Token {
	return ast.Token{
		Type:    ast.TokenType(token),
		Literal: literal,
		Position: []int{start.GetLine(),
			start.GetColumn(),
			stop.GetLine(),
			stop.GetColumn(),
		},
	}
}

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

func MergeStrSlices(sl1 []string, sl2 []string) []string {
	var results []string
	skip := false
	results = append(results, sl1...)
	for _, v2 := range sl2 {
		for _, v1 := range sl1 {
			if v2 == v1 {
				skip = true
				break
			}
		}
		if !skip {
			results = append(results, v2)
		} else {
			skip = false
		}
	}
	return results
}

func CaptureState(id string) (string, bool, bool) {
	var a, c bool
	raw := strings.Split(id, "_")
	if len(raw) > 2 { //Not a constant
		c = false
		a = true
	} else {
		c = true
		a = false
	}

	_, err := strconv.Atoi(raw[len(raw)-1])
	if err != nil {
		return "", a, c
	} else {
		return raw[len(raw)-1], false, false
	}

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

func MaxInt16(nums []int16) int16 {
	var temp int16
	for _, i := range nums {
		if temp < i {
			temp = i
		}
	}
	return temp
}

func Combinations(l [][]string, n int) [][][]string {
	if len(l) <= n {
		return [][][]string{l}
	}

	var subset [][][]string
	for idx, itm := range l {
		if n == 1 {
			subset = append(subset, [][]string{itm})
		} else if len(l) > idx+1 {
			i := idx //0
			for {
				pos := i + n // 1
				if pos > len(l) {
					break
				}
				items := append([][]string{itm}, l[i+1:pos]...)
				subset = append(subset, items)
				i++
			}
		}
	}
	return subset
}

func NotInSet(o [][]string, c [][]string) [][]string {
	var s [][]string
	for _, r := range c {
		sw := true
	exit:
		for _, in := range o {
			if strings.Join(r, "") == strings.Join(in, "") {
				sw = false
				break exit
			}
		}
		if sw {
			s = append(s, r)
		}
	}
	return s
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

func DetectMode(filename string) string {
	switch filepath.Ext(filename) {
	case ".fspec":
		return "fspec"
	case ".fsystem":
		return "fsystem"
	default:
		return ""
	}
}
