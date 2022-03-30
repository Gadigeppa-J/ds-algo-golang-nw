package main

/*
38. Linked List Cycles
Question:​ Given a linked list, determine whether it contains a cycle.

https://leetcode.com/problems/linked-list-cycle/

Time Complexity: O(n)
Space Complexity: O(1)

*/

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

func hasCycle(head *utils.ListNode) bool {

	fast := head
	slow := head

	for fast != nil && fast.Next != nil {

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}

	}

	return false
}
