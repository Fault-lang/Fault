package unpack

import (
	"fault/generator/rules"
	"fault/generator/unroll"
	"fmt"
	"strings"

	"github.com/barkimedes/go-deepcopy"
)

// Step 2 in the Generation Process: Unpack
// the Rules in the LLUnits to produce a flat
// set of SMT that reflects all state branches, forks
// phis and parallell scenarios
type Unpacker struct {
	Inits    []*rules.Init
	SSA      *rules.SSA
	Phis     map[string][]int16
	PhiLevel int
	HaveSeen map[string]bool    // Have we seen this variable so far in this fork?
	OnEntry  map[string][]int16 // SSA of variables on entry to a fork
	VarTypes map[string]string
}

func NewUnpacker() *Unpacker {
	return &Unpacker{
		SSA:      rules.NewSSA(),
		Phis:     make(map[string][]int16),
		HaveSeen: make(map[string]bool),
		OnEntry:  make(map[string][]int16),
		VarTypes: make(map[string]string),
	}
}

func (u *Unpacker) Inherits(u1 *Unpacker) {
	u.SSA = u1.SSA.Clone()
	u.PhiLevel = u1.PhiLevel
	u.OnEntry = u1.OnEntry
	u.VarTypes = u1.VarTypes
}

func (u *Unpacker) NewLevel() {
	u.PhiLevel++
}

func (u *Unpacker) SetEntries(start *rules.SSA) {
	for var_name := range start.Iter() {
		if _, ok := u.OnEntry[var_name]; ok {
			u.OnEntry[var_name] = []int16{}
		}
		if len(u.OnEntry[var_name]) < u.PhiLevel {
			n := u.PhiLevel - len(u.OnEntry[var_name])
			filler := make([]int16, n)
			u.OnEntry[var_name] = append(u.OnEntry[var_name], filler...)
		}
		u.OnEntry[var_name] = append(u.OnEntry[var_name], start.Get(var_name))
	}
}

func (u *Unpacker) PopEntries() {
	for var_name := range u.OnEntry {
		u.OnEntry[var_name] = u.OnEntry[var_name][:len(u.OnEntry[var_name])-1]
	}
	u.PhiLevel--
}

func (u *Unpacker) GetEntry(var_name string) int16 {
	return u.OnEntry[var_name][len(u.OnEntry[var_name])-1]
}

func (u *Unpacker) SetPhis(start *rules.SSA, end *rules.SSA) {
	for var_name := range end.Iter() {
		u.UpsertPhi(var_name, end.Get(var_name))
	}
}

func (u *Unpacker) UpsertPhi(var_name string, val int16) {
	if _, ok := u.Phis[var_name]; !ok {
		u.Phis[var_name] = []int16{u.SSA.Get(var_name), val} //start value before function, end value after function
	} else {
		u.Phis[var_name][1] = val
	}
}

func (u *Unpacker) AddInit(inits []*rules.Init) {
	u.Inits = append(u.Inits, inits...)
}

func (u *Unpacker) InitVars() []string {
	smt := []string{}
	declareOnce := make(map[string]bool)
	for _, i := range u.Inits {
		if !declareOnce[i.Ident] {
			_, r, _ := i.WriteRule(u.SSA)
			smt = append(smt, r)
			declareOnce[i.Ident] = true
		}
	}
	return smt
}

func (u *Unpacker) Unpack(f *unroll.LLFunc) []string {
	// Unpack the rules
	r := u.unpackBlock(f.Start)

	function_rules := []string{}
	for _, ru := range f.Rules {
		line := u.FormatRule(ru, u.unpackRule(ru))
		function_rules = append(function_rules, line)
	}

	r = append(function_rules, r...)
	return r
}

func (u *Unpacker) unpackBlock(b *unroll.LLBlock) []string {
	smt := []string{}
	for _, r := range b.Rules {
		line := u.FormatRule(r, u.unpackRule(r))
		smt = append(smt, line)
	}

	if b.After != nil {
		next := u.unpackBlock(b.After)
		smt = append(smt, next...)
	}
	return smt
}

