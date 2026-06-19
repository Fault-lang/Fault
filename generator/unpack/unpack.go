package unpack

import (
	"fault/generator/rules"
	"fault/generator/scenario"
	"fault/generator/unroll"
	"fault/util"
	"fmt"
	"sort"
	"strings"
)

// Step 2 in the Generation Process: Unpack
// the Rules in the LLUnits to produce a flat
// set of SMT that reflects all state branches, forks
// phis and parallell scenarios

// Create an Event log that makes it easier to display Z3 output
// in a user friendly way

type Unpacker struct {
	Inits           []*rules.Init
	CurrentBlock    string
	Registry        map[string][][]string // current_round_current_block -> [(var_instance, ssa), (var_instance, ssa)]
	SSA             *rules.SSA
	Phis            map[string][]int16 // A tuple of [entry ssa, last updated ssa]
	PhiLevel        int
	HaveSeen        map[string]bool    // Have we seen this variable so far in this fork?
	OnEntry         map[string][]int16 // SSA of variables on entry to a fork
	VarTypes        map[string]string
	Whens           map[string][]map[string]string // list of variable combinations for when/then asserts
	Log             *scenario.Logger
	Round           int // Current round
	AssumeOverrides map[string]string // base var name → literal value from assume == constraint
	Warnings        []string
}

func NewUnpacker(block_id string) *Unpacker {
	return &Unpacker{
		SSA:             rules.NewSSA(),
		CurrentBlock:    block_id,
		Registry:        make(map[string][][]string),
		Phis:            make(map[string][]int16),
		HaveSeen:        make(map[string]bool),
		OnEntry:         make(map[string][]int16),
		VarTypes:        make(map[string]string),
		Whens:           make(map[string][]map[string]string),
		Log:             scenario.NewLogger(),
		Round:           0,
		AssumeOverrides: make(map[string]string),
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
	u.Log.EnterFunction(f.Ident, f.Env.CurrentRound)

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
		if initRule, ok := c.(*rules.Init); ok {
			initRule.Global = true
			initRule.SSA = fmt.Sprintf("%d", u.SSA.Get(initRule.Ident))
			u.Register([]*rules.Init{initRule})

			if overrideVal, overridden := u.AssumeOverrides[initRule.Ident]; overridden {
				initVal := initConstantLiteral(initRule)
				if initVal != "" && initVal != overrideVal {
					u.Warnings = append(u.Warnings, fmt.Sprintf(
						"assume overrides constant init for %s: was %s, assume sets %s",
						initRule.Ident, initVal, overrideVal))
				}
				initRule.SuppressValueAssertion = true
			}

			c = initRule
		}
		inits, finishedRules := u.unpackRule(c)
		line := u.FormatRule(c, finishedRules)
		r = append(r, line)
		u.AddInit(inits)
		u.Register(inits)
	}
	return r
}

