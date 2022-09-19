package main

import (
	"fmt"
	"strings"
)

func maxLengthBetweenEqualCharacters(s string) int {
	ans := -1
	for i := 0; i < 26; i++ {
		if now := strings.LastIndex(s, string(rune('a'+i))) - strings.Index(s, string(rune('a'+i))) - 1; now > ans {
			ans = now
		}
	}
	return ans
}
func main() {
	fmt.Println(maxLengthBetweenEqualCharacters("adasdafsdfdfafaada"))
}
