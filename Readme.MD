# Maximum Profit from Stock Trades

This project contains a Go program that calculates the maximum profit that can be made from a series of stock trades. The main function of interest is `maxProfit`.

## Function: maxProfit

The `maxProfit` function takes in two parameters:

- `prices`: An array of integers representing the stock prices on various days.
- `trxAmount`: The number of transactions that can be made.

The function returns the maximum profit that can be made from the given number of transactions.

```go

func main() {
    prices := []int{3, 2, 6, 5, 0, 3} // stock price sample in various days
    trxAmount := 2 // transaction amount
    fmt.Println(maxProfit(prices, trxAmount)) // 7
}
