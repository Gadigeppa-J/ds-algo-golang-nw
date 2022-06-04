package main

/*
no leetcode
TC: O(n)
SC: O(n)
*/

import (
	"fmt"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

func main() {
	nums := []int{1, 0, 3}
	fmt.Println(subarrayZeroSum(nums))
}

func subarrayZeroSum(nums []int) int {
	m := make(map[int]int)
	max := 0
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum == 0 {
			max = utils.Max(max, i+1)
		} else {
			if j, ok := m[sum]; ok {
				max = utils.Max(max, i-j)
			} else {
				m[sum] = i
			}
		}

	}
	return max
}
