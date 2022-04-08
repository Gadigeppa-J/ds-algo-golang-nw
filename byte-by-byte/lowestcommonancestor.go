package main

/*
18. Lowest Common Ancestor
Question:​ Given two nodes in a binary tree, write a function to find the lowest common ancestor.

https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/

Time Complexity: O(n)
Space Complexity: O(h)
*/

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

func lowestCommonAncestor(root, p, q *utils.TreeNode) *utils.TreeNode {

	if root == nil {
		return nil
	}

	if p.Val == root.Val || q.Val == root.Val {
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)

	if l != nil && r != nil {
		return root
	} else if l != nil {
		return l
	} else if r != nil {
		return r
	} else {
		return nil
	}

}
