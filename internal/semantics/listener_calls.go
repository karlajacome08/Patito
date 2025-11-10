package semantics

import (
	"fmt"

	p "Entrega1_GO/gen/grammar"
)

func (s *SemanticListener) ExitCall(ctx *p.CallContext) {
	name := ctx.ID().GetText()
	fn, ok := s.FD.Get(name)
	if !ok {
		s.err(ctx.GetStart(), "función no declarada: "+name)
		s.callRet.Push(TVoid)
		return
	}

	// Conteo y tipos de args
	argCount := len(ctx.AllExpr())
	if argCount != len(fn.Params) {
		s.err(ctx.GetStart(), fmt.Sprintf("llamada a %s: se esperaban %d args, llegaron %d",
			name, len(fn.Params), argCount))
	}
	for i := argCount - 1; i >= 0; i-- {
		argT, ok := s.types.Pop()
		if !ok {
			s.err(ctx.GetStart(), "faltan argumentos evaluados en pila de tipos")
			break
		}
		if i < len(fn.Params) {
			if !AssignCompatible(fn.Params[i].Type, argT) {
				s.err(ctx.GetStart(), fmt.Sprintf("argumento #%d de %s: se esperaba %s, llegó %s",
					i+1, name, fn.Params[i].Type, argT))
			}
		}
	}
	s.callRet.Push(fn.ReturnType)
}

func (s *SemanticListener) ExitCallPrim(ctx *p.CallPrimContext) {
	ret, ok := s.callRet.Pop()
	if !ok {
		s.err(ctx.GetStart(), "inconsistencia en llamada: falta tipo de retorno")
		ret = TVoid
	}
	if ret == TVoid {
		s.err(ctx.GetStart(), "una función NULA no puede usarse en una expresión")
	}
	s.types.Push(ret)
}
