package main

/*
https://leetcode.com/problems/symmetric-tree/
TC: O(n)
SC: O(h)
*/

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

func isSymmetric(root *utils.TreeNode) bool {

	if root == nil {
		return true
	}

	return isSymmetricHelper(root.Left, root.Right)

}

func isSymmetricHelper(left, right *utils.TreeNode) bool {

	if left == nil || right == nil {
		return left == right
	}

	if left.Val != right.Val {
		return false
	}

	return isSymmetricHelper(left.Left, right.Right) && isSymmetricHelper(left.Right, right.Left)

}
