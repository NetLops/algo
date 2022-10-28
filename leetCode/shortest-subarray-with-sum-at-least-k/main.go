package main

import "math"

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type Pair struct {
	Pre, Index int
}

// 参考：https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/solution/liang-zhang-tu-miao-dong-dan-diao-dui-li-9fvh/
func shortestSubarray(nums []int, k int) int {
	q := []Pair{{0, -1}}
	ans := math.MaxInt32
	cur := 0 // 当前前缀和
	for i, num := range nums {
		cur += num // 统计前缀和
		// 队列开始
		for len(q) > 0 && cur-q[0].Pre >= k {
			ans = Min(ans, i-q[0].Index)
			q = q[1:]
		}
		for len(q) > 0 && q[len(q)-1].Pre >= cur {
			q = q[:len(q)-1]
		}
		q = append(q, Pair{cur, i})
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}
