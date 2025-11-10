package semantics

import (
	p "Entrega1_GO/gen/grammar"
)

func (s *SemanticListener) ExitVarDecl(ctx *p.VarDeclContext) {
	fn := s.curScope()
	t := mapType(ctx.Type_())

	ids := ctx.IdList().AllID()
	for _, idTok := range ids {
		v := VariableInfo{
			Name: idTok.GetText(),
			Type: t,
			Kind: ifElse(fn == "global", KindGlobal, KindLocal),
		}
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
}

func (s *SemanticListener) ExitFuncDef(ctx *p.FuncDefContext) {
	s.popScope()
}

func (s *SemanticListener) ExitParam(ctx *p.ParamContext) {
	fn := s.curScope()
	name := ctx.ID().GetText()
	t := mapType(ctx.Type_())
	if err := s.FD.AddParam(fn, name, t); err != nil {
		s.err(ctx.GetStart(), err.Error())
	}
}
