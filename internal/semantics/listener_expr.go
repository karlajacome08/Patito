package semantics

import (
	p "Entrega1_GO/gen/grammar"

	"github.com/antlr4-go/antlr/v4"
)

func (s *SemanticListener) ExitIdent(ctx *p.IdentContext) {
	name := ctx.ID().GetText()
	if v, ok := s.lookup(name); ok {
		s.S.OnId(name, v.Type)
	} else {
		s.err(ctx.GetStart(), "variable no declarada: "+name)
		s.S.OnId(name, TVoid)
	}
}

// Constantes
func (s *SemanticListener) ExitInt(ctx *p.IntContext)     { s.S.OnIntConst(ctx.GetText()) }
func (s *SemanticListener) ExitFloat(ctx *p.FloatContext) { s.S.OnFloatConst(ctx.GetText()) }

func (s *SemanticListener) ExitAddSub(ctx *p.AddSubContext) {
	for _, ch := range ctx.GetChildren() {
		if t, ok := ch.(antlr.TerminalNode); ok {
			switch t.GetText() {
			case "+":
				s.S.OnAdd()
			case "-":
				s.S.OnSub()
			}
		}
	}
}

func (s *SemanticListener) ExitMulDiv(ctx *p.MulDivContext) {
	for _, ch := range ctx.GetChildren() {
		if t, ok := ch.(antlr.TerminalNode); ok {
			switch t.GetText() {
			case "*":
				s.S.OnMul()
			case "/":
				s.S.OnDiv()
			}
		}
	}
}

func (s *SemanticListener) ExitPrefix(ctx *p.PrefixContext) {
	txt := ctx.GetText()
	if txt == "" {
		return
	}
	switch txt[0] {
	case '-':
		s.S.OnUMinus()
	case '!':
		s.S.OnNot()
	}
}

func (s *SemanticListener) ExitRelExpr(ctx *p.RelExprContext) {
	if r := ctx.Relop(); r != nil {
		switch r.GetText() {
		case "==":
			s.S.OnEq()
		case "!=":
			s.S.OnNeq()
		case "<":
			s.S.OnLt()
		case "<=":
			s.S.OnLe()
		case ">":
			s.S.OnGt()
		case ">=":
			s.S.OnGe()
		}
	}
	s.S.OnEndExpr()
}

func (s *SemanticListener) EnterParens(_ *p.ParensContext) { s.S.OnLParen() }
func (s *SemanticListener) ExitParens(_ *p.ParensContext)  { s.S.OnRParen() }
