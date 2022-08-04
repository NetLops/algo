package main

import "fmt"

//
//func maxValue(grid [][]int) int {
//	ans := 0
//	rowLen := len(grid)
//	if rowLen == 0 {
//		return ans
//	}
//	columnLen := len(grid[0])
//	//f := make([][]int, rowLen)
//	//for i := range f {
//	//	f[i] = make([]int, columnLen)
//	//}
//	f := make([]int, columnLen)
//	f[0] = grid[0][0]
//	for i := 0; i < rowLen; i++ {
//		for j := 0; j < columnLen; j++ {
//			// 第一行不能下走
//			if i == 0 {
//				if j == 0 {
//					f[j] = grid[i][j]
//					continue
//				}
//				f[j] = f[j-1] + grid[i][j]
//				continue
//			}
//			// 第一列不能右走
//			if j == 0 {
//				f[j] = f[0] + grid[i][j]
//				continue
//			}
//			// 右走
//			right := f[j-1] + grid[i][j]
//			// 下走
//			bottom := f[j] + grid[i][j]
//			// 选择最大的
//			if right > bottom {
//				f[j] = right
//			} else {
//				f[j] = bottom
//			}
//
//		}
//	}
//
//	return f[columnLen-1]
//}

func maxValue(grid [][]int) int {
	ans := 0
	rowLen := len(grid)
	if rowLen == 0 {
		return ans
	}
	columnLen := len(grid[0])
	f := make([][]int, rowLen)
	for i := range f {
		f[i] = make([]int, columnLen)
	}
	f[0][0] = grid[0][0]
	for i := 0; i < rowLen; i++ {
		for j := 0; j < columnLen; j++ {
			// 第一行不能下走
			if i == 0 {
				if j == 0 {
					f[i][j] = grid[i][j]
					continue
				}
				f[i][j] = f[0][j-1] + grid[i][j]
				continue
			}
			// 第一列不能右走
			if j == 0 {
				f[i][j] = f[i-1][0] + grid[i][j]
				continue
			}
			// 右走
			right := f[i][j-1] + grid[i][j]
			// 下走
			bottom := f[i-1][j] + grid[i][j]
			// 选择最大的
			if right > bottom {
				f[i][j] = right
			} else {
				f[i][j] = bottom
			}

		}
	}

	//for i := rowLen - 1; i > 0; i-- {
	//	for j := columnLen - 1; j > 0; j-- {
	//		// 看看是不是上边的 或是右边的
	//		if f[i][j]-grid[i][j] == f[i-1][j] {
	//			fmt.Println(grid[i][j])
	//
	//		} else if f[i][j]-grid[i][j] == f[i][j-1] {
	//			fmt.Println(grid[i][j])
	//		}
	//	}
	//}
	i, j := rowLen-1, columnLen-1
	for {
		if i == 0 {
			for k := j; k >= 0; k-- {
				fmt.Printf("->%d", grid[0][k])
			}
			break
		}
		if j == 0 {
			for k := i; k >= 0; k-- {
				fmt.Printf("->%d", grid[i][0])
			}
		}
		//看看是不是上边的 或是右边的
		if f[i][j]-grid[i][j] == f[i-1][j] {
			fmt.Printf("->%d", grid[i][j])
			i--
		} else if f[i][j]-grid[i][j] == f[i][j-1] {
			fmt.Printf("->%d", grid[i][j])
			j--
		}
	}
	return f[rowLen-1][columnLen-1]
}

func main() {

}
