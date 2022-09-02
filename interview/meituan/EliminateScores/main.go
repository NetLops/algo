package main

import (
	"fmt"
	"sort"
)

func main() {
	n, x, y := 0, 0, 0

	if scanln, err := fmt.Scanln(&n, &x, &y); err != nil {
		fmt.Println("请输入数字")
	} else {
		if scanln != 3 {
			fmt.Println("请输入n,x,y三个数")
		}
	}
	score := make([]int, n)
	for i := range score {
		if _, err := fmt.Scanf("%d", &score[i]); err != nil {
			//fmt.Println("请正确输入整数", err)
			break
		}
	}
	fmt.Println(getM(n, x, y, score))
}

func getM(n int, x int, y int, score []int) int {
	// 输入的人数太大， 会出现 通过方/淘汰方 大于 y
	if len(score) > 2*y {
		return -1
	}
	// 如果人数太小，会出现 通过方/淘汰房 一方不够x
	if len(score) < 2*x {
		return -1
	}
	// 数组排序
	sort.Ints(score)
	left := x
	right := y + 1
	for i := left; i < right; i++ {
		j := n - i // 晋级人数
		if j >= x && j <= y {
			return score[i-1]
		}
	}
	return -1

}
