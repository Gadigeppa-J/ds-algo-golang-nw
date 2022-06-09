package main

/*
https://leetcode.com/problems/subsets-ii/
TC: O(n*2^n) recursive
SC: O(n) recursive
output SC: O(n*2^n)
*/

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	subsetsWithDupRecursive(nums, 0, []int{}, &res)
	return res
}

func subsetsWithDupRecursive(nums []int, index int, curr []int, res *[][]int) {

	tmp := make([]int, len(curr))
	copy(tmp, curr)
	*res = append(*res, tmp)

	for i := index; i < len(nums); i++ {

		if i > index && nums[i] == nums[i-1] {
			continue
		}

		curr = append(curr, nums[i])
		subsetsWithDupRecursive(nums, i+1, curr, res)
		curr = curr[:len(curr)-1]
	}

}
