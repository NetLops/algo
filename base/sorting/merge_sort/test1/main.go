package main

import (
	"fmt"
)

func Merge_sort(arr []int) {
	if arr == nil || len(arr) == 0 {
		return
	}
	Merge_sort_c(arr, 0, len(arr)-1)
}

func Merge_sort_c(arr []int, left int, right int) {
	if left >= right { // 终止条件 找到头了
		return
	}
	mid := left + (right-left)>>1
	Merge_sort_c(arr, left, mid)
	Merge_sort_c(arr, mid+1, right)
	Merge(arr, left, mid, right)
}

func Merge(arr []int, left int, mid int, right int) {
	i, j := left, mid+1

	temp := make([]int, 0, right-left+1)
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp = append(temp, arr[i])
			i++
		} else {
			temp = append(temp, arr[j])
			j++
		}
	}

	// 合并剩余的排序数组
	for start := i; start <= mid; start++ {
		temp = append(temp, arr[start])
	}
	for start := j; start <= right; start++ {
		temp = append(temp, arr[start])
	}

	for start := 0; start <= right-left; start++ {
		arr[start+left] = temp[start]
	}
}
func main() {
	arr := []int{1, 5, 6, 5, 657, 5675, 674, 5342, 31, 212, 134, 534, 564, 54, 6756, 4231, 1212, 121, 332, 54, 545, 45, 6, 5636, 456, 362, 52, 43, 31, 4, 1, 0}
	Merge_sort(arr)
	fmt.Println(arr)

}
