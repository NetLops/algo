package main

import (
	"fmt"
	"testing"
)

func coinChange(coins []int, amount int) int {
	n := len(coins)
	//length := amount
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, amount+1)
		for j := 0; j < amount+1; j++ {
			dp[i][j] = amount + 1 // 设置一个较大的数
		}
	}
	dp[0][0] = 0
	for i := 1; i < n+1; i++ {
		for j := 0; j < amount+1; j++ {
			if j < coins[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = min(dp[i][j-coins[i-1]]+1, dp[i-1][j])
			}

		}
	}
	fmt.Println(dp)
	ans := dp[n][amount]
	if ans == amount+1 {
		ans = -1
	}
	return ans
}

func coinChange3(coins []int, amount int) int {
	n := len(coins)
	//length := amount
	dp := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i < n+1; i++ {
		for j := 0; j < amount+1; j++ {
			//if j < coins[i-1] {
			//	dp[j] = dp[j]
			//} else {
			//	dp[j] = min(dp[j-coins[i-1]]+1, dp[j])
			//}
			if j >= coins[i-1] {
				dp[j] = min(dp[j-coins[i-1]]+1, dp[j])
			}

		}
	}
	fmt.Println(dp)
	ans := dp[amount]
	if ans == amount+1 {
		ans = -1
	}
	return ans
}
func TestName(t *testing.T) {
	//coins := []int{1, 2, 5}
	//amount := 6
	//
	//coins := []int{2}
	//amount := 3

	coins := []int{1, 2, 5}
	amount := 5
	fmt.Println(coinChange3(coins, amount))
}
