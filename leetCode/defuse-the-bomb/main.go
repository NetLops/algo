package main

func decrypt(code []int, k int) []int {

	n := len(code)
	ans := make([]int, len(code))
	if k == 0 {
		return ans
	}
	// copy(ans, code)
	for i := 0; i < len(code); i++ {
		if k > 0 {
			for j := 1; j <= k; j++ {
				ans[i] += code[(i+j)%n]
			}
		} else if k < 0 {
			for j := 1; j <= -k; j++ {
				ans[i] += code[(n+(i-j)%n)%n]
				// fmt.Println((n+(i-j)%n)%n,code[(n+(i-j)%n)%n],code,code[2])
			}
		}
	}
	return ans
}
func main() {
	//fmt.Println((4 - 7%4))
}