// initConstantLiteral returns the literal string value stored in an Init rule's
// Value field (a *rules.Wrap), or "" if it cannot be determined.
func initConstantLiteral(i *rules.Init) string {
	if w, ok := i.Value.(*rules.Wrap); ok {
		return w.Value
	}
	return ""
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
		if base, suppressed := u.suppressedRunInitInfix(r); suppressed {
			// An assume overrides this __run init assertion — keep the declaration
			// (via AddInit) but drop the value assertion so the assume wins.
			u.AddInit(inits)
			u.Register(inits)
			u.Warnings = append(u.Warnings, fmt.Sprintf(
				"assume overrides constant init for %s; suppressing init assertion", base))
			continue
		}
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
		u.SetRound(ru.GetRound())
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
	case *rules.HistoryWrap:
		inits, rule, u.SSA = ru.WriteRule(u.SSA)
	case *rules.FuncCall:
		inits, rule, u.SSA = ru.WriteRule(u.SSA)
	case *rules.Stay:
		u.Log.AddMessage("Stay in current state", u.Round)
	case *rules.SynthSlot:
		inits, rule = u.unpackSynthSlot(ru)
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
	case *rules.HistoryWrap:
		// Nothing to do
	case *rules.FuncCall:
		// Nothing to do
	case *rules.Stay:
		// Nothing to do
	case *rules.SynthSlot:
		// Nothing to do — candidates' when/then will be handled when they unpack
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
			//If the vals are the same as the last known value, we don't need to create a phi rule,
			// but we still record the branch in sync so the fill-in pass doesn't incorrectly
			// replace this branch's value with OnEntry.
			if last != nil {
				if last_vals, ok := last[var_name]; ok {
					if vals[len(vals)-1] == last_vals[len(last_vals)-1] {
						sync[var_name] = append(sync[var_name], i)
						continue
					}
				}
			}

			sync[var_name] = append(sync[var_name], i) // Store the branch the phi was found in

			if !hasPhi[var_name] {
				u.SSA.Update(var_name)
				hasPhi[var_name] = true
			}

			// ends is always vals[-1]: the branch's final SSA value for this variable.
			// The phi SSA (u.SSA.Get after Update) is always last_branch_end+1, so it
			// can never equal vals[-1] when SSA is managed correctly.
			ends := fmt.Sprintf("%s_%d", var_name, vals[len(vals)-1])
			phi := fmt.Sprintf("%s_%d", var_name, u.SSA.Get(var_name))

			i := rules.NewInit(var_name, u.VarTypes[var_name], int(u.SSA.Get(var_name)), nil, false, false)
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
					entryIdx := int16(0)
					if entry := u.OnEntry[var_name]; len(entry) > 0 {
						entryIdx = entry[len(entry)-1]
					}
					ends := fmt.Sprintf("%s_%d", var_name, entryIdx)
					phi := fmt.Sprintf("%s_%d", var_name, u.SSA.Get(var_name))
					i := rules.NewInit(var_name, u.VarTypes[var_name], int(u.SSA.Get(var_name)), nil, false, false)
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

			isNew := !hasPhi[var_name]
			if isNew {
				u.SSA.Update(var_name)
				hasPhi[var_name] = true
			}

			ends := fmt.Sprintf("%s_%d", var_name, vals[len(vals)-1])
			phi := fmt.Sprintf("%s_%d", var_name, u.SSA.Get(var_name))
			i := rules.NewInit(var_name, u.VarTypes[var_name], int(u.SSA.Get(var_name)), nil, false, false)
			i.SetRound(u.Round)
			inits = append(inits, i)
			u.Log.AddPhiOption(phi, ends)

			if isNew {
				// Log the phi output variable in the current function scope.
				// This must happen here (at creation time) rather than when the phi
				// is later read in a condition, which would misplace it in the next
				// round's scope.
				u.Log.UpdateVariable(phi, false)

				// Record per-round phi SSAs for HistoryWrap's value[now-N] resolution.
				// On the first phi ever for this variable, prepend the initial SSA (before
				// any rounds) so that RoundPhis[0] = initial, RoundPhis[N] = phi after round N.
				if _, exists := u.Log.RoundPhis[var_name]; !exists {
					initialSSA := u.OnEntry[var_name][len(u.OnEntry[var_name])-1]
					u.Log.AddRoundPhi(var_name, initialSSA)
				}
				u.Log.AddRoundPhi(var_name, u.SSA.Get(var_name))
			}

			rule_set = append(rule_set, fmt.Sprintf("(= %s %s)", phi, ends))
		}
		if len(rule_set) == 1 {
			caps = append(caps, rule_set...)
		} else if len(rule_set) > 1 {
			caps = append(caps, fmt.Sprintf("(and %s)", strings.Join(rule_set, " ")))
		}

	}
	return inits, caps, hasPhi
}

