package listener

import (
	"fmt"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type FaultErrorListener struct {
	*antlr.DefaultErrorListener
	Filename string
}

func (f *FaultErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	file := strings.Split(f.Filename, string(os.PathSeparator))

	sym, ok := offendingSymbol.(antlr.Token)
	if !ok {
		panic(fmt.Sprintf("Invalid spec syntax on line %d col %d in spec %s", line, column, file[len(file)-1]))
	} else {
		panic(fmt.Sprintf("Invalid spec syntax %s on line %d col %d in spec %s", sym.GetText(), line, column, file[len(file)-1]))
	}

}
