package main

func matrixReshape(mat [][]int, r int, c int) [][]int {
	rowLen := len(mat)
	if rowLen == 0 {
		return mat
	}
	columnLen := len(mat[0])
	amount := rowLen * columnLen
	if amount/r != c {
		return mat
	}
	ans := make([][]int, r)
	for i := range ans {
		ans[i] = make([]int, c)
	}
	x, y := 0, 0
	for i := 0; i < rowLen; i++ {
		for j := 0; j < columnLen; j++ {
			ans[x][y] = mat[i][j]
			y++
			if y == c {
				y = 0
				x++
			}
		}
	}

	return ans
}
func main() {

}
