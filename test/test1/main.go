package main

import (
	"fmt"
)

func getShuiXianHua(m, n int) (ans []int) {
	if m > n {
		return ans
	}
	for i := m; i <= n; i++ {
		three := i / 100
		two := (i / 10) % 10
		one := i % 10
		if three*three*three+two*two*two+one*one*one == i {
			ans = append(ans, i)
		}
	}
	return
}
func main() {
	m, n := 0, 0
	_, _ = fmt.Scanln(&m, &n)
	ans := getShuiXianHua(m, n)
	if len(ans) == 0 {
		fmt.Println("no")
	}
	for i := range ans {
		fmt.Print(ans[i], " ")
	}
}
