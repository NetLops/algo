package main

func subarraySum(nums []int, k int) int {
	ans := 0

	n := len(nums)
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			if sum == k {
				ans++
			}
		}
	}
	return ans
}
