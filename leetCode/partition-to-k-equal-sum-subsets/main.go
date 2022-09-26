package main

import (
	"fmt"
	"sort"
)

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%k != 0 {
		return false
	}
	// target 即每个子集需要满足的和
	target := sum / k // 必须是整数的，否则就是不会存在符合的答案
	bucket := make([]int, k+1)
	// 先降序处理，增加爆率 // bucket[i]+nums[index] > target
	sort.Slice(nums, func(i, j int) bool {
		if nums[i] > nums[j] {
			return true
		}
		return false
	})
	var travel func(index int) bool
	travel = func(index int) bool {
		// 已经处理完了所有的球
		if index == len(nums) {
			// 此时必定为true了
			//for i := 0; i < k; i++ {
			//	// 有那集合不符合就G了
			//	if bucket[i] != target {
			//		return false
			//	}
			//}
			return true
		}
		// 开始左选择
		for i := 0; i < k; i++ {
			// 如果当前桶和上一个桶内的元素和相等，则跳过
			// 原因：如果元素相等，那么 nums[index] 选择上一个桶和选择当前桶可以得到的结果是一致的
			if i > 0 && bucket[i] == bucket[i-1] {
				continue
			}

			// 剪枝： 放入求后超过 target 的值，选择一下桶
			// 如果 超过桶的 大小就直接跳过，看看下个是不是符合
			if bucket[i]+nums[index] > target {
				continue
			}

			// 做选择 放入i号桶
			bucket[i] += nums[index]
			// 处理下一个球，即 nums[index+ 1]
			if travel(index + 1) {
				return true
			}
			// 不符合就踢出桶 // 换一个看看
			bucket[i] -= nums[index]
		}
		return false
	}
	// 都不满足答案就直接返回false
	return travel(0)
}
func main() {
	fmt.Println(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
	fmt.Println(canPartitionKSubsets([]int{1, 2, 3, 4}, 3))
}
