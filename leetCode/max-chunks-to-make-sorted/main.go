package main

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func maxChunksToSorted(arr []int) int {
	ans := 0
	n := len(arr)
	// left 左边界 right 右边界, right另外也是移动的指针
	for right, left, min, max := 0, 0, n, -1; right < n; right++ {
		min = Min(min, arr[right])
		max = Max(max, arr[right])
		if left == min && right == max { // left 左 right 不能超过 当前进过的下标最大的
			ans++
			left = right + 1 // 到下一个左边界
			min, max = n, -1 // 重置
		}
	}
	return ans
}
func main() {

}
