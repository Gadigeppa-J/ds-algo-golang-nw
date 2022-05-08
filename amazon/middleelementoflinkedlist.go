package main

/*
https://leetcode.com/problems/middle-of-the-linked-list/
TC: O(n)
SC: O(1)
*/

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

func middleNode(head *utils.ListNode) *utils.ListNode {

	p1 := head
	p2 := head

	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}

	return p1
}
