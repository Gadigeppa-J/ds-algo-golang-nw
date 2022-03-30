package main

import (
	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

/*
42. Nth to the Last Element
Question:​ Given a linked list, and an input n, write a function that returns the nth-to-last element of the linked list.

Time Complexity: O(n)
Space Complexity: O(1)
*/

func nthToLast(l *utils.ListNode, index int) int {

	p1 := l
	p2 := l

	for i := 0; i < index; i++ {
		p2 = p2.Next
		if p2 == nil {
			panic("input Index out of bound")
		}
	}

	for p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p1.Val
}
