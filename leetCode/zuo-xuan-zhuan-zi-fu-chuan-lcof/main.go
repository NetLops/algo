package main

import "fmt"

//func reverseLeftWords(s string, n int) string {
//	return s[n:] + s[:n]
//}

//func reverseLeftWords(s string, n int) string {
//	res := make([]uint8, len(s))
//	for i := n; i < len(s)+n; i++ {
//		res[i-n] = s[i%len(s)]
//	}
//	return string(res)
//}

func reverseLeftWords(s string, n int) string {
	length := len(s)
	res := make([]byte, length)
	for i := n; i < len(s)+n; i++ {
		res[i-n] = s[i%length]
	}
	return string(res)
}
func main() {
	s := "abcdefg"
	k := 2
	fmt.Println(reverseLeftWords(s, k))
}
