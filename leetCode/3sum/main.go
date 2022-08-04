package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	length := len(nums)
	sort.Ints(nums)
	for i := 0; i < length-2; i++ { // 双指针法
		if nums[i] > 0 { // 第一个数大于0，后面的数都比它打，肯定不能成立
			break
		}
		if i > 0 && nums[i] == nums[i-1] { // 去掉重复的情况
			continue
		}
		target := -nums[i]
		left := i + 1
		right := length - 1
		for left < right {
			if nums[left]+nums[right] == target {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				// 增加left，减小right，但是不能重复
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++

				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if nums[left]+nums[right] < target {
				left++
			} else {
				right--
			}
		}

	}
	return ans
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
