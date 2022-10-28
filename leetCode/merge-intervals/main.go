package main

import "sort"

func merge(intervals [][]int) [][]int {
	// 先排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})
	ans := [][]int{}

	left, far := -1, -1
	for _, interval := range intervals {
		start := interval[0]
		end := interval[1]
		if start <= far {
			if end > far {
				far = end
			}

		} else {
			if far >= 0 {
				ans = append(ans, []int{left, far})
			}
			left = start
			far = end
		}
	}
	// 有可能intervals 啥都没有
	if far != -1 {
		ans = append(ans, []int{left, far})
	}
	return ans
}
