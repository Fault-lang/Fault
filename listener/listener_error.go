package listener

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type FaultErrorListener struct {
	antlr.ErrorListener
}

func (f *FaultErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	sym, ok := offendingSymbol.(string)
	if !ok{
		fmt.Printf("Invalid spec syntax on line %d col %d\n", line, column)
	}else{
		fmt.Printf("Invalid spec syntax %s on line %d col %d\n", sym, line, column)
	}
}
