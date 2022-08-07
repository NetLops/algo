package main

import "fmt"

//
//func missingNumber(nums []int) int {
//	left, right := 0, len(nums)-1
//	for left <= right {
//		m := left + (right-left)>>1
//		if nums[m] == m { // 往右走
//			left = m + 1
//		} else { // 只会左边缺少
//			right = m - 1
//		}
//	}
//	return left
//}
//func missingNumber(nums []int) int {
//	for i, num := range nums {
//		if i != num {
//			return i
//		}
//	}
//	return 0
//}
func missingNumber(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		m := left + (right-left)>>1
		if nums[m] == m {
			left = m + 1
		} else {
			right = m - 1
		}
	}
	return left
}

func main() {
	nums := []int{0, 1}
	fmt.Println(missingNumber(nums))
}
