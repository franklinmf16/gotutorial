package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}
	min := prices[0]
	largest := 0
	for i:=1;i< len(prices); i++  {
		if prices[i] < min{
			min = prices[i]
		}

		if (prices[i] - min) > largest {
			largest = prices[i] - min
		}
	}
	return largest
}

func main() {
	prices := []int{1,3}
	profit := maxProfit(prices)
	fmt.Println(profit)
}
