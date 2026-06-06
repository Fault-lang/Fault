package generator

import (
	"fault/ast"
	"fault/generator/asserts"
	"fault/generator/rules"
	"fault/generator/scenario"
	"fault/generator/unpack"
	"fault/generator/unroll"
	"fault/llvm"
	"fault/util"
	"fmt"
	"os"
	"strings"

	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir"
)

const (
	DefaultSMTTimeout       = 30_000 // milliseconds
	DefaultSMTMemoryMaxSize = 1_096  // MB
)

// GeneratorOptions controls optional SMT-LIB2 solver directives that are
// prepended to the generated formula.
type GeneratorOptions struct {
	Timeout       int // (set-option :timeout N) in milliseconds; 0 = omit
	MemoryMaxSize int // (set-option :memory_max_size N) in MB; 0 = omit
}

// Take LL IR and Generate SMTLib2
//
// Step 1: Translate LL IR into Rules by LLBlock
// and Function
//
// Step 2: Add Phi values necessary for the SMT model
// and flatten to a single set of rules

type Generator struct {
	constants   []rules.Rule
	functions   map[string]*ir.Func
	Env         *unroll.Env
	RawInputs   *llvm.RawInputs
	RunBlock    *unroll.LLFunc
	smt         []string
	ResultLog   *scenario.Logger
	StringRules map[string]string
	IsCompound  map[string]bool
}

func NewGenerator(ri *llvm.RawInputs, sr map[string]string, is map[string]bool, opts GeneratorOptions) *Generator {
	// Optimization pass: if every free numeric variable is whole(), declare them as
	// Int-sorted and use QF_NIA (standard, cross-compatible). This avoids is_int
	// assertions and the QF_NRA + is_int combination that is Z3-specific.
	// Note: whole vars are also tracked in Unknowns (they are free variables), so
	// the condition is simply: no non-whole free variables present.
	if len(ri.Wholes) > 0 && len(ri.Uncertains) == 0 {
		// Check that every entry in Unknowns is also in Wholes (i.e. all unknowns are whole)
		wholeSet := make(map[string]bool, len(ri.Wholes))
		for _, w := range ri.Wholes {
			wholeSet[w] = true
		}
		allWhole := true
		for _, u := range ri.Unknowns {
			if !wholeSet[u] {
				allWhole = false
				break
			}
		}
		if allWhole {
			ri.IntegerMode = true
		}
	}

	logic := "QF_NRA"
	if ri.IntegerMode {
		logic = "QF_NIA"
	}

	var preamble []string
	if opts.Timeout > 0 {
		preamble = append(preamble, fmt.Sprintf("(set-option :timeout %d)", opts.Timeout))
	}
	if opts.MemoryMaxSize > 0 {
		preamble = append(preamble, fmt.Sprintf("(set-option :memory_max_size %d)", opts.MemoryMaxSize))
	}
	preamble = append(preamble, fmt.Sprintf("(set-logic %s)", logic))
	return &Generator{
		functions:   make(map[string]*ir.Func),
		Env:         unroll.NewEnv(ri),
		smt:         preamble,
		RawInputs:   ri,
		StringRules: sr,
		IsCompound:  is,
	}
}

func Execute(compiler *llvm.Compiler, opts GeneratorOptions) *Generator {
	generator := NewGenerator(compiler.RawInputs, compiler.StringRules, compiler.IsCompound, opts)
	generator.Run(compiler.GetOptimizedIR())
	return generator
}

func (g *Generator) AppendSMT(new_smt []string) {
	for _, s := range new_smt {
		if s != "" {
			g.smt = append(g.smt, s)
		}
	}
}

func (g *Generator) Run(llopt string) {
	os.WriteFile("/tmp/fault_debug.ll", []byte(llopt), 0644)
	m, err := asm.ParseString("", llopt) //"" because ParseString has a path variable
	if err != nil {
		panic(err)
	}
	g.newCallgraph(m)

}

