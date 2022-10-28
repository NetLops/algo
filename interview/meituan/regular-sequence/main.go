package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	_, _ = fmt.Scanln(&n)
	seq := make([]int, n)
	for i := range seq {
		_, _ = fmt.Scanf("%d", &seq[i])
	}
	sort.Ints(seq)

	ans := 0
	for i, v := range seq {
		ans += abs(v - i - 1)
	}

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
