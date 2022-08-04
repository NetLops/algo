package main

import "fmt"

func change(amount int, coins []int) int {
	//n := len(coins)
	dp := make([]int, amount+1)
	//for i := 1; i < amount+1; i++ {
	//	dp[i] = amount + 1
	//}
	//temp := 0
	dp[0] = 1
	for _, coin := range coins {
		for j := coin; j < amount+1; j++ {
			//if j >= amount {
			//	dp[j] += dp[j-coin]
			//}
			dp[j] += dp[j-coin]
		}
	}
	ans := dp[amount]
	if dp[amount] == amount+1 {
		ans = -1
	}
	return ans
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func main() {
	coins := []int{1, 2, 5}
	amount := 5
	fmt.Println(change(amount, coins))
}
