package main

/*
https://leetcode.com/problems/single-number/
TC: O(n)
SC: O(1)
*/

func singleNumber(nums []int) int {

	ans := 0

	for _, n := range nums {
		ans = ans ^ n
	}

	return ans

}
