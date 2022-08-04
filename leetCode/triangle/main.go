package main

import (
	"fmt"
	"math"
)

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(triangle))
}

//func minimumTotal(triangle [][]int) int {
//	min := math.MaxInt
//	findMin(&triangle, 0, 0, 0, &min)
//	return min
//}
//
//func findMin(triangle *[][]int, i, j, temp int, min *int) {
//	if i >= len(*triangle) {
//		if *min > temp {
//			*min = temp
//		}
//		return
//	}
//	temp += (*triangle)[i][j]
//	findMin(triangle, i+1, j, temp, min)
//	findMin(triangle, i+1, j+1, temp, min)
//	// 最左的情况
//
//	// 中间部分情况
//
//	// 最右的情况
//}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//func minimumTotal(triangle [][]int) int {
//	n := len(triangle)
//	ans := math.MaxInt
//	f := make([][]int, n)
//	for i := range f {
//		f[i] = make([]int, n)
//	}
//	f[0][0] = triangle[0][0]
//	for i := 1; i < n; i++ {
//		for j := 0; j < i+1; j++ {
//			val := triangle[i][j]
//			f[i][j] = math.MaxInt
//			if j != 0 {
//				f[i][j] = min(f[i][j], f[i-1][j-1]+val)
//			}
//			if j != i {
//				f[i][j] = min(f[i][j], f[i-1][j]+val)
//			}
//		}
//	}
//	for i := 0; i < n; i++ {
//		ans = min(ans, f[n-1][i])
//	}
//	return ans
//}

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	ans := math.MaxInt
	f := make([][]int, 2)
	for i := range f {
		f[i] = make([]int, n)
	}
	f[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		for j := 0; j < i+1; j++ {
			val := triangle[i][j]
			f[i&1][j] = math.MaxInt
			if j != 0 {
				f[i&1][j] = min(f[i&1][j], f[(i-1)&1][j-1]+val)
			}
			if j != i {
				f[i&1][j] = min(f[i&1][j], f[(i-1)&1][j]+val)
			}
		}
	}
	for i := range triangle {
		ans = min(ans, f[(n-1)&1][i])
	}
	return ans
}