func (u *Unpacker) buildItePhis(tPhis []map[string][]int16, fPhis []map[string][]int16, tInit []*rules.Init, fInit []*rules.Init, blocknames map[string]string) ([]*rules.Init, *scenario.BranchSelector, *scenario.BranchSelector) {
	var tI, fI []*rules.Init
	var tRules, fRules []string
	var hasPhi map[string]bool

	// Save the exclusive branch inits (before phi outputs are appended) for use in
	// BranchSelector.Vars. Phi output variables are written by BOTH branches via the
	// ITE phi merge, so they are always live — including them in Vars would cause Kill()
	// to incorrectly mark them dead when the other branch selector is false.
	tExclusive := make([]*rules.Init, len(tInit))
	copy(tExclusive, tInit)
	fExclusive := make([]*rules.Init, len(fInit))
	copy(fExclusive, fInit)

	tI, tRules, hasPhi = u.buildPhis(tPhis, nil)
	tInit = append(tInit, tI...)

	if len(fPhis) > 0 {
		fI, fRules, _ = u.buildPhis(fPhis, hasPhi)
		fInit = append(fInit, fI...)

		// Liveness check: if a variable is modified in only one branch, the other
		// branch needs an identity phi (= entry value) to keep the model complete.
		tPhiVars := varsInPhis(tPhis)
		fPhiVars := varsInPhis(fPhis)
		for k := range tPhiVars {
			if !fPhiVars[k] {
				if entry := u.OnEntry[k]; len(entry) > 0 {
					fRules = append(fRules, fmt.Sprintf("(= %s_%d %s_%d)", k, u.SSA.Get(k), k, entry[len(entry)-1]))
				}
			}
		}
		for k := range fPhiVars {
			if !tPhiVars[k] {
				if entry := u.OnEntry[k]; len(entry) > 0 {
					tRules = append(tRules, fmt.Sprintf("(= %s_%d %s_%d)", k, u.SSA.Get(k), k, entry[len(entry)-1]))
				}
			}
		}
	} else {
		// If there are no rules in the false branch we still need the phis
		blocknames["false"] = fmt.Sprintf("%sfalse", blocknames["true"][0:len(blocknames["true"])-4])
		for l := range tPhis {
			for k := range tPhis[l] {
				fRules = append(fRules, fmt.Sprintf("(= %s_%d %s_%d)", k, u.SSA.Get(k), k, u.OnEntry[k][len(u.OnEntry[k])-1]))
			}
		}
	}

	// Create selectors after all rules are finalized (so complement rules are included).
	// Use the current round as the SSA index so selectors are unique per round —
	// reusing the same Bool variable across rounds causes UNSAT when the branch condition
	// evaluates differently in different rounds.
	// Use only exclusive branch inits (not phi outputs) in Vars so Kill() does not
	// incorrectly mark phi output variables as dead.
	tSelectorInit := rules.NewInit(tSelectorName(blocknames), "Bool", u.Round, nil, false, false)
	tInit = append(tInit, tSelectorInit)
	tSelectorRule := u.Log.NewBranchSelector(tSelectorName(blocknames), u.Round, tRules, InitsToList(append(tExclusive, tSelectorInit)))
	u.Log.AddBranchSelector(tSelectorRule)

	fSelectorInit := rules.NewInit(fSelectorName(blocknames), "Bool", u.Round, nil, false, false)
	fInit = append(fI, fSelectorInit)
	fSelectorRule := u.Log.NewBranchSelector(fSelectorName(blocknames), u.Round, fRules, InitsToList(append(fExclusive, fSelectorInit)))
	u.Log.AddBranchSelector(fSelectorRule)

	inits := append(tInit, fInit...)
	return inits, tSelectorRule, fSelectorRule
}

func tSelectorName(blocknames map[string]string) string {
	return util.FormatBlock(blocknames["true"])
}

func fSelectorName(blocknames map[string]string) string {
	return util.FormatBlock(blocknames["false"])
}

