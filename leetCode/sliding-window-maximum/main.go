package main

import "container/heap"

type pair struct {
	index, val int
}

func maxSlidingWindow(nums []int, k int) []int {
	ans := []int{}
	h := &PairHeap{}
	i := 0
	for ; i < k-1; i++ {
		heap.Push(h, pair{index: i, val: nums[i]})
	}
	for ; i < len(nums); i++ {
		heap.Push(h, pair{index: i, val: nums[i]})
		pop := heap.Pop(h).(pair)
		for pop.index <= i-k {
			pop = heap.Pop(h).(pair)
		}
		ans = append(ans, pop.val)
		heap.Push(h, pop)
	}
	return ans
}

type PairHeap []pair

func (m PairHeap) Len() int {
	return len(m)
}

func (m PairHeap) Less(i, j int) bool {
	if m[i].val > m[j].val {
		return true
	}
	return false
}

func (m PairHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *PairHeap) Push(x interface{}) {
	(*m) = append((*m), x.(pair))
}

func (m *PairHeap) Pop() interface{} {
	n := len(*m)
	x := (*m)[n-1]
	*m = (*m)[:n-1]
	return x
}
func main() {
	maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7, 4, 5, 6, 7, 8, 1}, 3)
}
