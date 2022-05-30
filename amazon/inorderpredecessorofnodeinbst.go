package main

/*
no leetcode
TC: O(h)
SC: O(1)
*/

import (
	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

/*
func main() {
	n20 := &utils.TreeNode{
		Val: 20,
	}

	n8 := &utils.TreeNode{
		Val: 8,
	}

	n4 := &utils.TreeNode{
		Val: 4,
	}

	n12 := &utils.TreeNode{
		Val: 12,
	}

	n10 := &utils.TreeNode{
		Val: 10,
	}

	n14 := &utils.TreeNode{
		Val: 14,
	}

	n22 := &utils.TreeNode{
		Val: 22,
	}

	n20.Left = n8
	n20.Right = n22

	n8.Left = n4
	n8.Right = n12

	n12.Left = n10
	n12.Right = n14

	n20.Right = n22

	fmt.Println(inorderPredecessor(n20, n8))

}
*/

func inorderPredecessor(root *utils.TreeNode, tar *utils.TreeNode) *utils.TreeNode {

	curr := root
	var lastRight *utils.TreeNode

	// find the node
	for curr != nil {

		if curr.Val > tar.Val {
			curr = curr.Left
		} else if curr.Val < tar.Val {
			lastRight = curr
			curr = curr.Right
		} else {
			break
		}

	}

	if curr != nil {

		if curr.Left != nil {

			temp := curr.Left

			for temp.Right != nil {
				curr = temp.Right
			}

			return temp
		} else {
			return lastRight
		}

	} else {

		return nil

	}

}
