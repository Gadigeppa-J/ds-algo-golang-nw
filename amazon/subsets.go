package main

import "fmt"

/*
https://leetcode.com/problems/subsets/
TC: O(n*2^n) recursive
SC: O(n) recursive
output SC: O(n*2^n)
*/

func main() {
	res := [][]int{}
	nums := []int{1, 2, 3}
	subsetsRecursive(nums, 0, []int{}, &res)
	fmt.Println(res)
}

func subsetsRecursive(nums []int, level int, curr []int, res *[][]int) {

	tmp := make([]int, len(curr))
	copy(tmp, curr)
	*res = append(*res, tmp)

	for i := level; i < len(nums); i++ {
		curr = append(curr, nums[i])
		subsetsRecursive(nums, i+1, curr, res)
		curr = curr[:len(curr)-1]
	}

}

func subsetsIterative(nums []int) [][]int {

	// TC - (n * 2^n)
	// SC - (n * 2^n)

	res := [][]int{}
	res = append(res, []int{})

	for i := 0; i < len(nums); i++ {
		n := len(res)
		for j := 0; j < n; j++ {
			dest := make([]int, len(res[j]))
			copy(dest, res[j])
			dest = append(dest, nums[i])
			res = append(res, dest)
		}
	}

	return res
}
