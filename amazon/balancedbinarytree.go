package main

/*
https://leetcode.com/problems/balanced-binary-tree/
TC: O(n)
SC: O(h)
*/

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

func isBalanced(root *utils.TreeNode) bool {

	if isBalancedHelper(root, 0) == -1 {
		return false
	}

	return true
}

func isBalancedHelper(root *utils.TreeNode, h int) int {

	if root == nil {
		return h
	}

	lh := isBalancedHelper(root.Left, h+1)
	rh := isBalancedHelper(root.Right, h+1)

	if lh == -1 || rh == -1 {
		return -1
	}

	diff := Abs(lh - rh)

	if diff > 1 {
		return -1
	}

	return Max(lh, rh)

}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Max(a, b int) int {

	if a > b {
		return a
	}

	return b

}
