package main

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

type MyQueue struct {
	s1 *utils.Stack
	s2 *utils.Stack
}

func Constructor() MyQueue {
	return MyQueue{
		s1: utils.NewStack(),
		s2: utils.NewStack(),
	}
}

func (this *MyQueue) Push(x int) {
	this.s1.Push(x)
}

func (this *MyQueue) Pop() int {
	if this.s2.IsEmpty() {
		for !this.s1.IsEmpty() {
			this.s2.Push(this.s1.Pop())
		}
	}
	return this.s2.Pop()
}

func (this *MyQueue) Peek() int {
	if this.s2.IsEmpty() {
		for !this.s1.IsEmpty() {
			this.s2.Push(this.s1.Pop())
		}
	}

	return this.s2.Peek()
}

func (this *MyQueue) Empty() bool {
	if this.s2.IsEmpty() {
		for !this.s1.IsEmpty() {
			this.s2.Push(this.s1.Pop())
		}
	}

	return this.s2.IsEmpty()
}
