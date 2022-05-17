package main

/*
https://leetcode.com/problems/copy-list-with-random-pointer/
TC: O(n)

SC: (n) PrintRandomLinkedList
SC (1) cloneRandomLinkedListWithoutExtraSpace if we don't consider the space consumed by resultant cloned list

*/

import (
	"fmt"
)

type RandomListNode struct {
	Val    int
	Next   *RandomListNode
	Random *RandomListNode
}

func PrintRandomLinkedList(h *RandomListNode) {

	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}

}

func PrintRandomLinkedListAddr(h *RandomListNode) {

	for h != nil {
		fmt.Printf("%p\n", h)
		h = h.Next
	}

}

func main() {

	n5 := &RandomListNode{
		Val: 5,
	}

	n4 := &RandomListNode{
		Val:  4,
		Next: n5,
	}

	n3 := &RandomListNode{
		Val:  3,
		Next: n4,
	}

	n2 := &RandomListNode{
		Val:    2,
		Random: n4,
		Next:   n3,
	}

	n1 := &RandomListNode{
		Val:  1,
		Next: n2,
	}

	PrintRandomLinkedList(cloneRandomLinkedListWithoutExtraSpace(n1))

}

func cloneRandomLinkedList(head *RandomListNode) *RandomListNode {

	curr := head
	var nprev *RandomListNode
	var ncurr *RandomListNode
	var result *RandomListNode
	m := make(map[*RandomListNode]*RandomListNode)

	for curr != nil {
		ncurr = &RandomListNode{
			Val: curr.Val,
		}
		m[curr] = ncurr

		if nprev != nil {
			nprev.Next = ncurr
		} else {
			result = ncurr
		}

		curr = curr.Next
		nprev = ncurr
	}

	curr = head
	ncurr = result

	for curr != nil {
		if curr.Random != nil {
			r := m[curr.Random]
			ncurr.Random = r
		}
		curr = curr.Next
		ncurr = ncurr.Next
	}
	return result

}

func cloneRandomLinkedListWithoutExtraSpace(head *RandomListNode) *RandomListNode {

	curr := head

	for curr != nil {

		node := &RandomListNode{
			Val: curr.Val,
		}
		temp := curr.Next
		curr.Next = node
		node.Next = temp
		curr = temp
	}

	curr = head

	for curr != nil {
		if curr.Random != nil {
			curr.Next.Random = curr.Random.Next
		}
		curr = curr.Next.Next
	}

	curr = head
	var nwHead *RandomListNode

	for curr != nil && curr.Next != nil {

		if nwHead == nil {
			nwHead = curr.Next
		}
		tmp := curr.Next
		curr.Next = curr.Next.Next
		curr = tmp
	}

	return nwHead
}
