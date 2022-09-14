package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	n := len(quality)
	ds := make([][]float64, n)
	for i := range ds {
		ds[i] = make([]float64, 2)
	}
	for i := 0; i < n; i++ {
		ds[i][0] = float64(wage[i]) / float64(quality[i])
		ds[i][1] = float64(i)
	}
	sort.Slice(ds, func(i, j int) bool {
		if ds[i][0] < ds[j][0] {
			return true
		}
		return false
	})

	fmt.Println(ds)
	maxpq := IntHeap{}
	ans := 1e18 * 1.0
	total := 0
	for i := 0; i < n; i++ {
		cur := quality[int(ds[i][1])]
		total += cur
		heap.Push(&maxpq, cur)
		if maxpq.Len() > k {
			total -= (heap.Pop(&maxpq)).(int)
		}
		if maxpq.Len() == k {
			ans = min(ans, float64(total)*ds[i][0])
		}
	}
	return ans
}

func min(a, b float64) float64 {
	if a > b {
		return b
	}
	return a
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

// 大顶堆嘛
func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = (*h)[:len(old)-1]
	return x
}
func main() {
	mincostToHireWorkers([]int{10, 20, 5}, []int{70, 50, 30}, 2)
}