func (g *Generator) newCallgraph(m *ir.Module) {
	g.Env.MutableVars = unroll.FindMutableVars(m.Funcs)
	g.Env.UsedVars = unroll.FindUsedVars(m.Funcs)
	g.Env.WriteSets = unroll.FindWriteSets(m.Funcs)
	g.Env.StringRules = g.StringRules
	g.constants = unroll.NewConstants(g.Env, m.Globals, g.RawInputs)
	g.sortFuncs(m.Funcs)

	g.Env.WhensThens = unroll.WhenThen(append(g.RawInputs.Asserts, g.RawInputs.Assumes...))

	g.RunBlock = unroll.NewLLFunc(g.Env, g.functions, g.functions["__run"])
	g.RunBlock.Unroll()

	p := unpack.NewUnpacker(g.RunBlock.Ident)
	p.LoadStringRules(g.StringRules, g.IsCompound)
	p.Log.Uncertains = g.RawInputs.Uncertains

	p.VarTypes = g.Env.VarTypes
	smt := p.Unpack(g.constants, g.RunBlock)

	// Upgrade Real-typed Inits to Bool for variables used in boolean contexts
	// in assume/assert statements (e.g. unknown() fields used with ||/&&/!).
	allStmts := append(g.RawInputs.Asserts, g.RawInputs.Assumes...)
	boolVars := inferBoolVarNames(allStmts)
	for _, init := range p.Inits {
		if boolVars[init.Ident] && init.Type == "Real" {
			init.Type = "Bool"
			g.Env.VarTypes[init.Ident] = "Bool"
		}
	}

	g.AppendSMT(p.InitVars())
	g.AppendSMT(smt)

	g.ResultLog = p.Log

	// ProcessUnfuncs first: it declares _available_N shadow variables that
	// assume/assert statements (e.g. "assume x available") may reference.
	unfuncSMT := g.ProcessUnfuncs(g.RawInputs.Unfuncs, g.Env.CurrentRound, p.Registry)
	g.AppendSMT(unfuncSMT)

	assertSMT := g.ProcessAsserts(g.RawInputs.Asserts, g.Env.CurrentRound, p.Registry, p.Whens)
	g.AppendSMT(assertSMT)
	assumeSMT := g.ProcessAsserts(g.RawInputs.Assumes, g.Env.CurrentRound, p.Registry, p.Whens)
	g.AppendSMT(assumeSMT)
}

func (g *Generator) ProcessAsserts(assertList []*ast.AssertionStatement, rounds int, registry map[string][][]string, whens map[string][]map[string]string) []string {
	var rules []string

	for _, as := range assertList {
		// "assume/assert x available" pins the initial availability of a field.
		// The _available_0 shadow variable is declared by ProcessUnfuncs (which
		// runs first), so we can safely assert against it here.
		if as.Temporal == "available" {
			rules = append(rules, availabilityAssertions(as)...)
			continue
		}
		if !asserts.IsRelevant(g.Env.VarTypes, as.Constraint) { //If the assert is on a variable that is not used, drop the assert
			continue
		}
		c, err := asserts.NewConstraint(as, rounds, registry, whens, g.Env.VarTypes)
		if err != nil {
			panic(err.Error())
		}
		rules = append(rules, c.Parse()...)
	}
	return rules
}

// availabilityAssertions generates SMT for "assume/assert x available".
// For each instance in the constraint's AssertVar, it emits an assertion that
// the field's _available_0 shadow variable is true (assume) or false (assert).
func availabilityAssertions(as *ast.AssertionStatement) []string {
	assertVar, ok := as.Constraint.Left.(*ast.AssertVar)
	if !ok {
		return nil
	}
	val := "true"
	if !as.Assume {
		val = "false"
	}
	var rules []string
	for _, inst := range assertVar.Instances {
		rules = append(rules, fmt.Sprintf("(assert (= %s_available_0 %s))", inst, val))
	}
	return rules
}

func (g *Generator) sortFuncs(funcs []*ir.Func) {
	//Iterate through all the function blocks and store them by
	// function call name.
	for _, f := range funcs {
		// Get function name.
		g.functions[util.FormatIdent(f.Ident())] = f
		continue
	}
}

