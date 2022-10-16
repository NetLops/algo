package main

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func maxChunksToSorted(arr []int) int {
	ans := 0
	maxInt := 0
	for i, num := range arr {
		maxInt = max(maxInt, num)
		if num >= maxInt && num <= i {
			ans++
		}
	}
	return ans
}
