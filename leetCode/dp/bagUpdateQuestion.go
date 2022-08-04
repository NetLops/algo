package main

import (
	"fmt"
	"math"
)

var (
	maxV  = math.MinInt
	items = []int{2, 2, 4, 6, 4}
	value = []int{3, 4, 8, 9, 6}
	n     = 5
	w     = 9
)

func find2(i, cw, cv int) {
	if cw == w || i == n {
		if cv > maxV {
			maxV = cv
		}
		return
	}
	find2(i+1, cw, cv)
	if (cw + items[i]) <= w {
		find2(i+1, cw+items[i], cv+value[i])
	}
}

func knapsack3(weight, value []int, n, w int) int {
	states := make([][]int, n)
	temp := 0
	for i := 0; i < n; i++ {
		states[i] = make([]int, w+1)
	}
	states[0][weight[0]] = value[0]
	for i := 1; i < n; i++ { // 动态规划，状态转移
		for j := 0; j <= w; j++ { // 不选择第 i 个物品
			if states[i-1][j] > 0 {
				states[i][j] = states[i-1][j]
			}
		}
		for j := 0; j <= w-weight[i]; j++ {

			if states[i-1][j] > 0 {
				temp = states[i-1][j] + value[i]
				if temp > states[i-1][j+weight[i]] {

					states[i][j+weight[i]] = temp
				}
			}
		}
	}
	maxValue := -1
	fmt.Println(states)
	for j := w; j >= 0; j-- {
		if states[n-1][j] > maxValue {
			maxValue = states[n-1][j]
		}
	}
	return maxValue
}

func knapsack4(weight, value []int, n, w int) int {
	states := make([]int, w+1)
	temp := 0
	count := 0
	states[weight[0]] = value[0]
	for i := 1; i < n; i++ {
		//for j := 0; j <= w-weight[i]; j++ {
		//	if states[j] > 0 {
		//		temp = states[j] + value[i]
		//		if temp > states[j+weight[i]] {
		//			states[j+weight[i]] = temp
		//		}
		//		count++
		//	}
		//
		//}
		for j := w - weight[i]; j >= 0; j-- {
			if states[j] > 0 {
				temp = states[j] + value[i]
				if temp > states[j+weight[i]] {
					states[j+weight[i]] = temp
				}
				count++
			}
		}
	}
	fmt.Println(states, count)
	max := 0
	for j := w; j >= 0; j-- {
		if states[j] > max {
			max = states[j]
		}
	}
	return max
}

//func main() {
//	find2(0, 0, 0)
//	fmt.Println(maxV)
//
//	fmt.Println(knapsack3(items, value, n, w))
//	fmt.Println(knapsack4(items, value, n, w))
//
//}
