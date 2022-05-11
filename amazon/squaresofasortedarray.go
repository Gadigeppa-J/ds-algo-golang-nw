package main

/*
https://leetcode.com/problems/squares-of-a-sorted-array/
TC: O(n)
SC: O(n)
*/

func sortedSquares(nums []int) []int {

	if len(nums) == 0 {
		return nums
	}

	res := make([]int, len(nums))

	p := 0
	q := len(nums) - 1
	r := q

	for r >= 0 {

		if nums[p]*nums[p] > nums[q]*nums[q] {
			res[r] = nums[p] * nums[p]
			p++
		} else {
			res[r] = nums[q] * nums[q]
			q--
		}
		r--
	}

	return res

}
