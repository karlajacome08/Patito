package main

import (
	"flag"
	"fmt"
	"os"

	grammar "Entrega1_GO/gen/grammar"
	"Entrega1_GO/internal/semantics"

	"github.com/antlr4-go/antlr/v4"
)

type syntaxErrListener struct {
	*antlr.DefaultErrorListener
	hadError bool
}

func (l *syntaxErrListener) SyntaxError(r antlr.Recognizer, sym interface{},
	line, column int, msg string, e antlr.RecognitionException) {
	l.hadError = true
	fmt.Fprintf(os.Stderr, "[Sintaxis] L%d:C%d %s\n", line, column, msg)
}

func main() {
	var dump bool
	flag.BoolVar(&dump, "dump", false, "Imprimir directorio de funciones y tablas de variables")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("Uso: patito [-dump] <archivo.pat>")
		os.Exit(2)
	}
	filename := flag.Arg(0)

	input, err := antlr.NewFileStream(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "No pude abrir %s: %v\n", filename, err)
		os.Exit(1)
	}

	lex := grammar.NewPatitoLexer(input)
	lex.RemoveErrorListeners()
	syn := &syntaxErrListener{}
	lex.AddErrorListener(syn)

	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	par := grammar.NewPatitoParser(tokens)
	par.RemoveErrorListeners()
	par.AddErrorListener(syn)

	tree := par.Program()

	// Si hay errores léxicos/sintácticos -> para
	if syn.hadError {
		os.Exit(1)
	}

	sem := semantics.NewSemanticListener()
	antlr.ParseTreeWalkerDefault.Walk(sem, tree)

	if err := sem.Error(); err != nil {
		fmt.Fprintln(os.Stderr, "Errores semánticos:\n"+err.Error())
		os.Exit(1)
	}

	fmt.Println("✔ OK: Scanner/Parser y análisis semántico completados")

	if dump {
		fmt.Println("── Directorio de funciones y variables ──")
		dumpSymbols(sem.FD)
	}
}

func dumpSymbols(fd *semantics.FunctionDirectory) {
	for name, fn := range fd.Fns {
		fmt.Printf("func %s(", name)
		for i, p := range fn.Params {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%s:%s", p.Name, p.Type)
		}
		fmt.Printf(") -> %s\n", fn.ReturnType)
		// locals
		if len(fn.Locals.Symbols) > 0 {
			fmt.Println("  locals:")
			for _, v := range fn.Locals.Symbols {
				fmt.Printf("    - %s : %s (%s)\n", v.Name, v.Type, v.Kind)
			}
		}
	}
}
