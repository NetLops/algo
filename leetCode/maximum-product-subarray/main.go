package main

func maxProduct(nums []int) int {
	ans := nums[0]
	mins := make([]int, len(nums))
	maxs := make([]int, len(nums))
	mins[0], maxs[0] = nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		// 找每次的最小值
		mins[i] = min(nums[i], min(nums[i]*maxs[i-1], mins[i-1]*nums[i]))
		// 找每次的最大值
		maxs[i] = max(nums[i], max(nums[i]*mins[i-1], maxs[i-1]*nums[i]))
		if ans < maxs[i] {
			ans = maxs[i]
		}
	}
	// fmt.Println(mins)
	// fmt.Println(maxs)
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func main() {

}
