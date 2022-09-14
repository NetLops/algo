package main

import (
	"fmt"
	"sort"
)

func specialArray(nums []int) int {
	temp := 1
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		if (i > 0 && nums[i-1] < temp) && nums[i] >= temp {
			return temp
		}
		if i == 0 {
			if len(nums) >= nums[i] {
				return len(nums)
			}
		}
		temp++
	}

	return -1

}
func main() {
	fmt.Println(specialArray([]int{0, 4, 3, 0, 4}))
}
