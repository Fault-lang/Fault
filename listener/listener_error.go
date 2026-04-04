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
	specName := file[len(file)-1]

	sym, ok := offendingSymbol.(antlr.Token)
	if !ok {
		hint := underscoreHint(msg, "")
		panic(fmt.Sprintf("Invalid spec syntax on line %d col %d in spec %s%s", line, column, specName, hint))
	} else {
		hint := underscoreHint(msg, sym.GetText())
		panic(fmt.Sprintf("Invalid spec syntax %s on line %d col %d in spec %s%s", sym.GetText(), line, column, specName, hint))
	}
}

// underscoreHint returns a hint string when the error looks like an underscore in an identifier.
func underscoreHint(msg string, symText string) string {
	if symText == "_" || strings.Contains(msg, "'_'") {
		return "\n  Hint: Fault identifiers cannot contain underscores — use camelCase instead (e.g. 'myVariable' not 'my_variable')"
	}
	return ""
}
