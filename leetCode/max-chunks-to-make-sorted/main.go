package main

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func maxChunksToSorted(arr []int) int {
	ans, max := 0, 0
	for i := 0; i < len(arr); i++ {
		max = Max(max, arr[i])
		if arr[i] >= max {
			ans++
		}
	}
	return ans
}
func main() {

}
