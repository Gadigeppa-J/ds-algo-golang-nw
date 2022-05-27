package main

/*
Flatten a Multi-Level Linked List
TC: O(n)
SC: O(n)
*/

import "fmt"

type MultiListNode struct {
	Val    int
	Next   *MultiListNode
	Bottom *MultiListNode
}

func main() {

	n5 := &MultiListNode{
		Val: 5,
	}

	n7 := &MultiListNode{
		Val: 7,
	}

	n8 := &MultiListNode{
		Val: 8,
	}

	n30 := &MultiListNode{
		Val: 30,
	}

	n5.Bottom = n7
	n7.Bottom = n8
	n8.Bottom = n30

	n10 := &MultiListNode{
		Val: 10,
	}

	n20 := &MultiListNode{
		Val: 20,
	}

	n10.Bottom = n20

	n19 := &MultiListNode{
		Val: 19,
	}

	n22 := &MultiListNode{
		Val: 22,
	}

	n50 := &MultiListNode{
		Val: 50,
	}

	n28 := &MultiListNode{
		Val: 28,
	}

	n35 := &MultiListNode{
		Val: 35,
	}

	n40 := &MultiListNode{
		Val: 40,
	}

	n45 := &MultiListNode{
		Val: 45,
	}

	n19.Bottom = n22
	n22.Bottom = n50

	n28.Bottom = n35
	n35.Bottom = n40
	n40.Bottom = n45

	n5.Next = n10
	n10.Next = n19
	n19.Next = n28

	res := flattenMultiLevelLinkedList(n5)

	for res != nil {
		fmt.Println(res.Val)
		res = res.Bottom
	}

}

func flattenMultiLevelLinkedList(head *MultiListNode) *MultiListNode {

	if head == nil {
		return nil
	}

	res := flattenMultiLevelLinkedList(head.Next)

	return mergeMultiListNode(head, res)
}

func mergeMultiListNode(l1, l2 *MultiListNode) *MultiListNode {

	res := &MultiListNode{} // dummy node
	tmp := res

	p1 := l1
	p2 := l2

	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			tmp.Bottom = p1
			tmp = tmp.Bottom
			p1 = p1.Bottom
		} else {
			tmp.Bottom = p2
			tmp = tmp.Bottom
			p2 = p2.Bottom
		}
	}

	if p1 == nil {
		tmp.Bottom = p2
	} else {
		tmp.Bottom = p1
	}

	if res.Bottom != nil {
		res.Bottom.Next = nil
	}

	return res.Bottom
}
