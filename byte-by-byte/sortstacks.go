package main

/*
28. Sort Stacks
Question:​ Given a stack, sort the elements in the stack using one additional stack. eg.

sortStack([​1​, ​3​, ​2​, ​4​]) = [​1​, ​2​, ​3​, ​4​]

Time Complexity: O(n^2)
Space Complexity: O(n)

*/

import (
	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

func sortStack(s *utils.Stack) *utils.Stack {

	if s == nil || s.IsEmpty() {
		return s
	}

	o := utils.NewStack()

	for !s.IsEmpty() {

		if o.IsEmpty() {
			o.Push(s.Pop())
		} else {

			temp := s.Pop()

			for true {

				if !o.IsEmpty() && o.Peek() > temp {
					s.Push(o.Pop())
				} else {
					o.Push(temp)
					break
				}
			}

		}

	}

	return o
}
