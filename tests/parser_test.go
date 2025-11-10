package tests

import (
	"strings"
	"testing"

	pg "Entrega1_GO/gen/grammar"

	antlr "github.com/antlr4-go/antlr/v4"
)

type errListener struct{ *antlr.DefaultErrorListener }

func (e *errListener) SyntaxError(rec antlr.Recognizer, sym interface{}, line, col int, msg string, _ antlr.RecognitionException) {
	panic(msg)
}

func parseOK(t *testing.T, src string) {
	t.Helper()

	is := antlr.NewInputStream(src)
	lex := pg.NewPatitoLexer(is)

	lex.RemoveErrorListeners()
	lex.AddErrorListener(&errListener{})

	ts := antlr.NewCommonTokenStream(lex, 0)
	p := pg.NewPatitoParser(ts)

	p.RemoveErrorListeners()
	p.AddErrorListener(&errListener{})

	p.Program()
}

func parseFAIL(t *testing.T, src string) {
	t.Helper()
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Esperaba error y parseó OK:\n%s", src)
		}
	}()
	parseOK(t, src)
}

func Test_Validos(t *testing.T) {
	casos := []string{
		`programa p; inicio { } fin`,
		`programa p; vars x, y: entero; inicio { x = 1 + 2 * 3; } fin`,
		`programa p; inicio { escribe("hola", 1+2); } fin`,
		`programa p; inicio { si (1<2) { } ; } fin`,
		`programa p; inicio { si (1<2) { } sino { } ; } fin`,
		`programa p; inicio { mientras (3!=4) haz { } ; } fin`,
		`programa p; inicio { x = (1 + (2*3)); } fin`,
		`programa p; inicio { [ x = 1; si (1<2) { } ; ] } fin`,
		`programa p; inicio { [ [ escribe("ok"); ] ] } fin`,
	}
	for _, c := range casos {
		parseOK(t, c)
	}
}

func Test_Invalidos(t *testing.T) {
	casos := []string{
		`programa p inicio { } fin`,                // falta ';'
		`programa p; inicio { x = ; } fin`,         // expr incompleta
		`programa p; inicio { si (1<) { } ; } fin`, // relacional roto
		`programa p; inicio { (1+2; } fin`,         // paréntesis sin cerrar
		`programa p; inicio { @ } fin`,             // token ilegal (error LÉXICO)
		`programa p; vars a: malo; inicio { } fin`, // tipo inexistente (error SINTÁCTICO)
		`programa p; inicio { [ x = 1; } fin`,      // falta ']'
	}
	for _, c := range casos {
		parseFAIL(t, c)
	}
}

func Test_CommentsAndWS(t *testing.T) {
	src := `
	// comentario
	programa demo;
	vars
	  x: entero; /* block
	  comment */
	inicio
	{
	  x = 1 + 2; escribe("ok"); // fin de línea
	}
	fin
	`
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("No debía fallar: %v\n%s", r, src)
		}
	}()
	parseOK(t, strings.TrimSpace(src))
}
