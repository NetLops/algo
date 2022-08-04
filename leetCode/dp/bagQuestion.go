package main

import (
	"math"
)

var (
	nowMaxW   = math.MinInt          // 最大重量
	weight    = []int{2, 2, 3, 6, 3} // 物品重量
	count     = 5                    // 物品个数
	maxWeight = 9                    // 背包承受的最大重量
	mem       [5][10]bool
)

func find(i, nowWeight int) {
	if nowWeight == maxWeight || i == count {
		if nowWeight > nowMaxW {
			nowMaxW = nowWeight
		}
		return
	}
	if mem[i][nowWeight] {
		return
	}
	mem[i][nowWeight] = true
	find(i+1, nowWeight)
	if nowWeight+weight[i] <= maxWeight {
		find(i+1, nowWeight+weight[i])
	}
}

//func main() {
//	find(0, 0)
//	fmt.Println(nowMaxW)
//
//}
