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
}

func NewSMTListener() *SMTListener {
	return &SMTListener{
		Results: make(map[string]Scenario),
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

func (l *SMTListener) ExitGet_model_response(c *parser.Get_model_responseContext) {

}

func (l *SMTListener) ExitModel_response(c *parser.Model_responseContext) {

}

func (l *SMTListener) ExitFunction_def(c *parser.Function_defContext) {
	term := l.pop()
	sort := l.pop()
	sym := l.pop()

	value := convertTerm(sort.(string), term.(string))
	key, id := splitIdent(sym.(string))
	i, err := strconv.ParseInt(key, 10, 16)
	k := int16(i)
	if err != nil {
		panic(fmt.Sprintf("symbol returned from model is malformed. got=%s", sym.(string)))
	}

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
	if c.GetChildCount() > 1 {
		l.pop() // when negative term is also a child of term :(
	}
	l.push(c.GetText())
}

func (l *SMTListener) ExitSort(c *parser.SortContext) {
	l.push(c.GetText())
}

func convertTerm(sort string, term string) interface{} {
	var value interface{}
	var err error

	if string(term[0]) == "(" {
		term = term[1 : len(term)-2]
	}

	switch sort {
	case "Real":
		value, err = strconv.ParseFloat(term, 64)
		if err != nil {
			panic(err)
		}
	case "Bool":
		if term == "true" {
			value = true
		} else if term == "false" {
			value = false
		} else {
			panic(fmt.Sprintf("bool not a valid bool. got=%s", term))
		}
	case "Int":
		value, err = strconv.ParseInt(term, 10, 64)
		if err != nil {
			panic(err)
		}
	default:
		value = term
	}
	return value
}

func splitIdent(ident string) (string, string) {
	s := strings.Split(ident, "_")
	return s[len(s)-1], strings.Join(s[0:len(s)-1], "_")
}
