package main

/*
https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
TC: O(n)
SC: O(1)
*/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {

	res := root
	c1 := root

	for c1 != nil && c1.Left != nil {

		c2 := c1
		c1 = c1.Left

		for c2 != nil {

			c2.Left.Next = c2.Right

			if c2.Next != nil {
				c2.Right.Next = c2.Next.Left
			}

			c2 = c2.Next
		}

	}

	return res

}
