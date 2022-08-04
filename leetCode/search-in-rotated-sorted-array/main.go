package main

import "fmt"

// 将数组一分为二，其中一定有一个是有序的，另一个可能是有序，也能是部分有序。
//此时有序部分用二分法查找。无序部分再一分为二，其中一个一定有序，另一个可能有序，可能无序。就这样循环.
func search(nums []int, target int) int {
	length := len(nums)
	left, right, middle := 0, length-1, 0
	for left <= right {
		middle = left + (right-left)>>1
		if nums[middle] == target {
			return middle
		}
		// 可能左侧递增，左开右闭区间
		if nums[left] <= nums[middle] {
			if nums[left] <= target && target <= nums[middle] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		} else { // 右侧一定递增 也是 左开右闭区间
			if nums[middle] < target && target <= nums[right] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
	}
	return -1
}
func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0

	//nums := []int{3, 1}
	//target := 1
	fmt.Println(search(nums, target))
}
