package main

import (
	"fmt"
	"sort"
)

//
//func Max(a, b int) int {
//	if a < b {
//		return b
//	}
//	return a
//}
//
//type Stack []int
//
//func (s Stack) Top() int {
//	if len(s) > 0 {
//		return s[len(s)-1]
//	}
//	return -1
//}
//
//func (s *Stack) Push(val int) {
//	*s = append(*s, val)
//}
//
//func (s *Stack) Pop() (val int) {
//	if len(*s) > 0 {
//		val = (*s)[len(*s)-1]
//		*s = (*s)[:len(*s)-1]
//		return val
//	}
//	return -1
//}
//
//func (s *Stack) Len() int {
//	return len(*s)
//}
//
//func maxChunksToSorted(arr []int) int {
//	stack := Stack{}
//	for _, num := range arr {
//		if stack.Len() != 0 && num < stack.Top() {
//			head := stack.Pop()
//			for stack.Len() != 0 && num < stack.Top() {
//				stack.Pop()
//			}
//			stack.Push(head)
//		} else {
//			stack.Push(num)
//		}
//	}
//	return stack.Len()
//}
func maxChunksToSorted(arr []int) int {
	clone := make([]int, len(arr))
	copy(clone, arr)
	sort.Ints(clone)
	m := map[int]int{}
	n, ans := len(arr), 0
	for i := 0; i < n; i++ {
		a, b := arr[i], clone[i]
		m[a]++
		if m[a] == 0 {
			delete(m, a)
		}
		m[b]--
		if m[b] == 0 {
			delete(m, b)
		}
		if len(m) == 0 {
			ans++
		}
		fmt.Println(m)
	}
	return ans
}

// 解法2
func main() {
	a := []int{1, 1, 2, 1, 1, 3, 4, 5, 1, 6}
	_ = []int{1, 1, 1, 1, 1, 2, 3, 4, 5, 6}
	fmt.Println(maxChunksToSorted(a))
}
