package main

/*
https://leetcode.com/problems/remove-nth-node-from-end-of-list/
TC: O(n)
SC: O(1)
*/

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

func removeNthFromEnd(head *utils.ListNode, n int) *utils.ListNode {

	if head == nil {
		return head
	}

	p := head
	q := head

	for i := 0; i < n; i++ {
		q = q.Next
	}

	if q == nil {
		head = head.Next
		return head
	}

	for q.Next != nil {
		p = p.Next
		q = q.Next
	}

	p.Next = p.Next.Next
	return head

}
