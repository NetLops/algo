package main

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func nextGreaterElement2(n int) int {
	itoa := strconv.Itoa(n)
	length := len(itoa)
	nums := make([]int, length, length)
	for i := range nums {
		nums[i] = int(itoa[i]) - 48
	}
	if length <= 1 {
		return -1
	}
	i, j, k := length-2, length-1, length-1

	// 从后往前找到 nums[i] 小于 num[j]的地方， 此时 [j ,length-1] 必为降序
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	result := -1
	// i >= 0 时 代表有戏
	if i >= 0 {
		result = 0
		// 找到 [j,length - 1] 处 最小k 大于 nums[i]的位置, 此处已降序
		for nums[k] <= nums[i] {
			k--
		}
		// k, i 的位置互换
		nums[k], nums[i] = nums[i], nums[k]

		// 将 [j, end]改成升序即为 下一个序列
		for i, j := j, length-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}

		for i, num := range nums {
			result += num * int(math.Pow10(length-i-1))
			if result > math.MaxInt32 {
				return -1
			}
		}

	}
	return result
}

func nextGreaterElement3(n int) int {
	str := strconv.Itoa(n)
	length := len(str)
	if length <= 1 {
		return -1
	}
	nums := make([]int, length, length)
	for i := range str {
		nums[i] = int(str[i]) - 48
	}
	i, j, k := length-2, length-1, length-1
	// 找到 nums[i] 小于 nums[j] 的位置， 此时 [j, length - 1] 必定为降序
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	result := -1
	// i >= 0  就代表有下一排序
	if i >= 0 {
		result = 0
		// k是 [j, length - 1]降序序列中最小一个大于i位置数的位置 找到k的位置，k用来替换i的位置 nums[k] > nums[i]
		for nums[k] <= nums[i] {
			k--
		}
		nums[k], nums[i] = nums[i], nums[k]
		// 将[j, length] 改为升序, 即原nums 的下一个排列
		for i, j = j, length-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
		// 此时的nums 就是下一序列
		for i, num := range nums {
			result += num * int(math.Pow10(length-i-1))
			if (result < 0) || (result > math.MaxInt32) {
				return -1
			}
		}
	}
	return result
}

func TestName(t *testing.T) {
	fmt.Println(nextGreaterElement3(123))
}
