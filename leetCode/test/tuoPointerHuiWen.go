package main

import "fmt"

func huiwenDel(s string) int {
	length := len(s)
	count := 0
	if valid(s, 0, length-1) {
		count++
		return 1
	} else {
		for i := 0; i < length-1; i++ {
			count++
			if valid(s, i, length-1) {
				return count
			}
			if valid(s, 0, length-1-i) {
				return count
			}
		}
	}
	return count
}
func valid(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func main() {
	del := huiwenDel("baabb")
	fmt.Println(del)
}
