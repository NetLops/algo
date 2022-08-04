package main

import "fmt"

func getPost(pos, m, n int) (i, j int) {
	if pos > (m*n)-1 {
		pos = pos % (m * n)
	}
	return pos / n, pos % n
}
func shiftGrid(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])
	copyGrid := make([][]int, m)
	for i := range copyGrid {
		copyGrid[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			nowPos := i*n + j + k
			x, y := getPost(nowPos, m, n)
			fmt.Println(x, y, " ", i, j, nowPos)
			copyGrid[x][y] = grid[i][j]
		}
	}
	return copyGrid
}
func main() {
	//grid := [][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}}
	grid := [][]int{{1}, {2}, {3}, {4}, {7}, {6}, {5}}
	target := 23
	fmt.Println(shiftGrid(grid, target))

	fmt.Println(getPost(0+23, 7, 1))
}
