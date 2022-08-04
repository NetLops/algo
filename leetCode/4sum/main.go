package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	ans := make([][]int, 0)
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left := j + 1
			right := n - 1
			// 双指针法
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					ans = append(ans, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return ans
}
func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	//nums := []int{2, 2, 2, 2, 2}
	//target := 8
	fmt.Println(fourSum(nums, target))
}
