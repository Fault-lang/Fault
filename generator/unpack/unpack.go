package unpack

import (
	"fault/generator/rules"
	"fault/generator/scenario"
	"fault/generator/unroll"
	"fmt"
	"strings"

	"github.com/barkimedes/go-deepcopy"
)

// Step 2 in the Generation Process: Unpack
// the Rules in the LLUnits to produce a flat
// set of SMT that reflects all state branches, forks
// phis and parallell scenarios

// Create an Event log that makes it easier to display Z3 output
// in a user friendly way

type Unpacker struct {
	Inits    []*rules.Init
	SSA      *rules.SSA
	Phis     map[string][]int16
	PhiLevel int
	HaveSeen map[string]bool    // Have we seen this variable so far in this fork?
	OnEntry  map[string][]int16 // SSA of variables on entry to a fork
	VarTypes map[string]string
	Log      *scenario.Logger
}

func NewUnpacker() *Unpacker {
	return &Unpacker{
		SSA:      rules.NewSSA(),
		Phis:     make(map[string][]int16),
		HaveSeen: make(map[string]bool),
		OnEntry:  make(map[string][]int16),
		VarTypes: make(map[string]string),
		Log:      scenario.NewLogger(),
	}
}

func (u *Unpacker) Inherits(u1 *Unpacker) {
	u.SSA = u1.SSA.Clone()
	u.PhiLevel = u1.PhiLevel
	u.OnEntry = u1.OnEntry
	u.VarTypes = u1.VarTypes
	u.Log = u1.Log
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
		if end.Get(var_name) != start.Get(var_name) {
			u.UpsertPhi(var_name, end.Get(var_name))
		}
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
	u.Log.EnterFunction(f.Ident, f.Env.CurrentRound)

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
	r.LoadContext(u.PhiLevel, u.HaveSeen, u.OnEntry, u.Log)

	switch ru := r.(type) {
	case *rules.Basic:
		inits, rule, u.SSA = ru.WriteRule(u.SSA)
	case *rules.Init:
		inits, rule, u.SSA = ru.WriteRule(u.SSA)
	case *rules.Ite:
		inits, rule = u.unPackIte(ru)
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

func (u *Unpacker) buildPhis(phis []map[string][]int16, hasPhi map[string]bool) ([]*rules.Init, []string, map[string]bool) {
	var inits []*rules.Init
	var caps []string
	if hasPhi == nil {
		hasPhi = make(map[string]bool)
	}

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
			u.Log.AddPhiOption(phi, ends)

			rule_set = append(rule_set, fmt.Sprintf("(= %s %s)", phi, ends))
		}
		if len(rule_set) == 1 {
			caps = append(caps, rule_set...)
		} else {
			caps = append(caps, fmt.Sprintf("(and %s)", strings.Join(rule_set, " ")))
		}
	}
	return inits, caps, hasPhi
}

func (u *Unpacker) buildItePhis(tPhis []map[string][]int16, fPhis []map[string][]int16) ([]*rules.Init, []string, []string) {
	var tInit, fInit []*rules.Init
	var tRules, fRules []string
	var hasPhi map[string]bool
	tInit, tRules, hasPhi = u.buildPhis(tPhis, nil)

	if len(fPhis) > 0 {
		fInit, fRules, _ = u.buildPhis(fPhis, hasPhi)
	} else {
		// If there are no rules in the false branch we still need the phis
		for k, _ := range u.Phis {
			fRules = append(fRules, fmt.Sprintf("(= %s_%d %s_%d)", k, u.SSA.Get(k), k, u.OnEntry[k][len(u.OnEntry[k])-1]))
		}
	}
	inits := append(tInit, fInit...)
	return inits, tRules, fRules
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
			u.Log.EnterFunction(call, p.Round)
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
	inits, caps, _ := u.buildPhis(phis, nil)
	u.Inits = append(u.Inits, inits...)
	capRulePhi := fmt.Sprintf("(assert (or %s))", strings.Join(caps, " "))
	rule_set = append(rule_set, capRulePhi)

	u.PopEntries()

	return u.Inits, fmt.Sprintf("%s", strings.Join(rule_set, "\n"))
}

func (u *Unpacker) unPackIte(ite *rules.Ite) ([]*rules.Init, string) {
	u.NewLevel()
	u.SetEntries(u.SSA)

	cond := u.unpackRule(ite.Cond)

	var t, f string
	var tPhis, fPhis []map[string][]int16
	var tRules, fRules, aRules []string
	var tEnds, fEnds []string
	var inits []*rules.Init
	var phis []map[string][]int16

	if len(ite.T) > 0 {
		u2 := NewUnpacker()
		u2.Inherits(u)
		for _, ru := range ite.T {
			line := u.FormatRule(ru, u2.unpackRule(ru))
			tRules = append(tRules, line)
		}

		u.SetPhis(u.SSA, u2.SSA)
		u.SSA = u2.SSA.Clone()
		PhiClone, err := deepcopy.Anything(u.Phis)

		if err != nil {
			panic(err)
		}

		tPhis = append(tPhis, PhiClone.(map[string][]int16))
		u.AddInit(u2.Inits)
	}

	if len(ite.F) > 0 {
		u2 := NewUnpacker()
		u2.Inherits(u)
		for _, ru := range ite.F {
			line := u.FormatRule(ru, u2.unpackRule(ru))
			fRules = append(fRules, line)
		}
		u.SetPhis(u.SSA, u2.SSA)
		u.SSA = u2.SSA.Clone()
		PhiClone, err := deepcopy.Anything(u.Phis)

		if err != nil {
			panic(err)
		}

		fPhis = append(phis, PhiClone.(map[string][]int16))
		u.AddInit(u2.Inits)
	}

	inits, tEnds, fEnds = u.buildItePhis(tPhis, fPhis)

	u.AddInit(inits)
	if len(tEnds) == 1 {
		t = tEnds[0]
	} else {
		t = fmt.Sprintf("(and %s)", strings.Join(tEnds, " "))
	}

	if len(fEnds) == 1 {
		f = fEnds[0]
	} else {
		f = fmt.Sprintf("(and %s)", strings.Join(fEnds, " "))
	}

	if len(ite.After) > 0 {
		u2 := NewUnpacker()
		u2.Inherits(u)
		for _, ru := range ite.After {
			line := u.FormatRule(ru, u2.unpackRule(ru))
			aRules = append(aRules, line)
		}
	}
	u.PopEntries()
	//True rules
	//False rules
	// ite cond t_phi f_phi
	ifAssert := fmt.Sprintf("(assert (ite %s %s %s))", cond, t, f)
	return u.Inits, fmt.Sprintf("%s\n%s\n%s", strings.Join(tRules, "\n"), strings.Join(fRules, "\n"), ifAssert)
}

func (u *Unpacker) FormatRule(r rules.Rule, rule string) string {
	switch r.(type) {
	case *rules.Parallels:
		return rule // Already formatted
	default:
		if rule[0:7] == "(assert" { //Already formatted
			return rule
		}

		return fmt.Sprintf("(assert %s)", rule)
	}
}
