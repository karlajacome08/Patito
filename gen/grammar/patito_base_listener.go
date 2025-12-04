// Code generated from grammar/Patito.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Patito

import "github.com/antlr4-go/antlr/v4"

// BasePatitoListener is a complete listener for a parse tree produced by PatitoParser.
type BasePatitoListener struct{}

var _ PatitoListener = &BasePatitoListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePatitoListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePatitoListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePatitoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePatitoListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasePatitoListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasePatitoListener) ExitProgram(ctx *ProgramContext) {}

// EnterVars is called when production vars is entered.
func (s *BasePatitoListener) EnterVars(ctx *VarsContext) {}

// ExitVars is called when production vars is exited.
func (s *BasePatitoListener) ExitVars(ctx *VarsContext) {}

// EnterVarDecl is called when production varDecl is entered.
func (s *BasePatitoListener) EnterVarDecl(ctx *VarDeclContext) {}

// ExitVarDecl is called when production varDecl is exited.
func (s *BasePatitoListener) ExitVarDecl(ctx *VarDeclContext) {}

// EnterIdList is called when production idList is entered.
func (s *BasePatitoListener) EnterIdList(ctx *IdListContext) {}

// ExitIdList is called when production idList is exited.
func (s *BasePatitoListener) ExitIdList(ctx *IdListContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BasePatitoListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BasePatitoListener) ExitType_(ctx *Type_Context) {}

// EnterFuncDef is called when production funcDef is entered.
func (s *BasePatitoListener) EnterFuncDef(ctx *FuncDefContext) {}

// ExitFuncDef is called when production funcDef is exited.
func (s *BasePatitoListener) ExitFuncDef(ctx *FuncDefContext) {}

// EnterParams is called when production params is entered.
func (s *BasePatitoListener) EnterParams(ctx *ParamsContext) {}

// ExitParams is called when production params is exited.
func (s *BasePatitoListener) ExitParams(ctx *ParamsContext) {}

// EnterParam is called when production param is entered.
func (s *BasePatitoListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BasePatitoListener) ExitParam(ctx *ParamContext) {}

// EnterBody is called when production body is entered.
func (s *BasePatitoListener) EnterBody(ctx *BodyContext) {}

// ExitBody is called when production body is exited.
func (s *BasePatitoListener) ExitBody(ctx *BodyContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BasePatitoListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BasePatitoListener) ExitStmt(ctx *StmtContext) {}

// EnterAssign is called when production assign is entered.
func (s *BasePatitoListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BasePatitoListener) ExitAssign(ctx *AssignContext) {}

// EnterPrint is called when production print is entered.
func (s *BasePatitoListener) EnterPrint(ctx *PrintContext) {}

// ExitPrint is called when production print is exited.
func (s *BasePatitoListener) ExitPrint(ctx *PrintContext) {}

// EnterPrintItemList is called when production printItemList is entered.
func (s *BasePatitoListener) EnterPrintItemList(ctx *PrintItemListContext) {}

// ExitPrintItemList is called when production printItemList is exited.
func (s *BasePatitoListener) ExitPrintItemList(ctx *PrintItemListContext) {}

// EnterPrintItem is called when production printItem is entered.
func (s *BasePatitoListener) EnterPrintItem(ctx *PrintItemContext) {}

// ExitPrintItem is called when production printItem is exited.
func (s *BasePatitoListener) ExitPrintItem(ctx *PrintItemContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BasePatitoListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BasePatitoListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterIfMarkThen is called when production ifMarkThen is entered.
func (s *BasePatitoListener) EnterIfMarkThen(ctx *IfMarkThenContext) {}

// ExitIfMarkThen is called when production ifMarkThen is exited.
func (s *BasePatitoListener) ExitIfMarkThen(ctx *IfMarkThenContext) {}

// EnterElseMark is called when production elseMark is entered.
func (s *BasePatitoListener) EnterElseMark(ctx *ElseMarkContext) {}

// ExitElseMark is called when production elseMark is exited.
func (s *BasePatitoListener) ExitElseMark(ctx *ElseMarkContext) {}

// EnterWhileStmt is called when production whileStmt is entered.
func (s *BasePatitoListener) EnterWhileStmt(ctx *WhileStmtContext) {}

// ExitWhileStmt is called when production whileStmt is exited.
func (s *BasePatitoListener) ExitWhileStmt(ctx *WhileStmtContext) {}

// EnterWhileStart is called when production whileStart is entered.
func (s *BasePatitoListener) EnterWhileStart(ctx *WhileStartContext) {}

// ExitWhileStart is called when production whileStart is exited.
func (s *BasePatitoListener) ExitWhileStart(ctx *WhileStartContext) {}

// EnterWhileCond is called when production whileCond is entered.
func (s *BasePatitoListener) EnterWhileCond(ctx *WhileCondContext) {}

// ExitWhileCond is called when production whileCond is exited.
func (s *BasePatitoListener) ExitWhileCond(ctx *WhileCondContext) {}

// EnterCall is called when production call is entered.
func (s *BasePatitoListener) EnterCall(ctx *CallContext) {}

// ExitCall is called when production call is exited.
func (s *BasePatitoListener) ExitCall(ctx *CallContext) {}

// EnterRelExpr is called when production RelExpr is entered.
func (s *BasePatitoListener) EnterRelExpr(ctx *RelExprContext) {}

// ExitRelExpr is called when production RelExpr is exited.
func (s *BasePatitoListener) ExitRelExpr(ctx *RelExprContext) {}

// EnterRelop is called when production relop is entered.
func (s *BasePatitoListener) EnterRelop(ctx *RelopContext) {}

// ExitRelop is called when production relop is exited.
func (s *BasePatitoListener) ExitRelop(ctx *RelopContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BasePatitoListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BasePatitoListener) ExitAddSub(ctx *AddSubContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BasePatitoListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BasePatitoListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterPrefix is called when production Prefix is entered.
func (s *BasePatitoListener) EnterPrefix(ctx *PrefixContext) {}

// ExitPrefix is called when production Prefix is exited.
func (s *BasePatitoListener) ExitPrefix(ctx *PrefixContext) {}

// EnterToPrimary is called when production ToPrimary is entered.
func (s *BasePatitoListener) EnterToPrimary(ctx *ToPrimaryContext) {}

// ExitToPrimary is called when production ToPrimary is exited.
func (s *BasePatitoListener) ExitToPrimary(ctx *ToPrimaryContext) {}

// EnterParens is called when production Parens is entered.
func (s *BasePatitoListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production Parens is exited.
func (s *BasePatitoListener) ExitParens(ctx *ParensContext) {}

// EnterCallPrim is called when production CallPrim is entered.
func (s *BasePatitoListener) EnterCallPrim(ctx *CallPrimContext) {}

// ExitCallPrim is called when production CallPrim is exited.
func (s *BasePatitoListener) ExitCallPrim(ctx *CallPrimContext) {}

// EnterIdent is called when production Ident is entered.
func (s *BasePatitoListener) EnterIdent(ctx *IdentContext) {}

// ExitIdent is called when production Ident is exited.
func (s *BasePatitoListener) ExitIdent(ctx *IdentContext) {}

// EnterInt is called when production Int is entered.
func (s *BasePatitoListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production Int is exited.
func (s *BasePatitoListener) ExitInt(ctx *IntContext) {}

// EnterFloat is called when production Float is entered.
func (s *BasePatitoListener) EnterFloat(ctx *FloatContext) {}

// ExitFloat is called when production Float is exited.
func (s *BasePatitoListener) ExitFloat(ctx *FloatContext) {}
