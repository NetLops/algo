package main

import (
	"fmt"
	"sort"
)

//
//func frequencySort(nums []int) []int {
//	var numSortNumber [101][]int
//	m := map[int]int{}
//
//	for _, num := range nums {
//		m[num]++
//	}
//
//	for k, v := range m {
//		if numSortNumber[v] == nil {
//			numSortNumber[v] = []int{}
//		}
//		numSortNumber[v] = append(numSortNumber[v], k)
//	}
//	ans := make([]int, len(nums))
//
//	t := 0
//	for i := 1; i <= 100; i++ {
//		if numSortNumber[i] != nil {
//			temp := numSortNumber[i]
//			if len(temp) > 1 {
//				sort.Slice(temp, func(i, j int) bool {
//					if temp[i] > temp[j] {
//						return true
//					}
//					return false
//				})
//			}
//			for k := 0; k < len(temp); k++ {
//				for j := 0; j < i; j++ {
//					ans[t] = temp[k]
//					t++
//				}
//			}
//		}
//	}
//	return ans
//}
type pair struct {
	val, count int
}

func frequencySort(nums []int) []int {
	m := map[int]int{}

	for _, num := range nums {
		m[num]++
	}
	pairs := []pair{}
	for k, v := range m {
		pairs = append(pairs, pair{val: k, count: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count < pairs[j].count {
			return true
		} else if pairs[i].count == pairs[j].count {
			if pairs[i].val > pairs[j].val {
				return true
			}
			return false
		}
		return false
	})
	ans := make([]int, len(nums))
	t := 0
	for _, p := range pairs {
		i := p.count
		for i > 0 {
			ans[t] = p.val
			t++
			i--
		}
	}
	return ans
}
func main() {
	a := []int{2, 45, 35, 57, 74, 2, 4454, 5656, 3, 3232, 0}
	sort.Ints(a)
	fmt.Println(a)
}
