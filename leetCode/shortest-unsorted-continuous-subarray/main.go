package main

import "fmt"

//
//func findUnsortedSubarray(nums []int) int {
//	arr := make([]int, len(nums))
//	copy(arr, nums)
//	sort.Ints(arr)
//	i, j := 0, len(nums)-1
//	for i <= j && nums[i] == arr[i] {
//		i++
//	}
//	for i <= j && nums[j] == arr[j] {
//		j--
//	}
//	return j - i + 1
//}
func findUnsortedSubarray(nums []int) int {
	MIN := -100005
	MAX := 100005
	n := len(nums)
	i, j := 0, n-1
	for i < j && nums[i] <= nums[i+1] {
		i++
	}
	for i < j && nums[j] >= nums[j-1] {
		j--
	}
	l, r := i, j
	min, max := nums[i], nums[j]
	// [2,6,4,10,8,9,15] 保证 子区间内 最小的值 的位置，最大值在最右边界 的最合适的位置
	for a := l; a <= r; a++ {
		if nums[a] < min {
			for i >= 0 && nums[i] > nums[a] {
				i--
			}
			if i >= 0 {
				min = nums[i]
			} else {
				min = MIN
			}
		}
		if nums[a] > max {
			for j < n && nums[j] < nums[a] {
				j++
			}
			if j < n {
				max = nums[j]
			} else {
				max = MAX
			}
		}
	}
	if j == i {
		return 0
	} else {
		return j - 1 - (i + 1) + 1
	}
}

func main() {
	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
}
