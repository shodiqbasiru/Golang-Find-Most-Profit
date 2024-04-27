# Maximum Profit from Stock Trades

This project contains a Go program that calculates the maximum profit that can be made from a series of stock trades. The main function of interest is `maxProfit`.

## Function: maxProfit

The `maxProfit` function takes in two parameters:

- `prices`: a slice of integers representing the stock prices on various days.
- `trxAmount`: an integer representing the number of transactions that can be made.

The function returns an integer representing the maximum profit that can be made from the stock trades.
```go
func main() {
    prices := []int{3, 2, 6, 5, 0, 3} // stock price sample in various days
    trxAmount := 2 // transaction amount
    fmt.Println(maxProfit(prices, trxAmount)) // 7
}
