package main

func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	ans := []int{}

	var dfs func(i, j int)
	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[0]))
	}
	column, row := len(matrix[0]), len(matrix)
	//flag := 'r'
	count := 0
	dfs = func(i, j int) {
		if row == 0 || column == 0 {
			return
		}
		ans = append(ans, matrix[i][j])
		count++
		if count == len(matrix)*len(matrix[0]) {
			return
		}
		// 往右走
		if i == len(matrix)-row && j < column-1 {
			dfs(i, j+1)
			return
		}
		// 往下走
		if j == column-1 && i < row-1 {
			dfs(i+1, j)
			return
		}
		// 往左走
		if i == row-1 && j > len(matrix[0])-column {
			if j-1 == len(matrix[0])-column {
				row--
			}
			dfs(i, j-1)
			return
		}
		// 往上走
		if j == len(matrix[0])-column && i > len(matrix)-row {
			if i-1 == len(matrix)-row {
				column--
			}
			dfs(i-1, j)
			return
		}
	}
	dfs(0, 0)
	return ans
}
func main() {

}
