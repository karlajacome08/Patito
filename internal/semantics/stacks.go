package semantics

import "fmt"

// -------- Pilas genéricas --------
type Stack[T any] struct{ data []T }

func (s *Stack[T]) Push(v T)    { s.data = append(s.data, v) }
func (s *Stack[T]) Empty() bool { return len(s.data) == 0 }
func (s *Stack[T]) Len() int    { return len(s.data) }
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.data) == 0 {
		return zero, false
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, true
}
func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if len(s.data) == 0 {
		return zero, false
	}
	return s.data[len(s.data)-1], true
}
func (s *Stack[T]) Clear() { s.data = s.data[:0] }

// -------- Quad queue --------
type Quad struct {
	Op     Op
	Arg1   string
	Arg2   string
	Result string
}

type QuadQueue struct{ list []Quad }

func (q *QuadQueue) Emit(op Op, a1, a2, r string) int {
	q.list = append(q.list, Quad{Op: op, Arg1: a1, Arg2: a2, Result: r})
	return len(q.list) - 1
}
func (q *QuadQueue) At(i int) Quad { return q.list[i] }
func (q *QuadQueue) Len() int      { return len(q.list) }
func (q *QuadQueue) String() string {
	out := "Idx\t(op, arg1, arg2, res)\n"
	for i, x := range q.list {
		out += fmt.Sprintf("%d\t(%s, %s, %s, %s)\n", i, x.Op.String(), x.Arg1, x.Arg2, x.Result)
	}
	return out
}

// -------- Estado semántico (pilas + quads) --------
type SemState struct {
	operands  Stack[Addr]
	operators Stack[Op]
	types     Stack[TypeTag]
	quads     QuadQueue
	tmpCount  int
}

func NewSemState() *SemState { return &SemState{} }

func (s *SemState) Quads() *QuadQueue { return &s.quads }
