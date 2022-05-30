package main

/*
https://leetcode.com/problems/longest-increasing-subsequence/
TC: O(n^2)
SC: O(n)

Can we solve it in O(n log(n))??
*/

import (
	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

/*
func main() {

	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))

}
*/

func lengthOfLIS(nums []int) int {

	longArr := make([]int, len(nums))

	for i := range longArr {
		longArr[i] = 1
	}

	for j := 0; j < len(nums); j++ {
		for i := 0; i < j; i++ {
			if nums[i] < nums[j] {
				v := longArr[i] + 1
				if v > longArr[j] {
					longArr[j] = v
				}
			}
		}
	}

	longest := 0

	for _, v := range longArr {
		longest = utils.Max(longest, v)
	}

	return longest
}
