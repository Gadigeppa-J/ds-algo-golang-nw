package utils

// stack implementation
type Stack struct {
	arr []int
}

func NewStack() *Stack {
	return &Stack{
		arr: []int{},
	}
}

func (s *Stack) Push(i int) {
	s.arr = append(s.arr, i)
}

func (s *Stack) Pop() int {

	if s.IsEmpty() {
		panic("stack is empty!")
	}

	i := s.arr[len(s.arr)-1]

	s.arr = s.arr[:len(s.arr)-1]

	return i
}

func (s *Stack) IsEmpty() bool {

	if len(s.arr) == 0 {
		return true
	}

	return false
}

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		panic("stack is empty!")
	}
	return s.arr[len(s.arr)-1]
}
