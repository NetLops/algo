package main

func nextPermutation(nums []int) {
	// 特殊情况
	if len(nums) <= 1 {
		return
	}
	i, j := 0, 0
	// 找到第一个 i 使得 nums[i] < nums[i+1]，此时 [i+1, n) 一定为降序区间
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	// 如果第1步找到了 i(未找到i=-1), 则在 [i+1, n) 区间中从后向前寻找第一个大于 nums[i] 的值 nums[j]，随后进行两个值的交换
	if i >= 0 {
		for j = len(nums) - 1; j > i; j-- {
			if nums[j] > nums[i] {
				break
			}
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	// 原地交换 [i+1, n) 区间内的元素，使其变为升序，无需对该区间进行排序
	for i, j = i+1, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
