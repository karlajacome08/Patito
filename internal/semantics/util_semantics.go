package semantics

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

func (s *SemanticListener) lookup(name string) (VariableInfo, bool) {
	if f, ok := s.FD.Get(s.curScope()); ok {
		if v, ok := f.Locals.Lookup(name); ok {
			return v, true
		}
	}
	if g, ok := s.FD.Get("global"); ok {
		if v, ok := g.Locals.Lookup(name); ok {
			return v, true
		}
	}
	return VariableInfo{}, false
}

func mapType(tctx antlr.ParserRuleContext) TypeTag {
	if tctx == nil {
		return TVoid
	}
	txt := strings.ToLower(tctx.GetText())
	switch txt {
	case "entero":
		return TInt
	case "flotante":
		return TFloat
	case "booleano", "bool", "logico", "lógico":
		return TBool
	default:
		return TVoid
	}
}

var prec = map[Op]int{
	OEq: 1, ONEq: 1, OLt: 1, OGt: 1,
	OAdd: 2, OSub: 2,
	OMul: 3, ODiv: 3}

func (s *SemState) newTemp(t TypeTag) Addr {
	s.tmpCount++
	addr := s.mem.AllocTemp(t)
	return Addr{
		Name:    fmt.Sprintf("t%d", s.tmpCount),
		Type:    t,
		Address: addr,
	}
}

func (s *SemState) PushOperand(name string, t TypeTag, addr int) {
	a := Addr{
		Name:    name,
		Type:    t,
		Address: addr,
	}
	s.operands.Push(a)
	s.types.Push(t)
}

func (s *SemState) PushOperator(op Op) {
	for !s.operators.Empty() {
		top, ok := s.operators.Peek()
		if !ok || top == OParen || prec[top] < prec[op] {
			break
		}
		s.reduceOnce()
	}
	s.operators.Push(op)
}

func (s *SemState) OpenParen() { s.operators.Push(OParen) }

func (s *SemState) CloseParen() {
	for !s.operators.Empty() {
		top, _ := s.operators.Peek()
		if top == OParen {
			break
		}
		s.reduceOnce()
	}
	if !s.operators.Empty() {
		if t, _ := s.operators.Peek(); t == OParen {
			s.operators.Pop()
		}
	}
}

func (s *SemState) EndExpression() {
	for !s.operators.Empty() {
		top, _ := s.operators.Peek()
		if top == OParen {
			break
		}
		s.reduceOnce()
	}
}

func (s *SemState) reduceOnce() {
	op, ok := s.operators.Pop()
	if !ok {
		return
	}
	rb, ok := s.operands.Pop()
	if !ok {
		panic(fmt.Errorf("operador %s sin RHS", op.String()))
	}
	tb, _ := s.types.Pop()

	ra, ok := s.operands.Pop()
	if !ok {
		panic(fmt.Errorf("operador %s sin LHS", op.String()))
	}
	ta, _ := s.types.Pop()

	tRes, ok := ResultType(op, ta, tb)
	if !ok {
		panic(fmt.Errorf("tipos inválidos: %v %s %v", ta, op.String(), tb))
	}
	tmp := s.newTemp(tRes)
	s.quads.Emit(op, ToAddrString(ra.Address), ToAddrString(rb.Address), ToAddrString(tmp.Address))
	s.operands.Push(tmp)
	s.types.Push(tmp.Type)
}

func (s *SemState) DoAssign(lhsAddr int, lhsType TypeTag) {
	s.EndExpression()
	rhs, ok := s.operands.Pop()
	if !ok {
		panic(fmt.Errorf("asignación sin RHS"))
	}
	_, _ = s.types.Pop()

	if !IsCompatibleAssign(lhsType, rhs.Type) {
		panic(fmt.Errorf("asignación incompatible: %v ← %v", lhsType, rhs.Type))
	}

	s.quads.Emit(
		OAssign,
		ToAddrString(rhs.Address),
		"",
		ToAddrString(lhsAddr),
	)
}

func (s *SemState) DoWriteExpr() {
	s.EndExpression()
	val, ok := s.operands.Pop()
	if !ok {
		panic(fmt.Errorf("write sin expresión"))
	}
	_, _ = s.types.Pop()

	arg := val.Name
	if val.Address > 0 {
		arg = ToAddrString(val.Address)
	}
	s.quads.Emit(OWrite, arg, "", "")
}

func (s *SemState) DoReadID(idAddr int) {
	s.quads.Emit(ORead, "", "", ToAddrString(idAddr))
}

func (s *SemState) OnId(name string, t TypeTag, addr int) {
	s.PushOperand(name, t, addr)
}

func (s *SemState) OnIntConst(lit string) {
	addr := s.mem.AllocConstInt(lit)
	s.PushOperand(lit, TInt, addr)
}

func (s *SemState) OnFloatConst(lit string) {
	addr := s.mem.AllocConstFloat(lit)
	s.PushOperand(lit, TFloat, addr)
}

func (s *SemState) OnBoolConst(lit string) {
	s.PushOperand(lit, TBool, -1)
}

func (s *SemState) OnStringConst(lit string) {
	s.PushOperand(lit, TString, -1)
}

// Operadores
func (s *SemState) OnAdd() { s.PushOperator(OAdd) }
func (s *SemState) OnSub() { s.PushOperator(OSub) }
func (s *SemState) OnMul() { s.PushOperator(OMul) }
func (s *SemState) OnDiv() { s.PushOperator(ODiv) }
func (s *SemState) OnLt()  { s.PushOperator(OLt) }
func (s *SemState) OnGt()  { s.PushOperator(OGt) }
func (s *SemState) OnEq()  { s.PushOperator(OEq) }
func (s *SemState) OnNeq() { s.PushOperator(ONEq) }

func (s *SemState) OnLParen()  { s.OpenParen() }
func (s *SemState) OnRParen()  { s.CloseParen() }
func (s *SemState) OnEndExpr() { s.EndExpression() }

func (s *SemState) OnAssign(lhsAddr int, lhsType TypeTag) { s.DoAssign(lhsAddr, lhsType) }
func (s *SemState) OnRead(idAddr int)                     { s.DoReadID(idAddr) }
func (s *SemState) OnWriteExpr()                          { s.DoWriteExpr() }

func ifElse[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

func (s *SemState) EmitERA(funcName string) { s.quads.Emit(OEra, funcName, "", "") }
func (s *SemState) EmitParam(arg Addr, index int) {
	s.quads.Emit(OParam, ToAddrString(arg.Address), "", fmt.Sprintf("%d", index))
}

func (s *SemanticListener) EmitGoSub(funcName string, startIndex int) {
	s.S.Quads().Emit(OGoSub, funcName, "", fmt.Sprintf("%d", startIndex))
}

func (s *SemanticListener) EmitEndFunc() { s.S.Quads().Emit(OEndFunc, "", "", "") }
