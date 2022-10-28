package main

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func coinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}

	dp := make([]int, amount+1)

	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 {
				dp[i] = Min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
