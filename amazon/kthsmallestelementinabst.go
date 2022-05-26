package main

/*
https://leetcode.com/problems/kth-smallest-element-in-a-bst/
TC: O(n)
SC: O(n)
*/

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

func kthSmallest(root *utils.TreeNode, k int) int {

	kv := k

	r := kthSmallestHelp(root, &kv)

	if r == nil {
		return -1
	}

	return r.Val

}

func kthSmallestHelp(root *utils.TreeNode, k *int) *utils.TreeNode {

	if root == nil {
		return nil
	}

	l := kthSmallestHelp(root.Left, k)

	if l != nil {
		return l
	}

	*k = *k - 1

	if *k == 0 {
		return root
	}

	r := kthSmallestHelp(root.Right, k)

	return r

}
