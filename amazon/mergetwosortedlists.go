package main

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

func main() {
	pn3 := &utils.ListNode{
		Val:  4,
		Next: nil,
	}

	pn2 := &utils.ListNode{
		Val:  2,
		Next: pn3,
	}

	pn1 := &utils.ListNode{
		Val:  1,
		Next: pn2,
	}

	qn3 := &utils.ListNode{
		Val:  4,
		Next: nil,
	}

	qn2 := &utils.ListNode{
		Val:  3,
		Next: qn3,
	}

	qn1 := &utils.ListNode{
		Val:  1,
		Next: qn2,
	}

	utils.PrintLinkedList(mergeTwoLists(pn1, qn1))

}

func mergeTwoLists(list1 *utils.ListNode, list2 *utils.ListNode) *utils.ListNode {

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var ret *utils.ListNode
	var s *utils.ListNode
	var p *utils.ListNode
	var q *utils.ListNode

	if list1.Val <= list2.Val {
		s = list1
		p = list1.Next
		q = list2
	} else {
		s = list2
		p = list1
		q = list2.Next
	}
	ret = s
	for true {

		if p.Val <= q.Val {
			s.Next = p
			s = s.Next
			p = p.Next
		} else {
			s.Next = q
			s = s.Next
			q = q.Next
		}

		if p == nil {
			s.Next = q
			break
		}

		if q == nil {
			s.Next = p
			break
		}

	}

	return ret

}
