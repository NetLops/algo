package main

import (
	"fmt"
)

func quick_pow(a, b int) int {
	ans := 1
	for b != 0 {
		if b&1 == 1 {
			ans *= a
		}
		fmt.Println(a)
		a *= a
		b = b >> 1
	}
	return ans
}

func main() {
	fmt.Println(quick_pow(3, 89))
}
