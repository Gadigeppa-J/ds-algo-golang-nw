package main

/*
https://leetcode.com/problems/reverse-linked-list-ii/
TC: O(n)
SC: O(n)
*/

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

/*
func main() {

	/*
		n1 := &utils.ListNode{
			Val: 1,
		}

		n2 := &utils.ListNode{
			Val: 2,
		}

*/

/*
	n3 := &utils.ListNode{
		Val: 3,
	}

	/*
		n4 := &utils.ListNode{
			Val: 4,
		}
*/

/*
	n5 := &utils.ListNode{
		Val: 5,
	}

	//n1.Next = n2
	//n2.Next = n3
	//n3.Next = n4
	//n4.Next = n5

	n3.Next = n5

	utils.PrintLinkedList(reverseBetween(n3, 1, 2))

}
*/

func reverseBetween(head *utils.ListNode, left int, right int) *utils.ListNode {

	if right-left < 1 {
		return head
	}

	prev := &utils.ListNode{
		Next: head,
	}

	dummy := prev
	curr := head

	for i := 1; i < left; i++ {
		prev = curr
		curr = curr.Next
	}

	prev.Next = nil
	nxt := curr

	for i := 0; i < (right - left); i++ {
		nxt = nxt.Next
	}

	temp := nxt
	nxt = nxt.Next

	temp.Next = nil

	curr = reveseLinkedList(curr)
	prev.Next = curr

	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = nxt

	return dummy.Next

}

func reveseLinkedList(head *utils.ListNode) *utils.ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	nwHead := reveseLinkedList(head.Next)

	head.Next.Next = head
	head.Next = nil

	return nwHead
}
