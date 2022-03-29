package main

/*
29. Stack from Queues
Question:​ Implement a LIFO stack with basic functionality (push and pop) using FIFO queues to store the data.

Push operation will take O(n) time complexity

*/

import (
	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

type StackFromQueue struct {
	p *utils.Queue
	s *utils.Queue
}

func NewStackFromQueue() *StackFromQueue {
	return &StackFromQueue{
		p: utils.NewQueue(),
		s: utils.NewQueue(),
	}
}

func (s *StackFromQueue) swapRefs() {
	temp := s.p
	s.p = s.s
	s.s = temp
}

func (s *StackFromQueue) Push(i int) {

	s.s.Enqueue(i)

	for !s.p.IsEmpty() {
		s.s.Enqueue(s.p.Dequeue())
	}

	s.swapRefs()
}

func (s *StackFromQueue) Pop() int {
	if s.p.IsEmpty() {
		panic("Stack is empty!")
	}

	return s.p.Dequeue()
}

func (s *StackFromQueue) Top() int {
	if s.p.IsEmpty() {
		panic("Stack is empty!")
	}
	return s.p.Peek()
}

func (s *StackFromQueue) IsEmpty() bool {
	if s.p.IsEmpty() {
		return true
	}
	return false
}
