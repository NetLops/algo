package main

import "fmt"

func binaryLen(x int) (int, int) {
	ans := 0
	count := 0
	for x|0 != 0 {
		if x&1 == 1 {
			ans++
		}
		x >>= 1
		count++
	}
	return count, ans
}

func main() {
	var n int
	_, _ = fmt.Scanln(&n)
	fmt.Println(binaryLen(n))
}
