package main

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//
//func canPartition(nums []int) bool {
//	n := len(nums)
//
//	sum := 0
//	for _, num := range nums {
//		sum += num
//	}
//	middle := sum / 2
//	// 同为奇数 就不可能了
//	if middle*2 != sum {
//		return false
//	}
//
//	p := make([]int, middle+1)
//
//	for i := 0; i < n; i++ {
//		t := nums[i]
//		for j := middle; j >= 0; j-- {
//			// 不选第i件物品
//			a := p[j]
//			// 选第i件物品
//			b := 0
//			if j >= t {
//				b = p[j-t] + t
//			}
//			p[j] = max(a, b)
//		}
//	}
//	return p[middle] == middle
//}

//func canPartition(nums []int) bool {
//	n := len(nums)
//	sum := 0
//	for _, num := range nums {
//		sum += num
//	}
//	if sum&1 != 0 { // 等价于 sum % 2
//		return false
//	}
//	target := sum >> 1 //  等价于 sum / 2
//	// guard
//	paths := map[string]bool{}
//	var dfs func(index, cur int) bool
//	dfs = func(index, cur int) bool {
//		if index == n || cur > target {
//			return false
//		}
//		if cur == target {
//			return true
//		}
//		key := fmt.Sprintf("%d_%d", index, cur)
//		if val, ok := paths[key]; ok {
//			return val
//		}
//		// select/ no select
//		val := dfs(index+1, cur) || dfs(index+1, cur+nums[index])
//		paths[key] = val
//		return val
//	}
//	return dfs(0, 0)
//}

func canPartition(nums []int) bool {
	sum, target := 0, 0
	maxSum := 0
	for i := range nums {
		sum += nums[i]
		maxSum = Max(maxSum, nums[i])
	}
	if sum&1 != 0 { //  sum % 2
		return false
	}
	target = sum >> 1
	// 最大值 大于 target 就无法 分割了
	if maxSum > target {
		return false
	}
	dp := make([][]bool, len(nums)+1) // 代表 条件 第几个元素，从1 开始
	for i := range dp {
		dp[i] = make([]bool, target+1) // j 要有 target 的情况🐴
	}
	dp[0][0] = true // 默认选0 为true 虚拟值
	for i := 1; i <= len(nums); i++ {
		for j := 0; j <= target; j++ {
			v := nums[i-1] // 第i个物品
			if v > j {     // j 为此时 最大的承受数， 可以看做 数组中的值 为 重量，j 为此时 最大的承受重量	// 画图很明显哈哈
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v] // 选/不选第i - 1 个数组元素
			}
		}
	}
	return dp[len(nums)][target]
}

