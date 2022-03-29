package main

import "fmt"

func reverseStack(s *stack) {

	if s.isEmpty() {
		return
	}

	v, err := s.pop()

	if err != nil {
		panic(err)
	}

	reverseStack(s)
	pushToBottom(s, v)

}

func pushToBottom(s *stack, i int) {

	if s.isEmpty() {
		s.push(i)
		return
	}

	v, err := s.pop()

	if err != nil {
		panic(err)
	}

	pushToBottom(s, i)

	s.push(v)

}

// stack implementation
type stack struct {
	arr []int
}

func NewStack() *stack {
	return &stack{
		arr: []int{},
	}
}

func (s *stack) push(i int) {
	s.arr = append(s.arr, i)
}

func (s *stack) pop() (int, error) {

	if s.isEmpty() {
		return 0, fmt.Errorf("stack is empty!")
	}

	i := s.arr[len(s.arr)-1]

	s.arr = s.arr[:len(s.arr)-1]

	return i, nil
}

func (s *stack) isEmpty() bool {

	if len(s.arr) == 0 {
		return true
	}

	return false
}
