package main

/*
	https://leetcode.com/problems/longest-common-subsequence/
	TC of DP: O(m*n)
	SC of DP: O(m*n)
*/

import (
	"fmt"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

func main() {
	//fmt.Println(longestCommonSubsequence("abcde", "ace"))
	fmt.Println(lcsDP("abcde", "ace"))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1))

	for i := range dp {
		dp[i] = make([]int, len(text2))

		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return lcsRecursive(text1, text2, 0, 0, dp)
}

func lcsRecursive(t1 string, t2 string, i int, j int, dp [][]int) int {

	if i == len(t1) || j == len(t2) {
		return 0
	}

	if dp[i][j] > -1 {
		return dp[i][j]
	}

	var ans int
	if t1[i] == t2[j] {
		ans = 1 + lcsRecursive(t1, t2, i+1, j+1, dp)
	} else {
		ans = utils.Max(lcsRecursive(t1, t2, i+1, j, dp), lcsRecursive(t1, t2, i, j+1, dp))
	}
	dp[i][j] = ans
	return ans

}

func lcsDP(t1 string, t2 string) int {

	dp := make([][]int, len(t2)+1)

	for i := range dp {
		dp[i] = make([]int, len(t1)+1)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if t2[i-1] == t1[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = utils.Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(t2)][len(t1)]
}
