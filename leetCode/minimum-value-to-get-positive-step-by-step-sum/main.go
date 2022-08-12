package main

import "math"

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a

}
func minStartValue(nums []int) int {
	min := math.MaxInt
	sum := 1
	for _, num := range nums {
		sum += num
		min = minInt(min, sum)
	}
	if min < 1 {
		return 1 - min
	} else {
		return 1
	}
}
func main() {

}
