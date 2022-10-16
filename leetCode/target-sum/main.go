package main

//func findTargetSumWays(nums []int, target int) int {
//	var dfs func(next, cur int) int
//	next, cur := 0, 0
//	dfs = func(next, cur int) int {
//		if next == len(nums) {
//			if cur == target {
//				return 1
//			} else {
//				return 0
//			}
//		}
//		left := dfs(next+1, cur+nums[next])
//		right := dfs(next+1, cur-nums[next])
//		return left + right
//	}
//	return dfs(next, cur)
//}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

//
//func findTargetSumWays(nums []int, target int) int {
//	sum := 0
//	for _, num := range nums {
//		sum += num
//	}
//	// 绝对值范围超过了 sum 的绝对值返回则无法得到
//	if Abs(target) > Abs(sum) {
//		return 0
//	}
//	n := len(nums)
//	rang := sum*2 + 1 // 因为要包含负数所以要两倍，又要加上0
//	dp := make([][]int, n)
//	for i := range dp {
//		dp[i] = make([]int, rang)
//	}
//	// 加上sum 纯粹是因为下标问题，赋第二维的值的时候 都要加上sum
//	// 初始化 第一个数只能分别组成 +-nums[i] 的一种强开
//	dp[0][sum+nums[0]] += 1
//	dp[0][sum-nums[0]] += 1
//	for i := 1; i < n; i++ {
//		for j := -sum; j <= sum; j++ {
//			if j+nums[i] > sum { // +不成立 加上当前数大于了sum，减去当前的数
//				dp[i][j+sum] = dp[i-1][j-nums[i]+sum] + 0
//			} else if j-nums[i] < -sum { // - 不成立 减去当前数小于-sum 只能加上当前的数
//				dp[i][j+sum] = dp[i-1][j+nums[i]+sum] + 0
//			} else { // 加减都可以
//				dp[i][j+sum] = dp[i-1][j+nums[i]+sum] + dp[i-1][j-nums[i]+sum]
//			}
//		}
//	}
//	return dp[n-1][sum+target]
//}
//
//// 参考：https://leetcode.cn/problems/target-sum/solution/dong-tai-gui-hua-si-kao-quan-guo-cheng-by-keepal/
//// 转移方程 dp[i][j] = dp[i][j + nums[i]] + dp[i][j - nums[i]]
//func findTargetSumWays(nums []int, target int) int {
//	sum := 0
//	for i := range nums {
//		sum += nums[i]
//	}
//	// 如果 sum < target 就不成立了
//	if Abs(target) > Abs(sum) {
//		return 0
//	}
//	// 因为 target有 负数情况，但数组没有 负数下标
//	rang := 2*sum + 1 // [-sum....0.....sum]
//	dp := make([][]int, len(nums))
//	for i := range dp {
//		dp[i] = make([]int, rang)
//	}
//	// 第 0 层初始化
//
//	dp[0][sum+nums[0]] = 1
//	dp[0][sum-nums[0]] += 1 // 毕竟 nums[0] 还可能是 0 ，所以两手考虑用 +=  就是两种情况了
//	for i := 1; i < len(nums); i++ {
//		for j := -sum; j <= sum; j++ {
//			// 考虑边界问题
//			if j+nums[i] > sum { // 加号的情况就不存在了
//				dp[i][j+sum] = dp[i-1][j+sum-nums[i]] + 0
//			} else if j-nums[i] < -sum { // 负号的情况情况就不存在了
//				dp[i][j+sum] = dp[i-1][j+sum+nums[i]] + 0
//			} else {
//				dp[i][j+sum] = dp[i-1][j+sum+nums[i]] + dp[i-1][j+sum-nums[i]]
//			}
//		}
//	}
//	return dp[len(nums)-1][target+sum]
//}

func findTargetSumWays(nums []int, target int) int {
	var dfs func(index, cur int) int
	dfs = func(index, cur int) int {
		if len(nums) == index {
			if cur == target {
				return 1
			} else {
				return 0
			}
		}
		left := dfs(index+1, cur+nums[index])
		right := dfs(index+1, cur-nums[index])
		return left + right
	}
	return dfs(0, 0)
}
