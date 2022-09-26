package main

import (
	"container/heap"
	"fmt"
)

func busiestServers(k int, arrival []int, load []int) []int {
	used, available, cnts := &ListHeap{}, &IntHeap{}, make([]int, k)
	for i := 0; i < k; i++ {
		heap.Push(available, i)
	}
	for i := 0; i < len(arrival); i++ {
		arr, duration := arrival[i], load[i]
		for used.Len() > 0 && (*used)[0][0] <= arr {
			cur := heap.Pop(used).([]int)
			heap.Push(available, i+((cur[1]-i)%k+k)%k)
		}
		if available.Len() > 0 {
			idx := heap.Pop(available).(int) % k
			cnts[idx]++
			heap.Push(used, []int{arr + duration, idx})
		}
	}
	ans, m := []int{}, 0
	for i := 0; i < k; i++ {
		if cnts[i] > m {
			m = cnts[i]
			ans = []int{i}
		} else if cnts[i] == m {
			ans = append(ans, i)
		}
	}
	return ans
}

type ListHeap [][]int

func (h ListHeap) Len() int           { return len(h) }
func (h ListHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h ListHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *ListHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *ListHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func main() {
	fmt.Println(busiestServers(8, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{5, 2, 3, 3, 3, 5, 7, 8, 9, 10}))

	const (
		NAME1 = 1 << iota
		NAME2
		NAME3
		NAME4
		NAME5
	)
	fmt.Println(NAME1, NAME2, NAME3, NAME4, NAME5)

}
