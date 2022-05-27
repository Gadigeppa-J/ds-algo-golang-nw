package main

/*
https://leetcode.com/problems/linked-list-cycle-ii/
TC: O(n)
SC: O(1)
*/

import (
	"fmt"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

func main() {

	n3 := &utils.ListNode{
		Val: 3,
	}

	n2 := &utils.ListNode{
		Val: 2,
	}

	n0 := &utils.ListNode{
		Val: 0,
	}

	n04 := &utils.ListNode{
		Val: -4,
	}

	n3.Next = n2
	n2.Next = n0
	n0.Next = n04
	n04.Next = n2

	fmt.Println(detectCycle(n3))

}

func detectCycle(head *utils.ListNode) *utils.ListNode {

	m := meetPoint(head)

	if m == nil {
		return m
	}

	s := head

	for s != m {
		s = s.Next
		m = m.Next
	}

	return m
}

func meetPoint(head *utils.ListNode) *utils.ListNode {

	s := head
	f := head

	for f != nil && f.Next != nil {

		s = s.Next
		f = f.Next.Next

		if s == f {
			return s
		}
	}

	return nil
}
