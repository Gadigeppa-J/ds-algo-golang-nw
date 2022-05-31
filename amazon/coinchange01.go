package main

import (
	"fmt"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

/*
https://leetcode.com/problems/coin-change/
You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.

TC: O(amount * noOfCoins)
SC: O(amount)

*/

func main() {

	coins := []int{1, 2, 5}
	amount := 11

	fmt.Println(coinChange(coins, amount))
}

func coinChange(coins []int, amount int) int {
	/*
		m := make(map[int]int)
		return coinChangeRecursive(coins, amount, m)
	*/

	return coinChangeBottomUp(coins, amount)
}

// Top Down approach
func coinChangeRecursive(coins []int, amount int, mem map[int]int) int {

	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}

	if v, ok := mem[amount]; ok {
		return v
	}

	minCoin := amount + 1
	for i := 0; i < len(coins); i++ {
		res := coinChangeRecursive(coins, amount-coins[i], mem)

		if res != -1 {
			minCoin = utils.Min(minCoin, res+1)
		}

	}

	if minCoin == amount+1 {
		minCoin = -1
	}

	mem[amount] = minCoin

	return minCoin
}

// Bottom Up approach
func coinChangeBottomUp(coins []int, amount int) int {

	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}

	for i := 1; i < len(dp); i++ {
		for _, c := range coins {
			diff := i - c
			if diff >= 0 {
				dp[i] = utils.Min(dp[i], dp[diff]+1)
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}
