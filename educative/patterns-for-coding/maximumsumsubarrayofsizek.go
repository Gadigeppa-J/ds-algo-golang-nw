package main

import (
	"fmt"
	"math"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

func main() {
	//fmt.Println(maxSubArray([]int{2, 1, 5, 1, 3, 2}, 3))
	fmt.Println(maxSubArray([]int{2, 3, 4, 1, 5}, 2))
}

func maxSubArray(nums []int, k int) int {

	max := math.MinInt16
	winSum := 0

	for i := 0; i < k; i++ {
		winSum += nums[i]
	}

	max = winSum
	for q := k; q < len(nums); q++ {
		winSum += nums[q] - nums[q-k]
		max = utils.Max(max, winSum)
	}

	return max
}
