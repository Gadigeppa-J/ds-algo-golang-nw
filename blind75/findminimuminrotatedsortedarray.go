package main

/*
https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
*/

import (
	"fmt"
	"math"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

func main() {
	fmt.Println(findMin([]int{3, 1, 2}))
}

func findMin(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	if nums[0] <= nums[len(nums)-1] {
		return nums[0]
	}

	l := 0
	h := len(nums) - 1

	min := math.MaxInt

	for h >= l {

		if nums[l] <= nums[h] {
			return utils.Min(min, nums[l])
		}

		m := (h + l) / 2

		min = utils.Min(min, nums[m])

		if nums[m] >= nums[l] {
			l = m + 1
		} else {
			h = m - 1
		}

	}
	return min
}
