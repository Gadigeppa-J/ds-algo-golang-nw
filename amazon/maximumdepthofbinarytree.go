package main

/*
https://leetcode.com/problems/maximum-depth-of-binary-tree/
TC: O(n)
SC: O(h)
*/

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

func maxDepth(root *utils.TreeNode) int {

	if root == nil {
		return 0
	}

	lh := maxDepth(root.Left)
	rh := maxDepth(root.Right)

	return Max(lh+1, rh+1)
}

/*
func Max(a, b int) int {

	if a > b {
		return a
	}

	return b
}
*/
