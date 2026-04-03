//go:build ignore

package main

import (
	"fmt"
	"fault/ast"
	"fault/lexer"
	"fault/parser"
	"fault/preprocess"
)

func prepTest(test string, initialPass bool) *preprocess.Processor {
	l := lexer.New(test)
	p := parser.New(l)
	tree := p.ParseProgram()
	process := preprocess.NewProcesser()
	process.Run(tree, initialPass)
	return process
}

func main() {
	test := `spec test1;
const a = 2;
for 1 run {
	if a == 2{
		if a != 0{
			3;
		}else if a < 1 {
			if a >= 2 {
				true;
			}
		}
	}else if a !=5 {
		true;
	}else{
		if a > 4 {
			false;
		}
	}
};
`
	process := prepTest(test, true)
	tree := process.Processed
	spec := tree.Statements

	if1 := spec[2].(*ast.ForStatement).Body.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.IfExpression)
	fmt.Printf("if1.Condition = %s\n", if1.Condition)
	if if1.Elif != nil {
		fmt.Printf("if1.Elif.Condition = %s\n", if1.Elif.Condition)
		if if1.Elif.Elif != nil {
			fmt.Printf("if1.Elif.Elif.Condition = %s\n", if1.Elif.Elif.Condition)
			if if1.Elif.Elif.Elif != nil {
				fmt.Printf("if1.Elif.Elif.Elif.Condition = %s\n", if1.Elif.Elif.Elif.Condition)
			}
		}
	}
}