func varsInPhis(phis []map[string][]int16) map[string]bool {
	vars := make(map[string]bool)
	for _, p := range phis {
		for k := range p {
			vars[k] = true
		}
	}
	return vars
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
			u.AddInit(init) // declare branch-internal vars
			// Do NOT register branch-internal inits: they are unconstrained when
			// this branch is not selected, so including them in the registry would
			// let temporal constraints reference free variables.
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
	var branchAssertions []string

	for i, rs := range rule_set {
		selectorName := fmt.Sprintf("%s_%d", o.BranchName, i)
		inits = append(inits, rules.NewInit(o.BranchName, "Bool", i, nil, false, false))
		selectors = append(selectors, selectorName)
		selectorRule := u.Log.NewBranchSelector(o.BranchName, i, caps[i], queue[i])
		u.Log.AddBranchSelector(selectorRule)
		if rule := selectorRule.WriteRule(); rule != "" {
			ret = append(ret, fmt.Sprintf("(assert %s)", rule))
		}

		// Wrap branch rules in an implication from the selector
		if len(rs) > 0 {
			// Filter out empty rules
			var nonEmptyRules []string
			for _, r := range rs {
				if r != "" {
					nonEmptyRules = append(nonEmptyRules, r)
				}
			}

			if len(nonEmptyRules) > 0 {
				var branchRule string
				if len(nonEmptyRules) == 1 {
					branchRule = nonEmptyRules[0]
				} else {
					branchRule = fmt.Sprintf("(and %s)", strings.Join(nonEmptyRules, "\n"))
				}
				branchAssertions = append(branchAssertions, fmt.Sprintf("(assert (=> %s %s))", selectorName, branchRule))
			}
		}
	}

	u.AddInit(inits)
	u.PopEntries()

	cap := fmt.Sprintf("(assert %s)", strictOr(selectors))

	// Return only the canonical post-phi inits (phi-merge vars + selector vars).
	// Branch-internal inits are already in u.Inits for declaration but must not
	// be registered into the parent registry — they are unconstrained when their
	// branch is inactive, which would make temporal constraints trivially satisfiable.
	return inits, fmt.Sprintf("%s\n%s\n%s", strings.Join(branchAssertions, "\n"), strings.Join(ret, "\n"), cap)
}

