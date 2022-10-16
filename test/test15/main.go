package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	_, _ = fmt.Scanf("%d %d", &n, &k)
	c := make([]int, n)
	i := 0
	for {
		var temp int
		if _, err := fmt.Scanf("%d", &temp); err != nil {
			//fmt.Println(err)
			break
		} else {
			c[i] = temp
			i++
		}
	}

	sort.Slice(c, func(i, j int) bool {
		return c[i] > c[j]
	})
	//start := 0
	//last := 0
	ans := 0
	for i = 0; i < n; i++ {
		if k == 0 {
			ans -= c[i]
		} else {
			// 使用圣光回复
			k--
		}
	}

	fmt.Println(0 - (ans + n - 1))
}

//5 2
//7 3 1 6 8

// 8 7 6 3 1
