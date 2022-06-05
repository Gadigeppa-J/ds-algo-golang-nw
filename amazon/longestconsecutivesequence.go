package main

/*
https://leetcode.com/problems/longest-consecutive-sequence/
TC: O(n)
SC: O(n)
*/

import (
	"fmt"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

func main() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(nums))
}

func longestConsecutive(nums []int) int {

	s := make(map[int]struct{})
	for _, n := range nums {
		s[n] = struct{}{}
	}

	max := 0

	for k := range s {

		count := 1
		if _, ok := s[k-1]; ok {
			continue
		}

		curr := k

		for true {
			_, ok := s[curr+1]
			if ok {
				count++
				curr = curr + 1
			} else {
				break
			}
		}
		max = utils.Max(max, count)
	}

	return max
}
