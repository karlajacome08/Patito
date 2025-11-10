package semantics

import (
	"fmt"

	p "Entrega1_GO/gen/grammar"
)

func (s *SemanticListener) ExitAssign(ctx *p.AssignContext) {
	lhs := ctx.ID().GetText()

	v, ok := s.lookup(lhs)
	if !ok {
		s.err(ctx.GetStart(), "variable no declarada: "+lhs)
		return
	}
	rT, ok := s.types.Pop()
	if !ok {
		s.err(ctx.GetStart(), "expresión vacía en asignación")
		return
	}
	if !AssignCompatible(v.Type, rT) {
		s.err(ctx.GetStart(), fmt.Sprintf("type mismatch: %s = %s", v.Type, rT))
	}
}

func (s *SemanticListener) ExitIfStmt(ctx *p.IfStmtContext) {
	t, ok := s.types.Pop()
	if !ok {
		s.err(ctx.GetStart(), "falta condición en if")
		return
	}
	if t != TBool {
		s.err(ctx.GetStart(), "la condición de if debe ser bool (usa un relop como ==, <, etc.)")
	}
}

func (s *SemanticListener) ExitWhileStmt(ctx *p.WhileStmtContext) {
	t, ok := s.types.Pop()
	if !ok {
		s.err(ctx.GetStart(), "falta condición en while")
		return
	}
	if t != TBool {
		s.err(ctx.GetStart(), "la condición de while debe ser bool (usa un relop como ==, <, etc.)")
	}
}

func (s *SemanticListener) ExitStmt(ctx *p.StmtContext) {
	if c := ctx.Call(); c != nil {
		_, _ = s.callRet.Pop()
	}
}
