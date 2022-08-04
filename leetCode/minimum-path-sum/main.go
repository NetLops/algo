package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func minPathSum(grid [][]int) int {
	f := make([][]int, len(grid))
	column := len(grid[0])
	for i := range f {
		f[i] = make([]int, column)
	}
	left, top := 0, 0
	i, j := 0, 0
	for i = 0; i < len(grid); i++ {
		for j = 0; j < column; j++ {
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
			}
		}
	}
	return f[i-1][j-1]
}

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum(grid))
}
