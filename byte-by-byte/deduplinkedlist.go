package main

/*
40. Dedup Linked List
Question:​ Given an unsorted linked list, write a function to remove all the duplicates.

eg.
dedup(​1​ -> ​2​ -> ​3​ -> ​2​ -> ​1​) = ​1​ -> ​2​ -> ​3

Time Complexity: O(n)
Space Complexity: O(n)

*/

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

func dedup(head *utils.ListNode) *utils.ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	m := make(map[int]struct{})

	curr := head
	var prev *utils.ListNode

	for curr != nil {

		if _, ok := m[curr.Val]; ok {
			prev.Next = curr.Next
			curr.Next = nil
			curr = prev.Next
		} else {
			m[curr.Val] = struct{}{}
			temp := curr.Next
			prev = curr
			curr = temp
		}

	}

	return head
}
