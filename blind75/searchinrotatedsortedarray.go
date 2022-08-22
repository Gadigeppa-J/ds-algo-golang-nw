package main

/*
https://leetcode.com/problems/search-in-rotated-sorted-array/
*/
import "fmt"

func main() {
	fmt.Println(search([]int{1}, 0))
}

func search(nums []int, target int) int {

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	if nums[len(nums)-1] >= nums[0] {
		return binarySearch(nums, target, 0, len(nums)-1)
	}

	pivotIndex := findPivotIndex(nums)

	if target >= nums[0] && target <= nums[pivotIndex] {
		return binarySearch(nums, target, 0, pivotIndex)
	} else {
		return binarySearch(nums, target, pivotIndex+1, len(nums)-1)
	}

}

func binarySearch(nums []int, target, l, h int) int {

	for h >= l {

		m := (h + l) / 2

		if nums[m] == target {
			return m
		} else if nums[m] > target {
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

	for h >= l {

		m := (h + l) / 2

		if nums[m] > nums[m+1] {
			return m
		} else if nums[m] >= nums[l] {
			l = m + 1
		} else {
			h = m - 1
		}

	}

	return -1
}
