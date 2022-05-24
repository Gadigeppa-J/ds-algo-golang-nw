package main

/*

 */

import (
	"fmt"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

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

	n6 := &utils.TreeNode{
		Val: 6,
	}

	n7 := &utils.TreeNode{
		Val: 7,
	}

	n1.Left = n2
	n1.Right = n3

	n2.Left = n4
	n2.Right = n5

	n3.Left = n6
	n3.Right = n7

	fmt.Println(printBTZigZagLevelOrder(n1))

}

func printBTLevelOrder(root *utils.TreeNode) {

	q := utils.NewGenericQueue()

	q.Enqueue(root)

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

func printBTZigZagLevelOrder(root *utils.TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	s1 := utils.NewGenericStack()
	s2 := utils.NewGenericStack()

	s1.Push(root)

	for !s1.IsEmpty() || !s2.IsEmpty() {

		tmp := []int{}

		for !s1.IsEmpty() {

			n := s1.Pop().(*utils.TreeNode)
			//fmt.Println(n.Val)
			tmp = append(tmp, n.Val)

			if n.Left != nil {
				s2.Push(n.Left)
			}

			if n.Right != nil {
				s2.Push(n.Right)
			}

		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		tmp = []int{}

		for !s2.IsEmpty() {

			n := s2.Pop().(*utils.TreeNode)
			//fmt.Println(n.Val)
			tmp = append(tmp, n.Val)

			if n.Right != nil {
				s1.Push(n.Right)
			}

			if n.Left != nil {
				s1.Push(n.Left)
			}

		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

	}

	return result
}
