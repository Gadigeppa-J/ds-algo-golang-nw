package main

/*
Remove Loop From Linked List
no leetcode
*/

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

func main() {

	n1 := &utils.ListNode{
		Val: 1,
	}

	n2 := &utils.ListNode{
		Val: 2,
	}

	n3 := &utils.ListNode{
		Val: 3,
	}

	n4 := &utils.ListNode{
		Val: 4,
	}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	//n4.Next = n1

	removeLoop(n1)

	utils.PrintLinkedList(n1)
}

func removeLoop(head *utils.ListNode) {

	if mp, prev := meetPoint(head); mp != nil {
		sp := head

		for true {

			if sp == mp {
				prev.Next = nil
				return
			}

			prev = mp
			sp = sp.Next
			mp = mp.Next
		}

	}
}

func meetPoint(head *utils.ListNode) (*utils.ListNode, *utils.ListNode) {

	s := head
	f := head
	var p *utils.ListNode

	for f != nil && f.Next != nil {
		p = s
		s = s.Next
		f = f.Next.Next
		if s == f {
			return s, p
		}
	}

	return nil, nil
}
