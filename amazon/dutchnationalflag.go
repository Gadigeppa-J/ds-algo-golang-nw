package main

/*
https://leetcode.com/problems/sort-colors/
TC: O(n)
SC: O(1)
*/

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}

func sortColors(nums []int) {

	l, m := 0, 0
	h := len(nums) - 1

	for m <= h {

		if nums[m] == 0 {
			// when m = 0
			swapInArray(nums, l, m)
			l++
			m++
		} else if nums[m] == 1 {
			// when m = 1
			m++
		} else if nums[m] == 2 {
			// when m = 2
			swapInArray(nums, h, m)
			h--
		}
	}
}

func swapInArray(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}
