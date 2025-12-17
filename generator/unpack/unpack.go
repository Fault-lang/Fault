package unpack

import (
	"fault/generator/rules"
	"fault/generator/scenario"
	"fault/generator/unroll"
	"fault/util"
	"fmt"
	"strings"
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
	Phis         map[string][]int16 // A tuple of [entry ssa, last updated ssa]
	PhiLevel     int
	HaveSeen     map[string]bool    // Have we seen this variable so far in this fork?
	OnEntry      map[string][]int16 // SSA of variables on entry to a fork
	VarTypes     map[string]string
	Whens        map[string][]map[string]string // list of variable combinations for when/then asserts
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
		Whens:        make(map[string][]map[string]string),
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

	for k, v := range u1.Whens {
		u.Whens[k] = v
	}
}

func (u *Unpacker) NewLevel() {
	u.PhiLevel++
}

func (u *Unpacker) SetRound(round int) {
	u.Round = round
}

func (u *Unpacker) SetEntries(start *rules.SSA) {
	for var_name := range start.Iter() {
		if _, ok := u.OnEntry[var_name]; !ok {
			u.OnEntry[var_name] = []int16{}
		}

		if len(u.OnEntry[var_name]) < u.PhiLevel {
			n := u.PhiLevel - len(u.OnEntry[var_name]) //Calculate how many entries we need to add
			filler := make([]int16, n)                 //Generate a slice of n 0s
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

func (u *Unpacker) GetPhis(start *rules.SSA, end *rules.SSA) map[string][]int16 {
	phis := make(map[string][]int16)
	for var_name := range end.Iter() {
		if end.Get(var_name) != start.Get(var_name) {
			phis[var_name] = []int16{start.Get(var_name), end.Get(var_name)}
		}
	}
	return phis
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
	if len(inits) == 0 {
		return
	}
	key := fmt.Sprintf("%s-%d_%s", "round", inits[0].GetRound(), u.CurrentBlock)
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
	//Unpack the constants
	r := u.unpackConstants(con)
	u.Log.EnterFunction(f.Ident, u.Round)

	// Unpack the rules
	r0 := u.unpackBlock(f.Start)
	r = append(r, r0...)

	function_rules := []string{}
	for _, ru := range f.Rules {
		u.InspectRule(ru)
		inits, finishedRules := u.unpackRule(ru)
		line := u.FormatRule(ru, finishedRules)
		function_rules = append(function_rules, line)
		u.AddInit(inits)
		u.Register(inits)
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
		inits, finishedRules := u.unpackRule(c)
		line := u.FormatRule(c, finishedRules)
		r = append(r, line)
		u.AddInit(inits)
		u.Register(inits)
	}
	return r
}

func (u *Unpacker) LoadStringRules(StringRules map[string]string, IsCompound map[string]bool) {
	u.Log.IsCompound = IsCompound
	u.Log.StringRules = StringRules
	for k := range StringRules {
		u.Log.IsStringRule[k] = true
	}
}

func (u *Unpacker) unpackBlock(b *unroll.LLBlock) []string {
	u.SetRound(b.Round)

	smt := []string{}
	for _, r := range b.Rules {
		u.InspectRule(r)
		inits, finishedRules := u.unpackRule(r)
		line := u.FormatRule(r, finishedRules)
		smt = append(smt, line)
		u.AddInit(inits)
		u.Register(inits)
	}

	if b.After != nil {
		next := u.unpackBlock(b.After)
		smt = append(smt, next...)
	}
	return smt
}

func (u *Unpacker) unpackRule(r rules.Rule) ([]*rules.Init, string) {
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
		inits, rule = u.unpackIte(ru)
	case *rules.Prefix:
		inits, rule, u.SSA = ru.WriteRule(u.SSA)
	case *rules.Infix:
		inits, rule, u.SSA = ru.WriteRule(u.SSA)
	case *rules.Parallels:
		inits, rule = u.unpackParallel(ru)
	case *rules.Ands:
		inits, rule, u.SSA = ru.WriteRule(u.SSA)
	case *rules.Ors:
		inits, rule = u.unpackOrs(ru)
	case *rules.Wrap:
		inits, rule, u.SSA = ru.WriteRule(u.SSA)
	case *rules.Vwrap:
		inits, rule, u.SSA = ru.WriteRule(u.SSA)
	case *rules.FuncCall:
		inits, rule, u.SSA = ru.WriteRule(u.SSA)
	case *rules.Stay:
		u.Log.AddMessage("Stay in current state", u.Round)
	default:
		panic(fmt.Sprintf("Unknown rule type %T", ru))
	}
	u.Whens = u.unpackWhenThen(r, u.Whens)
	//u.AddInit(inits)
	//u.Register(inits)
	return inits, rule
}

func (u *Unpacker) unpackWhenThen(r rules.Rule, whens map[string][]map[string]string) map[string][]map[string]string {
	switch ru := r.(type) {
	case *rules.Basic:
		whens = u.unpackWhenThen(ru.X, whens)
		whens = u.unpackWhenThen(ru.Y, whens)
	case *rules.Init:
		// Nothing to do
	case *rules.Ite:
		for _, t := range ru.T {
			whens = u.unpackWhenThen(t, whens)
		}
		for _, f := range ru.F {
			whens = u.unpackWhenThen(f, whens)
		}
	case *rules.Prefix:
		whens = u.unpackWhenThen(ru.X, whens)
	case *rules.Infix:
		whens = u.unpackWhenThen(ru.X, whens)
		whens = u.unpackWhenThen(ru.Y, whens)
	case *rules.Parallels:
		// Nothing to do
	case *rules.Ands:
		for _, ru := range ru.X {
			whens = u.unpackWhenThen(ru, whens)
		}
	case *rules.Ors:
		for _, branch := range ru.X {
			for _, ru := range branch {
				whens = u.unpackWhenThen(ru, whens)
			}
		}
	case *rules.Wrap:
		// Unlike other asserts, we do this every time we see a new init of a variable
		// so that we capture overlapping states correctly
		if len(ru.Whens) == 0 {
			return whens
		}

		if !ru.Init {
			return whens
		}

		for a, when := range ru.Whens {
			assert_combo := make(map[string]string) // base => ssa instance
			kssa := u.SSA.Get(ru.Value)
			assert_combo[ru.Value] = fmt.Sprintf("%s_%d", ru.Value, kssa)
			for _, v := range when {
				current := u.SSA.Get(v)
				assert_combo[v] = fmt.Sprintf("%s_%d", v, current)
			}
			// Only add the assert combo if it's not already present
			if len(whens[a]) > 0 && util.CompareStringMaps(whens[a][len(whens[a])-1], assert_combo) {
				continue
			}
			whens[a] = append(whens[a], assert_combo)
		}

	case *rules.Vwrap:
		// Nothing to do
	case *rules.FuncCall:
		// Nothing to do
	case *rules.Stay:
		// Nothing to do
	default:
		panic(fmt.Sprintf("Unknown rule type %T", ru))
	}
	return whens
}

func (u *Unpacker) buildPhisOrs(phis []map[string][]int16, hasPhi map[string]bool) ([]*rules.Init, [][]string, map[string]bool) {
	var inits []*rules.Init
	var caps [][]string
	sync := make(map[string][]int) // Phis in one branch but not the other

	if hasPhi == nil {
		hasPhi = make(map[string]bool)
	}

	var last map[string][]int16

	for i, p := range phis {
		var rule_set []string
		for var_name, vals := range p {
			//If the vals are the same as the last known value, we don't need to create a phi
			if last != nil {
				if last_vals, ok := last[var_name]; ok {
					if vals[len(vals)-1] == last_vals[len(last_vals)-1] {
						continue // No need to create a phi, the value is the same as the last known value
					}
				}
			}

			sync[var_name] = append(sync[var_name], i) // Store the branch the phi was found in

			if !hasPhi[var_name] {
				u.SSA.Update(var_name)
				hasPhi[var_name] = true
			}

			var idx int
			if vals[len(vals)-1] == u.SSA.Get(var_name) { // I actually don't know what's wrong here
				idx = len(vals) - 2 // Bug in phis for ORs. The Phi cannot be the same as the last known value
			} else {
				idx = len(vals) - 1 // The last value is the current value
			}

			ends := fmt.Sprintf("%s_%d", var_name, vals[idx])
			phi := fmt.Sprintf("%s_%d", var_name, u.SSA.Get(var_name))

			i := &rules.Init{
				Ident: var_name,
				SSA:   fmt.Sprintf("%d", u.SSA.Get(var_name)),
				Type:  u.VarTypes[var_name],
				//Value: &rules.Wrap{Value: rules.DefaultValue(u.VarTypes[var_name])},
				Value: nil,
			}
			i.SetRound(u.Round)
			inits = append(inits, i)
			u.Log.AddPhiOption(phi, ends)

			rule_set = append(rule_set, fmt.Sprintf("(= %s %s)", phi, ends))
		}

		caps = append(caps, rule_set)
		last = p
	}

	for j, _ := range caps {
		if len(caps[j]) != len(sync) { // This branch is missing some phis
			for var_name, branches := range sync {
				found := false
				for _, branch := range branches {
					if branch == j { // This branch is missing the phi
						found = true
					}
				}
				if !found {
					ends := fmt.Sprintf("%s_%d", var_name, u.OnEntry[var_name][len(u.OnEntry[var_name])-1])
					phi := fmt.Sprintf("%s_%d", var_name, u.SSA.Get(var_name))

					i := &rules.Init{
						Ident: var_name,
						SSA:   fmt.Sprintf("%d", u.SSA.Get(var_name)),
						Type:  u.VarTypes[var_name],
						//Value: &rules.Wrap{Value: rules.DefaultValue(u.VarTypes[var_name])},
						Value: nil,
					}
					i.SetRound(u.Round)
					inits = append(inits, i)
					u.Log.AddPhiOption(phi, ends)

					caps[j] = append(caps[j], fmt.Sprintf("(= %s %s)", phi, ends))
				}
			}
		}
	}

	return inits, caps, hasPhi
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

			ends := fmt.Sprintf("%s_%d", var_name, vals[len(vals)-1])
			phi := fmt.Sprintf("%s_%d", var_name, u.SSA.Get(var_name))

			i := &rules.Init{
				Ident: var_name,
				SSA:   fmt.Sprintf("%d", u.SSA.Get(var_name)),
				Type:  u.VarTypes[var_name],
				//Value: &rules.Wrap{Value: rules.DefaultValue(u.VarTypes[var_name])},
				Value: nil,
			}
			i.SetRound(u.Round)
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

func (u *Unpacker) buildItePhis(tPhis []map[string][]int16, fPhis []map[string][]int16, blocknames map[string]string) ([]*rules.Init, *scenario.BranchSelector, *scenario.BranchSelector) {
	var tInit, fInit []*rules.Init
	var tRules, fRules []string
	var hasPhi map[string]bool
	tInit, tRules, hasPhi = u.buildPhis(tPhis, nil)
	tSelectorName := util.FormatBlock(blocknames["true"])
	tInit = append(tInit, rules.NewInit(tSelectorName, "Bool", 0))
	tSelectorRule := u.Log.NewBranchSelector(tSelectorName, 0, tRules, InitsToList((tInit)))
	u.Log.AddBranchSelector(tSelectorRule)
	//u.Log.QueueFork(InitsToList(tInit))

	if len(fPhis) > 0 {
		fInit, fRules, _ = u.buildPhis(fPhis, hasPhi)
	} else {
		// If there are no rules in the false branch we still need the phis
		blocknames["false"] = fmt.Sprintf("%sfalse", blocknames["true"][0:len(blocknames["true"])-4])
		for l := range tPhis {
			for k, _ := range tPhis[l] {
				fRules = append(fRules, fmt.Sprintf("(= %s_%d %s_%d)", k, u.SSA.Get(k), k, u.OnEntry[k][len(u.OnEntry[k])-1]))
			}
		}
	}
	fSelectorName := util.FormatBlock(blocknames["false"])
	fInit = append(fInit, rules.NewInit(fSelectorName, "Bool", 0))
	fSelectorRule := u.Log.NewBranchSelector(fSelectorName, 0, fRules, InitsToList((fInit)))
	u.Log.AddBranchSelector(fSelectorRule)

	//u.Log.QueueFork(InitsToList(fInit))
	inits := append(tInit, fInit...)
	return inits, tSelectorRule, fSelectorRule
}

func (u *Unpacker) unpackOrs(o *rules.Ors) ([]*rules.Init, string) {
	var rule_set [][]string
	var ret []string
	var hasPhi map[string]bool
	var inits []*rules.Init
	var queue [][]string //All the vars that have been initialized
	var caps [][]string
	var branches []map[string][]int16
	u.NewLevel()
	u.SetEntries(u.SSA)

	u.Log.EnterFunction(o.BranchName, o.Round)

	for _, ru := range o.X {
		var lines []string
		u2 := NewUnpacker(fmt.Sprintf("%s-%s", u.CurrentBlock, o.BranchName)) //Creating a unique block name for each or scenario
		u2.Inherits(u)

		var initQue []string
		for _, l := range ru {
			init, line := u2.unpackRule(l)
			lines = append(lines, line)
			initQue = append(initQue, InitsToList(init)...)
			u.AddInit(init)
			u.Register(init)
		}
		rule_set = append(rule_set, lines)
		PhiClone := u.GetPhis(u.SSA, u2.SSA)
		branches = append(branches, PhiClone)
		u.SSA = u2.SSA.Clone()

		//u.Log.QueueFork(InitsToList(u2.Inits))
		queue = append(queue, initQue)
		u.AddInit(u2.Inits)
		u.UpdateRegistry(u2.Registry)
	}

	inits, caps, hasPhi = u.buildPhisOrs(branches, hasPhi)
	var selectors []string

	for i, _ := range rule_set {
		selectorName := fmt.Sprintf("%s_%d", o.BranchName, i)
		inits = append(inits, rules.NewInit(o.BranchName, "Bool", i))
		selectors = append(selectors, selectorName)
		selectorRule := u.Log.NewBranchSelector(o.BranchName, i, caps[i], queue[i])
		u.Log.AddBranchSelector(selectorRule)
		ret = append(ret, fmt.Sprintf("(assert %s", selectorRule.WriteRule()))
	}

	u.AddInit(inits)
	u.PopEntries()

	return u.Inits, strictOr(selectors)
}

func (u *Unpacker) unpackParallel(p *rules.Parallels) ([]*rules.Init, string) {
	var rule_set []string
	var phis []map[string][]int16
	var selectors []string

	u.NewLevel()
	u.SetEntries(u.SSA)

	for i, perm := range p.Permutations {
		SelectorName := fmt.Sprintf("%s_%d", u.CurrentBlock, i)
		selectors = append(selectors, SelectorName)

		u2 := NewUnpacker(SelectorName) //Creating a unique block name for each parallel scenario
		u2.Inherits(u)

		for _, call := range perm {
			u.Log.EnterFunction(call, p.Round)
			function_rules := []string{}
			for _, ru := range p.Calls[call] {
				inits, finishedRules := u2.unpackRule(ru)
				line := u.FormatRule(ru, finishedRules)
				function_rules = append(function_rules, line)
				u.AddInit(inits)
				u.Register(inits)
			}
			rules := strings.Join(function_rules, "\n")
			rule_set = append(rule_set, rules)
		}

		PhiClone := u.GetPhis(u.SSA, u2.SSA)
		u.SSA = u2.SSA.Clone()

		phis = append(phis, PhiClone)
		inits, caps, _ := u.buildPhis(phis, nil)
		SelectorRule := u.Log.NewBranchSelector(u.CurrentBlock, i, caps, InitsToList((inits)))
		rule_set = append(rule_set, fmt.Sprintf("(assert %s", SelectorRule.WriteRule()))
		u.Log.AddBranchSelector(SelectorRule)
		u2.Inits = append(u2.Inits, rules.NewInit(u.CurrentBlock, "Bool", i))

		u.AddInit(u2.Inits)
		u.UpdateRegistry(u2.Registry)
	}
	// inits, caps, _ := u.buildPhis(phis, nil)
	// u.Inits = append(u.Inits, inits...)
	capRulePhi := fmt.Sprintf("(assert %s)", strictOr(selectors))
	rule_set = append(rule_set, capRulePhi)

	u.PopEntries()

	return u.Inits, fmt.Sprint(strings.Join(rule_set, "\n"))
}

func (u *Unpacker) unpackIteBlock(blockName string, block []rules.Rule) ([]string, []map[string][]int16) {
	var bPhis []map[string][]int16
	var bRules []string
	u2 := NewUnpacker(blockName)
	u2.Inherits(u)
	for _, ru := range block {
		if _, ok := ru.(*rules.FuncCall); ok {
			continue
		}

		inits, finishedRules := u2.unpackRule(ru)
		line := u.FormatRule(ru, finishedRules)
		bRules = append(bRules, line)
		u.AddInit(inits)
		u.Register(inits)
	}

	PhiClone := u.GetPhis(u.SSA, u2.SSA)
	u.SSA = u2.SSA.Clone()

	bPhis = append(bPhis, PhiClone)
	u.AddInit(u2.Inits)
	u.UpdateRegistry(u2.Registry)
	return bRules, bPhis
}

func (u *Unpacker) unpackIte(ite *rules.Ite) ([]*rules.Init, string) {
	u.NewLevel()
	u.SetEntries(u.SSA)

	//If this is just a stay(); Then we don't need to do anything
	_, isStay := ite.T[0].(*rules.Stay)
	if len(ite.T) == 1 && isStay && len(ite.F) == 0 && len(ite.After) == 0 {
		return []*rules.Init{}, ""
	}

	_, cond := u.unpackRule(ite.Cond)

	var t, f string
	var tPhis, fPhis []map[string][]int16
	var tRules, fRules, aRules []string
	var tEnds, fEnds *scenario.BranchSelector
	var inits []*rules.Init

	if len(ite.T) > 0 {
		tRules, tPhis = u.unpackIteBlock(ite.BlockNames["true"], ite.T)
	}

	if len(ite.F) > 0 {
		fRules, fPhis = u.unpackIteBlock(ite.BlockNames["false"], ite.F)
	}

	inits, tEnds, fEnds = u.buildItePhis(tPhis, fPhis, ite.BlockNames)

	endRule := fmt.Sprintf("(assert %s)", strictOr([]string{tEnds.Id(), fEnds.Id()}))

	u.AddInit(inits)
	u.Register(inits)
	t = tEnds.WriteRule()
	f = fEnds.WriteRule()

	u.PopEntries()

	if len(ite.After) > 0 {
		aRules, _ = u.unpackIteBlock(ite.BlockNames["after"], ite.After)
	}
	aRules = append(aRules, endRule)

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

func strictOr(rules []string) string {
	//Return rules where only ONE possible rules can be true
	var choice []string
	for i, _ := range rules {
		var ands []string
		for j := 0; j < len(rules); j++ {
			if i != j {
				ands = append(ands, fmt.Sprintf("(not %s)", rules[j]))
			} else {
				ands = append(ands, rules[j])
			}
		}
		var a string
		if len(ands) == 1 {
			a = ands[0]
		} else {
			a = fmt.Sprintf("(and %s)", strings.Join(ands, "\n"))
		}
		choice = append(choice, a)
	}

	if len(choice) == 1 {
		return choice[0]
	}

	return fmt.Sprintf("(or %s)", strings.Join(choice, "\n"))
}

func InitsToList(inits []*rules.Init) []string {
	var init_list []string
	for _, i := range inits {
		init_list = append(init_list, i.FullVar())
	}
	return init_list
}
