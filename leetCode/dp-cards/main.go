package main

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func win1(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	before := f1(arr, 0, len(arr)-1)
	after := g1(arr, 0, len(arr)-1)
	return max(before, after)
}

/**
 * 以先手状态获得的最大分数
 */

func f1(arr []int, L, R int) int {
	if L == R {
		return arr[L]
	}
	// 还剩余不止一张牌
	p1 := arr[L] + g1(arr, L+1, R)
	p2 := arr[R] + g1(arr, L, R+1)
	return max(p1, p2)
}

func g1(arr []int, L, R int) int {
	if L == R {
		return 0
	}
	// 对手拿走L位置的数
	p1 := f1(arr, L+1, R)
	p2 := f1(arr, L, R+1)
	return min(p1, p2)
}
func main() {

}
