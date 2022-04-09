package main

import (
	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

/*
func main() {

	root := &utils.TreeNode{
		Val: 1,

		Left: &utils.TreeNode{
			Val: 2,
			Left: &utils.TreeNode{
				Val: 4,
			},
			Right: &utils.TreeNode{
				Val: 5,
			},
		},

		Right: &utils.TreeNode{
			Val: 3,
			Left: &utils.TreeNode{
				Val: 6,
			},
			Right: &utils.TreeNode{
				Val: 7,
			},
		},
	}

	list := treeToDoublyList(root)

}
*/

func concatenate(a *utils.TreeNode, b *utils.TreeNode) *utils.TreeNode {

	if a == nil {
		return b
	}

	if b == nil {
		return a
	}

	aEnd := a.Left
	bEnd := b.Left

	a.Left = bEnd
	bEnd.Right = a

	aEnd.Right = b
	b.Left = aEnd

	return a
}

func treeToDoublyList(root *utils.TreeNode) *utils.TreeNode {

	if root == nil {
		return nil
	}

	l := treeToDoublyList(root.Left)
	r := treeToDoublyList(root.Right)

	root.Left = root
	root.Right = root

	root = concatenate(l, root)
	root = concatenate(root, r)

	return root
}
