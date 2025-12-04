package semantics

import (
	"fmt"

	p "Entrega1_GO/gen/grammar"
)

func (s *SemanticListener) ExitReturnStmt(ctx *p.ReturnStmtContext) {
	fnName := s.curScope()
	fn, ok := s.FD.Get(fnName)
	if !ok {
		s.err(ctx.GetStart(), "regresa fuera de una función")
		return
	}

	if fn.ReturnType == TVoid {
		s.err(ctx.GetStart(), "una función NULA no puede usar 'regresa'")
		return
	}

	// Aseguramos que la expresión esté reducida
	s.S.EndExpression()

	rhs, ok := s.S.operands.Pop()
	if !ok {
		s.err(ctx.GetStart(), "regresa sin expresión de retorno")
		return
	}
	_, _ = s.S.types.Pop()

	if !IsCompatibleAssign(fn.ReturnType, rhs.Type) {
		s.err(ctx.GetStart(), fmt.Sprintf(
			"tipo de retorno incompatible en %s: se esperaba %s, llegó %s",
			fnName, fn.ReturnType, rhs.Type,
		))
	}

	if fn.ReturnAddr <= 0 {
		s.err(ctx.GetStart(), "función sin dirección de retorno asignada: "+fnName)
		return
	}

	s.S.quads.Emit(
		OReturn,
		ToAddrString(rhs.Address),
		"",
		ToAddrString(fn.ReturnAddr),
	)
}
