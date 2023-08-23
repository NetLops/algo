package main

func maxCoins(nums []int) int {
	n := len(nums)
	process := make([]int, n+2)
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}
	process[0], process[n+1] = 1, 1
	// 处理一波边界问题
	for i := 0; i < n; i++ {
		process[i+1] = nums[i]
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j < n+2; j++ {
			for k := i + 1; k < j; k++ {
				temp := process[i]*process[j]*process[k] + dp[i][k] + dp[k][j]
				dp[i][j] = MAX(dp[i][j], temp)
			}
		}
	}
	return dp[0][n+1]

}

func MAX(a, b int) int {
	if a < b {
		return b

	}
	return a
}
