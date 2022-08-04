package main

import "fmt"

func nextPermutation(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}

	// i, j, k 都是从后往前, k 到 end 都为降序
	i, j := length-2, length-1
	// 找到nums[i], nums[j]的位置
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	k := j
	// 不是最后一个排序
	if i >= 0 {
		for k <= length-1 && nums[k] > nums[i] {
			k++
		}
		k--
		// 将a[i] 与 a[j] 替换
		nums[i], nums[k] = nums[k], nums[i]
	}
	for i, j := j, length-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func main() {
	//nums := []int{1, 2, 3}
	nums := []int{2, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}
