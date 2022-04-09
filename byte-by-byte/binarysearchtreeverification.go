package main

/*
25. Binary Search Tree Verification
Question:​ Given a binary tree, write a function to test if the tree is a binary search tree.

https://leetcode.com/problems/validate-binary-search-tree/submissions/

*/

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

func isValidBST(root *utils.TreeNode) bool {

	return isValid(root, -9223372036854775808, 9223372036854775807)

}

func isValid(root *utils.TreeNode, min, max int) bool {

	if root == nil {
		return true
	}

	if !(root.Val > min && root.Val < max) {
		return false
	} else {
		if !isValid(root.Left, min, root.Val) {
			return false
		}

		if !isValid(root.Right, root.Val, max) {
			return false
		}
	}

	return true

}
