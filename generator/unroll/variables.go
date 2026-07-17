package unroll

import (
	"fault/generator/rules"
	"fault/llvm"
	"fault/util"
	"fmt"
	"strconv"
	"strings"

	"github.com/llir/llvm/ir/constant"
	irtypes "github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// IsTemp checks if an SSA value string is a temporary variable (starts with "%" followed by a digit).
func IsTemp(id string) bool {
	if string(id[0]) == "%" && IsNumeric(string(id[1])) {
		return true
	}
	return false
}

// IsGlobal checks if an SSA value string is a global variable (starts with "@").
func IsGlobal(id string) bool {
	return string(id[0]) == "@"
}

// IsInt checks if a string parses as an integer.
func IsInt(char string) bool {
	if _, err := strconv.Atoi(char); err == nil {
		return true
	}
	return false
}

// IsNumeric checks if a string parses as a float or integer.
func IsNumeric(char string) bool {
	if _, err := strconv.ParseFloat(char, 64); err == nil {
		return true
	}
	if _, err := strconv.Atoi(char); err == nil {
		return true
	}
	return false
}

// IsBoolean checks if an SSA value string is a boolean literal.
func IsBoolean(id string) bool {
	if id == "true" || id == "false" {
		return true
	}
	return false
}

// IsClocked checks if an SSA value string represents a clocked (time-stepped) variable, identified by containing "(".
func IsClocked(id string) bool {
	return strings.Contains(id, "(")
}

// IsIndexed checks if an SSA value string references a specific version of a variable (e.g. example.var[2] — the value after its second state change).
func IsIndexed(id string) bool {
	rawid := strings.Split(id, "_")
	_, err := strconv.Atoi(rawid[len(rawid)-1])
	if e, ok := err.(*strconv.NumError); ok && e.Err == strconv.ErrSyntax {
		return false
	}
	return err == nil
}

// GetClockBase strips the clock index suffix from a clocked variable id, returning the base variable name.
func GetClockBase(id string) string {
	v := strings.Split(id, "_")
	v[0] = v[0][1:]
	return strings.Join(v[0:len(v)-1], "_")
}

// IsStaticValue checks if an SSA value string is a literal (boolean or numeric constant).
func IsStaticValue(id string) bool {
	if IsBoolean(id) || IsNumeric(id) {
		return true
	}
	return false
}

// LookupType infers the SMT type ("Bool" or "Real") from an LLVM IR value. Panics if the type cannot be determined.
func LookupType(id string, value value.Value) string {

	if _, ok := value.(*constant.ExprAnd); ok {
		return "Bool"
	}

	if _, ok := value.(*constant.ExprOr); ok {
		return "Bool"
	}

	if _, ok := value.(*constant.ExprFNeg); ok {
		return "Bool"
	}

	if value.Type().Equal(llvm.DoubleP) {
		return "Real"
	}
	if value.Type().Equal(llvm.I1P) {
		return "Bool"
	}

	switch value.Type().(type) {
	case *irtypes.FloatType:
		return "Real"
	case *irtypes.IntType: // LLVM doesn't have a bool type
		return "Bool" // since all Fault numbers are floats,
	// ints are probably bools
	case *irtypes.ArrayType:
		return "Bool"
	}

	panic(fmt.Sprintf("smt generation error, value for %s not found", id))
}

// isASolvable checks if id is an unknown or uncertain variable (a free SMT variable to be solved).
func isASolvable(id string, RawInputs *llvm.RawInputs) bool {
	for _, v := range RawInputs.Unknowns {
		if v == id {
			return true
		}
	}
	for k := range RawInputs.Uncertains {
		if k == id {
			return true
		}
	}
	return false
}

// isAWhole checks if id is a whole-number constrained variable.
func isAWhole(id string, RawInputs *llvm.RawInputs) bool {
	for _, v := range RawInputs.Wholes {
		if v == id {
			return true
		}
	}
	return false
}

// isAParam checks if id is a parameter variable.
func isAParam(id string, RawInputs *llvm.RawInputs) bool {
	for _, v := range RawInputs.Params {
		if v == id {
			return true
		}
	}
	return false
}

// FormatValue extracts the value portion from an LLVM IR value string (e.g. "float 1.5" → "1.5").
func FormatValue(val value.Value) string {
	v := strings.Split(val.String(), " ")
	return v[1]
}

// ConvertIdent resolves an SSA value string to its base variable name, following temp variable references through VarLoads.
func (b *LLBlock) ConvertIdent(f string, val string) string {
	if IsTemp(val) {
		refname := fmt.Sprintf("%s-%s", f, val)
		if v, ok := b.Env.VarLoads[refname]; ok {
			id := util.FormatIdent(v.Ident())
			return id
		} else {
			panic(fmt.Sprintf("variable %s not initialized", val))
		}
	} else {
		id := val
		if string(id[0]) == "%" || IsGlobal(id) {
			id = util.FormatIdent(id)
			return id
		}
		return id //Is a value, not an identifier
	}
}

// LookupCondPart returns the Rule associated with a temp variable in the current block's IR refs, used when resolving conditional expressions.
func (b *LLBlock) LookupCondPart(f string, val string) rules.Rule {
	if IsTemp(val) {
		refname := fmt.Sprintf("%s-%s", f, val)
		if v, ok := b.irRefs[refname]; ok {
			return v
		}
	}
	return nil
}
