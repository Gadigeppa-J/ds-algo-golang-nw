package main

/*
https://leetcode.com/problems/binary-tree-inorder-traversal/
TC: O(n)
SC: O(h)
*/

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

func inorderTraversal(root *utils.TreeNode) []int {

	if root == nil {
		return []int{}
	}

	larr := inorderTraversal(root.Left)
	larr = append(larr, root.Val)
	rarr := inorderTraversal(root.Right)
	larr = append(larr, rarr...)
	return larr
}
