package main

import "fmt"

//func numWays(n int) int {
//	// dp[i] = dp[i-1] + dp[i+1]
//	if n == 0 {
//		return 1
//	}
//	a, b := 1, 1
//	for i := 0; i < n; i++ {
//		a, b = b, (a+b)%(1e9+7)
//	}
//	return a
//}
func numWays(n int) int {
	a, b := 1, 1
	i := 1
	for i < n {
		a, b = b, a+b
		i++
		if b > 1e9+7 {
			b %= 1e9 + 7
		}
	}
	return b

}
func main() {
	fmt.Println(numWays(2))
	fmt.Println(numWays(7))
}
