package main

import "math"

/*
https://leetcode.com/problems/maximum-subarray/
TC: O(n)
SC: O(1)
*/

func maxSubArray(nums []int) int {

	sum := 0
	max := math.MinInt32

	for _, n := range nums {

		currSum := sum + n

		if currSum <= 0 {
			sum = 0
		} else {
			sum = currSum
		}

		if currSum > max {
			max = currSum
		}
	}

	return max
}
