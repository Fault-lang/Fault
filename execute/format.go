package execute

import (
	"bytes"
	"fmt"
	"strings"
)

/*
 Set of functions for formating the results from the model
 checker in user friendly ways
*/

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

// func definePath(results map[string]Scenario, trails []smt.fork) map[string]Scenario {
// 	/*
// 	 Removes states from the state graph that are not actually
// 	 in the failure scenario (Z3 will return solutions for
// 	 variables set in all branches of a conditional for example)

// 	 Compare the final variable in each branch to the phi state
// 	 to identify the winning branch.
// 	*/
// 	for _, v := range trails {
// 		b := selectBranch(results, v)
// 		for k, br := range v {
// 			if k != b {
// 				results = removeBranch(results, br)
// 			}
// 		}
// 	}
// 	return results
// }

// func removeBranch(results map[string]Scenario, trail []string) map[string]Scenario {
// 	for _, id := range trail {
// 		p := strings.Split(id, "_")
// 		name := strings.Join(p[0:len(p)-1], "_")
// 		i, err := strconv.ParseInt(p[len(p)-1], 0, 64)
// 		if err != nil {
// 			panic(err)
// 		}
// 		switch s := results[name].(type) {
// 		case *FloatTrace:
// 			s.Remove(i)
// 			results[name] = s
// 		case *IntTrace:
// 			s.Remove(i)
// 			results[name] = s
// 		case *BoolTrace:
// 			s.Remove(i)
// 			results[name] = s
// 		}
// 	}
// 	return results
// }

// func endStatesBranch(trail []string) map[string]int64 {
// 	// Return map[variable_name] = [last SSA]
// 	ret := make(map[string]int64)
// 	for _, v := range trail {
// 		p := strings.Split(v, "_")
// 		name := strings.Join(p[0:len(p)-1], "_")
// 		curr, err := strconv.ParseInt(p[len(p)-1], 0, 64)
// 		if err != nil {
// 			panic(err)
// 		}
// 		if n, ok := ret[name]; ok {
// 			if n < curr {
// 				ret[name] = curr
// 			}
// 		} else {
// 			ret[name] = curr
// 		}
// 	}
// 	return ret
// }

// func selectBranch(results map[string]Scenario, v map[string][]string) string {
// 	options := make(map[string]map[string]int64)
// 	for k, b := range v {
// 		options[k] = endStatesBranch(b)
// 	}

// 	phis := generatePhis(options)

// 	for k, o := range options {
// 	option:
// 		for k2, e := range o {
// 			switch s := results[k2].(type) {
// 			case *FloatTrace:
// 				i1, _ := s.Index(phis[k2])
// 				i2, _ := s.Index(e)
// 				//fmt.Println(i1, i2)
// 				if i1 != i2 {
// 					//	fmt.Println("break")
// 					break option
// 				}
// 			case *IntTrace:
// 				i1, _ := s.Index(phis[k2])
// 				i2, _ := s.Index(e)
// 				if i1 != i2 {
// 					break option
// 				}
// 			case *BoolTrace:
// 				i1, _ := s.Index(phis[k2])
// 				i2, _ := s.Index(e)
// 				if i1 != i2 {
// 					break option
// 				}
// 			}
// 		}
// 		return k
// 	}
// 	panic("could not identify the branch selected by the solver")
// }

// func generatePhis(o map[string]map[string]int64) map[string]int64 {
// 	phis := make(map[string]int64)
// 	for _, v := range o {
// 		for k2, e := range v {
// 			if p, ok := phis[k2]; !ok {
// 				phis[k2] = e + 1
// 			} else if p < e {
// 				phis[k2] = e + 1

// 			}
// 		}
// 	}
// 	return phis
// }

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
