package main

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	// 外层for 循环再遍历所有状态的所有取值
	for i := 0; i < len(dp); i++ {
		// 内层 for 循环再求所有选择的最小值
		for _, coin := range coins {
			// 无解的子问题就跳过
			if i-coin < 0 {
				continue
			}
			dp[i] = Min(dp[i], 1+dp[i-coin])
		}
	}
	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}
