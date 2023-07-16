package forks

import "sort"

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

type Fork2 struct {
	Choices  map[string][]string //slice of Branchid
	Branches map[string][]string //slice of variables in the branch
	Vars     map[string]*Var
	Bases    map[string]map[string]bool // Is there an instance of this variable in the branch?
}

// type Choice2 struct {
// 	BranchIds     []string
// 	WinningBranch string
// }

type Var struct {
	Last bool
	Phi  string
}

func InitFork() *Fork2 {
	return &Fork2{
		Choices:  make(map[string][]string),
		Branches: make(map[string][]string),
		Vars:     make(map[string]*Var),
		Bases:    make(map[string]map[string]bool),
	}
}

func (f *Fork2) AddVar(branch string, base string, id string, v *Var) {
	f.Branches[branch] = append(f.Branches[branch], id)
	f.Vars[id] = v
	if _, ok := f.Bases[branch]; ok {
		f.Bases[branch][base] = true
	} else {
		f.Bases[branch] = make(map[string]bool)
		f.Bases[branch][base] = true
	}
}

// func NewChoice() *Choice2 {
// 	return &Choice2{}
// }

func NewVar() *Var {
	return &Var{}
}

type Fork map[string]map[int16]map[string][]int16

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

func GetForkEndPoints(c []*Choice) []int16 {
	var ends []int16
	for _, v := range c {
		e := v.Values[len(v.Values)-1]
		ends = append(ends, e)
	}
	return ends
}

type Choice struct {
	Base   string  // What variable?
	Branch string  // For conditionals, is this the true block or false block?
	SSA    []int   // All the SSA assignment in this branch
	Values []int16 // All the versions of this variable in this branch
	Phi    int16   // The phi value associated with this branch
}

func (c *Choice) AddChoiceValue(n int16) *Choice {
	c.Values = append(c.Values, n)
	sort.Slice(c.Values, func(i, j int) bool { return c.Values[i] < c.Values[j] })
	return c
}

func (c *Choice) GetEnd() int16 {
	return c.Values[len(c.Values)-1]
}
