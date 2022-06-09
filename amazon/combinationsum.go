package main

/*
https://leetcode.com/problems/combination-sum/
*/

import "fmt"

func main() {
	nums := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(nums, target))
}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	combinationSumRecursive(candidates, target, 0, []int{}, &res)
	return res
}

func combinationSumRecursive(candidates []int, target int, index int, curr []int, res *[][]int) {

	if target == 0 {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		*res = append(*res, tmp)
		return
	}

	if target < 0 {
		return
	}

	for i := index; i < len(candidates); i++ {
		curr = append(curr, candidates[i])
		combinationSumRecursive(candidates, target-candidates[i], i, curr, res)
		curr = curr[:len(curr)-1]
	}

}

func sumInt(nums []int) int {

	sum := 0

	for _, n := range nums {
		sum += n
	}

	return sum
}
