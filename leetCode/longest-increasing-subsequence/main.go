package main

//
//func max(a, b int) int {
//	if a < b {
//		return b
//	}
//	return a
//}
//func lengthOfLIS(nums []int) int {
//	if nums == nil || len(nums) == 0 {
//		return 0
//	}
//
//	maxAns := 0
//
//	dp := make([]int, len(nums))
//	// 每个位置自己就是子集
//	for i := range dp {
//		dp[i] = 1
//	}
//	for i := 0; i < len(nums); i++ {
//		// 范围 [0, j)
//		for j := 0; j < i; j++ {
//			if nums[i] > nums[j] {
//				// dp [0,i-1]
//				dp[i] = max(dp[j]+1, dp[i])
//			}
//		}
//		maxAns = max(dp[i], maxAns)
//	}
//	// fmt.Println(dp)
//	return maxAns
//}
func lengthOfLIS(nums []int) int {

	if nums == nil || len(nums) == 0 {
		return 0
	}
	tail := []int{}
	j := 0 // 数组长度
	for i := 0; i < len(nums); i++ {
		if j == 0 {
			tail = append(tail, nums[i])
			j++
		} else {
			left, right := 0, j-1
			for left <= right {
				m := left + (right-left)>>1
				if tail[m] < nums[i] {
					left = m + 1
				} else {
					right = m - 1
				}
			}
			// 超过就代表往后加一手了
			if left >= j {
				tail = append(tail, nums[i])
				j++
			} else { // 没超过就替换那个最大的
				tail[left] = nums[i]
			}
		}
		// fmt.Println(tail)
	}

	return j
}
func main() {

}
