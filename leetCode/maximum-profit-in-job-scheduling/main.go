package main

import "sort"

// 参考：https://leetcode.cn/problems/maximum-profit-in-job-scheduling/solution/by-ac_oier-rgup/
func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	list := [][]int{}
	ans := make([]int, n+10)
	for i := 0; i < n; i++ {
		list = append(list, []int{startTime[i], endTime[i], profit[i]})
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i][1] < list[j][1]
	})
	for i := 1; i <= n; i++ {
		info := list[i-1]
		a, _, c := info[0], info[1], info[2]
		ans[i] = Max(ans[i-1], c)
		l, r := 0, i-1
		//  二分法优化一下前驱
		for l < r {
			mid := (l + r + 1) >> 1 // 边界向右
			if list[mid][1] <= a {
				l = mid
			} else {
				r = mid - 1
			}
		}
		if list[r][1] <= a {
			ans[i] = Max(ans[i], ans[r+1]+c)
		}
	}
	return ans[n]
}
