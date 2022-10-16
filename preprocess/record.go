package preprocess

import (
	"fault/ast"
	"fault/util"
	"fmt"
	"strings"

	deepcopy "github.com/barkimedes/go-deepcopy"
)

type SpecRecord struct {
	SpecName   string
	Stocks     map[string]map[string]ast.Node
	Flows      map[string]map[string]ast.Node
	Components map[string]map[string]ast.Node
	Constants  map[string]ast.Node
	// Because the order in which structs are declared matters
	Order [][]string // ("STOCK", this_var_name)
}

func NewSpecRecord() *SpecRecord {
	return &SpecRecord{
		Stocks:     make(map[string]map[string]ast.Node),
		Flows:      make(map[string]map[string]ast.Node),
		Components: make(map[string]map[string]ast.Node),
		Constants:  make(map[string]ast.Node),
	}
}

func (sr *SpecRecord) Id() string {
	return sr.SpecName
}

func (sr *SpecRecord) AddStock(name string, v map[string]ast.Node) {
	sr.Stocks[name] = v
}

func (sr *SpecRecord) AddFlow(name string, v map[string]ast.Node) {
	sr.Flows[name] = v
}

func (sr *SpecRecord) AddComponent(name string, v map[string]ast.Node) {
	sr.Components[name] = v
}

func (sr *SpecRecord) AddConstant(name string, v ast.Node) {
	sr.Constants[name] = v
}

func (sr *SpecRecord) AddInstance(name string, v map[string]ast.Node, ty string) {
	// When creating an instance of a struct need to deep copy the data
	v2, err := deepcopy.Anything(v)

	if err != nil {
		panic(fmt.Sprintf("failed to clone struct into instance %s", name))
	}

	switch ty {
	case "STOCK":
		sr.AddStock(name, v2.(map[string]ast.Node))
	case "FLOW":
		sr.AddFlow(name, v2.(map[string]ast.Node))
	case "COMPONENT":
		sr.AddComponent(name, v2.(map[string]ast.Node))
	}
}

func (sr *SpecRecord) GetStructType(rawid []string) string {
	if len(rawid) == 2 && sr.FetchConstant(rawid[1]) != nil {
		return "CONSTANT"
	}

	id := strings.Join(rawid[1:], "_")

	for _, v := range sr.Order {
		if v[1] == id {
			return v[0]
		}
	}
	return "NIL"
}

func (sr *SpecRecord) Fetch(name string, ty string) map[string]ast.Node {
	switch ty {
	case "STOCK":
		return sr.FetchStock(name)
	case "FLOW":
		return sr.FetchFlow(name)
	case "COMPONENT":
		return sr.FetchComponent(name)
	case "CONSTANT":
		return map[string]ast.Node{name: sr.FetchConstant(name)}
	default:
		panic(fmt.Sprintf("Cannot fetch a variable %s of type %s", name, ty))
	}
}

func (sr *SpecRecord) FetchStock(name string) map[string]ast.Node {
	return sr.Stocks[name]
}

func (sr *SpecRecord) FetchFlow(name string) map[string]ast.Node {
	return sr.Flows[name]
}

func (sr *SpecRecord) FetchComponent(name string) map[string]ast.Node {
	return sr.Components[name]
}

func (sr *SpecRecord) FetchConstant(k string) ast.Node {
	return sr.Constants[k]
}

func (sr *SpecRecord) FetchOrder() [][]string {
	return sr.Order
}
func (sr *SpecRecord) FetchAll() map[string]ast.Node {
	var all map[string]ast.Node
	for _, v := range sr.Stocks {
		all = util.MergeNodeMaps(all, v)
	}
	for _, v := range sr.Flows {
		all = util.MergeNodeMaps(all, v)
	}
	for _, v := range sr.Components {
		all = util.MergeNodeMaps(all, v)
	}
	for k, v := range sr.Constants {
		all[k] = v
	}
	return all
}

func (sr *SpecRecord) FetchVar(rawid []string, ty string) ast.Node {
	var str string
	var br string

	if len(rawid) >= 2 { //Otherwise this is a constant
		str = strings.Join(rawid[1:len(rawid)-1], "_")
		br = rawid[len(rawid)-1]
	}

	switch ty {
	case "STOCK":
		s := sr.FetchStock(str)
		return s[br]
	case "FLOW":
		f := sr.FetchFlow(str)
		return f[br]
	case "COMPONENT":
		c := sr.FetchComponent(str)
		return c[br]
	case "CONSTANT":
		return sr.FetchConstant(rawid[1])
	default:
		panic(fmt.Sprintf("Cannot fetch a variable %s of type %s", rawid, ty))
	}
}

func (sr *SpecRecord) Update(rawid []string, val map[string]ast.Node) error {
	var err error
	name := strings.Join(rawid[1:], "_")
	ty := sr.GetStructType(rawid)
	switch ty {
	case "STOCK":
		sr.UpdateStock(name, val)
	case "FLOW":
		sr.UpdateFlow(name, val)
	case "COMPONENT":
		sr.UpdateComponent(name, val)
	case "NIL":
		return fmt.Errorf("cannot find the struct value %s", rawid)
	}
	return err
}

func (sr *SpecRecord) UpdateStock(name string, val map[string]ast.Node) {
	sr.Stocks[name] = val
}

func (sr *SpecRecord) UpdateFlow(name string, val map[string]ast.Node) {
	sr.Flows[name] = val
}

func (sr *SpecRecord) UpdateComponent(name string, val map[string]ast.Node) {
	sr.Components[name] = val
}

func (sr *SpecRecord) UpdateConstant(name string, val ast.Node) {
	sr.Constants[name] = val
}

func (sr *SpecRecord) UpdateVar(rawid []string, ty string, val ast.Node) error {
	var err error
	name := strings.Join(rawid[1:len(rawid)-1], "_")
	pr := rawid[len(rawid)-1]
	switch ty {
	case "STOCK":
		sr.Stocks[name][pr] = val
	case "FLOW":
		sr.Flows[name][pr] = val
	case "COMPONENT":
		sr.Components[name][pr] = val
	case "CONSTANT":
		name = strings.Join(rawid[1:], "_")
		sr.Constants[name] = val
	case "NIL":
		return fmt.Errorf("cannot find the struct value %s", rawid)
	}
	return err
}

func (sr *SpecRecord) Index(ty string, name string) {
	sr.pushOrder(ty, name)

}

func (sr *SpecRecord) pushOrder(ty string, name string) {
	sr.Order = append(sr.Order, []string{ty, name})
}
