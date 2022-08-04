package main

import (
	"fmt"
	"strconv"
)

func translateNum(num int) int {
	s := strconv.Itoa(num)
	// pre -> f(n - 2) q -> f(n - 1)
	pre, q, ans := 0, 0, 1
	for i := 0; i < len(s); i++ {
		pre, q, ans = q, ans, 0
		ans += q
		if i == 0 { // 这个就直接跳过了
			continue
		}
		segment := s[i-1 : i+1]
		if segment >= "10" && segment <= "25" {
			ans += pre // 加上 f(n - 2) 组成 f(n - 2) + f( n - 1 )
		}
	}
	return ans
}
func main() {
	fmt.Println(translateNum(12258))
}
