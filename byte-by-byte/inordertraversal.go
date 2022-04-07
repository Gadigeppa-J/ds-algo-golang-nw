package main

import (
	"fmt"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

/*
27. Inorder Traversal
Question:​ Given a binary search tree, print out the elements of the tree in order without using recursion.

https://leetcode.com/problems/binary-tree-inorder-traversal/

Time Complexity: O(n)
Space Complexity: O(h) where h is the height of the binary tree

*/

func inorderTraversal(treeNode *utils.TreeNode) {

	s := NewTreeNodeStack()
	curr := treeNode

	for true {

		for curr != nil {
			s.Push(curr)
			curr = curr.Left
		}

		if s.IsEmpty() {
			break
		}

		n := s.Pop()
		fmt.Println(n.Val)

		curr = n.Right
	}

}

type TreeNodeStack struct {
	arr []*utils.TreeNode
}

func NewTreeNodeStack() *TreeNodeStack {
	return &TreeNodeStack{
		arr: []*utils.TreeNode{},
	}
}

func (s *TreeNodeStack) Push(i *utils.TreeNode) {
	s.arr = append(s.arr, i)
}

func (s *TreeNodeStack) Pop() *utils.TreeNode {

	if s.IsEmpty() {
		panic("stack is empty!")
	}

	i := s.arr[len(s.arr)-1]

	s.arr = s.arr[:len(s.arr)-1]

	return i
}

func (s *TreeNodeStack) IsEmpty() bool {

	if len(s.arr) == 0 {
		return true
	}

	return false
}

func (s *TreeNodeStack) Peek() *utils.TreeNode {
	if s.IsEmpty() {
		panic("stack is empty!")
	}
	return s.arr[len(s.arr)-1]
}

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

	inorderTraversal(n5)   // 1, 2, 3, 5, 6, 7, 8
	preorderTraversal(n5)  // 5, 2, 1, 3, 7, 6, 8
	postorderTraversal(n5) // 1, 3, 2, 6, 8, 7, 5
}

/*
func inorderTraversal(treeNode *utils.TreeNode) {

	if treeNode == nil {
		return
	}

	inorderTraversal(treeNode.Left)
	fmt.Println(treeNode.Val)
	inorderTraversal(treeNode.Right)
}

func preorderTraversal(treeNode *utils.TreeNode) {

	if treeNode == nil {
		return
	}

	fmt.Println(treeNode.Val)
	preorderTraversal(treeNode.Left)
	preorderTraversal(treeNode.Right)
}

func postorderTraversal(treeNode *utils.TreeNode) {

	if treeNode == nil {
		return
	}

	postorderTraversal(treeNode.Left)
	postorderTraversal(treeNode.Right)
	fmt.Println(treeNode.Val)
}
*/
