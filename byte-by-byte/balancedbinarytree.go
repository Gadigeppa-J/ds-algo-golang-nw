package main

/*
24. Balanced Binary Tree
Question:​ Given a binary tree, write a function to determine whether the tree is balanced.

https://leetcode.com/problems/balanced-binary-tree/

Time Complexity: O(n)
Space Complexity: O(h)
*/

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

func isBalanced(root *utils.TreeNode) bool {
	if height(root) < 0 {
		return false
	}
	return true
}

func height(root *utils.TreeNode) int {

	if root == nil {
		return 0
	}

	l := height(root.Left)
	r := height(root.Right)

	if (l-r) > 1 || (l-r) < -1 || l < 0 || r < 0 {
		return -1
	}

	if l >= r {
		return l + 1
	} else {
		return r + 1
	}

}
