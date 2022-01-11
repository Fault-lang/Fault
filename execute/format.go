package execute

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

/*
 Set of functions for formating the results from the model
 checker in user friendly ways
*/

func (mc *ModelChecker) Format(results map[string]Scenario) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Variable", "State (Weight)"})
	var row []string
	results = definePath(results, mc.branchTrail)
	for k, v := range results {
		row = append(row, k)
		r := generateRows(v)
		row = append(row, strings.Join(r, " "))
		table.Append(row)
	}
	table.Render()
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

func definePath(results map[string]Scenario, trails map[string]map[string][]string) map[string]Scenario {
	/*
	 Removes states from the state graph that are not actually
	 in the failure scenario (Z3 will return solutions for
	 variables set in all branches of a conditional for example)

	 Compare the final variable in each branch to the phi state
	 to identify the winning branch.
	*/
	for _, v := range trails {
		b := selectBranch(results, v)
		for k, br := range v {
			if k != b {
				results = removeBranch(results, br)
			}
		}
	}
	return results
}

func removeBranch(results map[string]Scenario, trail []string) map[string]Scenario {
	for _, id := range trail {
		p := strings.Split(id, "_")
		name := strings.Join(p[0:len(p)-1], "_")
		i, err := strconv.ParseInt(p[len(p)-1], 0, 64)
		if err != nil {
			panic(err)
		}
		switch s := results[name].(type) {
		case *FloatTrace:
			s.Remove(i)
			results[name] = s
		case *IntTrace:
			s.Remove(i)
			results[name] = s
		case *BoolTrace:
			s.Remove(i)
			results[name] = s
		}
	}
	return results
}

func endStatesBranch(trail []string) map[string]int64 {
	// Return map[variable_name] = [last SSA]
	ret := make(map[string]int64)
	for _, v := range trail {
		p := strings.Split(v, "_")
		name := strings.Join(p[0:len(p)-1], "_")
		curr, err := strconv.ParseInt(p[len(p)-1], 0, 64)
		if err != nil {
			panic(err)
		}
		if n, ok := ret[name]; ok {
			if n < curr {
				ret[name] = curr
			}
		} else {
			ret[name] = curr
		}
	}
	return ret
}

func selectBranch(results map[string]Scenario, v map[string][]string) string {
	options := make(map[string]map[string]int64)
	for k, b := range v {
		options[k] = endStatesBranch(b)
	}

	phis := generatePhis(options)

	for k, o := range options {
	option:
		for k2, e := range o {
			switch s := results[k2].(type) {
			case *FloatTrace:
				i1, _ := s.Index(phis[k2])
				i2, _ := s.Index(e)
				//fmt.Println(i1, i2)
				if i1 != i2 {
					//	fmt.Println("break")
					break option
				}
			case *IntTrace:
				i1, _ := s.Index(phis[k2])
				i2, _ := s.Index(e)
				if i1 != i2 {
					break option
				}
			case *BoolTrace:
				i1, _ := s.Index(phis[k2])
				i2, _ := s.Index(e)
				if i1 != i2 {
					break option
				}
			}
		}
		return k
	}
	panic("could not identify the branch selected by the solver")
}

func generatePhis(o map[string]map[string]int64) map[string]int64 {
	phis := make(map[string]int64)
	for _, v := range o {
		for k2, e := range v {
			if p, ok := phis[k2]; !ok {
				phis[k2] = e + 1
			} else if p < e {
				_, _, _ = phis, k2, e

			}
		}
	}
	return phis
}

func hWToString(n interface{}, i int64, weights map[int64]float64) string {
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
