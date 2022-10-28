package main

func kthGrammar(n int, k int) int {
	var dfs func(i, j, cur int) int
	dfs = func(i, j, cur int) int {
		if i == 1 {
			return cur
		}
		if (j%2 == 0 && cur == 0) || (j%2 == 1 && cur == 1) {
			return dfs(i-1, (j-1)/2+1, 1)
		} else {
			return dfs(i-1, (j-1)/2+1, 0)
		}

	}
	if dfs(n, k, 1) == 0 {
		return 1
	}
	return 0
}
