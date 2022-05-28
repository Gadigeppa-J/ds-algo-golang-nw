package main

/*
Add One to Linked List
TC: O(n)
SC: O(n)
*/

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

/*
func main() {

	n1 := &utils.ListNode{
		Val: 8,
	}

	n2 := &utils.ListNode{
		Val: 8,
	}

	n3 := &utils.ListNode{
		Val: 9,
	}

	n1.Next = n2
	n2.Next = n3

	utils.PrintLinkedList(addOne(n1))

}
*/

func addOne(head *utils.ListNode) *utils.ListNode {

	if head == nil {
		return head
	}

	reversed := reverseListRecursive(head)

	curr := reversed

	for true {

		if curr.Val+1 >= 10 {
			curr.Val = 0
			if curr.Next == nil {
				curr.Next = &utils.ListNode{
					Val: 1,
				}
				break
			}
			curr = curr.Next
		} else {
			curr.Val += 1
			break
		}

	}

	return reverseListRecursive(reversed)

}

func reverseListRecursive(head *utils.ListNode) *utils.ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	nwHead := reverseListRecursive(head.Next)

	head.Next.Next = head
	head.Next = nil

	return nwHead
}
