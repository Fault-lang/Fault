package util

import (
	"fmt"
	"os"
	ospath "path/filepath"
	"sort"
	"strconv"
	"strings"
)

var OP_NEGATE = map[string]string{
	"==":   "!=",
	">=":   "<",
	">":    "<=",
	"<=":   ">",
	"!=":   "==",
	"<":    ">=",
	"&&":   "||",
	"||":   "&&",
	"then": "then",
	//"=": "!=",
}

func PlainLangOp(op string) string {
	//Use plain language instead of logic operators
	switch op {
	case "&&":
		return "and"
	case "||":
		return "or"
	case ">":
		return "is greater than"
	case "<":
		return "is less than"
	case ">=":
		return "is greater than or equal to"
	case "<=":
		return "is less than or equal to"
	case "==":
		return "is equal to"
	case "!=":
		return "is not equal to"
	case "!":
		return "not"
	case "-":
		return "not"
	default:
		return op
	}
}

type StringSet struct {
	base map[string]bool
	vals []string
}

func NewStrSet() *StringSet {
	return &StringSet{
		base: make(map[string]bool),
	}
}

func (s *StringSet) Add(str string) {
	if !s.In(str) {
		s.base[str] = true
		s.vals = append(s.vals, str)
	}
}

func (s *StringSet) Merge(strs []string) {
	for _, str := range strs {
		s.Add(str)
	}
}

func (s *StringSet) In(str string) bool {
	return s.base[str]
}

func (s *StringSet) Len() int {
	return len(s.base)
}

func (s *StringSet) Values() []string {
	return s.vals
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

func Filepath(filepath string) string {
	if host, ok := os.LookupEnv("FAULT_HOST"); ok {
		if strings.Contains(filepath, "~") {
			return home(host, filepath)
		}
		for strings.Contains(filepath, "..") {
			idx := strings.Index(filepath, "..")
			if idx == 0 {
				host = uplevel(host, true)
				filepath = filepath[3:]
				continue
			}

			left := filepath[:idx]
			right := filepath[idx+2:]
			path := uplevel(left, false)

			if path == "" {
				filepath = right
			} else {
				filepath = ospath.Join(path, right)
			}

		}

		if len(filepath) < len(host) || host != filepath[0:len(host)] {
			filepath = ospath.Join(host, filepath)
		}

		dup := fmt.Sprintf("%s%s", string(ospath.Separator), string(ospath.Separator))
		if strings.Contains(filepath, dup) {
			path := strings.Split(filepath, dup)
			filepath = ospath.Join(path...)
		}

	}
	return filepath
}

func home(host string, filepath string) string {
	path := strings.Split(filepath, "~")
	if string(path[1][0]) == string(ospath.Separator) {
		filepath = path[1][1:]
	} else {
		filepath = path[1]
	}
	return ospath.Join(host, filepath)
}

func FormatBlock(blockName string) string {
	if len(blockName) > 0 && blockName[0] == '%' {
		blockName = blockName[1:]
	}
	parts := strings.Split(blockName, "-")
	return strings.Join(parts, "")
}

func uplevel(path string, host bool) string {
	parts := strings.Split(path, string(ospath.Separator))
	parts = trimSlashes(parts, host)

	if len(parts) > 0 {
		return ospath.Join(parts[0 : len(parts)-1]...)
	}
	return ""
}

func trimSlashes(parts []string, host bool) []string {
	if len(parts) == 0 {
		return parts
	}

	if parts[0] == "" && !host { //Leading slashes
		parts = parts[1:]
		return trimSlashes(parts, host)
	}

	if parts[len(parts)-1] == "" { //Trailing slashes
		parts = parts[0 : len(parts)-1]
		return trimSlashes(parts, host)
	}

	return parts
}

func FormatIdent(id string) string {
	//Removes LLVM IR specific leading characters
	if string(id[0]) == "@" {
		return id[1:]
	} else if string(id[0]) == "%" {
		return id[1:]
	} else if string(id[0:2]) == "c\"" {
		//Trim the c" "
		id = id[2 : len(id)-1]
	}
	return id
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

func CompareStringMaps(m1 map[string]string, m2 map[string]string) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if val, ok := m2[k]; !ok || val != v {
			return false
		}
	}
	return true
}

func MergeStringMaps(m1 map[string]string, m2 map[string]string) map[string]string {
	for k, v := range m2 {
		m1[k] = v
	}
	return m1
}

func MergeStringSliceMaps(m1 map[string][][]string, m2 map[string][][]string) map[string][][]string {
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

func MergeStringSets(m1 map[string]*StringSet, m2 map[string]*StringSet) map[string]*StringSet {
	for k, v := range m2 {
		if _, ok := m1[k]; ok {
			for _, val := range v.Values() {
				m1[k].Add(val)
			}
		} else {
			m1[k] = v
		}
	}
	return m1
}

func MergeIntSliceMaps(m1 map[string][]int16, m2 map[string][]int16) map[string][]int16 {
	// For Phis in unpacker
	for k, v := range m2 {
		if _, ok := m1[k]; ok {
			// If key exists, append the value
			m1[k] = append(m1[k], v...)
		} else {
			m1[k] = v
		}
	}
	return m1
}

func SliceOfIndex(l int) []int {
	if l < 0 {
		panic("length cannot be negative")
	}
	inverse := make([]int, l)
	for i := 0; i < l; i++ {
		inverse[i] = i
	}
	return inverse
}

func RemoveFromStringSlice(sl []string, sub string) []string {
	var new []string
	for _, s := range sl {
		if s != sub {
			new = append(new, s)
		}
	}
	return new
}

func InStringSlice(sl []string, sub string) bool {
	for _, s := range sl {
		if s == sub {
			return true
		}
	}
	return false
}

func StableSortKeys(keys []string) []string {
	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
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
	ret = append(ret, callstack...)
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

func DetectMode(filename string) string {
	switch ospath.Ext(filename) {
	case ".fspec":
		return "fspec"
	case ".fsystem":
		return "fsystem"
	default:
		return ""
	}
}

func GetVarBase(id string) (string, int) {
	v := strings.Split(id, "_")
	num, err := strconv.Atoi(v[len(v)-1])
	if err != nil {
		panic(fmt.Sprintf("improperly formatted variable SSA name %s", id))
	}
	return strings.Join(v[0:len(v)-1], "_"), num
}

func Difference(s1 map[string]bool, s2 map[string]bool) []string {
	vars := []string{}
	if len(s2) == 0 {
		for k := range s1 {
			vars = append(vars, k)
		}
		return vars
	}

	for k := range s2 {
		if _, ok := s1[k]; !ok {
			vars = append(vars, k)
		}
	}
	return vars
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
	if len(str) < offset {
		return str
	}
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
