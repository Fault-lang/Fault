package listener

import (
	"fault/parser"
	"fmt"
	"os"
	"strings"
	"unicode"

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
		symText := sym.GetText()
		hint := underscoreHint(msg, symText) + keywordHint(msg, symText)
		panic(fmt.Sprintf("Invalid spec syntax %s on line %d col %d in spec %s%s", symText, line, column, specName, hint))
	}
}

// underscoreHint returns a hint string when the error looks like an underscore in an identifier.
func underscoreHint(msg string, symText string) string {
	if symText == "_" || strings.Contains(msg, "'_'") {
		return "\n  Hint: Fault identifiers cannot contain underscores — use camelCase instead (e.g. 'myVariable' not 'my_variable')"
	}
	return ""
}

// reservedKeywords is derived at init time from the generated lexer's LiteralNames.
// All tokens defined before IDENT in the grammar whose literals contain only
// letters, digits, hyphens, or underscores are keywords that cannot be used as identifiers.
var reservedKeywords map[string]bool

func init() {
	parser.FaultLexerInit()
	literals := parser.FaultLexerLexerStaticData.LiteralNames
	reservedKeywords = make(map[string]bool, parser.FaultLexerIDENT)
	// LiteralNames is 1-indexed (index 0 is empty); stop before the IDENT token type.
	for i := 1; i < parser.FaultLexerIDENT && i < len(literals); i++ {
		lit := literals[i]
		if len(lit) < 3 || lit[0] != '\'' || lit[len(lit)-1] != '\'' {
			continue
		}
		word := lit[1 : len(lit)-1]
		if isKeywordLiteral(word) {
			reservedKeywords[word] = true
		}
	}
}

// isKeywordLiteral reports whether s looks like a keyword (letters, digits, hyphens, underscores only).
func isKeywordLiteral(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '-' && r != '_' {
			return false
		}
	}
	return true
}

// keywordHint returns a hint string when a reserved keyword is used where an identifier is expected.
func keywordHint(msg string, symText string) string {
	if reservedKeywords[symText] && (strings.Contains(msg, "expecting IDENT") || strings.Contains(msg, "mismatched input")) {
		return fmt.Sprintf("\n  Hint: %q is a reserved keyword and cannot be used as a variable name", symText)
	}
	return ""
}
