package main

import (
	"sort"

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

	n22 := &utils.TreeNode{
		Val: 22,
	}

	n5 := &utils.TreeNode{
		Val: 5,
	}

	n3 := &utils.TreeNode{
		Val: 3,
	}

	n4 := &utils.TreeNode{
		Val: 4,
	}

	n25 := &utils.TreeNode{
		Val: 25,
	}

	n10 := &utils.TreeNode{
		Val: 10,
	}

	n14 := &utils.TreeNode{
		Val: 14,
	}

	n20.Left = n8
	n20.Right = n22

	n8.Left = n5
	n8.Right = n3

	n3.Left = n10
	n3.Right = n14

	n22.Right = n25
	n22.Left = n4

	result := bottomViewOfBinaryTree(n20)

	for _, r := range result {
		fmt.Println(r.Val)
	}

}
*/

type TreeNodeWrapper struct {
	n  *utils.TreeNode
	hd int
}

func bottomViewOfBinaryTree(root *utils.TreeNode) []*utils.TreeNode {

	q := utils.NewGenericQueue()

	q.Enqueue(&TreeNodeWrapper{
		n:  root,
		hd: 0,
	})

	m := make(map[int]*utils.TreeNode)

	for !q.IsEmpty() {
		node := q.Dequeue().(*TreeNodeWrapper)
		m[node.hd] = node.n

		if node.n.Left != nil {
			q.Enqueue(&TreeNodeWrapper{
				n:  node.n.Left,
				hd: node.hd - 1,
			})
		}

		if node.n.Right != nil {
			q.Enqueue(&TreeNodeWrapper{
				n:  node.n.Right,
				hd: node.hd + 1,
			})
		}

	}

	result := []*utils.TreeNode{}
	tmpList := make([]int, len(result))

	for k := range m {
		tmpList = append(tmpList, k)
	}

	sort.Ints(tmpList)

	for _, k := range tmpList {
		result = append(result, m[k])
	}

	return result

}
