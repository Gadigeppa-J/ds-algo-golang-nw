package main

/*
Reverse linked list using stack
TC: O(n)
SC: O(n)
*/

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

func main() {

	n1 := utils.NewListNode(1)
	n2 := utils.NewListNode(2)
	n3 := utils.NewListNode(3)
	n4 := utils.NewListNode(4)
	n5 := utils.NewListNode(5)

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	rev := reverseListUsingStack(n1)
	utils.PrintLinkedList(rev)
}

func reverseStackRecursively(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseStackRecursively(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}

func reverseListUsingStack(head *utils.ListNode) *utils.ListNode {

	s := utils.NewGenericStack()

	curr := head

	for curr != nil {
		s.Push(curr)
		curr = curr.Next
	}

	var nwHead *utils.ListNode

	for !s.IsEmpty() {
		h := s.Pop().(*utils.ListNode)

		if nwHead == nil {
			nwHead = h
		} else {
			h.Next.Next = h
			h.Next = nil
		}
	}

	return nwHead
}
