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
//		if nums[left] == nums[middle] && nums[middle] == nums[right] {
//			left++
//			right--
//			continue
//		}
//		if nums[left] <= nums[middle] {
//			if nums[left] < min {
//				right = middle - 1
//			} else {
//				left = middle + 1
//			}
//		} else { // 有序递增
//			right = middle - 1
//		}
//	}
//	return min
//}

func findMin(nums []int) int {
	length := len(nums)
	left, right, middle := 0, length-1, 0
	for left <= right {
		middle = left + (right-left)>>1
		if nums[middle] > nums[right] {
			left = middle + 1
		} else if nums[middle] < nums[right] {
			right = middle
		} else {
			right--
		}
	}
	return nums[left]
}
func main() {
	//nums := []int{1, 3, 5}
	nums := []int{10, 1, 10, 10, 10}
	//nums := []int{2, 0, 1, 1, 1}
	fmt.Println(findMin(nums))
}
