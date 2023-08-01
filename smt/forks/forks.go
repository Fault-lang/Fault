package forks

import (
	"fault/smt/rules"
	"fmt"
	"strconv"
)

// Key is the base variable name
//type Fork map[string][]*Choice
// map[base_var]map[phi_SSA]map[branch_id]
// base_var
// -> Whats the Phi value
// -> Whats the last defined state in the chain?
// -> What's the full chain?
// -> What's the unique branch id?
// -> What's the choice id?
// -> have we picked a winning branch for this choice?

// Choice[id] = {branch_ids, winning_branch}
// branch[id] = []full_var_names
// Var[base_var][branch_id] = {[]full_var_names, phi}

type Fork struct {
	Choices  map[string][]string //slice of Branchid
	Branches map[string][]string //slice of variables in the branch
	Vars     map[string]*Var
	Bases    map[string]map[string]bool // Is there an instance of this variable in the branch?
	ToKill   map[string]bool            //Quick lookup around which variables are listed as to be killed
}

type Var struct {
	Base string
	Last map[string]bool
	SSA  string
	Phi  map[string]string // map[choiceID] = phi (handles nestled phis)
}

func InitFork() *Fork {
	return &Fork{
		Choices:  make(map[string][]string),
		Branches: make(map[string][]string),
		Vars:     make(map[string]*Var),
		Bases:    make(map[string]map[string]bool),
		ToKill:   make(map[string]bool),
	}
}

func (f *Fork) MarkMany(ids []string) {
	for _, id := range ids {
		f.Mark(id)
	}
}

func (f *Fork) Mark(id string) {
	f.ToKill[id] = true
}

func (f *Fork) MarkedForDeath(id string) bool {
	return f.ToKill[id]
}

func (f *Fork) InBranch(branch string, id string) bool {
	for _, v := range f.Branches[branch] {
		if id == v {
			return true
		}
	}
	return false
}

func (f *Fork) AddToBranch(branch string, id string) {
	if !f.InBranch(branch, id) {
		f.Branches[branch] = append(f.Branches[branch], id)
	}
}

func (f *Fork) AddVar(branch string, base string, id string, v *Var) {
	f.AddToBranch(branch, id)

	if _, ok := f.Vars[id]; !ok {
		f.Vars[id] = v
	} else {
		for k, vlast := range v.Last {
			f.Vars[id].Last[k] = vlast
		}
		for k, vphi := range v.Phi {
			f.Vars[id].Phi[k] = vphi
		}
	}

	if _, ok := f.Bases[branch]; ok {
		f.Bases[branch][base] = true
	} else {
		f.Bases[branch] = make(map[string]bool)
		f.Bases[branch][base] = true
	}
}

func (f *Fork) IsPhi(r *rules.Wrap) bool {
	if f.Vars[r.Value] == nil {
		return false
	}
	return f.Vars[r.Value].IsPhi(r.Choice())
}

func NewVar(base string, last bool, ssa string, choice string, phi string) *Var {
	v := &Var{Base: base, Last: make(map[string]bool), SSA: ssa, Phi: make(map[string]string)}
	v.Last[choice] = last
	v.Phi[choice] = phi
	return v
}

func (v *Var) IsPhi(choice string) bool {
	return v.SSA == v.Phi[choice]
}

func (v *Var) FullPhi(choice string) string {
	return fmt.Sprintf("%s_%s", v.Base, v.Phi[choice])
}

func (v *Var) PhiInt(choice string) int {
	i, err := strconv.ParseInt(v.Phi[choice], 10, 32)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func (v *Var) PhiInt16(choice string) int16 {
	i, err := strconv.ParseInt(v.Phi[choice], 10, 32)
	if err != nil {
		panic(err)
	}
	return int16(i)
}

type PhiState struct {
	levels int
}

func NewPhiState() *PhiState {
	return &PhiState{
		levels: 0,
	}
}

func (p *PhiState) Check() bool {
	return p.levels > 0
}

func (p *PhiState) Level() int {
	return p.levels
}

func (p *PhiState) In() {
	p.levels = p.levels + 1
}

func (p *PhiState) Out() {
	if p.levels != 0 {
		p.levels = p.levels - 1
	}
}
