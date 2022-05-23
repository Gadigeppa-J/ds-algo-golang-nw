package main

/*
https://leetcode.com/problems/3sum/
TC: O(n^2)
SC: O(1)
*/

import (
	"sort"
)

/*
func main() {

	//[-2,0,1,1,2]

	nums := []int{-2, 0, 1, 1, 2}

	fmt.Println(threeSum(nums))
}
*/

func threeSum(nums []int) [][]int {

	result := [][]int{}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	// [-4 -1 -1 0 1 2]

	for i := 0; i < len(nums)-2; i++ {

		if (i == 0) || nums[i] != nums[i-1] {

			low := i + 1
			high := len(nums) - 1
			sum := 0 - nums[i]

			for low < high {

				sum2 := nums[low] + nums[high]

				if sum2 == sum {
					result = append(result, []int{nums[i], nums[low], nums[high]})

					for (low < high) && (nums[low] == nums[low+1]) {
						low++
					}

					for (low < high) && (nums[high] == nums[high-1]) {
						high--
					}

					low++
					high--
				} else if sum2 > sum {
					high--
				} else {
					low++
				}

			}
		}

	}

	return result
}
