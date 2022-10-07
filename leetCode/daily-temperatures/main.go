package main

import (
	"container/heap"
	"fmt"
)

func dailyTemperatures(temperatures []int) []int {
	m := map[[2]int]int{}

	ans := make([]int, len(temperatures))
	intHeap := IntHeap{}
	for i, num := range temperatures {
		for len(intHeap) > 0 && intHeap[0][1] < num {
			x := heap.Pop(&intHeap).([2]int)
			m[x]++
			ans[x[0]] = i - x[0]
		}

		heap.Push(&intHeap, [2]int{i, num})
	}

	return ans
}

type IntHeap [][2]int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	(*h) = append((*h), x.([2]int))
}

func (h *IntHeap) Pop() interface{} {
	n := len((*h))
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

//func dailyTemperatures(temperatures []int) []int {
//	n := len(temperatures)
//	q := []int{}
//	ans := make([]int, n)
//	for i := range temperatures {
//		for len(q) != 0 && temperatures[q[len(q)-1]] < temperatures[i] {
//			idx := q[len(q)-1]
//			q = q[:len(q)-1]
//			ans[idx] = i - idx
//		}
//		q = append(q, i)
//	}
//	return ans
//}

func main() {
	intHeap := IntHeap{}

	s := []int{4, 546, 456, 45, 6345, 3, 452, 34, 2, 341, 234, 1213, 2343425, 1}
	for i := range s {
		heap.Push(&intHeap, [2]int{i, s[i]})
	}

	s2 := []int{-23, -34}
	for i := range s2 {
		heap.Push(&intHeap, [2]int{i, s2[i]})
		fmt.Println(intHeap[0])
		fmt.Println(heap.Pop(&intHeap).([2]int))
	}
	fmt.Println(intHeap)
}
