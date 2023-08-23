package main

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	ans := [][]int{}
	if n < 3 {
		return nil
	}
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			if left > i+1 && nums[left] == nums[left-1] {
				left++
				continue
			}
			temp := nums[i] + nums[left] + nums[right] // 计算三个数的
			if temp == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				left, right = left+1, right-1
			} else if temp < 0 {
				left++
			} else {
				right--
			}

		}
	}
	return ans
}
