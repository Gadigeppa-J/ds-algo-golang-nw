package utils

import "fmt"

func PrintLinkedList(h *ListNode) {

	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}

}

func CreateLinkedList(v ...int) *ListNode {

	var head *ListNode
	var prev *ListNode

	for _, i := range v {

		curr := &ListNode{
			Val:  i,
			Next: nil,
		}

		if head == nil {
			head = curr
		}

		if prev != nil {
			prev.Next = curr
		}

		prev = curr

	}

	return head
}

func Max(a, b int) int {

	if a > b {
		return a
	}

	return b

}

func Min(a, b int) int {

	if a < b {
		return a
	}

	return b

}
