package main

import (
	"fmt"
)

func coinChange2(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i, c := range coins {
		for j := 0; j <= amount; j++ {
			// 当前最佳方案（amount + 1表示无效）
			temp := amount + 1
			// 整数倍：获得一种组合方案
			if j > 0 && j%c == 0 {
				temp = j / c
			}
			// 排除第一种面值：后面才有前驱
			if i > 0 {
				// 当前币取count 个 [0, amount/c]
				for count := 0; count <= amount/c; count++ {
					sum := c * count
					preSum := j - sum
					if preSum < 0 {
						break
					}
					// 再动规位置i - 1时， 能组合出preSum最少数量
					preTemp := dp[preSum]
					if preTemp > 0 {
						temp = min(temp, preTemp+count)
					}
				}
			}
			if temp > amount {
				dp[j] = 0
			} else {
				dp[j] = temp
			}
		}

	}
	result := dp[amount]
	if result == 0 {
		return -1
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	nums := []int{1, 2147483647}
	target := 2
	fmt.Println(coinChange2(nums, target))
}
