package main

//
func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//func numSquares(n int) int {
//	INF := 0x3f3f3f3f
//	cur := 1
//	list := []int{}
//	for i := 1; cur < n; i++ {
//		cur = i * i
//		list = append(list, cur)
//	}
//
//	m := len(list)
//	dp := make([][]int, m+1)
//	for i := range dp {
//		dp[i] = make([]int, n+1)
//	}
//	// 当没有任何数的时候，出了dp[0][0]（花费0个数值凑出0），其他均为无效值
//	for i := 0; i < m+1; i++ {
//		for j := 0; j < n+1; j++ {
//			dp[i][j] = INF
//		}
//	}
//	dp[0][0] = 0
//
//	// 核心处理
//	for i := 1; i < m+1; i++ {
//		x := list[i-1]
//		for j := 0; j < n+1; j++ {
//			// 当不选第i个数的情况
//			dp[i][j] = dp[i-1][j]
//			// 对于选k次第i个数的情况
//			for k := 0; k*x <= j; k++ {
//				// 能够选择k个x的前提剩余的数字j - k * x也能被凑出
//				if dp[i-1][j-k*x] != INF {
//					dp[i][j] = Min(dp[i][j], dp[i-1][j-k*x]+k)
//				}
//			}
//		}
//	}
//	return dp[m][n]
//}
func numSquares(n int) int {
	INF := 0x3f3f3f3f
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = INF
	}
	dp[0] = 0
	for i := 0; i*i <= n; i++ {
		x := i * i
		for j := x; j <= n; j++ {
			dp[j] = Min(dp[j], dp[j-x]+1)
		}
	}
	return dp[n]
}

func main() {
	numSquares(4)
}
