package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
}

func maxProduct(nums []int) int {

	maxP := nums[0]
	minP := nums[0]
	ans := nums[0]

	for i := 1; i < len(nums); i++ {

		if nums[i] < 0 {
			tmp := minP
			minP = maxP
			maxP = tmp
		}

		maxP = Max(nums[i], maxP*nums[i])
		minP = Min(nums[i], minP*nums[i])

		ans = Max(ans, maxP)

	}

	return ans
}

func Max(a, b int) int {

	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {

	if a < b {
		return a
	}

	return b
}
