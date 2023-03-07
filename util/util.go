package util

import (
	"fault/ast"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type StringSet struct {
	base map[string]bool
}

func NewStrSet() *StringSet {
	return &StringSet{
		base: make(map[string]bool),
	}
}

func (s *StringSet) Add(str string) {
	s.base[str] = true
}

func (s *StringSet) In(str string) bool {
	return s.base[str]
}

func (s *StringSet) Len() int {
	return len(s.base)
}

func DiffStrSets(s1 *StringSet, s2 *StringSet) *StringSet {
	s3 := NewStrSet()
	for k := range s1.base {
		if !s2.In(k) {
			s3.Add(k)
		}
	}

	for k := range s2.base {
		if !s1.In(k) {
			s3.Add(k)
		}
	}
	return s3
}

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
			if idx == 0 {
				break
			}
			path := strings.Split(filepath[0:idx], "/")
			if path[0] == "" { //Leading slashes
				path = path[1:]
			}
			if path[len(path)-1] == "" { //Trailing slashes
				path = path[0 : len(path)-1]
			}

			var pathstr string
			if len(path) > 1 {
				pathstr = strings.Join(path[0:len(path)-1], "/")
				filepath = strings.Join([]string{pathstr, filepath[idx+2:]}, "")
			} else {
				filepath = filepath[idx+2:]
			}
		}

		if len(filepath) < len(host) || host != filepath[0:len(host)] {
			filepath = strings.Join([]string{host, filepath}, "/")
		}

		if strings.Contains(filepath, "//") {
			path := strings.Split(filepath, "//")
			return strings.Join(path, "/")
		}

	}
	return filepath
}

func Preparse(pairs map[*ast.Identifier]ast.Expression) map[string]ast.Node {
	properties := make(map[string]ast.Node)
	for k, v := range pairs {
		id := strings.TrimSpace(k.String())
		properties[id] = v
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

func MergeNodeMaps(m1 map[string]ast.Node, m2 map[string]ast.Node) map[string]ast.Node {
	for k, v := range m2 {
		m1[k] = v
	}
	return m1
}

func MergeStringMaps(m1 map[string]string, m2 map[string]string) map[string]string {
	for k, v := range m2 {
		m1[k] = v
	}
	return m1
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

func InStringSlice(sl []string, sub string) bool {
	for _, s := range sl {
		if s == sub {
			return true
		}
	}
	return false
}

func Keys(m map[string]ast.Node) []string {
	var ret []string
	for k := range m {
		ret = append(ret, k)
	}
	return ret
}

func StableSortKeys(keys []string) []string {
	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

func ExtractBranches(b map[string]*ast.StructProperty) map[string]ast.Node {
	ret := make(map[string]ast.Node)
	for k, v := range b {
		ret[k] = v.Value
	}
	return ret
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

func Copy(callstack []string) []string {
	var ret []string
	for _, v := range callstack {
		ret = append(ret, v)
	}
	return ret
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

func PairCombinations(left []string, right []string) [][]string {
	var ret [][]string
	for _, l := range left {
		for _, r := range right {
			ret = append(ret, []string{l, r})
		}
	}
	return ret
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

func Intersection(s1 []string, s2 []string, init bool) []string {
	var s3 []string
	for _, s := range s1 {
		s3 = append(s3, s)
		for _, z := range s2 {
			if s == z {
				s3 = s3[0 : len(s3)-1]
			}
		}
	}
	if init {
		s4 := Intersection(s2, s1, false)
		s3 = append(s3, s4...)
	}
	return s3
}

func FromEnd(str string, offset int) string {
	return str[len(str)-offset:]
}

type ImportTrail []string

func (i ImportTrail) BaseSpec() string {
	if len(i) == 0 {
		panic(fmt.Sprintln("import trail is empty"))
	}
	return i[0]
}

func (i ImportTrail) CurrentSpec() string {
	if len(i) == 0 {
		panic(fmt.Sprintln("import trail is empty"))
	}
	return i[len(i)-1]
}

func (i ImportTrail) PushSpec(spec string) []string {
	i = append(i, spec)
	return i
}

func (i ImportTrail) PopSpec() (string, []string) {
	if len(i) == 0 {
		panic(fmt.Sprintln("import trail is empty"))
	}
	spec := i[len(i)-1]
	i = i[0 : len(i)-1]
	return spec, i
}