// unpackSynthSlot handles a synthesis step (__).
// It mirrors unpackOrs: one branch per candidate function, exactly-one selector constraint,
// conditional transitions, and phi merges for changed variables (frame conditions are
// provided automatically by the phi mechanism — unchanged vars get identity phis).
func (u *Unpacker) unpackSynthSlot(slot *rules.SynthSlot) ([]*rules.Init, string) {
	if len(slot.Candidates) == 0 {
		return nil, ""
	}

	// Sort candidate names for deterministic SMT output.
	names := make([]string, 0, len(slot.Candidates))
	for n := range slot.Candidates {
		names = append(names, n)
	}
	sort.Strings(names)

	var ruleSet [][]string
	var ret []string
	var hasPhi map[string]bool
	var inits []*rules.Init
	var queue [][]string
	var caps [][]string
	var branches []map[string][]int16

	slotName := fmt.Sprintf("synth_%d", slot.Round)
	u.NewLevel()
	u.SetEntries(u.SSA)
	u.Log.EnterFunction(slotName, slot.Round)

	for _, fname := range names {
		candidateRules := slot.Candidates[fname]
		var lines []string

		u2 := NewUnpacker(fmt.Sprintf("%s_%s", slotName, fname))
		u2.Inherits(u)

		var initQue []string
		for _, l := range candidateRules {
			u2.InspectRule(l)
			init, line := u2.unpackRule(l)
			lines = append(lines, line)
			initQue = append(initQue, InitsToList(init)...)
			u.AddInit(init) // declare candidate-internal vars
			// Do NOT register candidate-internal inits: they are unconstrained when
			// this candidate is not selected, so including them in the registry would
			// let temporal constraints (e.g. "eventually") reference free variables
			// and become trivially satisfiable without the function ever being called.
		}
		ruleSet = append(ruleSet, lines)
		phiClone := u.GetPhis(u.SSA, u2.SSA)
		branches = append(branches, phiClone)
		u.SSA = u2.SSA.Clone()

		queue = append(queue, initQue)
		u.AddInit(u2.Inits)
		u.UpdateRegistry(u2.Registry)
	}

	inits, caps, hasPhi = u.buildPhisOrs(branches, hasPhi)
	var selectors []string
	var branchAssertions []string

	for i, rs := range ruleSet {
		fname := names[i]
		selectorName := fmt.Sprintf("%s_%s", slotName, fname)
		fullSelectorName := fmt.Sprintf("%s_%d", selectorName, slot.Round)
		inits = append(inits, rules.NewInit(selectorName, "Bool", slot.Round, nil, false, false))
		selectors = append(selectors, fullSelectorName)
		selectorRule := u.Log.NewBranchSelector(selectorName, slot.Round, caps[i], queue[i])
		u.Log.AddBranchSelector(selectorRule)
		if rule := selectorRule.WriteRule(); rule != "" {
			ret = append(ret, fmt.Sprintf("(assert %s)", rule))
		}

		for _, r := range rs {
			if r == "" {
				continue
			}
			// r may be a single expression or a multi-line block of (assert ...) statements
			// (e.g. from an Ite). Split into individual assert lines and wrap each in
			// an implication, since (assert (=> sel (assert ...))) is invalid SMT.
			for _, line := range splitAsserts(r) {
				branchAssertions = append(branchAssertions, fmt.Sprintf("(assert (=> %s %s))", fullSelectorName, line))
			}
		}
	}

	u.AddInit(inits)
	u.PopEntries()
	u.Log.ExitFunction(slotName, slot.Round)

	cap := fmt.Sprintf("(assert %s)", strictOr(selectors))
	// Return only the canonical post-phi inits (phi-merge vars + selector vars).
	// Candidate-internal inits are already in u.Inits for declaration but must not
	// be registered into the parent registry — they are unconstrained when their
	// candidate is inactive, which would make temporal constraints trivially satisfiable.
	return inits, fmt.Sprintf("%s\n%s\n%s", strings.Join(branchAssertions, "\n"), strings.Join(ret, "\n"), cap)
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
		u.AddInit(inits)
		SelectorRule := u.Log.NewBranchSelector(u.CurrentBlock, i, caps, InitsToList((inits)))
		if rule := SelectorRule.WriteRule(); rule != "" {
			rule_set = append(rule_set, fmt.Sprintf("(assert %s)", rule))
		}
		u.Log.AddBranchSelector(SelectorRule)
		u2.Inits = append(u2.Inits, rules.NewInit(u.CurrentBlock, "Bool", i, nil, false, false))

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

func (u *Unpacker) unpackIteBlock(blockName string, block []rules.Rule) ([]*rules.Init, []string, []map[string][]int16) {
	var bPhis []map[string][]int16
	var bRules []string
	var bInits []*rules.Init
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
		bInits = append(bInits, inits...)
	}

	PhiClone := u.GetPhis(u.SSA, u2.SSA)
	u.SSA = u2.SSA.Clone()

	bPhis = append(bPhis, PhiClone)
	u.AddInit(u2.Inits)
	u.UpdateRegistry(u2.Registry)
	return bInits, bRules, bPhis
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

	var tPhis, fPhis []map[string][]int16
	var tRules, fRules, aRules []string
	var tEnds, fEnds *scenario.BranchSelector
	var inits, tInits, fInits, aInits []*rules.Init

	if len(ite.T) > 0 {
		tInits, tRules, tPhis = u.unpackIteBlock(ite.BlockNames["true"], ite.T)
	}

	if len(ite.F) > 0 {
		fInits, fRules, fPhis = u.unpackIteBlock(ite.BlockNames["false"], ite.F)
	}

	inits, tEnds, fEnds = u.buildItePhis(tPhis, fPhis, tInits, fInits, ite.BlockNames)

	endRule := fmt.Sprintf("(assert %s)", strictOr([]string{tEnds.Id(), fEnds.Id()}))

	u.AddInit(inits)
	u.Register(inits)

	u.PopEntries()

	if len(ite.After) > 0 {
		aInits, aRules, _ = u.unpackIteBlock(ite.BlockNames["after"], ite.After)
	}
	aRules = append(aRules, endRule)
	inits = append(inits, aInits...)

	// Build the ite assertion that sets block selectors and enforces phis
	var tPhiRules, fPhiRules string
	if len(tEnds.Cond) == 0 {
		tPhiRules = ""
	} else if len(tEnds.Cond) == 1 {
		tPhiRules = tEnds.Cond[0]
	} else {
		tPhiRules = fmt.Sprintf("(and %s)", strings.Join(tEnds.Cond, "\n"))
	}

	if len(fEnds.Cond) == 0 {
		fPhiRules = ""
	} else if len(fEnds.Cond) == 1 {
		fPhiRules = fEnds.Cond[0]
	} else {
		fPhiRules = fmt.Sprintf("(and %s)", strings.Join(fEnds.Cond, "\n"))
	}

	tBranch := fmt.Sprintf("(= %s true) (= %s false)", tEnds.Id(), fEnds.Id())
	if tPhiRules != "" {
		tBranch = fmt.Sprintf("%s %s", tBranch, tPhiRules)
	}
	fBranch := fmt.Sprintf("(= %s false) (= %s true)", tEnds.Id(), fEnds.Id())
	if fPhiRules != "" {
		fBranch = fmt.Sprintf("%s %s", fBranch, fPhiRules)
	}
	ifAssert := fmt.Sprintf("(assert (ite %s (and %s) (and %s)))",
		cond, tBranch, fBranch)
	var resultParts []string
	if t := strings.Join(tRules, "\n"); t != "" {
		resultParts = append(resultParts, t)
	}
	if f := strings.Join(fRules, "\n"); f != "" {
		resultParts = append(resultParts, f)
	}
	resultParts = append(resultParts, ifAssert)
	resultParts = append(resultParts, strings.Join(aRules, "\n"))
	return inits, strings.Join(resultParts, "\n")
}

