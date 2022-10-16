package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	n := len(nums)
	for k := 0; k < n-2; k++ {
		if nums[k] > 0 { // j > i > k // omit 后续就不存在了，毕竟不会再次出现了 = 0 的情况
			break
		}
		if k > 0 && nums[k] == nums[k-1] { // 去重
			continue
		}
		i, j := k+1, len(nums)-1
		for i < j {
			sum := nums[k] + nums[i] + nums[j]
			if sum < 0 {
				i++
				for i < j && nums[i] == nums[i-1] {
					i++
				}
			} else if sum > 0 {
				j--
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			} else {
				res = append(res, []int{nums[k], nums[i], nums[j]})
				i++
				j--
				for i < j && nums[i] == nums[i-1] {
					i++
				}
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			}
		}
	}
	return res
}
