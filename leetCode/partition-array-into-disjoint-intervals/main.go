package main

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func partitionDisjoint(nums []int) int {
	mins := make([]int, len(nums))
	mins[len(nums)-1] = nums[len(nums)-1]
	max := 0
	for i := len(nums) - 2; i >= 0; i-- {
		mins[i] = Min(nums[i], mins[i+1])
	}
	for i := 0; i < len(nums)-1; i++ {
		max = Max(max, nums[i])

		if max <= mins[i+1] {
			return i + 1
		}
	}
	return -1
}
