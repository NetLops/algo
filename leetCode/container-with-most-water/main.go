package main

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	ans := 0

	for i < j {
		if height[i] < height[j] {
			ans = max(ans, (j-i)*height[i])
			i++
		} else {
			ans = max(ans, (j-i)*height[j])
			j--
		}
	}
	return ans
}
