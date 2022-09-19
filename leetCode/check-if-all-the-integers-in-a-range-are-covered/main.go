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
func isCovered(ranges [][]int, left int, right int) bool {

	m := make(map[int]bool, 50)
	for _, ints := range ranges {
		for i := ints[0]; i <= ints[1]; i++ {
			m[i] = true
		}
	}

	for i := left; i <= right; i++ {
		if !m[i] {
			return false
		}
	}

	return true

}
func main() {

}
