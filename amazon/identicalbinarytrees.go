package main

/*
https://leetcode.com/problems/same-tree/
TC: O(n)
SC: O(h)
*/
import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

func isSameTree(p *utils.TreeNode, q *utils.TreeNode) bool {

	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)

}
