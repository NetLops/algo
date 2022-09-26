package main

func missingTwo(nums []int) []int {
	n := len(nums) + 2
	sum := (1 + n) * n / 2
	for i := 0; i < len(nums); i++ {
		sum -= nums[i]
	}
	t := sum / 2 // 看手中值
	cur := (t + 1) * t / 2
	for i := 0; i < len(nums); i++ {
		if nums[i] <= t {
			cur -= nums[i]
		}
	}
	return []int{cur, sum - cur}
}
