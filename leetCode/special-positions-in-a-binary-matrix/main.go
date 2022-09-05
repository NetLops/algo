package main

func numSpecial(mat [][]int) int {
	ans := 0
	if mat == nil || len(mat) == 0 {
		return 0
	}
	rowLen, columnLen := len(mat), len(mat[0])
	dpy := make([]int, columnLen)
	dpx := make([]int, rowLen)
	for i := 0; i < rowLen; i++ {
		for j := 0; j < columnLen; j++ {
			if mat[i][j] == 1 {
				dpx[i] += 1
				dpy[j] += 1
			}

		}
	}
	for i := 0; i < rowLen; i++ {
		for j := 0; j < columnLen; j++ {
			if mat[i][j] == 1 && dpx[i] == 1 && dpy[j] == 1 {
				ans++
			}

		}
	}
	return ans
}
func main() {

}
