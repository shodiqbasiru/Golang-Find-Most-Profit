package main

import "fmt"

func maxProfit(prices []int, trxAmount int) int {
	nPrices := len(prices)
	if nPrices == 0 {
		return 0
	}

	profit := 0
	if trxAmount >= nPrices/2 {
		for i := 1; i < nPrices; i++ {
			if prices[i] > prices[i-1] {
				profit += prices[i] - prices[i-1]
			}
		}
		return profit
	}

	buy := make([]int, trxAmount+1)
	sell := make([]int, trxAmount+1)
	for i := 0; i <= trxAmount; i++ {
		buy[i] = -prices[0]
	}

	for i := 1; i < nPrices; i++ {
		for j := 1; j <= trxAmount; j++ {
			buy[j] = max(buy[j], sell[j]-prices[i])
			sell[j] = max(sell[j], buy[j]+prices[i])
		}
	}

	return sell[trxAmount]
}

func main() {
	prices := []int{2, 4, 1}
	trxAmount := 2
	fmt.Println(maxProfit(prices, trxAmount)) // 2

	prices = []int{3, 2, 6, 5, 0, 3} // sample stock prices in
	trxAmount = 2
	fmt.Println(maxProfit(prices, trxAmount)) // 7

	prices = []int{3, 4, 2, 6, 5, 1}
	trxAmount = 2
	fmt.Println(maxProfit(prices, trxAmount)) // 5
}
