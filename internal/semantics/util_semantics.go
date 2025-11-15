package semantics

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

///////////////////////////////
// UTILIDADES NECESARIAS
///////////////////////////////

// lookup: busca una variable primero en el scope actual y luego en 'global'.
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

// mapType: mapea el texto del tipo léxico a nuestro TypeTag.
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

// MOTOR DE CUÁDRUPLOS

// Precedencias (mayor número = mayor precedencia)
var prec = map[Op]int{
	OOr:  0,
	OAnd: 1,
	OEq:  2, ONEq: 2, OLt: 2, OLe: 2, OGt: 2, OGe: 2, // relacionales
	OAdd: 3, OSub: 3,
	OMul: 4, ODiv: 4,
	OUMinus: 5, ONot: 5,
}

func isUnary(op Op) bool { return op == OUMinus || op == ONot }

// Generador de temporales t1, t2, ...
func (s *SemState) newTemp(t TypeTag) Addr {
	s.tmpCount++
	return Addr{Name: fmt.Sprintf("t%d", s.tmpCount), Type: t}
}

// Empuja un operando y su tipo (id, cte o temp)
func (s *SemState) PushOperand(name string, t TypeTag) {
	addr := Addr{Name: name, Type: t}
	s.operands.Push(addr)
	s.types.Push(t)
}

// Empuja operador con reducciones por precedencia
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

// Cierra una expresión (antes de ';', ')' o ',')
func (s *SemState) EndExpression() {
	for !s.operators.Empty() {
		top, _ := s.operators.Peek()
		if top == OParen {
			break
		}
		s.reduceOnce()
	}
}

// Reducción de un solo operador del tope (usa el cubo semántico)
func (s *SemState) reduceOnce() {
	op, ok := s.operators.Pop()
	if !ok {
		return // nada que reducir
	}

	// UNARIO
	if isUnary(op) {
		x, ok := s.operands.Pop()
		if !ok {
			panic(fmt.Errorf("unario %s sin operando", op.String()))
		}
		_, _ = s.types.Pop() // mantener alineada la pila de tipos

		tRes, ok := ResultTypeUnary(op, x.Type)
		if !ok {
			panic(fmt.Errorf("tipo inválido: %s %v", op.String(), x.Type))
		}
		tmp := s.newTemp(tRes)
		s.quads.Emit(op, x.Name, "", tmp.Name)
		s.operands.Push(tmp)
		s.types.Push(tmp.Type)
		return
	}

	// BINARIO
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
	s.quads.Emit(op, ra.Name, rb.Name, tmp.Name)
	s.operands.Push(tmp)
	s.types.Push(tmp.Type)
}

// ===== Estatutos lineales =====

// Asignación: LHS = (expr)
func (s *SemState) DoAssign(lhsName string, lhsType TypeTag) {
	s.EndExpression()
	rhs, ok := s.operands.Pop()
	if !ok {
		panic(fmt.Errorf("asignación sin RHS para %s", lhsName))
	}
	_, _ = s.types.Pop()

	if !IsCompatibleAssign(lhsType, rhs.Type) {
		panic(fmt.Errorf("asignación incompatible: %v ← %v", lhsType, rhs.Type))
	}
	s.quads.Emit(OAssign, rhs.Name, "", lhsName)
}

// write(expr);
func (s *SemState) DoWriteExpr() {
	s.EndExpression()
	val, ok := s.operands.Pop()
	if !ok {
		panic(fmt.Errorf("write sin expresión"))
	}
	_, _ = s.types.Pop()
	s.quads.Emit(OWrite, val.Name, "", "")
}

func (s *SemState) DoReadID(idName string) {
	s.quads.Emit(ORead, "", "", idName)
}

func (s *SemState) OnId(name string, t TypeTag) { s.PushOperand(name, t) }

func (s *SemState) OnIntConst(lit string)   { s.PushOperand(lit, TInt) }
func (s *SemState) OnFloatConst(lit string) { s.PushOperand(lit, TFloat) }
func (s *SemState) OnBoolConst(lit string)  { s.PushOperand(lit, TBool) }

func (s *SemState) OnAdd()    { s.PushOperator(OAdd) }
func (s *SemState) OnSub()    { s.PushOperator(OSub) }
func (s *SemState) OnMul()    { s.PushOperator(OMul) }
func (s *SemState) OnDiv()    { s.PushOperator(ODiv) }
func (s *SemState) OnLt()     { s.PushOperator(OLt) }
func (s *SemState) OnLe()     { s.PushOperator(OLe) }
func (s *SemState) OnGt()     { s.PushOperator(OGt) }
func (s *SemState) OnGe()     { s.PushOperator(OGe) }
func (s *SemState) OnEq()     { s.PushOperator(OEq) }
func (s *SemState) OnNeq()    { s.PushOperator(ONEq) }
func (s *SemState) OnAnd()    { s.PushOperator(OAnd) }
func (s *SemState) OnOr()     { s.PushOperator(OOr) }
func (s *SemState) OnUMinus() { s.PushOperator(OUMinus) }
func (s *SemState) OnNot()    { s.PushOperator(ONot) }

func (s *SemState) OnLParen()  { s.OpenParen() }
func (s *SemState) OnRParen()  { s.CloseParen() }
func (s *SemState) OnEndExpr() { s.EndExpression() }

// Estatutos
func (s *SemState) OnAssign(lhsName string, lhsType TypeTag) { s.DoAssign(lhsName, lhsType) }
func (s *SemState) OnRead(idName string)                     { s.DoReadID(idName) }
func (s *SemState) OnWriteExpr()                             { s.DoWriteExpr() }

func ifElse[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

// Empujar string literal como operando; luego OnWriteExpr lo imprime
func (s *SemState) OnStringConst(lit string) { s.PushOperand(lit, TString) }
