package main

/*
https://leetcode.com/problems/search-in-rotated-sorted-array/
TC: O(LogN)
SC: O(n)
*/

import "fmt"

func main() {
	nums := []int{8, 9, 2, 3, 4}
	fmt.Println(searchRotated(nums, 9))

}

func searchRotated(nums []int, target int) int {

	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	pivot := findPivotIndex(nums)

	l := 0
	h := len(nums) - 1

	if pivot >= 0 {
		if nums[pivot] <= target && nums[len(nums)-1] >= target {
			l = pivot
			h = len(nums) - 1
		} else {
			l = 0
			h = pivot - 1
		}
	}

	return performBinSearch(nums, l, h, target)
}

func performBinSearch(nums []int, l, h, v int) int {

	if len(nums) == 0 {
		return -1
	}

	for l <= h {
		m := (h + l) / 2

		if nums[m] == v {
			return m
		} else if nums[m] > v {
			h = m - 1
		} else {
			l = m + 1
		}
	}

	return -1
}

func findPivotIndex(nums []int) int {

	l := 0
	h := len(nums) - 1
	pivot := -1

	for l <= h {
		m := (h + l) / 2

		if m < len(nums)-1 && nums[m] > nums[m+1] {
			pivot = m + 1
			break
		} else {

			if nums[l] <= nums[m] {
				l = m + 1
			} else {
				h = m - 1
			}
		}
	}

	return pivot
}
