package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scanln(&n)
	treeArr := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &treeArr[i])
	}

	//fmt.Println(treeArr)
	maxPrice := treeArr[0]
	//queue := []int{treeArr[0]}
	//i := 0
	//for len(treeArr) != 0 {
	//	q := treeArr[0]
	//
	//}
	for i := 1; i < n; i++ {
		treeArr[i] += treeArr[(i-1)/2]
		maxPrice = max(maxPrice, treeArr[i])
	}
	fmt.Println(maxPrice)
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
