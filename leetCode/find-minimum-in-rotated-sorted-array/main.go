package main

import "fmt"

//func findMin(nums []int) int {
//	length := len(nums)
//	left, right, middle := 0, length-1, 0
//	min := nums[left]
//	for left <= right {
//		middle = left + (right-left)>>1
//		if min > nums[middle] {
//			min = nums[middle]
//		}
//		if nums[left] <= nums[middle] {
//			if nums[left] < min {
//				right = middle - 1
//			} else {
//				left = middle + 1
//			}
//		} else { // 右边保证有序
//			right = middle - 1
//		}
//	}
//	return min
//}

func findMin(nums []int) int {
	left, right := 0, len(nums)
	for left <= right {
		m := left + (right-left)>>1
		if nums[m] > nums[right] {
			left = m + 1
		} else {
			right = m
		}
	}
	return nums[left]
}

func main() {
	//nums := []int{3, 4, 5, 1, 2}
	//nums := []int{4, 5, 6, 7, 0, 1, 2}
	nums := []int{11, 13, 15, 17}
	fmt.Println(findMin(nums))
}
