package main

/*
22. Longest Consecutive Branch
Question:​ Given a tree, write a function to find the length of the longest branch of nodes in increasing consecutive order.

https://www.lintcode.com/problem/595/description

Time Complexity: O(n)
Space Complexity: O(h)

*/

import (
	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

/*
func main() {

	n0 := &utils.TreeNode{
		Val: 0,

		Left: &utils.TreeNode{
			Val: 1,
			Left: &utils.TreeNode{
				Val: 1,
			},
			Right: &utils.TreeNode{
				Val: 2,
			},
		},

		Right: &utils.TreeNode{
			Val: 2,
			Left: &utils.TreeNode{
				Val: 1,
			},
			Right: &utils.TreeNode{
				Val: 3,
			},
		},
	}

	fmt.Println(longestConsecutiveBranch(n0))

}
*/

func longestConsecutiveBranch(root *utils.TreeNode) int {
	l := longest(root.Left, root, 1)
	r := longest(root.Right, root, 1)
	return max(l, r)
}

func longest(root *utils.TreeNode, prev *utils.TreeNode, len int) int {

	if root == nil {
		return len
	}

	if root.Val == (prev.Val + 1) {
		l := longest(root.Left, root, len+1)
		r := longest(root.Right, root, len+1)
		return max(l, r)
	} else {
		l := longest(root.Left, root, 1)
		r := longest(root.Right, root, 1)
		return max(l, r, len)
	}

}

func max(l ...int) int {

	m := -9223372036854775808

	for _, i := range l {
		if i > m {
			m = i
		}
	}

	return m
}
