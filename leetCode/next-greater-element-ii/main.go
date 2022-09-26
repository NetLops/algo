package main

//func max(a, b int) int {
//	if a < b {
//		return b
//	}
//	return a
//}
func nextGreaterElements(nums []int) []int {
	max := 0 // 找到最大下标（靠右侧）
	m := map[int]int{}
	stack := []int{} // 单调栈
	for i := range nums {
		if nums[i] >= nums[max] {
			max = i
		}
	}
	for i := max; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			m[i] = stack[len(stack)-1]
		} else {
			m[i] = -1
		}
		stack = append(stack, i)
	}
	for i := len(nums) - 1; i >= max+1; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			m[i] = stack[len(stack)-1]
		} else {
			m[i] = -1
		}
		stack = append(stack, i)
	}
	ans := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		if m[i] != -1 {
			ans = append(ans, nums[m[i]])
		} else {
			ans = append(ans, -1)
		}
	}
	return ans
}
