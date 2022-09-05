package main

import "sort"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func findLongestChain(pairs [][]int) int {
	if pairs == nil || len(pairs) == 0 {
		return 0
	}
	if len(pairs) == 1 {
		return 1
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][0] < pairs[j][0] {
			return true
		}
		return false
	})
	dp := make([]int, len(pairs))
	for i, pair := range pairs {
		dp[i] = 1
		for j, q := range pairs[:i] { // 0...i-1
			if pair[0] > q[1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return dp[len(pairs)-1]
}
func main() {

}
