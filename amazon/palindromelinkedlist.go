package main

/*
https://leetcode.com/problems/palindrome-linked-list/
TC: O(n)
SC: O(1)
*/

import (
	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

/*
func main() {

	n3 := &utils.ListNode{
		Val:  0,
		Next: nil,
	}

	n2 := &utils.ListNode{
		Val:  0,
		Next: n3,
	}

	n1 := &utils.ListNode{
		Val:  1,
		Next: n2,
	}

	fmt.Println(isPalindrome(n1))


}
*/

func isPalindrome(head *utils.ListNode) bool {

	if head == nil || head.Next == nil {
		return true
	}

	p := head
	q := head

	list1 := head
	var list2 *utils.ListNode

	// find middle
	for true {

		if q.Next == nil || q.Next.Next == nil { // odd
			list2 = p.Next
			break
		}

		p = p.Next
		q = q.Next.Next
	}

	// Reverse list2
	var prev *utils.ListNode
	curr := list2

	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	for prev != nil {
		if list1.Val != prev.Val {
			return false
		}
		list1 = list1.Next
		prev = prev.Next

	}

	return true
}
