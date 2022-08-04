package main

import (
	"fmt"
)

//func maxSubArray(nums []int) int {
//	max := nums[0]
//	for i := 0; i < len(nums); i++ {
//		temp := nums[i]
//		for j := i + 1; j < len(nums); j++ {
//			temp += nums[j]
//			if temp > max {
//				max = temp
//			}
//		}
//	}
//	//for i := range maxs {
//	//	fmt.Println(maxs[i])
//	//}
//	return max
//}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
func maxSubArray(nums []int) int {
	m := nums[0]
	pre := 0
	for i := 0; i < len(nums); i++ {
		pre = max(pre+nums[i], nums[i])
		m = max(pre, m)
	}
	return m
}
func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
