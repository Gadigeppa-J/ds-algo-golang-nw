package main

/*
	31. Max Stacks
	Question:​ Implement a LIFO stack that has a push(), pop(), and max() function, where max() returns the maximum value in the stack. All of these functions should run in O(1) time.

	eg.
	push​(​1​)
	max​() = ​1
	push​(​2​)
	max​() = ​2
	push​(​1​)
	max​() = ​2
	pop​() = ​1
	max​() = ​2
	pop​() = ​2
	max​() = ​1

	Time Complexity O(1)

*/

import "fmt"

type StackElement struct {
	val int
	max int
}

type Stack struct {
	arr []StackElement
}

func (s *Stack) Push(i int) {

	max, err := s.Max()

	if err != nil || i > max {
		max = i
	}

	s.arr = append(s.arr, StackElement{i, max})
}

func (s *Stack) Pop() (int, error) {
	if len(s.arr) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}

	v := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]

	return v.val, nil
}

func (s *Stack) Max() (int, error) {
	if len(s.arr) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}

	return s.arr[len(s.arr)-1].max, nil
}
