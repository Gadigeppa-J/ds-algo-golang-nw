package main

/*
https://leetcode.com/problems/maximum-product-subarray/
TC: O(n)
SC: O(1)
*/

import (
	"fmt"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
}

func maxProduct(nums []int) int {

	ans := nums[0]
	maxProduct := nums[0]
	minProduct := nums[0]

	for i := 1; i < len(nums); i++ {

		if nums[i] < 0 {
			tmp := maxProduct
			maxProduct = minProduct
			minProduct = tmp
		}

		maxProduct = utils.Max(nums[i], nums[i]*maxProduct)
		minProduct = utils.Min(nums[i], nums[i]*minProduct)

		ans = utils.Max(ans, maxProduct)

	}

	return ans
}