// ProcessUnfuncs generates SMT constraints for unfunc states.
//
// Each field referenced in any unfunc's requires or emits gets an auto-generated
// shadow Bool variable named <field>_available_N. The numeric field itself is
// left free for the solver; only the shadow is constrained here.
//
// For each unfunc, it emits:
//   - shadow Bool declarations for every unique field across all unfuncs
//   - a declare-fun for the uninterpreted function
//   - per-step activation variables
//   - activation guards (active => requires_available_N)
//   - write effects on the shadow (active => field_available_N+1 = true)
//   - frame conditions on the shadow (not active => field_available_N+1 = field_available_N)
func (g *Generator) ProcessUnfuncs(unfuncs []*llvm.UnfuncInfo, rounds int, registry map[string][][]string) []string {
	if len(unfuncs) == 0 {
		return nil
	}
	// Bound the trace to at least the number of unfunc states so the solver
	// has enough steps to find a plan even when the run block has no explicit
	// rounds.
	steps := rounds
	if steps < len(unfuncs) {
		steps = len(unfuncs)
	}

	var smt []string

	// Declare _available shadow Bool variables for every unique field across all
	// unfuncs. Steps 0..steps inclusive because write effects reference n+1 where
	// n goes up to steps-1.
	for _, fieldBase := range collectAllUnfuncFields(unfuncs) {
		for n := 0; n <= steps; n++ {
			smt = append(smt, fmt.Sprintf("(declare-fun %s_available_%d () Bool)", fieldBase, n))
		}
	}

	varTypes := g.Env.VarTypes

	// For each unique LHS field in assume clauses, declare new versioned variables
	// for steps 1..steps. These represent the output field's value immediately after
	// the unfunc fires. Step 0 (the initial value) already exists in the main formula.
	assumeLHSTypes := collectAssumeLHSTypes(unfuncs, varTypes)
	for lhsBase, lhsType := range assumeLHSTypes {
		for n := 1; n <= steps; n++ {
			smt = append(smt, fmt.Sprintf("(declare-fun %s_%d () %s)", lhsBase, n, lhsType))
		}
	}

	for _, uf := range unfuncs {
		stateId := util.FormatIdent(uf.StateKey)
		// The registry uses the bare state variable name (without __state suffix).
		stateVarBase := strings.TrimSuffix(stateId, "__state")

		for n := 0; n < steps; n++ {
			// The activation variable IS the state's Bool variable at round n+1
			// (declared and constrained by the run block in the main formula).
			// Using it directly connects unfunc constraints to the run block
			// rather than inventing a free variable the solver can ignore.
			activeVar := registryBestVersion(registry, stateVarBase, n+1)

			// Activation guard: if active, required fields must be available
			if uf.Requires != nil {
				reqSMT := unfuncExprToSMTAvail(uf.Requires, n)
				smt = append(smt, fmt.Sprintf("(assert (=> %s %s))", activeVar, reqSMT))
			}

			// Write effects and frame conditions on _available shadow variables
			if uf.Emits != nil {
				for _, fieldBase := range collectUnfuncFields(uf.Emits) {
					next := fmt.Sprintf("%s_available_%d", fieldBase, n+1)
					curr := fmt.Sprintf("%s_available_%d", fieldBase, n)
					smt = append(smt, fmt.Sprintf("(assert (=> %s (= %s true)))", activeVar, next))
					smt = append(smt, fmt.Sprintf("(assert (=> (not %s) (= %s %s)))", activeVar, next, curr))
				}
			}

			// Assume constraints: postcondition arithmetic scoped to this exact firing.
			//
			// LHS is the output field at n+1 — the value it takes on after the unfunc
			// fires. RHS uses the registry to resolve the actual SSA names of input
			// fields at round n, so constraints reference the variables that already
			// exist in the main formula rather than hypothetical step-indexed names.
			//
			// Frame condition: when the unfunc does NOT fire at step n, the output
			// field is unchanged (carries its previous value forward).
			for _, assume := range uf.Assumes {
				infix, ok := assume.(*ast.InfixExpression)
				if !ok {
					continue
				}
				lhsPC, ok := infix.Left.(*ast.ParameterCall)
				if !ok {
					continue
				}
				lhsBase := resolveVarBase(lhsPC, varTypes)
				lhsNext := fmt.Sprintf("%s_%d", lhsBase, n+1)

				// "curr" for the frame condition:
				//   - at step 0, use the version declared in the main formula
				//   - at step n>0, use the version we declared at the prior step
				var lhsCurr string
				if n == 0 {
					lhsCurr = registryBestVersion(registry, lhsBase, 0)
				} else {
					lhsCurr = fmt.Sprintf("%s_%d", lhsBase, n)
				}

				rhsSMT := unfuncArithExprToSMT(infix.Right, n, registry, varTypes)
				smt = append(smt, fmt.Sprintf("(assert (=> %s (= %s %s)))", activeVar, lhsNext, rhsSMT))
				smt = append(smt, fmt.Sprintf("(assert (=> (not %s) (= %s %s)))", activeVar, lhsNext, lhsCurr))
			}
		}
	}
	return smt
}

