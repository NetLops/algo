package main

import (
	"fmt"
	"sort"
)

func Count(a int) int {
	count := 0
	for a != 0 {
		count++
		a >>= 1
	}
	return count
}

func main() {
	n, k := 0, 0
	fmt.Scan(&n, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)
	for i := 0; i < k; i++ {
		a[i] = Count(a[i])
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += a[i]
	}
	fmt.Println(sum)
}
