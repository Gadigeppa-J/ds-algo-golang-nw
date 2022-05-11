package main

/*
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
TC: O(n)
SC: O(1)
*/

func maxProfitII(prices []int) int {

	if len(prices) == 1 {
		return 0
	}

	profit := 0
	p := 0
	q := 1

	for q < len(prices) {
		if prices[p] < prices[q] {
			diff := prices[q] - prices[p]
			profit = profit + diff
		}
		p++
		q++
	}

	return profit

}
