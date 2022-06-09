package main

/*
https://leetcode.com/problems/combination-sum-iii/
*/

import "fmt"

func main() {
	fmt.Println(combinationSum3(4, 1))
}

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	combinationSum3Recursive(k, n, 1, []int{}, &res)
	return res
}

func combinationSum3Recursive(k int, n int, index int, curr []int, res *[][]int) {

	if n < 0 {
		return
	}

	if len(curr) >= k {
		if n == 0 {
			tmp := make([]int, len(curr))
			copy(tmp, curr)
			*res = append(*res, tmp)
		}
		return
	}

	for i := index; i <= 9; i++ {
		curr = append(curr, i)
		combinationSum3Recursive(k, n-i, i+1, curr, res)
		curr = curr[:len(curr)-1]
	}
}
