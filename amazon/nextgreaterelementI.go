package main

/*
https://leetcode.com/problems/next-greater-element-i/
TC: O(nums2*nums1)
SC: O(nums1)
*/

import (
	"fmt"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

func main() {

	nums1 := []int{2, 4}
	nums2 := []int{1, 2, 3, 4}

	fmt.Println(nextGreaterElement(nums1, nums2))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	res := make([]int, len(nums1))
	numToIndexMap := make(map[int]int)

	for i, n := range nums1 {
		numToIndexMap[n] = i
	}

	s := utils.NewStack()

	for i := 0; i < len(nums2); i++ {

		for !s.IsEmpty() && s.Peek() < nums2[i] {
			v := s.Pop()
			idx := numToIndexMap[v]
			res[idx] = nums2[i]
		}
		if _, ok := numToIndexMap[nums2[i]]; ok {
			s.Push(nums2[i])
		}
	}

	for !s.IsEmpty() {
		v := s.Pop()
		idx := numToIndexMap[v]
		res[idx] = -1
	}

	return res
}
