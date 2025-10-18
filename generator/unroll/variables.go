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

func IsTemp(id string) bool {
	if string(id[0]) == "%" && IsNumeric(string(id[1])) {
		return true
	}
	return false
}

func IsGlobal(id string) bool {
	return string(id[0]) == "@"
}

func IsInt(char string) bool {
	if _, err := strconv.Atoi(char); err == nil {
		return true
	}
	return false
}

func IsNumeric(char string) bool {
	if _, err := strconv.ParseFloat(char, 64); err == nil {
		return true
	}
	if _, err := strconv.Atoi(char); err == nil {
		return true
	}
	return false
}

func IsBoolean(id string) bool {
	if id == "true" || id == "false" {
		return true
	}
	return false
}

func IsClocked(id string) bool {
	if strings.Contains(id, "(") {
		return true
	}
	return false
}

func IsIndexed(id string) bool {
	rawid := strings.Split(id, "_")
	_, err := strconv.Atoi(rawid[len(rawid)-1])
	if err != nil {
		return false
	}
	return true
}

func GetClockBase(id string) string {
	v := strings.Split(id, "_")
	v[0] = v[0][1:]
	return strings.Join(v[0:len(v)-1], "_")
}

func IsStaticValue(id string) bool {
	if IsBoolean(id) || IsNumeric(id) {
		return true
	}
	return false
}

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

func FormatValue(val value.Value) string {
	v := strings.Split(val.String(), " ")
	return v[1]
}

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

func (b *LLBlock) LookupCondPart(f string, val string) rules.Rule {
	if IsTemp(val) {
		refname := fmt.Sprintf("%s-%s", f, val)
		if v, ok := b.irRefs[refname]; ok {
			return v
		}
	}
	return nil
}
