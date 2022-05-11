package main

/*
https://leetcode.com/problems/merge-sorted-array/submissions/
TC: O(n)
SC: O(1)
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	p := m - 1
	q := n - 1
	r := len(nums1) - 1

	for q > -1 {
		if p > -1 && nums1[p] > nums2[q] {
			nums1[r] = nums1[p]
			p--
		} else {
			nums1[r] = nums2[q]
			q--
		}
		r--
	}

}
