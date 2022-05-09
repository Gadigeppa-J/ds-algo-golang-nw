package main

/*
https://leetcode.com/problems/binary-tree-postorder-traversal/
TC: O(n)
SC: O(h)
*/

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

func postorderTraversal(root *utils.TreeNode) []int {

	if root == nil {
		return []int{}
	}

	larr := postorderTraversal(root.Left)
	rarr := postorderTraversal(root.Right)

	return append(append(larr, rarr...), root.Val)

}
