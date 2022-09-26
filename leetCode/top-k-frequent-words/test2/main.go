package main

import (
	"container/heap"
	"fmt"
)

func topKFrequent(words []string, k int) []string {
	m := map[string]int{}
	for _, word := range words {
		m[word]++
	}

	pairHeap := PairHeap{}
	for key, value := range m {
		heap.Push(&pairHeap, &Pair{key, value})
	}

	ans := make([]string, 0, len(m))
	n := pairHeap.Len()
	for i := 0; i < n && i < k; i++ {
		pair := heap.Pop(&pairHeap).(*Pair)
		ans = append(ans, pair.word)
	}
	return ans
}

type Pair struct {
	word  string
	count int
}

type PairHeap []*Pair

func (h PairHeap) Len() int {
	return len(h)
}

func (h PairHeap) Less(i, j int) bool {
	if h[i].count > h[j].count || (h[i].count == h[j].count && h[i].word < h[j].word) {
		return true
	}
	return false
}

func (h PairHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PairHeap) Push(x interface{}) {
	(*h) = append((*h), x.(*Pair))
}

func (h *PairHeap) Pop() interface{} {
	n := len((*h))
	x := (*h)[n-1]
	(*h) = (*h)[:n-1]
	return x
}
func main() {
	fmt.Println("i" > "love")
}
