package semantics

import (
	p "Entrega1_GO/gen/grammar"
	"fmt"
)

func (s *SemanticListener) ExitVarDecl(ctx *p.VarDeclContext) {
	fn := s.curScope()
	t := mapType(ctx.Type_())

	ids := ctx.IdList().AllID()
	for _, idTok := range ids {
		kind := ifElse(fn == "global", KindGlobal, KindLocal)
		v := VariableInfo{
			Name: idTok.GetText(),
			Type: t,
			Kind: kind,
		}

		v.Address = s.Mem.AllocVar(v.Kind, v.Type)

		if err := s.FD.AddLocal(fn, v); err != nil {
			s.err(ctx.GetStart(), err.Error())
		}
	}
}

func (s *SemanticListener) EnterFuncDef(ctx *p.FuncDefContext) {
	name := ctx.ID().GetText()
	ret := TVoid
	if ctx.NULA() == nil {
		ret = mapType(ctx.Type_())
	}
	if err := s.FD.Declare(name, ret); err != nil {
		s.err(ctx.GetStart(), err.Error())
	}
	s.pushScope(name)

	if fn, ok := s.FD.Get(name); ok {
		if ret != TVoid {
			fn.ReturnAddr = s.Mem.AllocVar(KindLocal, ret)
			if fn.ReturnAddr == -1 {
				s.err(ctx.GetStart(), fmt.Sprintf("no se pudo asignar dirección a variable de retorno de función: %s", name))
			}
		}
		fn.StartQuad = s.S.Quads().Len()
	}
}

func (s *SemanticListener) ExitFuncDef(ctx *p.FuncDefContext) {
	s.S.Quads().Emit(OEndFunc, "", "", "")
	s.popScope()
}

func (s *SemanticListener) ExitParam(ctx *p.ParamContext) {
	fn := s.curScope()
	name := ctx.ID().GetText()
	t := mapType(ctx.Type_())

	v := VariableInfo{
		Name:    name,
		Type:    t,
		Kind:    KindParam,
		Address: s.Mem.AllocVar(KindParam, t),
	}

	if v.Address == -1 {
		s.err(ctx.GetStart(), fmt.Sprintf("no se pudo asignar dirección a parámetro: %s", name))
	}

	if err := s.FD.AddParam(fn, v); err != nil {
		s.err(ctx.GetStart(), err.Error())
	}
}
