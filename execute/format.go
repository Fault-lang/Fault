package execute

import (
	"bytes"
	"fault/smt"
	"fmt"
	"strings"
)

/*
 Set of functions for formating the results from the model
 checker in user friendly ways
*/

func (mc *ModelChecker) Mermaid() {
	var out bytes.Buffer
	out.WriteString("flowchart LR\n")
	for k, l := range mc.Results {
		out.WriteString(mc.writeObjects(k, l))
		out.WriteString("\n")
	}
	fmt.Println(out.String())
}

func (mc *ModelChecker) writeObjects(k string, objects []*smt.VarChange) string {
	var objs []string
	for _, o := range objects {
		if o.Parent != "" {
			objs = append(objs, mc.writeObject(o))
		}
	}
	last := objects[len(objects)-1]
	value := mc.ResultValues[last.Id]
	cap := fmt.Sprintf("\t% s--> %s(%s)", last.Id, k, value)
	objs = append(objs, cap)
	return strings.Join(objs, "\n")
}

func (mc *ModelChecker) writeObject(o *smt.VarChange) string {
	value, ok := mc.ResultValues[o.Parent]
	if ok {
		return fmt.Sprintf("\t% s--> |%s| %s", o.Parent, value, o.Id)
	} else {
		return fmt.Sprintf("\t%s --> %s", o.Parent, o.Id)
	}
}

func (mc *ModelChecker) Format(results map[string]Scenario) {
	var out bytes.Buffer
	//results = definePath(results, mc.forks)
	for k, v := range results {
		out.WriteString(k + "\n")
		filtered := deadBranches(k, v, mc.forks)
		r := generateRows(filtered)
		out.WriteString(strings.Join(r, " ") + "\n\n")
	}
	fmt.Println(out.String())
}

func generateRows(v Scenario) []string {
	switch s := v.(type) {
	case *FloatTrace:
		var r []string
		weights := s.GetWeights()
		for i, n := range s.Get() {
			if int(i) == len(r) {
				r = append(r, hWToString(n, i, weights))
			} else if int(i) > len(r) {
				for j := len(r) - 1; j < int(i)-1; j++ {
					r = append(r, "")
				}
				r = append(r, hWToString(n, i, weights))

			} else if int(i) < len(r) {
				r2 := r[0:i]
				r2 = append(r2, hWToString(n, i, weights))
				r = append(r2, r[i+1:]...)
			}
		}
		return r
	case *IntTrace:
		var r []string
		weights := s.GetWeights()
		for i, n := range s.Get() {
			if int(i) == len(r) {
				r = append(r, hWToString(n, i, weights))
			} else if int(i) > len(r) {
				for j := len(r) - 1; j < int(i); j++ {
					r = append(r, "")
				}
				r = append(r, hWToString(n, i, weights))

			} else if int(i) < len(r) {
				r[i] = hWToString(n, i, weights)
			}
		}
		return r
	case *BoolTrace:
		var r []string
		weights := s.GetWeights()
		for i, n := range s.Get() {
			if int(i) == len(r) {
				r = append(r, hWToString(n, i, weights))
			} else if int(i) > len(r) {
				for j := len(r) - 1; j < int(i); j++ {
					r = append(r, "")
				}
				r = append(r, hWToString(n, i, weights))

			} else if int(i) < len(r) {
				r[i] = hWToString(n, i, weights)
			}
		}
		return r
	}
	return nil
}

func deadBranches(id string, variable Scenario, branches map[string][]*Branch) Scenario {
	// Iterates through branches, determines the branches not needed by the model
	// and removes them from the Scenario
	//
	// Question: what to do in the situation where two
	// branches have the same end value as the phi but different
	// intermediate values?
	var phis []int16
	for _, b := range branches[id] {
		e := b.End()
		phis = append(phis, b.phi)
		switch v := variable.(type) {
		case *FloatTrace:
			endValue, ok := v.Index(e)
			if !ok {
				panic(fmt.Sprintf("end value for variable %s not found", id))
			}

			phiValue, ok := v.Index(b.phi)
			if !ok {
				panic(fmt.Sprintf("phi value for variable %s not found", id))
			}

			if endValue != phiValue {
				//Remove this branch from results
				for k := range v.results {
					if b.InTrail(k) {
						v.Remove(k)
					}
				}
			}
			variable = v
		case *IntTrace:
			endValue, ok := v.Index(e)
			if !ok {
				panic(fmt.Sprintf("end value for variable %s not found", id))
			}

			phiValue, ok := v.Index(b.phi)
			if !ok {
				panic(fmt.Sprintf("phi value for variable %s not found", id))
			}

			if endValue != phiValue {
				//Remove this branch from results
				for k := range v.results {
					if b.InTrail(k) {
						v.Remove(k)
					}
				}
			}
			variable = v
		case *BoolTrace:
			endValue, ok := v.Index(e)
			if !ok {
				panic(fmt.Sprintf("end value for variable %s not found", id))
			}

			phiValue, ok := v.Index(b.phi)
			if !ok {
				panic(fmt.Sprintf("phi value for variable %s not found", id))
			}

			if endValue != phiValue {
				//Remove this branch from results
				for k := range v.results {
					if b.InTrail(k) {
						v.Remove(k)
					}
				}
			}
			variable = v
		}
	}

	// Remove phis
	for _, i := range phis {
		switch v := variable.(type) {
		case *FloatTrace:
			v.Remove(i)
		case *IntTrace:
			v.Remove(i)
		case *BoolTrace:
			v.Remove(i)
		}
	}
	return variable
}

func hWToString(n interface{}, i int16, weights map[int16]float64) string {
	switch h := n.(type) {
	case float64:
		if val, ok := weights[i]; ok {
			return fmt.Sprintf("-> %f (%f)", h, val)
		} else {
			return fmt.Sprintf("-> %f", h)
		}
	case int64:
		if val, ok := weights[i]; ok {
			return fmt.Sprintf("-> %d (%f)", h, val)
		} else {
			return fmt.Sprintf("-> %d", h)
		}
	case bool:
		if val, ok := weights[i]; ok {
			return fmt.Sprintf("-> %v (%f)", h, val)
		} else {
			return fmt.Sprintf("-> %v", h)
		}
	default:
		panic(fmt.Sprintf("type %T not allowed", n))
	}
}
