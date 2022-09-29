package main

import (
	"fmt"
	"strings"
)

//
//func isFlipedString(s1 string, s2 string) bool {
//	if len(s1) != len(s2) {
//		return false
//	}
//	if len(s1) == 0 && len(s2) == 0 {
//		return true
//	}
//	start := 0
//	s3 := s2 + s2
//	for i := 0; i < len(s3); i++ {
//		if start == len(s1) {
//			return true
//		}
//		if s1[start] == s3[i] {
//			start++
//		} else {
//			// fmt.Println(start,i)
//			start = 0
//			i = i - start
//		}
//	}
//	// fmt.Println(start, s3)
//	return start == len(s1)
//}
func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s2 = s2 + s2
	return strings.Contains(s2, s1)
}
func main() {
	fmt.Println(isFlipedString("aba", "bab"))

	fmt.Println(strings.IndexByte("aba", 'a'))
}
