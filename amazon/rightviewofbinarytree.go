package main

/*
https://leetcode.com/problems/binary-tree-right-side-view/
TC: O(n)
SC: O()
*/

import (
	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

/*
func main() {

	n1 := &utils.TreeNode{
		Val: 1,
	}

	n2 := &utils.TreeNode{
		Val: 2,
	}

	n3 := &utils.TreeNode{
		Val: 3,
	}

	n4 := &utils.TreeNode{
		Val: 4,
	}

	n5 := &utils.TreeNode{
		Val: 5,
	}

	n1.Left = n2
	n1.Right = n3
	n2.Right = n5
	n3.Right = n4

	fmt.Println(rightSideView(n1))

}
*/
type TreeNodeWrapper2 struct {
	nd  *utils.TreeNode
	lvl int
}

func rightSideView(root *utils.TreeNode) []int {

	result := []int{}
	if root == nil {
		return result
	}

	q := utils.NewGenericQueue()

	q.Enqueue(&TreeNodeWrapper2{
		nd:  root,
		lvl: 1,
	})

	for !q.IsEmpty() {
		node := q.Dequeue().(*TreeNodeWrapper2)

		if len(result) < node.lvl {
			result = append(result, node.nd.Val)
		}

		if node.nd.Right != nil {
			q.Enqueue(&TreeNodeWrapper2{
				nd:  node.nd.Right,
				lvl: node.lvl + 1,
			})
		}

		if node.nd.Left != nil {
			q.Enqueue(&TreeNodeWrapper2{
				nd:  node.nd.Left,
				lvl: node.lvl + 1,
			})
		}

	}

	return result

}