// unfuncVarBase returns the SMT base variable name for a ParameterCall,
// using ProcessedName when available, falling back to Spec+Value.
func unfuncVarBase(pc *ast.ParameterCall) string {
	if len(pc.ProcessedName) > 0 {
		return strings.Join(pc.ProcessedName, "_")
	}
	parts := append([]string{pc.Spec}, pc.Value...)
	return strings.Join(parts, "_")
}

// resolveVarBase finds the actual SMT base name for a ParameterCall by
// checking VarTypes. The LLVM compiler sometimes drops intermediate path
// components (e.g. a component property calc.a in spec test becomes test_a,
// not test_calc_a). We try the full path first, then progressively shorter
// suffixes until we find a key present in VarTypes.
func resolveVarBase(pc *ast.ParameterCall, varTypes map[string]string) string {
	// Full path: spec + all value components
	full := unfuncVarBase(pc)
	if _, ok := varTypes[full]; ok {
		return full
	}
	// Try spec + last value component only (common for component properties)
	if len(pc.Value) > 1 {
		short := strings.Join([]string{pc.Spec, pc.Value[len(pc.Value)-1]}, "_")
		if _, ok := varTypes[short]; ok {
			return short
		}
	}
	return full
}

// registryBestVersion finds the full SSA variable name for baseName at the most
// recent round at or before maxRound. Registry keys have the form "round-N_blockId".
// Falls back to baseName_0 if nothing is found.
func registryBestVersion(registry map[string][][]string, baseName string, maxRound int) string {
	best := ""
	bestRound := -1
	for key, vars := range registry {
		var round int
		if _, err := fmt.Sscanf(key, "round-%d_", &round); err != nil {
			continue
		}
		if round > maxRound {
			continue
		}
		for _, varSSA := range vars {
			if varSSA[0] == baseName && round >= bestRound {
				bestRound = round
				best = strings.Join(varSSA, "_")
			}
		}
	}
	if best != "" {
		return best
	}
	return fmt.Sprintf("%s_0", baseName)
}

// collectAssumeLHSTypes returns a map from resolved base variable name to SMT
// type for every unique LHS ParameterCall across all unfunc assume clauses.
// Used to declare the new versioned output variables in ProcessUnfuncs.
func collectAssumeLHSTypes(unfuncs []*llvm.UnfuncInfo, varTypes map[string]string) map[string]string {
	result := make(map[string]string)
	for _, uf := range unfuncs {
		for _, assume := range uf.Assumes {
			infix, ok := assume.(*ast.InfixExpression)
			if !ok {
				continue
			}
			lhsPC, ok := infix.Left.(*ast.ParameterCall)
			if !ok {
				continue
			}
			base := resolveVarBase(lhsPC, varTypes)
			if _, seen := result[base]; seen {
				continue
			}
			ty := varTypes[base]
			if ty == "" {
				ty = "Real"
			}
			result[base] = ty
		}
	}
	return result
}

// unfuncExprToSMT converts an unfunc requires/emits expression to an SMT
// string, versioning ParameterCall leaves with the given step index.
func unfuncExprToSMT(expr ast.Expression, step int) string {
	switch e := expr.(type) {
	case *ast.ParameterCall:
		return fmt.Sprintf("%s_%d", unfuncVarBase(e), step)
	case *ast.InfixExpression:
		left := unfuncExprToSMT(e.Left, step)
		right := unfuncExprToSMT(e.Right, step)
		switch e.Operator {
		case "&&":
			return fmt.Sprintf("(and %s %s)", left, right)
		case "||":
			return fmt.Sprintf("(or %s %s)", left, right)
		default:
			return fmt.Sprintf("(%s %s %s)", e.Operator, left, right)
		}
	case *ast.PrefixExpression:
		inner := unfuncExprToSMT(e.Right, step)
		return fmt.Sprintf("(not %s)", inner)
	default:
		return ""
	}
}

