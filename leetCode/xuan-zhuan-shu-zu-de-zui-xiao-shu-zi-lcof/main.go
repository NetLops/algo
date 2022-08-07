package main

func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	for left <= right {
		middle := left + (right-left)>>1
		if numbers[middle] > numbers[right] {
			left = middle + 1
		} else if numbers[middle] < numbers[right] {
			right = middle
		} else {
			right -= 1
		}
	}
	return numbers[left]
}
func main() {

}
