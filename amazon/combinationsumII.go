package main

/*
https://leetcode.com/problems/combination-sum-ii/
*/

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
	combinationSum2Recursive(candidates, target, 0, []int{}, &res)
	return res
}

func combinationSum2Recursive(candidates []int, target int, index int, curr []int, res *[][]int) {

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

		if i > index && candidates[i] == candidates[i-1] {
			continue
		}

		curr = append(curr, candidates[i])
		combinationSum2Recursive(candidates, target-candidates[i], i+1, curr, res)
		curr = curr[:len(curr)-1]
	}

}
