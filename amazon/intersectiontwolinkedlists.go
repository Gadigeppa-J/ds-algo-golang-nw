package main

/*
https://leetcode.com/problems/intersection-of-two-linked-lists/
TC: O(n)
SC: O(1)
*/

import (
	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

/*
func main() {

	//[4,1,8,4,5]
	//[5,6,1,8,4,5]

	n5 := &utils.ListNode{
		Val:  5,
		Next: nil,
	}

	n4 := &utils.ListNode{
		Val:  4,
		Next: n5,
	}

	n3 := &utils.ListNode{
		Val:  8,
		Next: n4,
	}

	n2 := &utils.ListNode{
		Val:  1,
		Next: n3,
	}

	n1 := &utils.ListNode{
		Val:  4,
		Next: n2,
	}

	n8 := &utils.ListNode{
		Val:  1,
		Next: n3,
	}

	n7 := &utils.ListNode{
		Val:  6,
		Next: n8,
	}

	n6 := &utils.ListNode{
		Val:  5,
		Next: n7,
	}

	fmt.Println(getIntersectionNode(n1, n6))

}
*/

func getIntersectionNode(headA, headB *utils.ListNode) *utils.ListNode {

	h1Ptr := headA
	h2Ptr := headB

	for true {

		if h1Ptr == nil && h2Ptr == nil {
			break
		}

		if h1Ptr == h2Ptr {
			return h1Ptr
		}

		if h1Ptr == nil {
			h1Ptr = headB
		} else {
			h1Ptr = h1Ptr.Next
		}

		if h2Ptr == nil {
			h2Ptr = headA
		} else {
			h2Ptr = h2Ptr.Next
		}

	}

	return nil
}
