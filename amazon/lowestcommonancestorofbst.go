package main

/*
https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
TC: O(n)
SC: O(h)
*/

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

func lowestCommonAncestor(root, p, q *utils.TreeNode) *utils.TreeNode {

	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	lv := lowestCommonAncestor(root.Left, p, q)
	rv := lowestCommonAncestor(root.Right, p, q)

	if lv != nil && rv != nil {
		return root
	} else if lv != nil {
		return lv
	} else if rv != nil {
		return rv
	}

	return nil
}
