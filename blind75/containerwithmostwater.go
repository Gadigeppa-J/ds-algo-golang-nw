package main

import (
	"fmt"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {

	l := 0
	h := len(height) - 1
	max := 0

	for h >= l {
		area := utils.Min(height[l], height[h]) * (h - l)

		if area > max {
			max = area
		}

		if height[l] < height[h] {
			l++
		} else {
			h--
		}

	}

	return max
}
