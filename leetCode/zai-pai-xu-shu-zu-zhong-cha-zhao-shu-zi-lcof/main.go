package main

import "fmt"

//
//func search(nums []int, target int) int {
//	count := 0
//	length := len(nums)
//	left, right := 0, length-1
//	//if length == 1 {
//	//	if nums[0] == target {
//	//		return 1
//	//	} else {
//	//		return 0
//	//	}
//	//}
//	for left <= right {
//		mid := left + (right-left)>>1
//		if target < nums[mid] {
//			right = mid - 1
//		} else if target > nums[mid] {
//			left = mid + 1
//		} else {
//
//			left, right = mid, mid
//			count++
//			// 左移动
//			for left > 0 {
//				left--
//				if nums[left] == target {
//					count++
//				} else {
//					break
//				}
//			}
//			// 右移动
//			for right < length-1 {
//				right++
//				if nums[right] == target {
//					count++
//				} else {
//					break
//				}
//			}
//			break
//		}
//
//	}
//	return count
//}

func search(nums []int, target int) int {
	ans := 0
	if nums == nil || len(nums) == 0 {
		return ans
	}
	left, right := 0, len(nums)-1

	for left <= right {
		middle := left + (right-left)>>1
		if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] > target {
			right = middle - 1
		} else { //  也就是找到了
			ans++
			left, right = middle, middle
			left--
			right++
			for left >= 0 {
				if nums[left] == target {
					ans++
					left--
					continue
				}
				break
			}
			for right <= len(nums)-1 {
				if nums[right] == target {
					ans++
					right++
					continue
				}
				break
			}
			break
		}

	}
	return ans
}
func main() {
	//nums := []int{5, 7, 7, 8, 8, 10}
	nums := []int{5, 7, 7, 8, 8, 10}
	//nums := []int{1}
	//nums := []int{1, 4}
	target := 8
	fmt.Println(search(nums, target))
}
