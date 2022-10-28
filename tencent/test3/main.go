package main

import (
	"fmt"
)

func main() {
	var n int
	l := []int{}
	_, _ = fmt.Scan(&n)
	for i := 0; i < n; i++ {
		a := 0
		_, _ = fmt.Scan(&a)
		l = append(l, a)
	}
	ans := make([]int, n)
	index := 0
	for i := 1; i <= n; i++ {
		prev := l[0]
		back := l[len(l)-1]
		if i&1 == 1 {
			if prev > back {
				l = l[1:]
				ans[index] = prev

			} else {
				l = l[:len(l)-1]
				ans[index] = back
			}
		} else {
			if prev < back {
				l = l[1:]
				ans[index] = prev
			} else {
				l = l[:len(l)-1]
				ans[index] = back
			}
		}
		index++

	}
	for i := range ans {
		fmt.Printf("%d ", ans[i])
	}
}
