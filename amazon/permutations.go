package main

/*
https://leetcode.com/problems/permutations/
TC:
SC:
*/

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}

func permute(nums []int) [][]int {
	res := [][]int{}
	permuteRecursive(nums, make([]bool, len(nums)), []int{}, &res)
	return res
}

func permuteRecursive(nums []int, used []bool, curr []int, res *[][]int) {

	if len(curr) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, curr)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {

		if used[i] {
			continue
		}

		used[i] = true
		curr = append(curr, nums[i])
		permuteRecursive(nums, used, curr, res)
		curr = curr[:len(curr)-1]
		used[i] = false
	}

}
