package main

/*
https://leetcode.com/problems/container-with-most-water/
*/
import "fmt"

func main() {
	h := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(h))
}

func maxArea(height []int) int {

	l := 0
	h := len(height) - 1
	maxArea := 0

	for l <= h {
		v := min(height[l], height[h])

		if maxArea < v*(h-l) {
			maxArea = v * (h - l)
		}
		if height[l] < height[h] {
			l++
		} else {
			h--
		}
	}

	return maxArea
}

func min(a, b int) int {

	if a < b {
		return a
	}
	return b
}
