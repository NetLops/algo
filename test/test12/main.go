package main

import (
	"fmt"
	"sort"
)

func main() {
	n, m := 0, 0
	fmt.Scanln(&n, &m)

	a := make([]int, n)
	b := make([]int, m)

	for i := 0; i < n; i++ {
		_, err := fmt.Scanf("%d", &a[i])
		if err != nil {
			i = -1
		}
	}
	fmt.Scanln()
	for i := 0; i < m; i++ {
		_, err := fmt.Scanf("%d", &b[i])
		if err != nil {
			i = -1
		}
	}
	sort.Slice(a, func(i, j int) bool {
		return i > j
	})
	//sort.Ints(b)
	sort.Slice(b, func(i, j int) bool {
		return i > j
	})
	//fmt.Println(b)

	ans := 0
	if len(a) < len(b) {
		i := 0
		for ; i < len(a); i++ {
			ans += abs(a[i] - b[i])
		}
		for ; i < len(b); i++ {
			ans += abs(b[i])
		}
	} else {
		i := 0
		for ; i < len(b); i++ {
			ans += abs(a[i] - b[i])
		}
		for ; i < len(a); i++ {
			ans += abs(a[i])
		}
	}
	fmt.Println(ans)
	//fmt.Println(n, m)
	//fmt.Println(a)
	//fmt.Println(b)
}
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

//4 3
//-9821 -2234234 98 1
//7742 888 8989
