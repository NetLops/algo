package main

import (
	"fmt"
	"sort"
)

//
//func find(val int, arr *[]int) bool {
//	for _, i := range *arr {
//		if i == val {
//			return true
//		}
//	}
//	return false
//}
//func arrayRankTransform(arr []int) []int {
//	length := len(arr)
//	ans := make([]int, length)
//	for i := 0; i < length; i++ {
//		ans[i] = 1
//		for j := 0; j < length; j++ {
//			if i != j && arr[i] > arr[j] {
//				ans[i]++
//			}
//		}
//	}
//	// 对排名进行一波过滤
//	//result := []int{}
//	for i := 1; i <= length; i++ {
//		if !find(i, &ans) {
//			for j := 0; j < length; j++ {
//				if ans[j] > i {
//					ans[j]--
//				}
//			}
//		}
//	}
//	return ans
//}
func arrayRankTransform(arr []int) []int {
	length := len(arr)
	ans, clone := make([]int, length), make([]int, length)
	copy(clone, arr)
	sort.Ints(clone)
	m := make(map[int]int, length)
	idx := 0
	for _, val := range clone {
		if m[val] == 0 {
			idx++
			m[val] = idx
		}
	}
	for i, val := range arr {
		ans[i] = m[val]
	}
	return ans
}
func main() {
	arr := []int{37, 12, 28, 9, 100, 56, 80, 5, 12}
	fmt.Println(arrayRankTransform(arr))
}
