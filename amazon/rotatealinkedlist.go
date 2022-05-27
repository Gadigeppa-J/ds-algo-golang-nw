package main

/*
https://leetcode.com/problems/rotate-list/
TC: O(n)
SC: O(1)
*/

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

func rotateRight(head *utils.ListNode, k int) *utils.ListNode {

	if head == nil || head.Next == nil || k <= 0 {
		return head
	}

	// cal length
	leng := 0

	curr := head

	for curr != nil {
		leng++
		curr = curr.Next
	}

	k = k % leng

	if k == 0 {
		return head
	}

	curr = head

	for i := 0; i < (leng - k - 1); i++ {
		curr = curr.Next
	}

	nwHead := curr.Next
	curr.Next = nil
	curr = nwHead

	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = head

	return nwHead

}
