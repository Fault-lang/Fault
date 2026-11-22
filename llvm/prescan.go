package llvm

import "fault/ast"

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
		// Instances are OR'd alternatives (same property across spec instantiations).
		// If any instance is Bool, the expression is Bool.
		for _, inst := range e.Instances {
			if bools[inst] {
				return true
			}
		}
		return false
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

// findAssertVarInstances collects all AssertVar instance names referenced in
// the given assume/assert statements. These must not be pruned from the SMT
// model even if no flow function directly loads them.
func findAssertVarInstances(stmts []*ast.AssertionStatement) map[string]bool {
	result := make(map[string]bool)
	for _, a := range stmts {
		collectAssertVarInstances(a.Constraint.Left, result)
		collectAssertVarInstances(a.Constraint.Right, result)
	}
	return result
}

// findAnnotationEvidenceInstances collects all AssertVar instance names used
// in annotation evidence conditions. These must not be pruned from the SMT
// model even if no flow function directly loads the corresponding variable.
func findAnnotationEvidenceInstances(annotations []*ast.AnnotationStatement) map[string]bool {
	result := make(map[string]bool)
	for _, ann := range annotations {
		for _, clause := range ann.TrueWhen {
			collectAssertVarInstances(clause.Condition, result)
		}
		for _, clause := range ann.FalseWhen {
			collectAssertVarInstances(clause.Condition, result)
		}
	}
	return result
}

func collectAssertVarInstances(expr ast.Expression, out map[string]bool) {
	if expr == nil {
		return
	}
	switch e := expr.(type) {
	case *ast.AssertVar:
		for _, inst := range e.Instances {
			out[inst] = true
		}
	case *ast.InfixExpression:
		collectAssertVarInstances(e.Left, out)
		collectAssertVarInstances(e.Right, out)
	case *ast.PrefixExpression:
		collectAssertVarInstances(e.Right, out)
	}
}
