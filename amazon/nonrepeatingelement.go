package main

/*
https://leetcode.com/problems/single-element-in-a-sorted-array/
TC: O(Logn)
SC: O(1)
*/

import "fmt"

func main() {
	nums := []int{3, 3, 7, 7, 10, 11, 11}
	fmt.Println(singleNonDuplicate(nums))
}

func singleNonDuplicate(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	l := 0
	h := len(nums) - 1

	// boundary condition check
	if nums[l] != nums[l+1] {
		return nums[l]
	}

	if nums[h] != nums[h-1] {
		return nums[h]
	}

	for l <= h {

		mid := (h + l) / 2

		// identify patterns
		if mid == 0 || mid%2 == 0 {
			// even index
			if nums[mid] != nums[mid+1] {
				h = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			// odd index
			if nums[mid] != nums[mid-1] {
				h = mid - 1
			} else {
				l = mid + 1
			}
		}

	}

	return nums[l]
}
