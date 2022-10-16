package main

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func scoreOfParentheses(s string) int {
	res, k := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			k++
		} else {
			k--
			if s[i-1] == '(' {
				res += 1 << k
			}
		}
	}
	return res
}
func main() {

}
