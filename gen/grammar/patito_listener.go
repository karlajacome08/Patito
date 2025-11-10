// Code generated from grammar/Patito.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Patito

import "github.com/antlr4-go/antlr/v4"

// PatitoListener is a complete listener for a parse tree produced by PatitoParser.
type PatitoListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterVars is called when entering the vars production.
	EnterVars(c *VarsContext)

	// EnterVarDecl is called when entering the varDecl production.
	EnterVarDecl(c *VarDeclContext)

	// EnterIdList is called when entering the idList production.
	EnterIdList(c *IdListContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterFuncDef is called when entering the funcDef production.
	EnterFuncDef(c *FuncDefContext)

	// EnterParams is called when entering the params production.
	EnterParams(c *ParamsContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterBody is called when entering the body production.
	EnterBody(c *BodyContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterPrint is called when entering the print production.
	EnterPrint(c *PrintContext)

	// EnterPrintItemList is called when entering the printItemList production.
	EnterPrintItemList(c *PrintItemListContext)

	// EnterPrintItem is called when entering the printItem production.
	EnterPrintItem(c *PrintItemContext)

	// EnterWhileStmt is called when entering the whileStmt production.
	EnterWhileStmt(c *WhileStmtContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterCall is called when entering the call production.
	EnterCall(c *CallContext)

	// EnterRelExpr is called when entering the RelExpr production.
	EnterRelExpr(c *RelExprContext)

	// EnterRelop is called when entering the relop production.
	EnterRelop(c *RelopContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterPrefix is called when entering the Prefix production.
	EnterPrefix(c *PrefixContext)

	// EnterToPrimary is called when entering the ToPrimary production.
	EnterToPrimary(c *ToPrimaryContext)

	// EnterParens is called when entering the Parens production.
	EnterParens(c *ParensContext)

	// EnterCallPrim is called when entering the CallPrim production.
	EnterCallPrim(c *CallPrimContext)

	// EnterIdent is called when entering the Ident production.
	EnterIdent(c *IdentContext)

	// EnterInt is called when entering the Int production.
	EnterInt(c *IntContext)

	// EnterFloat is called when entering the Float production.
	EnterFloat(c *FloatContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitVars is called when exiting the vars production.
	ExitVars(c *VarsContext)

	// ExitVarDecl is called when exiting the varDecl production.
	ExitVarDecl(c *VarDeclContext)

	// ExitIdList is called when exiting the idList production.
	ExitIdList(c *IdListContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitFuncDef is called when exiting the funcDef production.
	ExitFuncDef(c *FuncDefContext)

	// ExitParams is called when exiting the params production.
	ExitParams(c *ParamsContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitBody is called when exiting the body production.
	ExitBody(c *BodyContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitPrint is called when exiting the print production.
	ExitPrint(c *PrintContext)

	// ExitPrintItemList is called when exiting the printItemList production.
	ExitPrintItemList(c *PrintItemListContext)

	// ExitPrintItem is called when exiting the printItem production.
	ExitPrintItem(c *PrintItemContext)

	// ExitWhileStmt is called when exiting the whileStmt production.
	ExitWhileStmt(c *WhileStmtContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitCall is called when exiting the call production.
	ExitCall(c *CallContext)

	// ExitRelExpr is called when exiting the RelExpr production.
	ExitRelExpr(c *RelExprContext)

	// ExitRelop is called when exiting the relop production.
	ExitRelop(c *RelopContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitPrefix is called when exiting the Prefix production.
	ExitPrefix(c *PrefixContext)

	// ExitToPrimary is called when exiting the ToPrimary production.
	ExitToPrimary(c *ToPrimaryContext)

	// ExitParens is called when exiting the Parens production.
	ExitParens(c *ParensContext)

	// ExitCallPrim is called when exiting the CallPrim production.
	ExitCallPrim(c *CallPrimContext)

	// ExitIdent is called when exiting the Ident production.
	ExitIdent(c *IdentContext)

	// ExitInt is called when exiting the Int production.
	ExitInt(c *IntContext)

	// ExitFloat is called when exiting the Float production.
	ExitFloat(c *FloatContext)
}
