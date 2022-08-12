package main

import "math"

//func maxProfit(prices []int) int {
//	min := math.MaxInt // 求最小的一天
//	ans := 0           // 每天的最大利润
//	for i := 0; i < len(prices); i++ {
//		if prices[i] < min {
//			min = prices[i]
//		}
//		temp := prices[i] - min
//		if ans < temp {
//			ans = temp
//		}
//	}
//	return ans
//}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func maxProfit(prices []int) int {
	ans := 0
	min := math.MaxInt
	for _, price := range prices {
		min = minInt(min, price)
		ans = maxInt(ans, price-min)
	}
	return ans
}
func main() {

}
