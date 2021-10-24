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
	s, l.stack = l.stack[len(l.stack)-1], l.stack[:len(l.stack)-1]
	return s
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
	k, err := strconv.ParseInt(key, 10, 64)
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

func (l *SMTListener) ExitSymbol(c *parser.SymbolContext) {
	sym := c.GetText()
	l.push(sym)
}

func (l *SMTListener) ExitTerm(c *parser.TermContext) {
	term := c.GetText()
	if term != "true" && term != "false" {
		//Like the sort, if this is a Boolean
		//it's a symbol too.
		l.push(term)
	}
}

// Sorts are also symbols so this results in duplicate
// stuff in the stack
/*func (l *SMTListener) ExitSort(c *parser.SortContext) {
	sort := c.GetText()
	l.push(sort)
}*/

func convertTerm(sort string, term string) interface{} {
	var value interface{}
	var err error

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
