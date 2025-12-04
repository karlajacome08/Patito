package semantics

import (
	p "Entrega1_GO/gen/grammar"
)

func (s *SemanticListener) ExitAssign(ctx *p.AssignContext) {
	lhs := ctx.ID().GetText()
	v, ok := s.lookup(lhs)
	if !ok {
		s.err(ctx.GetStart(), "variable no declarada: "+lhs)
		// pasamos addr inválido para no romper el flujo
		s.S.OnAssign(-1, TVoid)
		return
	}
	if v.Address == 0 {
		s.err(ctx.GetStart(), "variable sin dirección asignada: "+lhs)
	}
	s.S.OnAssign(v.Address, v.Type)
}

func (s *SemanticListener) ExitPrintItem(ctx *p.PrintItemContext) {
	if str := ctx.STRING(); str != nil {
		// STRING directo
		s.S.OnStringConst(str.GetText())
		s.S.OnWriteExpr()
		return
	}
	s.S.OnWriteExpr()
}
