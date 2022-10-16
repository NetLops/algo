package main

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func maxAscendingSum(nums []int) int {
	// 单调栈

	ans := 0
	max := 0
	for i := range nums {
		if i > 0 && nums[i] < nums[i-1] {
			ans = 0
		}
		ans += nums[i]
		max = maxInt(max, ans)
	}
	return max
}
