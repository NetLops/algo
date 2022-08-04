package main

//func findNumberIn2DArray(matrix [][]int, target int) bool {
//	if matrix == nil || len(matrix) == 0 || (len(matrix) > 0 && len(matrix[0]) == 0) {
//		return false
//	}
//	// 找到target 处于那行上
//	columnLen, rowLen := len(matrix), len(matrix[0])
//	top, bottom := 0, columnLen-1
//	m := 0
//	for top <= bottom {
//		m = top + (bottom-top)>>1
//		//fmt.Println(target, m, top)
//		if matrix[m][0] < target {
//			top = m + 1
//		} else if matrix[m][0] > target {
//			bottom = m - 1
//		} else {
//			return true
//		}
//	}
//	fmt.Println(top, top)
//	if top == 0 {
//		top = 1
//	}
//	for i := 0; i < top; i++ {
//		left, right := 0, rowLen-1
//		for left <= right {
//			m := left + (right-left)>>1
//			if matrix[i][m] < target {
//				left = m + 1
//			} else if matrix[i][m] > target {
//				right = m - 1
//			} else {
//				return true
//			}
//		}
//	}
//
//	return false
//}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || (len(matrix) > 0 && len(matrix[0]) == 0) {
		return false
	}
	rowLen, columnLen := len(matrix), len(matrix[0])
	left, bottom := columnLen-1, 0
	for bottom <= rowLen-1 && left >= 0 {
		if matrix[bottom][left] > target {
			left--
		} else if matrix[bottom][left] < target {
			bottom++
		} else {
			return true
		}
	}
	return false
}
func main() {
	martrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	//martrix := [][]int{{}}
	findNumberIn2DArray(martrix, 5)

}

//
//
//[1,2,3,4,5]
//[6,7,8,9,10]
//[11,12,13,14,15]
//[16,17,18,18,20]
