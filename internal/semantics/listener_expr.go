package semantics

import (
	"fmt"

	p "Entrega1_GO/gen/grammar"
)

func (s *SemanticListener) ExitIdent(ctx *p.IdentContext) {
	name := ctx.ID().GetText()
	v, ok := s.lookup(name)
	if !ok {
		s.err(ctx.GetStart(), "variable no declarada: "+name)
		s.types.Push(TVoid)
		return
	}
	s.types.Push(v.Type)
}
func (s *SemanticListener) ExitInt(ctx *p.IntContext)     { s.types.Push(TInt) }
func (s *SemanticListener) ExitFloat(ctx *p.FloatContext) { s.types.Push(TFloat) }

func (s *SemanticListener) ExitRelExpr(ctx *p.RelExprContext) {
	if ctx.Relop() == nil {
		return
	}
	rT, ok := s.types.Pop()
	if !ok {
		s.err(ctx.GetStart(), "expresión relacional incompleta (falta RHS)")
		return
	}
	lT, ok := s.types.Pop()
	if !ok {
		s.err(ctx.GetStart(), "expresión relacional incompleta (falta LHS)")
		return
	}
	if !isNumeric(lT) || !isNumeric(rT) {
		s.err(ctx.GetStart(), fmt.Sprintf("relop requiere numéricos: (%s ? %s)", lT, rT))
	}
	s.types.Push(TBool)
}

func (s *SemanticListener) ExitAddSub(ctx *p.AddSubContext) {
	n := len(ctx.AllProd())
	if n <= 1 {
		return
	}
	hasFloat := false
	for i := 0; i < n; i++ {
		t, ok := s.types.Pop()
		if !ok {
			s.err(ctx.GetStart(), "expresión +|- incompleta")
			return
		}
		if t == TFloat {
			hasFloat = true
		}
		if !isNumeric(t) {
			s.err(ctx.GetStart(), "operación +|- requiere operandos numéricos")
		}
	}
	if hasFloat {
		s.types.Push(TFloat)
	} else {
		s.types.Push(TInt)
	}
}

func (s *SemanticListener) ExitMulDiv(ctx *p.MulDivContext) {
	n := len(ctx.AllUnary())
	if n <= 1 {
		return
	}
	hasFloat := false
	for i := 0; i < n; i++ {
		t, ok := s.types.Pop()
		if !ok {
			s.err(ctx.GetStart(), "expresión *|/ incompleta")
			return
		}
		if t == TFloat {
			hasFloat = true
		}
		if !isNumeric(t) {
			s.err(ctx.GetStart(), "operación *|/ requiere operandos numéricos")
		}
	}
	if hasFloat {
		s.types.Push(TFloat)
	} else {
		s.types.Push(TInt)
	}
}

func (s *SemanticListener) ExitPrefix(ctx *p.PrefixContext) {
	txt := ctx.GetText()
	if txt == "" {
		return
	}
	op := rune(txt[0])

	t, ok := s.types.Pop()
	if !ok {
		s.err(ctx.GetStart(), "unario sin operando")
		return
	}

	switch op {
	case '+', '-':
		if !isNumeric(t) {
			s.err(ctx.GetStart(), "operador unario +|- requiere numérico")
		}
		s.types.Push(t)
	case '!':
		if t != TBool {
			s.err(ctx.GetStart(), "operador ! requiere bool (por ejemplo, una comparación)")
		}
		s.types.Push(TBool)
	default:
		s.types.Push(t)
	}
}
