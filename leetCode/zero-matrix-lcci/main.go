package main

func setZeroes(matrix [][]int) {
	n := len(matrix)
	m := len(matrix[0])
	col0 := false
	for i := 1; i < n; i++ {
		if matrix[i][0] == 0 {
			col0 = true
		}
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 {
			matrix[i][0] = 0
		}
	}

}
func main() {

}
