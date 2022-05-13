package main

import (
	"math"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

/*
func main() {
	//fmt.Println(minSizeSubArraySum([]int{2, 1, 5, 2, 3, 2}, 7))
	//fmt.Println(minSizeSubArraySum([]int{2, 1, 5, 2, 8}, 7))
	fmt.Println(minSizeSubArraySum([]int{3, 4, 1, 1, 6}, 8))
}
*/

func minSizeSubArraySum(nums []int, s int) int {

	winSum := 0
	min := math.MaxInt32
	winStart := 0

	for winEnd := 0; winEnd < len(nums); winEnd++ {
		winSum += nums[winEnd]

		for winSum >= s {
			min = utils.Min(min, winEnd-winStart+1)
			winSum -= nums[winStart]
			winStart++
		}
	}

	return min

}
