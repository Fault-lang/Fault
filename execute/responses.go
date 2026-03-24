package execute

import (
	"fault/execute/parser"
	"fmt"
	"strconv"
	"strings"
)

type SMTListener struct {
	*parser.BaseSMTLIBv2Listener
	stack   []interface{}
	Results map[string]Scenario
	Values  map[string]string
	err     error // first error encountered during the walk; halts further processing
}

func NewSMTListener() *SMTListener {
	return &SMTListener{
		Results: make(map[string]Scenario),
		Values:  make(map[string]string),
	}
}

func (l *SMTListener) push(n interface{}) {
	l.stack = append(l.stack, n)
}

func (l *SMTListener) pop() interface{} {
	var s interface{}
	if len(l.stack) > 0 {
		s, l.stack = l.stack[len(l.stack)-1], l.stack[:len(l.stack)-1]
		return s
	}
	return nil
}

func (l *SMTListener) peek() interface{} {
	if len(l.stack) > 0 {
		return l.stack[len(l.stack)-1]
	}
	return nil
}

// popString pops from the stack and asserts the value is a string.
// On failure it records l.err and returns ("", false).
func (l *SMTListener) popString(ctx string) (string, bool) {
	v := l.pop()
	s, ok := v.(string)
	if !ok {
		l.err = fmt.Errorf("SMT parse error in %s: expected string, got %T", ctx, v)
		return "", false
	}
	return s, true
}

func mergeTermParts(parts []string) (string, error) {
	if len(parts) == 1 {
		return parts[0], nil
	}

	if len(parts) > 2 {
		return "", fmt.Errorf("SMT parse error: too many term parts (%d), expected 1 or 2", len(parts))
	}

	value1, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return strings.Join(parts, ""), nil // a negative value represented as "-" + digits
	}

	value2, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return "", fmt.Errorf("SMT parse error: second term part %q is not a number", parts[1])
	}

	return fmt.Sprintf("%f", value1/value2), nil
}

func (l *SMTListener) ExitGet_model_response(c *parser.Get_model_responseContext) {

}

func (l *SMTListener) ExitModel_response(c *parser.Model_responseContext) {

}

func (l *SMTListener) ExitFunction_def(c *parser.Function_defContext) {
	if l.err != nil {
		return
	}

	termVal, ok := l.popString("function_def/term")
	if !ok {
		return
	}
	sortVal, ok := l.popString("function_def/sort")
	if !ok {
		return
	}
	symVal, ok := l.popString("function_def/sym")
	if !ok {
		return
	}

	t := termVal
	if len(t) > 0 && t[0] == '(' {
		t = t[1 : len(t)-2]
	}

	l.Values[symVal] = t

	value, err := convertTerm(sortVal, t)
	if err != nil {
		l.err = err
		return
	}

	key, id := splitIdent(symVal)
	i, err := strconv.ParseInt(key, 10, 16)
	if err != nil {
		l.err = fmt.Errorf("symbol returned from model is malformed: %s", symVal)
		return
	}
	k := int16(i)

	switch v := value.(type) {
	case float64:
		if l.Results[id] != nil {
			l.Results[id].(*FloatTrace).Add(k, v)
		} else {
			l.Results[id] = NewFloatTrace()
			l.Results[id].(*FloatTrace).Add(k, v)
		}
	case bool:
		if l.Results[id] != nil {
			l.Results[id].(*BoolTrace).Add(k, v)
		} else {
			l.Results[id] = NewBoolTrace()
			l.Results[id].(*BoolTrace).Add(k, v)
		}
	case int64:
		if l.Results[id] != nil {
			l.Results[id].(*IntTrace).Add(k, v)
		} else {
			l.Results[id] = NewIntTrace()
			l.Results[id].(*IntTrace).Add(k, v)
		}
	}
}

func (l *SMTListener) ExitVariable(c *parser.VariableContext) {
	l.push(c.GetText())
}

func (l *SMTListener) ExitTerm(c *parser.TermContext) {
	if l.err != nil {
		return
	}

	term := c.GetText()

	if c.GetChildCount() > 1 {
		parts := []string{}
		for range c.AllTerm() {
			p, ok := l.popString("term/part")
			if !ok {
				return
			}
			parts = append([]string{p}, parts...)
		}
		merge, err := mergeTermParts(parts)
		if err != nil {
			l.err = err
			return
		}
		if strings.Contains(term, "-") {
			term = fmt.Sprintf("-%s", merge)
		} else {
			term = merge
		}
	}
	l.push(term)
}

func (l *SMTListener) ExitSort(c *parser.SortContext) {
	l.push(c.GetText())
}

func convertTerm(sort string, term string) (interface{}, error) {
	switch sort {
	case "Real":
		v, err := strconv.ParseFloat(term, 64)
		if err != nil {
			return nil, fmt.Errorf("SMT parse error: Real value %q is not a valid float: %w", term, err)
		}
		return v, nil
	case "Bool":
		switch term {
		case "true":
			return true, nil
		case "false":
			return false, nil
		default:
			return nil, fmt.Errorf("SMT parse error: Bool value %q is not 'true' or 'false'", term)
		}
	case "Int":
		v, err := strconv.ParseInt(term, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("SMT parse error: Int value %q is not a valid integer: %w", term, err)
		}
		return v, nil
	default:
		return term, nil
	}
}

func splitIdent(ident string) (string, string) {
	s := strings.Split(ident, "_")
	return s[len(s)-1], strings.Join(s[0:len(s)-1], "_")
}
