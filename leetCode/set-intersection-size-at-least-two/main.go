package main

import (
	"fmt"
	"sort"
)

func intersectionSizeTwo(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		a, b := intervals[i], intervals[j]
		if a[1] == b[1] {
			return a[0] > b[0]
		} else {
			return a[1] < b[1]
		}
	})
	fmt.Println(intervals)
	a, b, ans := -1, -1, 0
	for _, interval := range intervals {
		if interval[0] > b {
			a, b = interval[1]-1, interval[1]
			ans += 2
		} else if interval[0] > a {
			a, b = b, interval[1]
			ans++
		}
	}
	return ans
}
func main() {
	intervals := [][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}}
	fmt.Println(intersectionSizeTwo(intervals))
}
