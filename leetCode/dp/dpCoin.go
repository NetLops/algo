package main

import (
	"fmt"
	"math"
)

var minCount = math.MaxInt

func coinChange(coins []int, amount int) int {
	memTest(coins, amount, 0)
	if minCount == math.MaxInt {
		return -1
	}
	return minCount
}

func memTest(coins []int, amount, count int) {
	if amount == 0 {
		if count < minCount {
			minCount = count
		}
		return
	}
	for i := 0; i < len(coins); i++ {
		if amount-coins[i] >= 0 {
			memTest(coins, amount-coins[i], count+1)
		} else {

		}
	}
}

func countMoney(moneyItems []int, resultMoney int) int {
	if resultMoney == 0 {
		return resultMoney
	}
	length := len(moneyItems)
	if nil == moneyItems || length < 1 {
		return -1
	}
	if resultMoney < 1 {
		return -1
	}
	// 计算遍历的层数，此按最小金额来支付即为最大层数
	levelNum := resultMoney / moneyItems[0]

	// 付款金额太小了, 喵了个咪
	if levelNum == 0 {
		return -1
	}

	status := make([][]int, levelNum)
	for i := 0; i < levelNum; i++ {
		status[i] = make([]int, resultMoney+1)
	}
	//// 初始化状态数组
	//for i := 0; i < levelNum; i++ {
	//	for j := 0; j < resultMoney; j++ {
	//		status[]
	//	}
	//}
	// 将第一层的数据填充
	for i := 0; i < length; i++ {
		status[0][moneyItems[i]] = moneyItems[i]
	}

	minNum := -1

	// 计算推导状态
	for i := 1; i < levelNum; i++ {
		// 推导当前状态
		for j := 0; j < resultMoney; j++ {
			if status[i-1][j] != 0 {
				// 遍历元素，进行累加
				for k := 0; k < length; k++ {
					if j+moneyItems[k] <= resultMoney {
						status[i][j+moneyItems[k]] = moneyItems[k]
					}
				}
			}
			// 找到最小的张数
			if status[i][resultMoney] > 0 {
				minNum = i + 1
				break
			}
		}
		if minNum > 0 {
			break
		}
	}
	if length == 1 && moneyItems[0] == resultMoney {
		minNum = 1
	}
	return minNum
}

func coinChange2(nums []int, target int) int {
	dp := make([]int, target+1)
	for i := 1; i < target; i++ {
		dp[i] = math.MaxInt32
	}
	for _, num := range nums {
		for i := 0; i < target-num; i++ {
			if dp[i] == math.MaxInt {
				continue
			}
			dp[i+num] = min(dp[i+num], dp[i]+1)
		}
	}
	if dp[target] == math.MaxInt32 {
		return -1
	}
	return dp[target]
}

func main() {
	//coins := []int{1, 3, 5}
	//amount := 9

	//coins := []int{1}
	//amount := 2
	//coins := []int{2}
	//amount := 1
	coins := []int{1, 2147483647}
	amount := 2
	change := coinChange(coins, amount)
	fmt.Println(change)

	//fmt.Println(countMoney(coins, amount))
}
