package main

/*
https://leetcode.com/problems/combinations/
TC: O(kn^k)
SC: O(k)
*/

import "fmt"

func main() {
	res := [][]int{}
	combinationRecursive(1, 4, 2, []int{}, &res)
	fmt.Println(res)
}

func combinationRecursive(index int, n int, k int, curr []int, res *[][]int) {

	if len(curr) == k {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		*res = append(*res, tmp)
		return
	}

	for i := index; i <= n; i++ {
		curr = append(curr, i)
		combinationRecursive(i+1, n, k, curr, res)
		curr = curr[:len(curr)-1]
	}

}
