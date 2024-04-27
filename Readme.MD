# Maximum Profit from Stock Trades

This project contains a Go program that calculates the maximum profit that can be made from a series of stock trades. The main function of interest is `maxProfit`.

## Usage
```go

func main() {
    prices := []int{3, 2, 6, 5, 0, 3} // stock price sample in various days
    trxAmount := 2 // transaction amount
    fmt.Println(maxProfit(prices, trxAmount)) // 7
}
