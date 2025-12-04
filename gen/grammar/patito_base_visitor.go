// Code generated from grammar/Patito.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Patito

import "github.com/antlr4-go/antlr/v4"

type BasePatitoVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePatitoVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitVars(ctx *VarsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitVarDecl(ctx *VarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitIdList(ctx *IdListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitType_(ctx *Type_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitFuncDef(ctx *FuncDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitParams(ctx *ParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitParam(ctx *ParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitBody(ctx *BodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitAssign(ctx *AssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitPrint(ctx *PrintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitPrintItemList(ctx *PrintItemListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitPrintItem(ctx *PrintItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitIfMarkThen(ctx *IfMarkThenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitElseMark(ctx *ElseMarkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitWhileStmt(ctx *WhileStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitWhileStart(ctx *WhileStartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitWhileCond(ctx *WhileCondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitCall(ctx *CallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitRelExpr(ctx *RelExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitRelop(ctx *RelopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitPrefix(ctx *PrefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitToPrimary(ctx *ToPrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitParens(ctx *ParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitCallPrim(ctx *CallPrimContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitIdent(ctx *IdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitInt(ctx *IntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitFloat(ctx *FloatContext) interface{} {
	return v.VisitChildren(ctx)
}
