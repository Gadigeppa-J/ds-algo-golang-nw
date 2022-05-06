package main

/*
https://leetcode.com/problems/max-consecutive-ones/
TC: O(n)
SC: O(1)
*/

func findMaxConsecutiveOnes(nums []int) int {

	max := 0
	sum := 0

	for _, n := range nums {
		if n == 1 {
			sum = sum + 1
		} else {
			sum = 0
		}

		if sum > max {
			max = sum
		}
	}

	return max

}
