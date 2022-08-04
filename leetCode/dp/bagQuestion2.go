package main

import "fmt"

func knapsack(weight []int, n, w int) int {
	state := make([][]bool, n)
	for i := 0; i < n; i++ {
		state[i] = make([]bool, w+1)
	}
	state[0][0] = true // 第一行特殊处理，利用哨兵优化
	state[0][weight[0]] = true
	for i := 1; i < n; i++ { // 动态规划状态转移
		for j := 0; j < w+1; j++ { // 不把第i个物品放入背包
			if state[i-1][j] {
				state[i][j] = state[i-1][j]
			}
		}
		for j := 0; j <= w-weight[i]; j++ { // 把第i个物品放入背包
			if state[i-1][j] {
				state[i][j+weight[i]] = true
			}
		}
	}
	for i := w; i >= 0; i-- { // 输出结果
		if state[n-1][i] == true {
			return i
		}
	}
	return 0
}

func knapsapck2(weight []int, n, w int) int {
	states := make([]bool, w+1)
	states[0] = true
	count := 0
	states[weight[0]] = true
	for i := 0; i < n; i++ {
		//for j := 0; j <= w-weight[i]; j++ {
		//	count = 0
		//	if states[j] {
		//		states[j+weight[i]] = true
		//		count++
		//	}
		//
		//	fmt.Println(count, i, w-weight[i], states)
		//
		//}

		for j := w - weight[i]; j >= 0; j-- {
			count = 0
			if states[j] {
				states[j+weight[i]] = true
				count++
			}

			fmt.Println(count, i, j, w-weight[i], states)
		}
	}
	for i := w; i >= 0; i++ {
		if states[i] {
			return i
		}
	}
	return 0
}

//func main() {
//	fmt.Println(knapsapck2([]int{2, 2, 4, 6, 3}, 5, 9))
//}
