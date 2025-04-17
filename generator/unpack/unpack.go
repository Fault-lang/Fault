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
	Inits        []*rules.Init
	CurrentBlock string
	Registry     map[string][][]string // current_round_current_block -> [(var_instance, ssa), (var_instance, ssa)]
	SSA          *rules.SSA
	Phis         map[string][]int16
	PhiLevel     int
	HaveSeen     map[string]bool    // Have we seen this variable so far in this fork?
	OnEntry      map[string][]int16 // SSA of variables on entry to a fork
	VarTypes     map[string]string
	Log          *scenario.Logger
	Round        int // Current round
}

func NewUnpacker(block_id string) *Unpacker {
	return &Unpacker{
		SSA:          rules.NewSSA(),
		CurrentBlock: block_id,
		Registry:     make(map[string][][]string),
		Phis:         make(map[string][]int16),
		HaveSeen:     make(map[string]bool),
		OnEntry:      make(map[string][]int16),
		VarTypes:     make(map[string]string),
		Log:          scenario.NewLogger(),
		Round:        0,
	}
}

func (u *Unpacker) Inherits(u1 *Unpacker) {
	u.SSA = u1.SSA.Clone()
	u.PhiLevel = u1.PhiLevel
	u.OnEntry = u1.OnEntry
	u.VarTypes = u1.VarTypes
	u.Log = u1.Log
	u.Round = u1.Round
}

func (u *Unpacker) NewLevel() {
	u.PhiLevel++
}

func (u *Unpacker) SetRound(round int) {
	u.Round = round
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

func (u *Unpacker) Register(inits []*rules.Init) {
	// We need to know in which scope variables are declared
	// in order to correctly write asserts.
	// for example assert A > B should ONLY evaluate As and Bs
	// that occur at the same time in the model. Otherwise the
	// rule becomes every possible value of A is greater than
	// every possible value of B, which probably isn't what people
	// expect.
	key := fmt.Sprintf("%s-%d_%s", "round", u.Round, u.CurrentBlock)
	if u.Registry[key] == nil {
		u.Registry[key] = [][]string{}
	}

	for _, i := range inits {
		u.Registry[key] = append(u.Registry[key], i.Tuple())
	}
}

func (u *Unpacker) UpdateRegistry(reg map[string][][]string) {
	for k, v := range reg {
		if u.Registry[k] == nil {
			u.Registry[k] = [][]string{}
		}
		u.Registry[k] = append(u.Registry[k], v...)
	}
}

func (u *Unpacker) InspectRule(ru rules.Rule) {
	if r, ok := ru.(*rules.FuncCall); ok {
		if r.Type == "Enter" {
			u.Log.EnterFunction(r.FunctionName, r.Round)
		} else {
			u.Log.ExitFunction(r.FunctionName, r.Round)
		}
	}
}

func (u *Unpacker) InitVars() []string {
	smt := []string{}
	declareOnce := make(map[string]bool)
	for _, i := range u.Inits {
		if !declareOnce[i.FullVar()] {
			_, r, _ := i.WriteRule(u.SSA)
			smt = append(smt, r)
			declareOnce[i.FullVar()] = true
		}
	}
	return smt
}

func (u *Unpacker) Unpack(con []rules.Rule, f *unroll.LLFunc) []string {
	round := fmt.Sprintf("%d", f.Env.CurrentRound)

	//Unpack the constants
	r := u.unpackConstants(con)

	u.Log.EnterFunction(f.Ident, round)

	// Unpack the rules
	r0 := u.unpackBlock(f.Start)
	r = append(r, r0...)

	function_rules := []string{}
	for _, ru := range f.Rules {
		u.InspectRule(ru)
		line := u.FormatRule(ru, u.unpackRule(ru))
		function_rules = append(function_rules, line)
	}

	r = append(r, function_rules...)
	return r
}

func (u *Unpacker) unpackConstants(con []rules.Rule) []string {
	r := []string{}
	for _, c := range con {
		if con, ok := c.(*rules.Init); ok {
			con.Global = true
			con.SSA = fmt.Sprintf("%d", u.SSA.Get(con.Ident))
			u.Register([]*rules.Init{con})
			c = con
		}
		line := u.FormatRule(c, u.unpackRule(c))
		r = append(r, line)
	}
	return r
}

func (u *Unpacker) LoadStringRules(StringRules map[string]string) {
	u.Log.StringRules = StringRules
	for k := range StringRules {
		u.Log.IsStringRule[k] = true
	}
}

func (u *Unpacker) unpackBlock(b *unroll.LLBlock) []string {
	u.SetRound(b.Env.CurrentRound)

	smt := []string{}
	for _, r := range b.Rules {
		u.InspectRule(r)
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
		_, ok := ru.Value.(*rules.Wrap)
		if u.Log.IsStringRule[ru.Ident] && ok {
			// u.SSA.Update(ru.Ident)
			// ru.SSA = fmt.Sprintf("%d", u.SSA.Get(ru.Ident))
			ru.Value = nil //Assumed false for LLVM but we don't need the value
		}

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
	case *rules.FuncCall:
		inits, rule, u.SSA = ru.WriteRule(u.SSA)
	default:
		panic(fmt.Sprintf("Unknown rule type %T", ru))
	}
	u.AddInit(inits)
	u.Register(inits)
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
			i := &rules.Init{
				Ident: var_name,
				SSA:   fmt.Sprintf("%d", u.SSA.Get(var_name)),
				Type:  u.VarTypes[var_name],
				//Value: &rules.Wrap{Value: rules.DefaultValue(u.VarTypes[var_name])},
				Value: nil,
			}
			inits = append(inits, i)
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
		for k := range u.Phis {
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

	for i, perm := range p.Permutations {
		u2 := NewUnpacker(fmt.Sprintf("%s-v%d", u.CurrentBlock, i)) //Creating a unique block name for each parallel scenario
		u2.Inherits(u)

		for _, call := range perm {
			round := fmt.Sprintf("%d", p.Round)
			u.Log.EnterFunction(call, round)
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
		u.UpdateRegistry(u2.Registry)
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
		u2 := NewUnpacker(ite.BlockNames["true"])
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
		u.UpdateRegistry(u2.Registry)
	}

	if len(ite.F) > 0 {
		u2 := NewUnpacker(ite.BlockNames["false"])
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
		u.UpdateRegistry(u2.Registry)
	}

	inits, tEnds, fEnds = u.buildItePhis(tPhis, fPhis)

	u.AddInit(inits)
	u.Register(inits)
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
		u2 := NewUnpacker(ite.BlockNames["after"])
		u2.Inherits(u)
		for _, ru := range ite.After {
			line := u.FormatRule(ru, u2.unpackRule(ru))
			aRules = append(aRules, line)
		}
		u.AddInit(u2.Inits)
		u.UpdateRegistry(u2.Registry)
	}
	u.PopEntries()
	ifAssert := fmt.Sprintf("(assert (ite %s %s %s))", cond, t, f)
	return u.Inits, fmt.Sprintf("%s\n%s\n%s\n%s", strings.Join(tRules, "\n"), strings.Join(fRules, "\n"), ifAssert, strings.Join(aRules, "\n"))
}

func (u *Unpacker) FormatRule(r rules.Rule, rule string) string {
	if rule == "" {
		return ""
	}

	if rule[0:7] == "(assert" || rule[0:8] == "\n(assert" { //Already formatted
		return rule
	}

	if _, ok := r.(*rules.Init); ok {
		return rule
	}

	return fmt.Sprintf("(assert %s)", rule)
}
