package main

import (
	"fmt"
	"reflect"
)

func fib(n int) int {
	a, b := 0, 1
	if n == 0 {
		return a
	}
	i := 1
	for i <= n {
		a, b = b, a+b
		i++
		if b >= 1e9+7 {
			b %= 1e9 + 7
		}
	}
	return b
}
func main() {
	var s int
	s = 1e9 + 7
	fmt.Println(s, reflect.TypeOf(s))
}
