package main

import "fmt"

func longestPalindrome(s string) int {
	ints := make([]int, 128, 128)
	ans := 0
	for _, i := range s {
		if ints[i] == 0 {
			ints[i]++
		} else {
			ints[i] = 0
			ans += 2
		}
	}
	for _, count := range ints {
		if count != 0 {
			ans++
			break
		}
	}
	return ans
}
func main() {
	s := "abccccdd"
	fmt.Println(longestPalindrome(s))
}
