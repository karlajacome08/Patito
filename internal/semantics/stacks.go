package semantics

type Stack[T any] struct{ data []T }

func (s *Stack[T]) Push(v T)    { s.data = append(s.data, v) }
func (s *Stack[T]) Empty() bool { return len(s.data) == 0 }
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
