package main

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; i-j*j >= 0; j-- {
			dp[i] = Min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func canPartition(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	mid := sum / 2
	dp := make([]bool, mid+1)
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		for j := mid; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[mid]
}
