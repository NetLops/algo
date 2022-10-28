package main

func combinationSum(candidates []int, target int) (res [][]int) {
	item := []int{}
	var dfs func(index, sum int)
	dfs = func(index, sum int) {
		if sum == target {
			res = append(res, append([]int{}, item...))
			return
		} else if sum > target || index >= len(candidates) {
			return
		}
		dfs(index+1, sum)
		item = append(item, candidates[index])
		dfs(index, sum+candidates[index])
		item = item[:len(item)-1]
	}
	dfs(0, 0)
	return
}

func generateParenthesis(n int) (ans []string) {
	var dfs func(int, int, int, string)
	dfs = func(index, left, right int, now string) {
		if left > n || right > left {
			return
		}
		if index == 2*n {
			ans = append(ans, now)
			return
		}
		dfs(index+1, left+1, right, now+"(")
		dfs(index+1, left, right+1, now+")")
	}
	dfs(0, 0, 0, "")
	return
}
