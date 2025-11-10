package semantics

import (
	"strings"
	"unicode"

	"github.com/antlr4-go/antlr/v4"
)

func (s *SemanticListener) lookup(name string) (VariableInfo, bool) {
	if f, ok := s.FD.Get(s.curScope()); ok {
		if v, ok := f.Locals.Lookup(name); ok {
			return v, true
		}
	}
	if g, ok := s.FD.Get("global"); ok {
		if v, ok := g.Locals.Lookup(name); ok {
			return v, true
		}
	}
	return VariableInfo{}, false
}

func mapType(tctx antlr.ParserRuleContext) TypeTag {
	if tctx == nil {
		return TVoid
	}
	txt := strings.ToLower(tctx.GetText())
	switch txt {
	case "entero":
		return TInt
	case "flotante":
		return TFloat
	default:
		return TVoid
	}
}

func ifElse[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

func isNumeric(t TypeTag) bool { return t == TInt || t == TFloat }

func looksIdentifier(s string) bool {
	if s == "" {
		return false
	}
	runes := []rune(s)
	if !(unicode.IsLetter(runes[0]) || runes[0] == '_') {
		return false
	}
	for _, r := range runes[1:] {
		if !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_') {
			return false
		}
	}
	return true
}
