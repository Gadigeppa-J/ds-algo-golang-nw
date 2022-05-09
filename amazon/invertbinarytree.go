package main

/*
https://leetcode.com/problems/invert-binary-tree/
TC: O(n)
SC: O(1)
*/
import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

func invertTree(root *utils.TreeNode) *utils.TreeNode {

	if root == nil {
		return nil
	}

	// step1
	invertTree(root.Left)

	// step2
	invertTree(root.Right)

	// step3
	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	return root
}
