package main

import "fmt"

func Quick_sort(nums []int) {
	QuickSort(nums, 0, len(nums)-1)
}

func QuickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	partition := Partition(nums, left, right)
	QuickSort(nums, left, partition-1)
	QuickSort(nums, partition+1, right)
}
func Partition(nums []int, left, right int) int {
	prior := nums[right]
	i, j := left, left
	for ; j < right; j++ {
		if nums[j] < prior {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[j] = nums[j], nums[i]
	return i
}

func main() {
	arr := []int{1, 5, 6, 5, 657, 5675, 674, 5342, 31, 212, 134, 534, 564, 54, 6756, 4231, 1212, 121, 332, 54, 545, 45, 6, 5636, 456, 362, 52, 43, 31, 4, 1, 0}
	Quick_sort(arr)
	fmt.Println(arr)

}
