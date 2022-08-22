package main

/*
https://leetcode.com/problems/product-of-array-except-self/
*/

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

func productExceptSelf(nums []int) []int {

	if len(nums) == 0 {
		return nums
	}

	// left
	left := make([]int, len(nums))

	left[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		left[i] = nums[i] * left[i-1]
	}

	// right
	right := make([]int, len(nums))

	right[len(nums)-1] = nums[len(nums)-1]

	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = nums[i] * right[i+1]
	}

	// result
	res := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			res[i] = right[i+1]
		} else if i == len(nums)-1 {
			res[i] = left[i-1]
		} else {
			res[i] = left[i-1] * right[i+1]
		}
	}

	return res
}
