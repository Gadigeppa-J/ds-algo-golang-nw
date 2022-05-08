package main

/*
https://leetcode.com/problems/reverse-linked-list/
TC: O(n)
SC: O(1)
*/

import (
	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

func reverseList(head *utils.ListNode) *utils.ListNode {

	var prev *utils.ListNode = nil
	curr := head

	for curr != nil {
		temp := curr.Next
		curr.Next = prev

		prev = curr
		curr = temp
	}

	return prev

}
