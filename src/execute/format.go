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
					r = append(r, "[branch]")
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
					r = append(r, "[branch]")
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
					r = append(r, "[branch]")
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
		for _, id := range v[b] {
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
	var end int64
	var branch string
	t := endStatesBranch(v["true"])
	f := endStatesBranch(v["false"])
	//Only need to test one variable to ID the correct branch
	for k, _ := range t {
		if t[k] > f[k] {
			end = t[k]
			branch = "false" //remove false branch if end is equal to phi
		} else {
			end = f[k]
			branch = "true"
		}

		phi := end + 1
		switch s := results[k].(type) {
		case *FloatTrace:
			i1, _ := s.Index(phi)
			i2, _ := s.Index(end)
			if i1 == i2 {
				return branch
			} else {
				if branch == "true" {
					return "false"
				}
			}
			return "true"

		case *IntTrace:
			i1, _ := s.Index(phi)
			i2, _ := s.Index(end)
			if i1 == i2 {
				return branch
			} else {
				if branch == "true" {
					return "false"
				}
			}
			return "true"
		case *BoolTrace:
			i1, _ := s.Index(phi)
			i2, _ := s.Index(end)
			if i1 == i2 {
				return branch
			} else {
				if branch == "true" {
					return "false"
				}
			}
			return "true"
		}
	}
	return ""
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
