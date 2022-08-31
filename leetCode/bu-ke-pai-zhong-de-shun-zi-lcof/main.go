package main

import (
	"sort"
)

func isStraight(nums []int) bool {
	sort.Ints(nums)

	kingCount := 0
	i := 0
	for i = 0; i < len(nums)-1; i++ {
		// 大小王最多只有两张
		if nums[0] == 0 {
			kingCount++
			if kingCount > 2 {
				// 哪儿2张以上的大小王
				return false
			}
			continue
		} else {
			// 有两个数 相同的情况也会排除
			if nums[i] == nums[i+1] {
				return false
			}
		}
	}
	// 小于5 就代表 能成顺子(因为没有重复的)
	return nums[4]-nums[i-1] < 5

}
func main() {

}
