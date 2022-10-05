package preprocess

import "fault/ast"

type SpecRecord struct {
	Stocks     map[string]map[string]ast.Node
	Flows      map[string]map[string]ast.Node
	Components map[string]map[string]ast.Node
	Constants  map[string]ast.Node
}

func NewSpecRecord() *SpecRecord {
	return &SpecRecord{
		Stocks:     make(map[string]map[string]ast.Node),
		Flows:      make(map[string]map[string]ast.Node),
		Components: make(map[string]map[string]ast.Node),
		Constants:  make(map[string]ast.Node),
	}
}

func (sr *SpecRecord) AddStock(k string, v map[string]ast.Node) {
	sr.Stocks[k] = v
}

func (sr *SpecRecord) AddFlow(k string, v map[string]ast.Node) {
	sr.Flows[k] = v
}

func (sr *SpecRecord) AddComponent(k string, v map[string]ast.Node) {
	sr.Components[k] = v
}

func (sr *SpecRecord) AddConstant(k string, v ast.Node) {
	sr.Constants[k] = v
}

func (sr *SpecRecord) FetchStock(k string) map[string]ast.Node {
	return sr.Stocks[k]
}

func (sr *SpecRecord) FetchFlow(k string) map[string]ast.Node {
	return sr.Flows[k]
}

func (sr *SpecRecord) FetchComponent(k string) map[string]ast.Node {
	return sr.Components[k]
}

func (sr *SpecRecord) FetchConstant(k string) ast.Node {
	return sr.Constants[k]
}
