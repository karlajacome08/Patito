package semantics

import (
	p "Entrega1_GO/gen/grammar"
	"fmt"
)

// Al terminar la condición y justo antes del cuerpo THEN
func (s *SemanticListener) ExitIfMarkThen(ctx *p.IfMarkThenContext) {
	// ya se debió haber evaluado la expr y generado sus temporales
	s.S.EndExpression()
	cond, ok := s.S.operands.Pop()
	if !ok {
		s.err(ctx.GetStart(), "IF sin expresión de condición")
		return
	}
	t, _ := s.S.types.Pop()
	if t != TBool {
		s.err(ctx.GetStart(), "condición de IF debe ser bool")
	}

	// Emitir GOTOF y guardar índice en pila de saltos
	quadIdx := s.S.quads.Emit(OGotoF, ToAddrString(cond.Address), "", "")
	s.S.jumpStack.Push(quadIdx)
}

// Al encontrar 'sino' antes de else-body
func (s *SemanticListener) ExitElseMark(ctx *p.ElseMarkContext) {
	// salto para saltarse el else
	gotoIdx := s.S.quads.Emit(OGoto, "", "", "")

	// parchear el GOTOF anterior para que apunte al inicio del ELSE
	ifIdx, ok := s.S.jumpStack.Pop()
	if !ok {
		s.err(ctx.GetStart(), "pila de saltos vacía en ELSE")
		return
	}
	// el destino del GOTOF es el siguiente quad
	s.S.quads.list[ifIdx].Result = fmt.Sprintf("%d", s.S.quads.Len())

	// guardar el GOTO en la pila para parcharlo al final del ELSE
	s.S.jumpStack.Push(gotoIdx)
}

// Al final del ifStmt completo
func (s *SemanticListener) ExitIfStmt(ctx *p.IfStmtContext) {
	// si hubo ELSE, aquí parchas el GOTO final
	if !s.S.jumpStack.Empty() {
		idx, _ := s.S.jumpStack.Pop()
		s.S.quads.list[idx].Result = fmt.Sprintf("%d", s.S.quads.Len())
	}
}

func (s *SemanticListener) ExitWhileStart(ctx *p.WhileStartContext) {
	start := s.S.quads.Len()
	s.S.jumpStack.Push(start)
}

func (s *SemanticListener) ExitWhileCond(ctx *p.WhileCondContext) {
	s.S.EndExpression()
	cond, ok := s.S.operands.Pop()
	if !ok {
		s.err(ctx.GetStart(), "WHILE sin condición")
		return
	}
	t, _ := s.S.types.Pop()
	if t != TBool {
		s.err(ctx.GetStart(), "condición de WHILE debe ser bool")
	}

	startIdx, ok := s.S.jumpStack.Pop()
	if !ok {
		s.err(ctx.GetStart(), "pila de saltos vacía en WHILE")
		return
	}
	s.S.jumpStack.Push(startIdx)

	gotofIdx := s.S.quads.Emit(OGotoF, ToAddrString(cond.Address), "", "")
	s.S.jumpStack.Push(gotofIdx)
}

// al final del whileStmt (después del body)
func (s *SemanticListener) ExitWhileStmt(ctx *p.WhileStmtContext) {
	// sacamos GOTOF y start
	gotofIdx, _ := s.S.jumpStack.Pop()
	startIdx, _ := s.S.jumpStack.Pop()

	// GOTO de regreso al inicio
	s.S.quads.Emit(OGoto, "", "", fmt.Sprintf("%d", startIdx))

	// parchear GOTOF para que brinque al siguiente quad (después del ciclo)
	s.S.quads.list[gotofIdx].Result = fmt.Sprintf("%d", s.S.quads.Len())
}
