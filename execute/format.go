package execute

import (
	"bytes"
	"fault/smt/variables"
	"fault/util"
	"fmt"
	"strings"
)

/*
 Set of functions for formating the results from the model
 checker in user friendly ways
*/

func (mc *ModelChecker) Mermaid() {
	if len(mc.Results) > 0 {
		var out bytes.Buffer
		out.WriteString("flowchart LR\n")
		for k, l := range mc.Results {
			out.WriteString(mc.writeObjects(k, l))
			out.WriteString("\n")
		}
		fmt.Println(out.String())
	}
}

func (mc *ModelChecker) writeObjects(k string, objects []*variables.VarChange) string {
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

func (mc *ModelChecker) writeObject(o *variables.VarChange) string {
	value, ok := mc.ResultValues[o.Parent]
	if ok {
		return fmt.Sprintf("\t% s--> |%s| %s", o.Parent, value, o.Id)
	} else {
		return fmt.Sprintf("\t%s --> %s", o.Parent, o.Id)
	}
}

func (mc *ModelChecker) Format(results map[string]Scenario) {
	var out bytes.Buffer
	for k, v := range results {
		out.WriteString(k + "\n")
		deadVars := mc.DeadVariables()
		filtered := deadBranches(k, v, deadVars)
		r := generateRows(filtered)
		out.WriteString(strings.Join(r, " ") + "\n\n")
	}
	fmt.Println(out.String())
}

func (mc *ModelChecker) EventLog(results map[string]Scenario) {
	var out bytes.Buffer
	for k, v := range results {
		mc.mapToLog(k, v)
	}

	deadVars := mc.DeadVariables()
	mc.Log.FilterOut(deadVars)

	out.WriteString(mc.Log.String())

	fmt.Println(out.String())
}

func (mc *ModelChecker) mapToLog(k string, vals Scenario) {
	switch v := vals.(type) {
	case *BoolTrace:
		for idx, s := range v.results {
			name := fmt.Sprintf("%s_%d", k, idx)
			j := mc.Log.Index(name)
			if j >= 0 {
				mc.Log.UpdateCurrent(j, fmt.Sprintf("%v", s))
			}
		}
	case *FloatTrace:
		for idx, s := range v.results {
			name := fmt.Sprintf("%s_%d", k, idx)
			j := mc.Log.Index(name)
			if j >= 0 {
				mc.Log.UpdateCurrent(j, fmt.Sprintf("%v", s))
			}
		}
	case *IntTrace:
		for idx, s := range v.results {
			name := fmt.Sprintf("%s_%d", k, idx)
			j := mc.Log.Index(name)
			if j >= 0 {
				mc.Log.UpdateCurrent(j, fmt.Sprintf("%v", s))
			}
		}
	}

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

func (mc *ModelChecker) DeadVariables() []string {
	var dead []string
	for choiceId, branchIds := range mc.Forks.Choices {
		winner, phi := mc.pickWinner(choiceId, branchIds)
		for _, b := range branchIds {
			if b != winner {
				dead = append(dead, mc.Forks.Branches[b]...)
				mc.Forks.MarkMany(mc.Forks.Branches[b])

				for _, p := range phi {
					if !mc.Forks.MarkedForDeath(p) {
						dead = append(dead, p) //Kill the phis too
						mc.Forks.Mark(p)
					}
				}
			}
		}
	}
	return dead
}

func (mc *ModelChecker) pickWinner(choiceId string, branchIds []string) (string, map[string]string) {
	var winner string
	var phiID = make(map[string]string)
	for _, branch := range branchIds { // Go through all the branches
		var candidate string
		var candidatePhi = make(map[string]string)
		var fail bool
		DeclaredVars := mc.Forks.Branches[branch] // Variables declared in this branch
		if mc.allPhis(choiceId, DeclaredVars) && winner != "" {
			return winner, phiID
		}

		// if mc.endStateEqualPhis(choiceId, DeclaredVars)
		for _, dvars := range DeclaredVars {
			if mc.Forks.Vars[dvars].Last[choiceId] { // Is this variable SSA the last one assigned in the branch?
				last := mc.ResultValues[dvars]
				phi := mc.ResultValues[mc.Forks.Vars[dvars].FullPhi(choiceId)]
				if last != phi { // Does it's returned value match the Phi?
					fail = true
					break // If not this can't be a winning branch
				}
				candidate = branch
				candidatePhi[dvars] = mc.Forks.Vars[dvars].FullPhi(choiceId)

				// If the only variables defined in the branch are phis
				// branch will default to true
			} else if mc.Forks.Vars[dvars].Phi[choiceId] == mc.Forks.Vars[dvars].SSA { //Is this the Phi?
				last := mc.ResultValues[mc.Forks.Vars[dvars].Previous[branch]] // What was the previous value?
				phi := mc.ResultValues[dvars]
				if last != phi {
					fail = true
					break
				}
				candidate = branch
				candidatePhi[dvars] = mc.Forks.Vars[dvars].FullPhi(choiceId)
			}
		}

		if !fail {
			winner = candidate
			phiID = candidatePhi
		}
	}
	if winner == "" { //This should never happen
		var message []string
		for _, branch := range branchIds {
			b := mc.Forks.Branches[branch]
			message = append(message, strings.Join(b, ","))
		}
		panic(fmt.Sprintf("event log corrupted, can't decide between branches %s", strings.Join(message, " or ")))
	}
	return winner, phiID
}

func (mc *ModelChecker) allPhis(choiceId string, vars []string) bool {
	var phis int
	for _, dvars := range vars {
		if mc.Forks.Vars[dvars].Phi[choiceId] == mc.Forks.Vars[dvars].SSA {
			phis++
		}
	}
	return phis == len(vars)
}

func deadBranches(id string, variable Scenario, deads []string) Scenario {
	// Iterates through branches, determines the branches not needed by the model
	// and removes them from the Scenario
	//
	// Question: what to do in the situation where two
	// branches have the same end value as the phi but different
	// intermediate values?

	for _, b := range deads {
		switch v := variable.(type) {
		case *FloatTrace:
			base, n := util.GetVarBase(b)
			if base != id {
				continue
			}
			v.Remove(int16(n))
			variable = v
		case *IntTrace:
			base, n := util.GetVarBase(b)
			if base != id {
				continue
			}
			v.Remove(int16(n))
			variable = v
		case *BoolTrace:
			base, n := util.GetVarBase(b)
			if base != id {
				continue
			}
			v.Remove(int16(n))
			variable = v
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
