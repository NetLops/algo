package main

import (
	"sort"
)

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	count := 0
	ans := []int{}
	for i := 0; i < len(ans); i++ {
		if count <= k {
			ans = append(ans, ans[i])
		}

	}
	return ans

}
func main() {

}
