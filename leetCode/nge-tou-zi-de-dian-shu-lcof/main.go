package main

import "fmt"

func dicesProbability(n int) []float64 {
	dp := make([]float64, 6)
	for i := range dp {
		dp[i] = 1 / 6.0
	}
	for i := 2; i <= n; i++ { // 外层循环次数
		temp := make([]float64, 5*i+1)
		for k := 0; k < len(dp); k++ {
			for j := 0; j < 6; j++ {
				temp[k+j] += dp[k] * (1 / 6.0)
			}
		}

		dp = temp
	}
	return dp
}
func main() {
	fmt.Println(dicesProbability(2))
}
