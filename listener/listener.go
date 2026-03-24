package listener

// From antlr parse tree create Fault AST

import (
	"fault/ast"
	"fault/parser"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type FaultListener struct {
	*parser.BaseFaultParserListener
	stack                []ast.Node
	AST                  *ast.Spec
	structscope          string
	scope                string
	currSpec             string
	specs                []string
	skipRun              bool
	Path                 string // The location of the main spec
	testing              bool   // bypass imports when we're running unit tests
	Uncertains           map[string][]float64
	Unknowns             []string
	StructsPropertyOrder map[string][]string
	instances            map[string]*ast.Instance
	swaps                map[string][]ast.Node
}

func NewListener(path string, testing bool, skipRun bool) *FaultListener {
	return &FaultListener{
		Path:                 path,
		testing:              testing,
		skipRun:              skipRun,
		Uncertains:           make(map[string][]float64),
		StructsPropertyOrder: make(map[string][]string),
		instances:            make(map[string]*ast.Instance),
		swaps:                make(map[string][]ast.Node),
	}
}

// Enter rules --> Validation
// Exit rules --> AST assembly

func Execute(spec string, path string, flags map[string]bool /*specType bool, testing bool*/) (l *FaultListener, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	is := antlr.NewInputStream(spec)
	lexer := parser.NewFaultLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFaultParser(stream)
	l = NewListener(path, flags["testing"], flags["skipRun"])

	if flags["specType"] {
		antlr.ParseTreeWalkerDefault.Walk(l, p.Spec())
	} else {
		antlr.ParseTreeWalkerDefault.Walk(l, p.SysSpec())
	}
	return l, nil
}

func (l *FaultListener) validate() {
	if l.testing { //will allow invalid specs during testing
		return
	}

	if len(l.stack) < 2 {
		panic(fmt.Sprintf("Malformed fspec or fsystem file. Too few statements (got %d).", len(l.stack)))
	}

	for _, v := range l.stack {

		if _, ok := v.(*ast.AssertionStatement); ok {
			return
		}

		if _, ok := v.(*ast.DefStatement); ok {
			return
		}

		if forS, ok := v.(*ast.ForStatement); ok {
			if len(forS.Inits.Statements) > 0 {
				return
			}
		}
	}

	panic("Malformed fspec or fsystem file. No model possible.")
}

func (l *FaultListener) push(n ast.Node) {
	l.stack = append(l.stack, n)
}

func (l *FaultListener) pushN(n []ast.Node) {
	l.stack = append(l.stack, n...)
}

func (l *FaultListener) pop() ast.Node {
	var s ast.Node
	s, l.stack = l.stack[len(l.stack)-1], l.stack[:len(l.stack)-1]
	return s
}

func (l *FaultListener) peek() ast.Node {
	return l.stack[len(l.stack)-1]
}
