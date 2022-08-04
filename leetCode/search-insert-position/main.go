package main

import "fmt"

func searchInsert(nums []int, target int) int {
	result := 0
	length := len(nums)
	if length <= 0 {
		return result
	}
	low := 0
	high := length - 1
	for i := low + (high-low)>>1; low <= high; i = low + (high-low)>>1 {
		if nums[i] < target {
			low = i + 1
		} else if nums[i] >= target {
			high = i - 1
		}
		result = i
	}
	if nums[result] != target && nums[result] < target {
		result++
	}
	return result
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 0
	result := searchInsert(nums, target)
	fmt.Println(result)
}
