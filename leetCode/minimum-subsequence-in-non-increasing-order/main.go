package main

import (
	"sort"
)

func minSubsequence(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		if nums[i] > nums[j] {
			return true
		}
		return false
	})
	sumIntsRight := make([]int, len(nums))
	sumIntsLeft := make([]int, len(nums))
	temp := 0
	for i := range nums {
		temp += nums[i]
		sumIntsRight[i] += temp
	}

	temp = 0
	for i := len(nums) - 1; i >= 0; i-- {
		temp += nums[i]
		sumIntsLeft[i] += temp
	}
	//fmt.Println(nums)
	//fmt.Println(sumIntsLeft)
	//fmt.Println(sumIntsRight)
	for i := 0; i < len(nums); i++ {
		if i < len(nums)-1 {
			if sumIntsRight[i] > sumIntsLeft[i+1] {
				return nums[:i+1]
			}
		} else {
			return nums[:]
		}
	}

	return nums
}
func main() {
	s := []int{4, 3, 10, 9, 8}
	minSubsequence(s)
}
