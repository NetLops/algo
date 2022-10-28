package main

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func trap(height []int) (sum int) {

	// 最左 最右 一定没有了 有就漏了
	for i := 1; i < len(height); i++ {
		maxLeft := 0
		// 找出左边最高
		for j := i - 1; j >= 0; j-- {
			if height[j] > maxLeft {
				maxLeft = height[j]
			}
		}

		maxRight := 0
		// 找出右边最高的
		for j := i + 1; j < len(height); j++ {
			if height[j] > maxRight {
				maxRight = height[j]
			}
		}

		// 找出两端最小的
		minMiddle := Min(maxLeft, maxRight)
		// 只有较小的一端大于当前列的高度才会有谁， 其他情况不会有水
		if minMiddle > height[i] {
			sum += minMiddle - height[i]
		}
	}
	return sum
}
