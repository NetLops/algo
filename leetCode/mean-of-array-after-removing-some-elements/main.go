package main

import (
	"fmt"
	"sort"
)

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	n := len(arr)
	left := float64(n) * 0.05
	avg := 0
	var i, v int
	for i, v = range arr[int(left) : n-int(left)] {
		avg += v
	}

	return float64(avg) / float64(i+1)
}
func main() {
	fmt.Println(trimMean([]int{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3}))
	fmt.Println(trimMean([]int{6, 0, 7, 0, 7, 5, 7, 8, 3, 4, 0, 7, 8, 1, 6, 8, 1, 1, 2, 4, 8, 1, 9, 5, 4, 3, 8, 5, 10, 8, 6, 6, 1, 0, 6, 10, 8, 2, 3, 4}))
}
