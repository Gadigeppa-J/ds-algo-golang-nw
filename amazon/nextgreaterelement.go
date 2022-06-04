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

type GreatItem struct {
	Val   int
	Index int
}

func nextGreaterElements(nums []int) []int {

	res := make([]int, len(nums))

	s := utils.NewGenericStack()

	for i := 0; i < len(nums); i++ {

		if s.IsEmpty() {
			s.Push(GreatItem{
				Val:   nums[i],
				Index: i,
			})
		} else {
			for !s.IsEmpty() && s.Peek().(GreatItem).Val < nums[i] {
				v := s.Pop().(GreatItem)
				res[v.Index] = nums[i]
			}
			s.Push(GreatItem{
				Val:   nums[i],
				Index: i,
			})
		}

	}

	for !s.IsEmpty() {
		v := s.Pop().(GreatItem)
		res[v.Index] = -1
	}

	return res
}
