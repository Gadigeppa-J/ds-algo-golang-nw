package main

/*
https://leetcode.com/problems/add-two-numbers/
TC: O(n)
TC: O(1)
*/

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

/*
func main() {

	/*
		l1_n2 := &utils.ListNode{
			Val: 2,
		}

		l1_n4 := &utils.ListNode{
			Val: 4,
		}

		l1_n3 := &utils.ListNode{
			Val: 3,
		}

		l1_n2.Next = l1_n4
		l1_n4.Next = l1_n3

		l2_n5 := &utils.ListNode{
			Val: 5,
		}

		l2_n6 := &utils.ListNode{
			Val: 6,
		}

		l2_n4 := &utils.ListNode{
			Val: 4,
		}

		l2_n5.Next = l2_n6
		l2_n6.Next = l2_n4
*/

/*
	l1_n0 := &utils.ListNode{
		Val: 5,
	}

	l2_n0 := &utils.ListNode{
		Val: 6,
	}

	l2_n1 := &utils.ListNode{
		Val: 6,
	}
	l2_n0.Next = l2_n1

	utils.PrintLinkedList(addTwoNumbers(l1_n0, l2_n0))
}
*/

func addTwoNumbers(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	lr1 := reverse(l1)
	lr2 := reverse(l2)

	lr1Len := listLen(lr1)
	lr2Len := listLen(lr2)

	if lr1Len < lr2Len {
		lenDiff := lr2Len - lr1Len
		addExtraNodes(lr1, lenDiff)
	} else if lr1Len > lr2Len {
		lenDiff := lr1Len - lr2Len
		addExtraNodes(lr2, lenDiff)
	}

	p1 := lr1
	p2 := lr2
	dummy := &utils.ListNode{}
	curr := dummy
	carry := 0

	for p1 != nil && p2 != nil {

		addtion := p1.Val + p2.Val + carry
		carry = addtion / 10

		if carry > 0 {
			remainder := addtion % 10
			curr.Next = &utils.ListNode{
				Val: remainder,
			}
		} else {
			curr.Next = &utils.ListNode{
				Val: addtion,
			}
		}

		curr = curr.Next
		p1 = p1.Next
		p2 = p2.Next

	}

	if carry > 0 {
		curr.Next = &utils.ListNode{
			Val: carry,
		}
	}

	return reverse(dummy.Next)

}

func reverse(head *utils.ListNode) *utils.ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	nwHead := reverse(head.Next)

	head.Next.Next = head
	head.Next = nil

	return nwHead

}

func listLen(head *utils.ListNode) int {

	count := 0

	for head != nil {
		head = head.Next
		count++
	}

	return count

}

func addExtraNodes(head *utils.ListNode, n int) {

	curr := head

	for curr.Next != nil {
		curr = curr.Next
	}

	for i := 0; i < n; i++ {
		curr.Next = &utils.ListNode{
			Val: 0,
		}
		curr = curr.Next
	}

}
