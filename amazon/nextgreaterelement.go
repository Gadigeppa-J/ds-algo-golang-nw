package main

/*
no leetcode
*/

import (
	"fmt"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

func main() {
	nums := []int{4, 3, 2, 1}
	fmt.Println(nextGreaterElements(nums))
}

func nextGreaterElements(nums []int) []int {

	res := make([]int, len(nums))
	s := utils.NewStack()

	for i := len(nums) - 1; i >= 0; i-- {

		for !s.IsEmpty() && s.Peek() <= nums[i] {
			s.Pop()
		}

		if !s.IsEmpty() {
			res[i] = s.Peek()
		} else {
			res[i] = -1
		}

		s.Push(nums[i])
	}

	return res
}
