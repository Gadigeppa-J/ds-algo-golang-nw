package main

import (
	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

func divide(ll *utils.ListNode) (*utils.ListNode, *utils.ListNode) {

	if ll == nil || ll.Next == nil {
		return ll, nil
	}

	slow := ll
	fast := ll.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	ll2 := slow.Next
	slow.Next = nil

	return ll, ll2
}
