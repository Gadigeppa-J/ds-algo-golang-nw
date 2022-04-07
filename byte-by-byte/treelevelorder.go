package main

import (
	"fmt"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

/*
44. Tree Level Order
Question:​ Given a tree, write a function that prints out the nodes of the tree in level order.

Time Complexity: O(n)
Space Complexity: O(n)

*/

/*
func main() {

	n1 := &utils.TreeNode{
		Val: 1,
	}

	n3 := &utils.TreeNode{
		Val: 3,
	}

	n6 := &utils.TreeNode{
		Val: 6,
	}

	n8 := &utils.TreeNode{
		Val: 8,
	}

	n2 := &utils.TreeNode{
		Val:   2,
		Left:  n1,
		Right: n3,
	}

	n7 := &utils.TreeNode{
		Val:   7,
		Left:  n6,
		Right: n8,
	}

	n5 := &utils.TreeNode{
		Val:   5,
		Left:  n2,
		Right: n7,
	}

	treeLevelOrder(n5)
}
*/

func treeLevelOrder(t *utils.TreeNode) {

	q := utils.NewGenericQueue()

	q.Enqueue(t)

	for !q.IsEmpty() {
		n := q.Dequeue().(*utils.TreeNode)
		fmt.Println(n.Val)
		if n.Left != nil {
			q.Enqueue(n.Left)
		}
		if n.Right != nil {
			q.Enqueue(n.Right)
		}
	}

}
