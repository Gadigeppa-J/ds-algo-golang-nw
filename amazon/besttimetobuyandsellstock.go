package main

/*
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
TC: O(n)
SC: O(1)
*/

import (
	"fmt"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

func main() {

	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))

}

func maxProfit(prices []int) int {

	if len(prices) == 1 {
		return 0
	}

	maxprofit := 0
	p := 0
	q := 1

	for q < len(prices) {

		if prices[p] >= prices[q] {
			p = q
		} else {
			diff := prices[q] - prices[p]
			maxprofit = utils.Max(maxprofit, diff)
		}

		q++
	}

	return maxprofit
}
