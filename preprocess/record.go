package preprocess

import (
	"fault/ast"
	"fmt"

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

func (sr *SpecRecord) AddConstant(k string, v ast.Node) {
	sr.Constants[k] = v
}

func (sr *SpecRecord) AddInstance(k string, v map[string]ast.Node, ty string) {
	// When creating an instance of a struct need to deep copy the data
	v2, err := deepcopy.Anything(v)

	if err != nil {
		panic(fmt.Sprintf("failed to clone struct into instance %s", k))
	}

	switch ty {
	case "STOCK":
		sr.AddStock(k, v2.(map[string]ast.Node))
	case "FLOW":
		sr.AddFlow(k, v2.(map[string]ast.Node))
	case "COMPONENT":
		sr.AddComponent(k, v2.(map[string]ast.Node))
	}
}

func (sr *SpecRecord) GetStructType(id []string) string {
	for _, v := range sr.Order {
		if v[0] == id[1] {
			return v[1]
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

func (sr *SpecRecord) Index(ty string, name string) {
	sr.pushOrder(ty, name)

}

func (sr *SpecRecord) pushOrder(ty string, name string) {
	sr.Order = append(sr.Order, []string{ty, name})
}
