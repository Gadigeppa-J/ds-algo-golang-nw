package main

/*
https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
TC: O(n)
SC: O(h)
*/

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

func main() {
	n1 := &utils.TreeNode{
		Val: 1,
	}
	n2 := &utils.TreeNode{
		Val: 2,
	}

	n1.Left = n2

	flatten(n1)
}

var prev *utils.TreeNode

func flatten(root *utils.TreeNode) {

	if root == nil {
		return
	}

	flatten(root.Right)
	flatten(root.Left)

	root.Right = prev
	root.Left = nil
	prev = root

}
