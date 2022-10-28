package main

import "fmt"

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func main() {
	n := 0
	fmt.Scan(&n)
	a := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		sum += a[i]
	}
	//x:= 0
	for x := -sum; x <= sum; x++ {
		//y := sum - x
	}
}
