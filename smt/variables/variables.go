package variables

import (
	"fault/llvm"
	"fault/smt/rules"
	"fmt"
	"strconv"
	"strings"

	irtypes "github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type VarChange struct {
	Id     string // SSA name of var
	Parent string // SSA name of proceeding var
}

type VarData struct {
	SSA   map[string]int16
	Ref   map[string]rules.Rule
	Loads map[string]value.Value
	Phis  map[string][][]int16
	Types map[string]string
}

func NewVariables() *VarData {
	return &VarData{
		SSA:   make(map[string]int16),
		Ref:   make(map[string]rules.Rule),
		Loads: make(map[string]value.Value),
		Phis:  make(map[string][][]int16),
		Types: make(map[string]string),
	}
}

func (vd *VarData) IsTemp(id string) bool {
	if string(id[0]) == "%" && vd.IsNumeric(string(id[1])) {
		return true
	}
	return false
}

func (vd *VarData) IsGlobal(id string) bool {
	return string(id[0]) == "@"
}

func (vd *VarData) IsNumeric(char string) bool {
	if _, err := strconv.ParseFloat(char, 64); err == nil {
		return true
	}
	if _, err := strconv.Atoi(char); err == nil {
		return true
	}
	return false
}

func (vd *VarData) IsBoolean(id string) bool {
	if id == "true" || id == "false" {
		return true
	}
	return false
}

func (vd *VarData) ConvertIdent(f string, val string) string {
	if vd.IsTemp(val) {
		refname := fmt.Sprintf("%s-%s", f, val)
		if v, ok := vd.Loads[refname]; ok {
			id := vd.FormatIdent(v.Ident())
			if v, ok := vd.SSA[id]; ok {
				return fmt.Sprint(id, "_", v)
			} else {
				panic(fmt.Sprintf("variable %s not initialized", id))
			}

		} else {
			panic(fmt.Sprintf("variable %s not initialized", val))
		}
	} else {
		id := val
		if string(id[0]) == "%" || vd.IsGlobal(id) {
			id = vd.FormatIdent(id)
			return fmt.Sprint(id, "_", vd.SSA[id])
		}
		return id //Is a value, not an identifier
	}
}

func (vd *VarData) FormatIdent(id string) string {
	//Removes LLVM IR specific leading characters
	if string(id[0]) == "@" {
		return id[1:]
	} else if string(id[0]) == "%" {
		return id[1:]
	}
	return id
}

func (vd *VarData) GetVarBase(id string) (string, int) {
	v := strings.Split(id, "_")
	num, err := strconv.Atoi(v[len(v)-1])
	if err != nil {
		panic(fmt.Sprintf("improperly formatted variable SSA name %s", id))
	}
	return strings.Join(v[0:len(v)-1], "_"), num
}

func (vd *VarData) LookupType(id string, value value.Value) string {
	if cache, ok := vd.Types[id]; ok { //If we've seen this one before
		return cache
	}

	val := vd.Loads[id]
	if val == nil { // A backup method
		switch value.Type().(type) {
		case *irtypes.FloatType:
			vd.Types[id] = "Real"
			return "Real"
		case *irtypes.IntType: // LLVM doesn't have a bool type
			vd.Types[id] = "Bool" // Just int type with a bitsize 1
			return "Bool"         // since all Fault numbers are floats,
		// ints are probably bools
		case *irtypes.ArrayType:
			vd.Types[id] = "Bool"
			return "Bool"
		}
	}

	if val.Type().Equal(llvm.DoubleP) {
		vd.Types[id] = "Real"
		return "Real"
	}
	if val.Type().Equal(llvm.I1P) {
		vd.Types[id] = "Bool"
		return "Bool"
	}

	panic(fmt.Sprintf("smt generation error, value for %s not found", id))
}

func (vd *VarData) LookupCondPart(f string, val string) rules.Rule {
	if vd.IsTemp(val) {
		refname := fmt.Sprintf("%s-%s", f, val)
		if v, ok := vd.Ref[refname]; ok {
			return v
		}
	}
	return nil
}

func (vd *VarData) FormatValue(val value.Value) string {
	v := strings.Split(val.String(), " ")
	return v[1]
}

func (vd *VarData) GetSSA(id string) string {
	if _, ok := vd.SSA[id]; ok {
		return fmt.Sprint(id, "_", vd.SSA[id])
	} else {
		vd.SSA[id] = 0
		return fmt.Sprint(id, "_0")
	}
}

func (vd *VarData) AdvanceSSA(id string) string {
	if i, ok := vd.SSA[id]; ok {
		vd.SSA[id] = i + 1
		return fmt.Sprint(id, "_", vd.SSA[id])
	} else {
		vd.SSA[id] = 0
		return fmt.Sprint(id, "_0")
	}
}

// When we have conditionals back to back (but not if elseif else)
// we need to make sure to track the phi
func (vd *VarData) InitPhis() {
	for k := range vd.Phis {
		vd.NewPhi(k, -1)
	}
}

func (vd *VarData) NewPhi(id string, init int16) {
	if _, ok := vd.Phis[id]; !ok {
		vd.Phis[id] = append(vd.Phis[id], []int16{0})
		return
	}

	if init != -1 {
		vd.Phis[id] = append(vd.Phis[id], []int16{init})
		return
	}

	init = vd.GetLastState(id)
	vd.Phis[id] = append(vd.Phis[id], []int16{init})
}

func (vd *VarData) PopPhis() {
	for k := range vd.Phis {
		vd.PopPhi(k)
	}
}

func (vd *VarData) PopPhi(id string) {
	if p, ok := vd.Phis[id]; ok {
		vd.Phis[id] = p[0 : len(p)-1]
	}
}

func (vd *VarData) GetLastState(id string) int16 {
	if p, ok := vd.Phis[id]; ok {
		last := p[len(p)-1]
		return last[len(last)-1]
	}
	return 0
}

func (vd *VarData) GetStartState(id string) int16 {
	if p, ok := vd.Phis[id]; ok {
		last := p[len(p)-1]
		return last[0]
	}
	return 0
}

func (vd *VarData) SaveState() map[string]int16 {
	state := make(map[string]int16)
	for k := range vd.Phis {
		f := vd.GetStartState(k)
		state[k] = f
	}
	return state
}

func (vd *VarData) LoadState(state map[string]int16) {
	for k, i := range state {
		vd.NewPhi(k, i)
	}
}

func (vd *VarData) AppendState(state map[string]int16) {
	for k, i := range state {
		vd.StoreLastState(k, i)
	}
}

func (vd *VarData) StoreLastState(id string, n int16) {
	if p, ok := vd.Phis[id]; ok {
		last := p[len(p)-1]
		updated := append(last, n)
		vd.Phis[id][len(p)-1] = updated
	} else {
		vd.NewPhi(id, 0) //Probably a bug but fixing it breaks a bunch of stuff haha
	}
}
