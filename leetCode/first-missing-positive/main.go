package main

func firstMissingPositive(nums []int) int {
	m := make([]int, 500001)
	for i := range nums {
		if nums[i] > 0 && nums[i] < 500001 {
			m[nums[i]] = 1
		}
	}
	for i := 1; i < 500001; i++ {
		if m[i] == 0 {
			return i
		}
	}
	return 0
}
