package main

import (
	"fmt"
	"sort"
)

func main() {
	n, m := 0, 0

	fmt.Scanln(&n, &m)
	repairs := make([]int, n)
	patchs := make([]int, m)

	for i := 0; i < n; i++ {
		_, err := fmt.Scanf("%d", &repairs[i])

		if err != nil {
			//fmt.Println("1", err)
			i = -1
		}
	}
	//fmt.Scanln()
	for i := 0; i < m; i++ {
		_, err := fmt.Scanf("%d", &patchs[i])

		if err != nil {
			//fmt.Println("2", err, i)
			i = -1
		}
	}

	sort.Ints(patchs)

	ans := 0
	for i := 0; i < n; i++ {
		temp := getPatch(repairs[i], patchs)
		if temp == -1 {
			fmt.Println(-1)
			return
		} else {
			ans += temp
		}
	}
	fmt.Println(ans)

}

func getPatch(a int, patch []int) int {
	for i := 0; i < len(patch); i++ {
		if patch[i] >= a {
			return patch[i]
		}
	}
	return -1
}
