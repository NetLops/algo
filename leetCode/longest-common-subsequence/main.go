package main

import "fmt"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func longestCommonSubsequence(text1 string, text2 string) int {
	n := len(text1)
	m := len(text2)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if text1[i-1] == text2[j-1] {
				f[i][j] = f[i-1][j-1] + 1
			} else {
				f[i][j] = max(f[i-1][j], f[i][j-1])
			}
		}
	}
	fmt.Println(f)
	return f[n][m]
}
func main() {
	text1 := "abcde"
	text2 := "ace"

	//text1 := "abcbdab"
	//text2 := "bdcaba"
	fmt.Println(longestCommonSubsequence(text1, text2))
}
