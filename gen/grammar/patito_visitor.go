// Code generated from grammar/Patito.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Patito

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by PatitoParser.
type PatitoVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PatitoParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by PatitoParser#vars.
	VisitVars(ctx *VarsContext) interface{}

	// Visit a parse tree produced by PatitoParser#varDecl.
	VisitVarDecl(ctx *VarDeclContext) interface{}

	// Visit a parse tree produced by PatitoParser#idList.
	VisitIdList(ctx *IdListContext) interface{}

	// Visit a parse tree produced by PatitoParser#type_.
	VisitType_(ctx *Type_Context) interface{}

	// Visit a parse tree produced by PatitoParser#funcDef.
	VisitFuncDef(ctx *FuncDefContext) interface{}

	// Visit a parse tree produced by PatitoParser#params.
	VisitParams(ctx *ParamsContext) interface{}

	// Visit a parse tree produced by PatitoParser#param.
	VisitParam(ctx *ParamContext) interface{}

	// Visit a parse tree produced by PatitoParser#body.
	VisitBody(ctx *BodyContext) interface{}

	// Visit a parse tree produced by PatitoParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by PatitoParser#returnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by PatitoParser#assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by PatitoParser#print.
	VisitPrint(ctx *PrintContext) interface{}

	// Visit a parse tree produced by PatitoParser#printItemList.
	VisitPrintItemList(ctx *PrintItemListContext) interface{}

	// Visit a parse tree produced by PatitoParser#printItem.
	VisitPrintItem(ctx *PrintItemContext) interface{}

	// Visit a parse tree produced by PatitoParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by PatitoParser#ifMarkThen.
	VisitIfMarkThen(ctx *IfMarkThenContext) interface{}

	// Visit a parse tree produced by PatitoParser#elseMark.
	VisitElseMark(ctx *ElseMarkContext) interface{}

	// Visit a parse tree produced by PatitoParser#whileStmt.
	VisitWhileStmt(ctx *WhileStmtContext) interface{}

	// Visit a parse tree produced by PatitoParser#whileStart.
	VisitWhileStart(ctx *WhileStartContext) interface{}

	// Visit a parse tree produced by PatitoParser#whileCond.
	VisitWhileCond(ctx *WhileCondContext) interface{}

	// Visit a parse tree produced by PatitoParser#call.
	VisitCall(ctx *CallContext) interface{}

	// Visit a parse tree produced by PatitoParser#RelExpr.
	VisitRelExpr(ctx *RelExprContext) interface{}

	// Visit a parse tree produced by PatitoParser#relop.
	VisitRelop(ctx *RelopContext) interface{}

	// Visit a parse tree produced by PatitoParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by PatitoParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by PatitoParser#Prefix.
	VisitPrefix(ctx *PrefixContext) interface{}

	// Visit a parse tree produced by PatitoParser#ToPrimary.
	VisitToPrimary(ctx *ToPrimaryContext) interface{}

	// Visit a parse tree produced by PatitoParser#Parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by PatitoParser#CallPrim.
	VisitCallPrim(ctx *CallPrimContext) interface{}

	// Visit a parse tree produced by PatitoParser#Ident.
	VisitIdent(ctx *IdentContext) interface{}

	// Visit a parse tree produced by PatitoParser#Int.
	VisitInt(ctx *IntContext) interface{}

	// Visit a parse tree produced by PatitoParser#Float.
	VisitFloat(ctx *FloatContext) interface{}
}
