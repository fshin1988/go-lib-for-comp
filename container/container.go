package container

type Stack struct {
	values []int
}

func (s *Stack) push(v int) {
	s.values = append(s.values, v)
}

func (s *Stack) pop() int {
	res := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return res
}
