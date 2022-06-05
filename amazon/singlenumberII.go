package main

/*
https://leetcode.com/problems/single-number-ii/
TC: O(nlogn)
SC: O(1)
*/

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{30000, 500, 100, 30000, 100, 30000, 100}
	fmt.Println(singleNumberII(nums))
}

func singleNumberII(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	sort.Ints(nums)

	// boundary conditions
	if nums[0] != nums[1] {
		return nums[0]
	}

	if nums[len(nums)-1] != nums[len(nums)-2] {
		return nums[len(nums)-1]
	}

	index := 1

	for true {
		if nums[index] != nums[index-1] {
			break
		}
		index = index + 3
	}

	return nums[index-1]
}
