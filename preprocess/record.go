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
	Globals    map[string]ast.Node //mostly for stand alone string rules
	// Because the order in which structs are declared matters
	Order [][]string // ("STOCK", this_var_name)
}

func NewSpecRecord() *SpecRecord {
	return &SpecRecord{
		Stocks:     make(map[string]map[string]ast.Node),
		Flows:      make(map[string]map[string]ast.Node),
		Components: make(map[string]map[string]ast.Node),
		Constants:  make(map[string]ast.Node),
		Globals:    make(map[string]ast.Node),
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

func (sr *SpecRecord) AddGlobal(name string, v ast.Node) {
	sr.Globals[name] = v
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

func (sr *SpecRecord) GetStructType(rawid []string) (string, []string) {
	if len(rawid) == 2 {
		_, err := sr.FetchConstant(rawid[1])
		if err == nil {
			return "CONSTANT", rawid
		}

		_, err = sr.FetchGlobal(rawid[1])
		if err == nil {
			return "GLOBAL", rawid
		}
	}

	id := strings.Join(rawid[1:], "_")
	for _, v := range sr.Order {
		if v[1] == id {
			return v[0], rawid
		}
	}

	rawid2 := rawid[0 : len(rawid)-1]
	if len(rawid2) > 1 {
		return sr.GetStructType(rawid2)
	}

	return "NIL", rawid
}

func (sr *SpecRecord) Fetch(name string, ty string) (map[string]ast.Node, error) {
	switch ty {
	case "STOCK":
		return sr.FetchStock(name)
	case "FLOW":
		return sr.FetchFlow(name)
	case "COMPONENT":
		return sr.FetchComponent(name)
	case "CONSTANT":
		ret, err := sr.FetchConstant(name)
		if err != nil {
			return nil, err
		}
		return map[string]ast.Node{name: ret}, nil
	case "GLOBAL":
		ret, err := sr.FetchGlobal(name)
		if err != nil {
			return nil, err
		}
		return map[string]ast.Node{name: ret}, nil
	default:
		return nil, fmt.Errorf("cannot fetch a variable %s of type %s", name, ty)
	}
}

func (sr *SpecRecord) FetchStock(name string) (map[string]ast.Node, error) {
	if sr.Stocks[name] != nil {
		return sr.Stocks[name], nil
	} else {
		return nil, fmt.Errorf("no stock found with name %s in spec %s", name, sr.SpecName)
	}
}

func (sr *SpecRecord) FetchFlow(name string) (map[string]ast.Node, error) {
	if sr.Flows[name] != nil {
		return sr.Flows[name], nil
	} else {
		return nil, fmt.Errorf("no flow found with name %s in spec %s", name, sr.SpecName)
	}
}

func (sr *SpecRecord) FetchInstanceStrMap(name string, parent, ty string) (map[string]string, error) {
	var inst map[string]ast.Node
	var err error
	switch ty {
	case "STOCK":
		inst, err = sr.FetchStock(name)
		if err != nil {
			return nil, err
		}
	case "FLOW":
		inst, err = sr.FetchFlow(name)
		if err != nil {
			return nil, err
		}
	}

	var children = make(map[string]string)
	for _, v := range inst {
		vname := v.(ast.Nameable).IdString()
		children[vname] = parent
		if str, ok := v.(*ast.StructInstance); ok {
			rawid := str.RawId()
			cname := strings.Join(rawid[1:], "_")
			ty, _ := sr.GetStructType(rawid)
			grandchildren, err := sr.FetchInstanceStrMap(cname, str.Parent[1], ty)
			if err != nil {
				return nil, err
			}

			children = util.MergeStringMaps(children, grandchildren)
		}
	}
	return children, nil
}

func (sr *SpecRecord) FetchComponent(name string) (map[string]ast.Node, error) {
	if sr.Components[name] != nil {
		return sr.Components[name], nil
	} else {
		return nil, fmt.Errorf("no component found with name %s in spec %s", name, sr.SpecName)
	}
}

func (sr *SpecRecord) FetchConstant(name string) (ast.Node, error) {
	if sr.Constants[name] != nil {
		return sr.Constants[name], nil
	} else {
		return nil, fmt.Errorf("no constant found with name %s in spec %s", name, sr.SpecName)
	}
}

func (sr *SpecRecord) FetchGlobal(name string) (ast.Node, error) {
	if sr.Globals[name] != nil {
		return sr.Globals[name], nil
	} else {
		return nil, fmt.Errorf("no global found with name %s in spec %s", name, sr.SpecName)
	}
}

func (sr *SpecRecord) FetchOrder() [][]string {
	return sr.Order
}
func (sr *SpecRecord) FetchAll() map[string]ast.Node {
	all := make(map[string]ast.Node)
	for _, v := range sr.Stocks {
		all = ast.MergeNodeMaps(all, v)
	}
	for _, v := range sr.Flows {
		all = ast.MergeNodeMaps(all, v)
	}
	for _, v := range sr.Components {
		all = ast.MergeNodeMaps(all, v)
	}
	for k, v := range sr.Constants {
		all[k] = v
	}

	for k, v := range sr.Globals {
		all[k] = v
	}
	return all
}

func (sr *SpecRecord) FetchVar(rawid []string, ty string) (ast.Node, error) {
	var str string
	var br string
	var ret ast.Node

	// if len(rawid) == 2 {
	// 	return sr.FetchConstant(rawid[1])
	// }

	if len(rawid) > 3 {
		str := strings.Join(rawid[1:len(rawid)-1], "_")
		return sr.FetchVar([]string{rawid[0], str, rawid[len(rawid)-1]}, ty)
	}

	str = rawid[1]
	br = rawid[len(rawid)-1]

	switch ty {
	case "STOCK":
		s, err := sr.FetchStock(str)
		if err != nil {
			return nil, err
		}

		ret = s[br]
		if ret != nil {
			return ret, nil
		}
		return nil, fmt.Errorf("no property named %s in stock %s", br, str)
	case "FLOW":
		f, err := sr.FetchFlow(str)
		if err != nil {
			return nil, err
		}

		ret = f[br]
		if ret != nil {
			return ret, nil
		}

		return nil, fmt.Errorf("no property named %s in flow %s", br, str)
	case "COMPONENT":
		c, err := sr.FetchComponent(str)
		if err != nil {
			return nil, err
		}
		ret = c[br]

		if ret != nil {
			return ret, nil
		}

		return nil, fmt.Errorf("no property named %s in component %s", br, str)
	case "CONSTANT":
		return sr.FetchConstant(rawid[1])
	case "GLOBAL":
		return sr.FetchGlobal(rawid[1])
	default:
		return nil, fmt.Errorf("cannot fetch a variable %s of type %s", rawid, ty)
	}
}

func (sr *SpecRecord) Update(rawid []string, val map[string]ast.Node) error {
	var err error
	ty, rawid2 := sr.GetStructType(rawid)
	name := strings.Join(rawid2[1:], "_")
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

func (sr *SpecRecord) UpdateStock(name string, val map[string]ast.Node) error {
	if sr.Stocks[name] != nil {
		sr.Stocks[name] = val
		return nil
	}
	return fmt.Errorf("no stock found with name %s in spec %s", name, sr.SpecName)
}

func (sr *SpecRecord) UpdateFlow(name string, val map[string]ast.Node) error {
	if sr.Flows[name] != nil {
		sr.Flows[name] = val
		return nil
	}
	return fmt.Errorf("no flow found with name %s in spec %s", name, sr.SpecName)
}

func (sr *SpecRecord) UpdateComponent(name string, val map[string]ast.Node) error {
	if sr.Components[name] != nil {
		sr.Components[name] = val
		return nil
	}
	return fmt.Errorf("no component found with name %s in spec %s", name, sr.SpecName)
}

func (sr *SpecRecord) UpdateGlobal(name string, val ast.Node) error {
	if sr.Globals[name] != nil {
		sr.Globals[name] = val
		return nil
	}
	return fmt.Errorf("no global found with name %s in spec %s", name, sr.SpecName)
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
	case "GLOBAL":
		name = strings.Join(rawid[1:], "_")
		sr.Globals[name] = val
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
