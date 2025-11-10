package semantics

import (
	p "Entrega1_GO/gen/grammar"
)

func (s *SemanticListener) EnterProgram(ctx *p.ProgramContext) {
	if len(s.scope) == 0 {
		s.pushScope("global")
	}
}
