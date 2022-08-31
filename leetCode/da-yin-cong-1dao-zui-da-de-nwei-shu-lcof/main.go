package main

func printNumbers(n int) []int {

	if n <= 0 {
		return []int{}
	}
	max := 0
	for j := 1; j <= n; j++ {
		max += 10*max + 9
	}
	ans := make([]int, max)
	for i := 0; i < max; i++ {
		ans[i] = i + 1
	}
	return ans
}
func main() {

}
