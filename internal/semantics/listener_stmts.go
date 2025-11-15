package semantics

import (
	p "Entrega1_GO/gen/grammar"
)

// a = expr ;
func (s *SemanticListener) ExitAssign(ctx *p.AssignContext) {
	lhs := ctx.ID().GetText()
	v, ok := s.lookup(lhs)
	if !ok {
		s.err(ctx.GetStart(), "variable no declarada: "+lhs)
		s.S.OnAssign(lhs, TVoid)
		return
	}
	s.S.OnAssign(lhs, v.Type)
}

// ES: print: ESCRIBE '(' printItemList? ')' ';'
// EN lugar de "write", aquí imprimimos cada item.
func (s *SemanticListener) ExitPrintItem(ctx *p.PrintItemContext) {
	if str := ctx.STRING(); str != nil {
		// STRING directo
		s.S.OnStringConst(str.GetText())
		s.S.OnWriteExpr()
		return
	}
	// Es una expr: ya se fueron apilando operandos/operadores;
	// aquí cerramos y emitimos el WRITE de esa expr.
	s.S.OnWriteExpr()
}
