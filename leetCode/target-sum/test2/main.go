package main

func Abs(target int) int {
	if target < 0 {
		return -target
	}
	return target
}
func findTargetSumWays(nums []int, target int) int {
	sum, n := 0, len(nums)

	for i := 0; i < n; i++ {
		sum += nums[i]
	}

	//  超出范围就无了
	if Abs(target) > Abs(sum) {
		return 0
	}

	// - 0 +
	t := sum*2 + 1
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, t)
	}

	// InitValue
	if nums[0] == 0 {
		dp[0][sum] = 2
	} else {
		dp[0][sum+nums[0]] = 1
		dp[0][sum-nums[0]] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < t; j++ {
			// 边界
			l := 0
			if j-nums[i] >= 0 {
				l = j - nums[i]
			} else {
				l = 0
			}
			r := 0
			if j+nums[i] < t {
				r = j + nums[i]
			} else {
				r = 0
			}
			dp[i][j] = dp[i-1][l] + dp[i-1][r]
		}
	}

	return dp[n-1][sum+target]

}
