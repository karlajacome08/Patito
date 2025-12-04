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
		return
	}

	argCount := len(ctx.AllExpr())
	if argCount != len(fn.Params) {
		s.err(ctx.GetStart(), fmt.Sprintf("llamada a %s: se esperaban %d args, llegaron %d",
			name, len(fn.Params), argCount))
	}

	s.S.Quads().Emit(OEra, name, "", "")

	//De abajo a arriba
	for i := argCount - 1; i >= 0; i-- {
		argVal, okOp := s.S.operands.Pop()
		if !okOp {
			s.err(ctx.GetStart(), "faltan operandos para los argumentos de la llamada")
			break
		}
		argT, ok := s.S.types.Pop()
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

		s.S.Quads().Emit(OParam, ToAddrString(argVal.Address), "", fmt.Sprintf("%d", i))
	}

	s.S.Quads().Emit(OGoSub, name, "", fmt.Sprintf("%d", fn.StartQuad))
}

func (s *SemanticListener) ExitCallPrim(ctx *p.CallPrimContext) {
	callCtx := ctx.Call()
	name := callCtx.ID().GetText()

	fn, ok := s.FD.Get(name)
	if !ok {
		s.err(ctx.GetStart(), "función no declarada en llamada: "+name)
		return
	}

	if fn.ReturnType == TVoid {
		s.err(ctx.GetStart(), "una función NULA no puede usarse en una expresión")
		return
	}

	if fn.ReturnAddr <= 0 {
		s.err(ctx.GetStart(), "función sin slot de retorno: "+name)
		return
	}

	s.S.PushOperand(name, fn.ReturnType, fn.ReturnAddr)
}
