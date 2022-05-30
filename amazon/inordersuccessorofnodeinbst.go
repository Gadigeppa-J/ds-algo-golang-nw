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

	fmt.Println(inorderSuccessor(n20, n22))

}
*/

func inorderSuccessor(node *utils.TreeNode, tar *utils.TreeNode) *utils.TreeNode {

	curr := node
	var lastLeft *utils.TreeNode

	// find node and capture last left
	for curr != nil {
		if curr.Val > tar.Val {
			lastLeft = curr
			curr = curr.Left
		} else if curr.Val < tar.Val {
			curr = curr.Right
		} else {
			break
		}
	}

	var temp *utils.TreeNode
	if curr != nil {
		if curr.Right != nil {
			temp = curr.Right
			for temp.Left != nil {
				temp = temp.Left
			}
			return temp
		} else {
			return lastLeft
		}
	} else {
		return nil
	}

}
