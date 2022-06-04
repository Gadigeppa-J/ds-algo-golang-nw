package main

/*
	https://leetcode.com/problems/next-greater-element-ii/
	TC: O(n)
	SC: O(n)
*/

import (
	"fmt"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

func main() {
	nums := []int{1, 2, 3, 4, 3}
	fmt.Println(nextGreaterElementsII(nums))
}

func nextGreaterElementsII(nums []int) []int {

	res := make([]int, len(nums))
	n := len(nums)
	s := utils.NewStack()

	for i := (2 * n) - 1; i >= 0; i-- {

		for !s.IsEmpty() && s.Peek() <= nums[i%n] {
			s.Pop()
		}

		if !s.IsEmpty() {
			res[i%n] = s.Peek()
		} else {
			res[i%n] = -1
		}

		s.Push(nums[i%n])

	}

	return res

}
