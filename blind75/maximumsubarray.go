package main

/*
https://leetcode.com/problems/maximum-subarray/
*/
import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubArray([]int{-1}))
}

func maxSubArray(nums []int) int {

	max := math.MinInt
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]

		if sum > max {
			max = sum
		}

		if sum < 0 {
			sum = 0
		}
	}

	return max
}
