package main

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func trap(height []int) int {
	sum := 0
	for i := 1; i < len(height)-1; i++ {
		max_left := 0
		// 找出左边最高
		for j := i - 1; j >= 0; j-- {
			if height[j] > max_left {
				max_left = height[j]
			}
		}
		max_right := 0
		// 找出右边最高
		for j := i + 1; j < len(height); j++ {
			if height[j] > max_right {
				max_right = height[j]
			}
		}

		// 找出两端较小的
		min_middle := min(max_left, max_right)
		// 只有较小的一段大于当前列的高度才会有水，其他情况不会有水
		if min_middle > height[i] {
			sum += min_middle - height[i]
		}
	}
	return sum
}