// unfuncExprToSMTAvail converts an unfunc expression to an SMT string
// referencing the _available shadow Bool variable (e.g. field_available_N).
// Used for requires guards so activation depends on availability, not value.
func unfuncExprToSMTAvail(expr ast.Expression, step int) string {
	switch e := expr.(type) {
	case *ast.ParameterCall:
		return fmt.Sprintf("%s_available_%d", unfuncVarBase(e), step)
	case *ast.InfixExpression:
		left := unfuncExprToSMTAvail(e.Left, step)
		right := unfuncExprToSMTAvail(e.Right, step)
		switch e.Operator {
		case "&&":
			return fmt.Sprintf("(and %s %s)", left, right)
		case "||":
			return fmt.Sprintf("(or %s %s)", left, right)
		default:
			return fmt.Sprintf("(%s %s %s)", e.Operator, left, right)
		}
	case *ast.PrefixExpression:
		inner := unfuncExprToSMTAvail(e.Right, step)
		return fmt.Sprintf("(not %s)", inner)
	default:
		return ""
	}
}

// unfuncArithExprToSMT converts an unfunc arithmetic expression (from an assume
// clause RHS) to an SMT string. ParameterCall leaves are resolved to their actual
// SSA names at round n using the registry and varTypes, so the constraint
// references variables that genuinely exist in the main formula.
func unfuncArithExprToSMT(expr ast.Expression, round int, registry map[string][][]string, varTypes map[string]string) string {
	switch e := expr.(type) {
	case *ast.ParameterCall:
		base := resolveVarBase(e, varTypes)
		return registryBestVersion(registry, base, round)
	case *ast.IntegerLiteral:
		return fmt.Sprintf("%d", e.Value)
	case *ast.FloatLiteral:
		return fmt.Sprintf("%g", e.Value)
	case *ast.InfixExpression:
		left := unfuncArithExprToSMT(e.Left, round, registry, varTypes)
		right := unfuncArithExprToSMT(e.Right, round, registry, varTypes)
		switch e.Operator {
		case "+":
			return fmt.Sprintf("(+ %s %s)", left, right)
		case "-":
			return fmt.Sprintf("(- %s %s)", left, right)
		case "*":
			return fmt.Sprintf("(* %s %s)", left, right)
		case "/":
			return fmt.Sprintf("(/ %s %s)", left, right)
		default:
			return fmt.Sprintf("(%s %s %s)", e.Operator, left, right)
		}
	default:
		return ""
	}
}

// collectUnfuncFields returns the base SMT variable names (without version
// suffix) for all ParameterCall leaves in an emits expression.
func collectUnfuncFields(expr ast.Expression) []string {
	switch e := expr.(type) {
	case *ast.ParameterCall:
		return []string{unfuncVarBase(e)}
	case *ast.InfixExpression:
		left := collectUnfuncFields(e.Left)
		right := collectUnfuncFields(e.Right)
		return append(left, right...)
	case *ast.PrefixExpression:
		return collectUnfuncFields(e.Right)
	default:
		return nil
	}
}

// collectAllUnfuncFields collects all unique field base names from the Requires
// and Emits expressions of every UnfuncInfo. Deduplication ensures each shadow
// variable is declared exactly once even if the same field appears in multiple
// unfuncs.
func collectAllUnfuncFields(unfuncs []*llvm.UnfuncInfo) []string {
	seen := make(map[string]bool)
	var result []string
	for _, uf := range unfuncs {
		for _, expr := range []ast.Expression{uf.Requires, uf.Emits} {
			if expr == nil {
				continue
			}
			for _, fieldBase := range collectUnfuncFields(expr) {
				if !seen[fieldBase] {
					seen[fieldBase] = true
					result = append(result, fieldBase)
				}
			}
		}
	}
	return result
}

func (g *Generator) SMT() string {
	return strings.Join(g.smt, "\n")
}

