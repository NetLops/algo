package main

import (
	"fmt"
	"strconv"
)

//func printNumbers(n int) []int {
//
//	if n <= 0 {
//		return []int{}
//	}
//	max := 0
//	for j := 1; j <= n; j++ {
//		max += 10*max + 9
//	}
//	ans := make([]int, max)
//	for i := 0; i < max; i++ {
//		ans[i] = i + 1
//	}
//	return ans
//}da-yin-cong-1dao-zui-da-de-nwei-shu-lcof

func printNumbers(n int) []int {
	ans := []int{}

	num := make([]byte, n) // 存每个位的数
	start := n - 1         // 默认个位数

	var dfs func(x int)
	nineCount := 0

	dfs = func(x int) { // 位数
		if x == n {
			s := num[start:]
			// 为0就不要了
			singleAns := string(s)
			if singleAns != "0" {
				// 转一手
				addAns, _ := strconv.Atoi(singleAns)
				ans = append(ans, addAns)
			}
			if n-start == nineCount { // 此时有9+了,该进位
				start-- // 位数上升
			}
		} else {
			// x = 0
			// num[x, x', x'', x''']
			// 0 0 0 0 ~ 0 9 9 9
			// 1 0 0 0 ~ 1 9 9 9
			// 2 0 0 0 ~ 2 9 9 9
			// 3 0 0 0 ~ 3 9 9 9
			// ...
			// 9 0 0 0 ~ 9 9 9 9
			for i := 0; i <= 9; i++ {
				if i == 9 {
					nineCount++
				}
				num[x] = byte(i + '0')
				// 内层套娃
				dfs(x + 1)
			}
			nineCount--
		}
	}

	dfs(0)

	return ans
}

func main() {
	fmt.Println(printNumbers(3))
}
