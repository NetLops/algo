package main

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maximalRectangle(matrix [][]byte) int {
	row := len(matrix)
	column := len(matrix[0])
	// 前缀和
	sum := make([][]int, row+10)
	for i := range sum {
		sum[i] = make([]int, column+10)
	}
	for i := 1; i <= row; i++ {
		for j := 1; j <= column; j++ {
			if matrix[i-1][j-1] == '0' {
				sum[i][j] = 0
			} else {
				sum[i][j] = sum[i-1][j] + 1
			}
		}
	}
	left, right := make([]int, column+10), make([]int, column+10)
	ans := 0
	for i := 1; i <= row; i++ {
		current := sum[i]
		for i2 := range left {
			left[i2] = 0
		}
		for i2 := range right {
			right[i2] = column + 1
		}
		q := []int{}
		for j := 1; j <= column; j++ {
			for len(q) != 0 && current[q[len(q)-1]] > current[j] {
				x := q[len(q)-1]
				right[x] = j
				q = q[:len(q)-1]
			}
			q = append(q, j)
		}
		q = []int{}
		for j := column; j >= 1; j-- {
			for len(q) != 0 && current[q[len(q)-1]] > current[j] {
				x := q[len(q)-1]
				left[x] = j
				q = q[:len(q)-1]
			}
			q = append(q, j)
		}
		for j := 1; j <= column; j++ {
			ans = max(ans, current[j]*(right[j]-left[j]-1))
		}
	}

	return ans
}
