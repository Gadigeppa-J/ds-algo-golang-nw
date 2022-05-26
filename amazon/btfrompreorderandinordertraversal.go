package main

/*
	https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
	TC: O(n)
	SC: O(h)

*/

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

func buildTree(preorder []int, inorder []int) *utils.TreeNode {

	// preorder: root - left - right
	// inorder: left - root -right

	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	node := &utils.TreeNode{
		Val: preorder[0],
	}

	mid := findIndexAt(inorder, preorder[0])

	node.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	node.Right = buildTree(preorder[mid+1:], inorder[mid+1:])

	return node

}

func findIndexAt(arr []int, val int) int {

	for i, v := range arr {

		if v == val {
			return i
		}

	}

	return -1

}
