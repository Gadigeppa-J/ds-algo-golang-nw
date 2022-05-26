package main

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

/*
https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
TC: O(n)
SC: O(h)
*/

func buildTreeIP(inorder []int, postorder []int) *utils.TreeNode {

	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	lastIndex := len(postorder) - 1

	node := &utils.TreeNode{
		Val: postorder[lastIndex],
	}

	mid := findIndexAt(inorder, postorder[lastIndex])

	node.Left = buildTreeIP(inorder[:mid], postorder[:mid])
	node.Right = buildTreeIP(inorder[mid+1:], postorder[mid:lastIndex])

	return node
}
