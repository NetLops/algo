package main

import "fmt"

func search(nums []int, target int) bool {
	length := len(nums)
	left, right, middle := 0, length-1, 0
	for left <= right {
		middle = left + (right-left)>>1
		if nums[middle] == target {

			return true
		}
		if nums[left] == nums[middle] && nums[right] == nums[middle] {
			left++
			right = middle - 1
			continue
		}
		// 划分界限
		if nums[left] <= nums[middle] {
			// 划分界限
			if nums[left] <= target && target <= nums[middle] {
				right = middle - 1

			} else {
				left = middle + 1
			}
		} else { // 得保证这个有序
			if nums[middle] <= target && target <= nums[right] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
	}

	return false
}

func main() {
	//nums := []int{2, 5, 6, 0, 0, 1, 2}
	//target := 0
	//nums := []int{1, 0, 1, 1, 1}
	//target := 0

	//nums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}
	//target := 2
	//nums := []int{1, 3, 3, 3, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	//target := 3
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}
	target := 2
	fmt.Println(search(nums, target))
}

//13333111111111