func (u *Unpacker) unpackRule(r rules.Rule) string {
	var rule string
	var inits []*rules.Init
	r.LoadContext(u.PhiLevel, u.HaveSeen, u.OnEntry)

	switch ru := r.(type) {
	case *rules.Basic:
		inits, rule, u.SSA = ru.WriteRule(u.SSA)
	case *rules.Init:
		inits, rule, u.SSA = ru.WriteRule(u.SSA)
	case *rules.Ite:
		inits, rule, u.SSA = ru.WriteRule(u.SSA)
	case *rules.Prefix:
		inits, rule, u.SSA = ru.WriteRule(u.SSA)
	case *rules.Infix:
		inits, rule, u.SSA = ru.WriteRule(u.SSA)
	case *rules.Parallels:
		inits, rule = u.unPackParallel(ru)
	case *rules.Ands:
		inits, rule, u.SSA = ru.WriteRule(u.SSA)
	case *rules.Wrap:
		inits, rule, u.SSA = ru.WriteRule(u.SSA)
	case *rules.Vwrap:
		inits, rule, u.SSA = ru.WriteRule(u.SSA)
	default:
		panic(fmt.Sprintf("Unknown rule type %T", ru))
	}
	u.AddInit(inits)
	return rule
}

func (u *Unpacker) buildPhis(phis []map[string][]int16) ([]*rules.Init, []string) {
	var inits []*rules.Init
	var caps []string
	hasPhi := make(map[string]bool)
	for _, p := range phis {
		var rule_set []string
		for var_name, vals := range p {

			if !hasPhi[var_name] {
				u.SSA.Update(var_name)
				hasPhi[var_name] = true
			}

			ends := fmt.Sprintf("%s_%d", var_name, vals[1])
			phi := fmt.Sprintf("%s_%d", var_name, u.SSA.Get(var_name))
			inits = append(inits, &rules.Init{Ident: phi, Type: u.VarTypes[var_name], Value: rules.DefaultValue(u.VarTypes[var_name])})

			rule_set = append(rule_set, fmt.Sprintf("(= %s %s)", phi, ends))
		}
		caps = append(caps, fmt.Sprintf("(and %s)", strings.Join(rule_set, " ")))
	}
	return inits, caps
}

func (u *Unpacker) unPackParallel(p *rules.Parallels) ([]*rules.Init, string) {
	var rule_set []string
	var phis []map[string][]int16
	u.NewLevel()
	u.SetEntries(u.SSA)

	for _, perm := range p.Permutations {
		u2 := NewUnpacker()
		u2.Inherits(u)

		for _, call := range perm {
			function_rules := []string{}
			for _, ru := range p.Calls[call] {
				line := u.FormatRule(ru, u2.unpackRule(ru))
				function_rules = append(function_rules, line)
			}
			capRule := strings.Join(function_rules, "\n")
			rule_set = append(rule_set, capRule)
		}

		u.SetPhis(u.SSA, u2.SSA)
		u.SSA = u2.SSA.Clone()
		PhiClone, err := deepcopy.Anything(u.Phis)

		if err != nil {
			panic(err)
		}
		phis = append(phis, PhiClone.(map[string][]int16))
		u.AddInit(u2.Inits)
	}
	inits, caps := u.buildPhis(phis)
	u.Inits = append(u.Inits, inits...)
	capRulePhi := fmt.Sprintf("(assert (or %s))", strings.Join(caps, " "))
	rule_set = append(rule_set, capRulePhi)

	u.PopEntries()

	return u.Inits, fmt.Sprintf("%s", strings.Join(rule_set, "\n"))
}

func (u *Unpacker) FormatRule(r rules.Rule, rule string) string {
	switch r.(type) {
	case *rules.Parallels:
		return rule // Already formatted
	default:
		return fmt.Sprintf("(assert %s)", rule)
	}
}
