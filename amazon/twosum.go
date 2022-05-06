package main

/*
Leetcode: https://leetcode.com/problems/two-sum/
TC: O(n)
SC: O(n)
*/

func twoSum(nums []int, target int) []int {

	m := make(map[int]int)

	for i, n := range nums {

		if l, ok := m[n]; ok {
			return []int{l, i}
		}

		m[target-n] = i

	}

	return []int{-1, -1}
}
