package main

import "sort"

// 回溯
func combinationSum(candidates []int, target int) [][]int {
	// 排序便于后续剪枝
	sort.Ints(candidates)
	var (
		ans [][]int
		cur []int
		dfs func(int, int)
		n   = len(candidates)
	)

	dfs = func(index, sum int) {
		if sum == target {
			t := make([]int, len(cur))
			copy(t, cur)
			ans = append(ans, t)
		} else if sum > target {
			return
		}
		for i := index; i < n; i++ {
			cur = append(cur, candidates[i])
			dfs(i, sum+candidates[i])
			cur = cur[:len(cur)-1]
		}
	}
	dfs(0, 0)
	return ans
}
