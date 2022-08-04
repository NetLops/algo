package main

import "fmt"

func uniquePaths(m int, n int) int {
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}
	// 三种情况 往下 往右 又能往下也能往右
	f[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 往下/往右
			if j > 0 && i > 0 {
				f[i][j] = f[i][j-1] + f[i-1][j]
			} else if i > 0 { // 往下
				f[i][j] = f[i-1][j]
			} else if j > 0 { // 往右
				f[i][j] = f[i][j-1]
			}

		}
	}
	return f[m-1][n-1]
}
func main() {
	m := 3
	n := 7
	fmt.Println(uniquePaths(m, n))
}
