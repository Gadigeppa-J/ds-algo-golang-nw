package main

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

func diameterOfBinaryTree(root *utils.TreeNode) int {

	maxi := 0
	height(root, &maxi)

	return maxi
}

func height(root *utils.TreeNode, maxi *int) int {

	if root == nil {
		return 0
	}

	lh := height(root.Left, maxi)
	rh := height(root.Right, maxi)

	*maxi = utils.Max(*maxi, lh+rh)

	return utils.Max(lh, rh) + 1
}
