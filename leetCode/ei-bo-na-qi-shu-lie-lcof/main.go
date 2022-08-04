package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, (a+b)%(1e9+7)
	}
	return a
}
func main() {
	fmt.Println(fib(5))
}
