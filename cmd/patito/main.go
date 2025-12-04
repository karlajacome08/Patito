package main

import (
	"fmt"
	"os"

	p "Entrega1_GO/gen/grammar"
	"Entrega1_GO/internal/semantics"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("uso: patito <archivo.pat>")
		os.Exit(1)
	}
	src, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	is := antlr.NewInputStream(string(src))
	lex := p.NewPatitoLexer(is)
	ts := antlr.NewCommonTokenStream(lex, 0)
	par := p.NewPatitoParser(ts)

	tree := par.Program()

	listener := semantics.NewSemanticListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	if err := listener.Error(); err != nil {
		fmt.Println("=== ERRORES ===")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("=== CUADRUPLOS ===")

	fmt.Print(listener.S.Quads().String())
}