// ParamManifest returns a map of __PARAM_name__ token → SMT sort (Real, Int, Bool)
// for every param() field declared in the spec. Used alongside SMT() in template mode
// so the caller knows which tokens exist and their expected types.
func (g *Generator) ParamManifest() map[string]string {
	manifest := make(map[string]string)
	for _, name := range g.RawInputs.Params {
		sort := "Real"
		// Prefer the TypeHint-derived type stored at compile time.
		if ty, ok := g.RawInputs.ParamTypes[name]; ok && ty != "" {
			sort = ty
		} else if ty, ok := g.Env.VarTypes[name]; ok {
			switch ty {
			case "i32", "i64", "Int":
				sort = "Int"
			case "i1", "Bool":
				sort = "Bool"
			}
		}
		manifest[name] = sort
	}
	return manifest
}

// inferBoolVarNames scans assume/assert constraints and returns a set of
// variable base names that should be declared as Bool in SMT. This covers:
//   - variables used directly as operands of ||, &&, or !
//   - variables in when/then positions (always Bool)
//   - variables compared to boolean expressions via == or !=
func inferBoolVarNames(stmts []*ast.AssertionStatement) map[string]bool {
	bools := make(map[string]bool)

	// Pass 1: mark direct operands of ||, &&, ! and both sides of when/then
	for _, a := range stmts {
		op := a.Constraint.Operator
		if op == "then" {
			// Both the when-condition and the then-consequent are Bool
			collectBoolOperands(a.Constraint.Left, true, bools)
			collectBoolOperands(a.Constraint.Right, true, bools)
		} else {
			collectBoolOperands(a.Constraint.Left, false, bools)
			collectBoolOperands(a.Constraint.Right, false, bools)
		}
	}

	// Pass 2: propagate Bool through == / != comparisons (fixed-point)
	changed := true
	for changed {
		changed = false
		for _, a := range stmts {
			op := a.Constraint.Operator
			if op != "==" && op != "!=" {
				continue
			}
			if exprIsBool(a.Constraint.Right, bools) {
				if markBoolVars(a.Constraint.Left, bools) {
					changed = true
				}
			}
			if exprIsBool(a.Constraint.Left, bools) {
				if markBoolVars(a.Constraint.Right, bools) {
					changed = true
				}
			}
		}
	}

	return bools
}

func collectBoolOperands(expr ast.Expression, inBoolCtx bool, bools map[string]bool) {
	switch e := expr.(type) {
	case *ast.InfixExpression:
		nextBool := e.Operator == "||" || e.Operator == "&&"
		collectBoolOperands(e.Left, nextBool, bools)
		collectBoolOperands(e.Right, nextBool, bools)
	case *ast.PrefixExpression:
		collectBoolOperands(e.Right, e.Operator == "!", bools)
	case *ast.AssertVar:
		if inBoolCtx {
			for _, inst := range e.Instances {
				bools[inst] = true
			}
		}
	}
}

// exprIsBool reports whether expr is a boolean expression given the current
// set of known-bool variables.
func exprIsBool(expr ast.Expression, bools map[string]bool) bool {
	switch e := expr.(type) {
	case *ast.InfixExpression:
		if e.Operator == "||" || e.Operator == "&&" {
			return true
		}
		return exprIsBool(e.Left, bools) && exprIsBool(e.Right, bools)
	case *ast.PrefixExpression:
		if e.Operator == "!" {
			return true
		}
		return exprIsBool(e.Right, bools)
	case *ast.AssertVar:
		for _, inst := range e.Instances {
			if !bools[inst] {
				return false
			}
		}
		return len(e.Instances) > 0
	case *ast.Boolean:
		return true
	}
	return false
}

// markBoolVars marks all AssertVar instances in expr as Bool. Returns true if
// any new variables were added.
func markBoolVars(expr ast.Expression, bools map[string]bool) bool {
	changed := false
	switch e := expr.(type) {
	case *ast.AssertVar:
		for _, inst := range e.Instances {
			if !bools[inst] {
				bools[inst] = true
				changed = true
			}
		}
	case *ast.InfixExpression:
		if markBoolVars(e.Left, bools) {
			changed = true
		}
		if markBoolVars(e.Right, bools) {
			changed = true
		}
	case *ast.PrefixExpression:
		if markBoolVars(e.Right, bools) {
			changed = true
		}
	}
	return changed
}