func (u *Unpacker) FormatRule(r rules.Rule, rule string) string {
	if rule == "" {
		return ""
	}

	trimmed := strings.TrimLeft(rule, "\n")
	if strings.HasPrefix(trimmed, "(assert") { //Already formatted
		return rule
	}

	if _, ok := r.(*rules.Init); ok {
		return rule
	}

	return fmt.Sprintf("(assert %s)", rule)
}

// splitAsserts takes a string that may contain one or more top-level (assert X)
// statements and returns the inner expressions X without the assert wrapper.
// If the input is not assert-wrapped, it is returned as-is in a single-element slice.
func splitAsserts(s string) []string {
	s = strings.TrimSpace(s)
	if !strings.HasPrefix(s, "(assert ") {
		return []string{s}
	}
	var results []string
	for len(s) > 0 {
		s = strings.TrimSpace(s)
		if !strings.HasPrefix(s, "(assert ") {
			break
		}
		// Find the matching closing paren for this (assert ...) by tracking depth
		depth := 0
		end := -1
		for i, ch := range s {
			if ch == '(' {
				depth++
			} else if ch == ')' {
				depth--
				if depth == 0 {
					end = i
					break
				}
			}
		}
		if end == -1 {
			// malformed — return remainder as-is
			results = append(results, s)
			break
		}
		inner := strings.TrimSpace(s[len("(assert ") : end])
		results = append(results, inner)
		s = s[end+1:]
	}
	return results
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

// suppressedRunInitInfix reports whether r is a __run init-store Infix for a
// variable overridden by an assume. When true, the assertion should be dropped
// (the assume provides the value) and only the declaration is emitted.
// Returns the base variable name and true when suppression applies.
func (u *Unpacker) suppressedRunInitInfix(r rules.Rule) (string, bool) {
	if len(u.AssumeOverrides) == 0 {
		return "", false
	}
	inf, ok := r.(*rules.Infix)
	if !ok {
		return "", false
	}
	wx, ok := inf.X.(*rules.Wrap)
	if !ok || !wx.Variable || !wx.Init || !wx.OmitFromOutput {
		return "", false
	}
	if _, exists := u.AssumeOverrides[wx.Value]; exists {
		return wx.Value, true
	}
	return "", false
}
