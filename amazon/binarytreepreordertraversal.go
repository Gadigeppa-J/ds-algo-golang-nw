package main

/*
https://leetcode.com/problems/binary-tree-preorder-traversal/
TC: O(n)
SC: O(h)
*/

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

func preorderTraversal(root *utils.TreeNode) []int {

	if root == nil {
		return []int{}
	}

	larr := preorderTraversal(root.Left)
	rarr := preorderTraversal(root.Right)

	return append(append([]int{root.Val}, larr...), rarr...)

}
