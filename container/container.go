package container

type Stack struct {
	arr []int
}

func (s *Stack) push(v int) {
	s.arr = append(s.arr, v)
}

func (s *Stack) pop() int {
	res := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return res
}

func (s *Stack) isEmpty() bool {
	return len(s.arr) == 0
}

type Queue struct {
	arr []int
}

func (q *Queue) push(v int) {
	q.arr = append(q.arr, v)
}

func (q *Queue) pop() int {
	res := q.arr[0]
	q.arr = q.arr[1:]
	return res
}

func (q *Queue) isEmpty() bool {
	return len(q.arr) == 0
}
