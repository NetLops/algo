package main

import "math"

func maxProfit(prices []int) int {
	ans := 0
	min := math.MaxInt
	for _, price := range prices {
		if min > price {
			min = price
		}
		temp := price - min
		if temp > ans {
			ans = temp
		}
	}

	return ans
}
func main() {

}
