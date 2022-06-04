package main

/*
https://leetcode.com/problems/reorder-list/
TC: O(n)
SC: O(1)
*/

import (
	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

func main() {

	n1 := &utils.ListNode{Val: 1}
	//n2 := &utils.ListNode{Val: 2}
	//n3 := &utils.ListNode{Val: 3}
	//n4 := &utils.ListNode{Val: 4}
	//n5 := &utils.ListNode{Val: 5}

	//n1.Next = n2
	//n2.Next = n3
	//n3.Next = n4
	//n4.Next = n5

	reorderList(nil)

	utils.PrintLinkedList(n1)

}

func reorderList(head *utils.ListNode) {

	if head == nil || head.Next == nil {
		return
	}

	mid := findMidNode(head)
	h1 := head
	h2 := mid.Next
	mid.Next = nil
	h2 = reverseLL(h2)
	mergeLL(h1, h2)
}

func reverseLL(head *utils.ListNode) *utils.ListNode {

	var prev *utils.ListNode
	curr := head

	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	return prev
}

func findMidNode(head *utils.ListNode) *utils.ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	s := head
	f := head.Next

	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}

	return s
}

func mergeLL(h1 *utils.ListNode, h2 *utils.ListNode) {

	p1 := h1
	p2 := h2

	for p1 != nil && p2 != nil {

		tmp1 := p1.Next
		tmp2 := p2.Next

		p1.Next = p2
		p2.Next = tmp1

		p1 = tmp1
		p2 = tmp2
	}

}
