package main

// 线搜树的查找方式
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	rowLen, columnLen := len(matrix), len(matrix[0])
	i, j := 0, columnLen-1
	for i < rowLen && j >= 0 {
		if matrix[i][j] < target {
			i++
		} else if matrix[i][j] > target {
			j--
		} else {
			return true
		}
	}
	return false
}
func main() {

}
