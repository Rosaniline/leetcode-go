package leetcode

import "golang.org/x/tools/container/intsets"

// Say you have an array for which the ith element is the price of a given stock on day i.
//
// If you were only permitted to complete at most one transaction
// (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.
//
// Note that you cannot sell a stock before you buy one.


func maxProfit(prices []int) int {
	var (
		profit	= 0
		buy		= intsets.MaxInt
	)

	for _, price := range prices {
		buy = min(buy, price)
		profit = max(profit, price - buy)
	}

	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}