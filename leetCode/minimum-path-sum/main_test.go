package main

import (
	"fmt"
	"math"
	"testing"
)

//func min(a, b int) int {
//	if a > b {
//		return b
//	}
//	return a
//}
func minPathSum2(grid [][]int) int {
	n := len(grid)
	f := make([][]int, n)
	m := len(grid[0])
	g := make([]int, n*m)
	for i := range f {
		f[i] = make([]int, m)
	}
	left, top := 0, 0
	i, j := 0, 0
	for i = 0; i < n; i++ {
		for j = 0; j < m; j++ {
			if i == 0 && j == 0 {
				f[0][0] = grid[0][0]
			} else {
				// 向下走，走不通 就不走
				if i > 0 {
					top = f[i-1][j] + grid[i][j]
				} else {
					top = math.MaxInt
				}
				// 向右走，走不通 就不走
				if j > 0 {
					left = f[i][j-1] + grid[i][j]
				} else {
					left = math.MaxInt
				}
				f[i][j] = min(left, top)
				if top < left {
					g[getIdx(i, j, n)] = getIdx(i-1, j, n)
				} else {
					g[getIdx(i, j, n)] = getIdx(i, j-1, n)
				}

			}
		}
	}

	// 从【结尾】开始，在g[] 数组中找 [上一步]
	idx := getIdx(m-1, n-1, n)
	path := make([][]int, m+n)
	for i2 := range path {
		path[i2] = make([]int, 2)
	}
	for i3 := 1; i3 < m+n; i3++ {
		path[m+n-1-i3] = parseIdx(g[idx], n)
		idx = g[idx]
	}
	fmt.Println(g)
	for i4 := 1; i4 < m+n; i4++ {
		x, y := path[i4][0], path[i4][1]
		fmt.Printf("(%d, %d)\n", x, y)
	}
	return f[i-1][j-1]
}
func getIdx(x, y, n int) int {
	return x*n + y
}
func parseIdx(idx, n int) []int {
	return []int{idx / n, idx % n}
}
func TestName(t *testing.T) {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum2(grid))
}
