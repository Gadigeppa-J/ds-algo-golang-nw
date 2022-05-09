package main

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

/*
https://leetcode.com/problems/two-sum-iv-input-is-a-bst/
TC: O(n)
SC: O(n)
*/

func findTarget(root *utils.TreeNode, k int) bool {

	return findTargetHelper(root, k, make(map[int]struct{}))

}

func findTargetHelper(root *utils.TreeNode, k int, m map[int]struct{}) bool {

	if root == nil {
		return false
	}

	l := findTargetHelper(root.Left, k, m)

	if l {
		return true
	} else if _, ok := m[root.Val]; ok {
		return true
	} else {
		m[(k - root.Val)] = struct{}{}
	}

	return findTargetHelper(root.Right, k, m)

}
