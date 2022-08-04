package main

import "fmt"

const (
	START = iota
	END
)

func searchRange(nums []int, target int) []int {
	position := make([]int, 2, 2)
	position[START] = -1
	position[END] = -1
	length := len(nums)
	if length <= 0 {
		return position
	}
	low := 0
	height := length - 1
	for i := (low + height) >> 1; i < length && i >= 0 && low <= height; i = (low + height) >> 1 {
		if nums[i] < target {
			low = i + 1
		} else if nums[i] > target {
			height = i - 1
		}
		if nums[i] == target {
			position[START], position[END] = i, i
			// 找开始结点
			for j := position[END] - 1; j >= 0; j-- {
				if nums[j] != target {
					break
				}
				position[START]--
			}
			// 找结束结点
			for j := position[END] + 1; j < length; j++ {
				if nums[j] != target {
					break
				}
				position[END]++
			}
			break
		}

	}
	return position
}

func searchRange2(nums []int, target int) []int {
	position := make([]int, 2, 2)
	position[START] = -1
	position[END] = -1
	length := len(nums)
	if length <= 0 {
		return position
	}
	low := 0
	height := length - 1
	for i := (low + height) >> 1; i < length && i >= 0 && low <= height; i = (low + height) >> 1 {
		if nums[i] < target {
			low = i + 1
		} else if nums[i] > target {
			height = i - 1
		}
		if nums[i] == target {
			position[START], position[END] = i, i
			// 找开始结点
			for j := position[END] - 1; j >= 0; j-- {
				if nums[j] != target {
					break
				}
				position[START]--
			}
			// 找结束结点
			for j := position[END] + 1; j < length; j++ {
				if nums[j] != target {
					break
				}
				position[END]++
			}
			break
		}

	}
	return position
}

func main() {
	//nums := []int{5, 7, 7, 8, 8, 10}
	//nums := []int{5, 7, 7, 8, 8, 10}
	nums := []int{1}
	target := 1
	result := searchRange(nums, target)
	fmt.Println(result)
}
