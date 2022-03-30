package main

/*
23. Print Reversed Linked List
Question:​ Given a linked list, write a function that prints the nodes of the list in reverse order.

reverseList(​1​ -> ​2​ -> ​3​) 3 -> 2 -> 1

https://leetcode.com/problems/reverse-linked-list/

Time complexity: O(n)
Space complexity: O(1)

*/

import (
	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

func reverseList(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	curr := head
	var prev *utils.ListNode

	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	return prev
}
