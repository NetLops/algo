package main

import "fmt"

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	for i := 1; i < n; i++ {

		if n%i == 0 { //  i必须能被n整除 才可能重复
			match := true
			for j := i; j < n; j++ {
				if s[j] != s[j-i] { // s[j] == s[j-i]	// i倍数相对应的
					match = false
					break
				}
			}
			if match {
				return true
			}
		}
	}
	return false
}
func main() {
	s := "abab"
	fmt.Println(repeatedSubstringPattern(s))
}
