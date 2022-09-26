package main

import "container/heap"

func eatenApples(apples []int, days []int) int {
	n := len(apples)
	pairHeap := PairHeap{}

	ans := 0
	i := 0
	for i < n || len(pairHeap) > 0 {
		if i < n && days[i] != 0 && apples[i] != 0 {
			heap.Push(&pairHeap, Pair{expire: i + days[i] - 1, count: apples[i]})
		}
		// 吃一手苹果
		for len(pairHeap) > 0 {
			nowEatPair := heap.Pop(&pairHeap).(Pair)
			if nowEatPair.expire < i {
				continue
			}
			// fmt.Println(nowEatPair, pairHeap)
			nowEatPair.count = nowEatPair.count - 1
			ans++ // 吃苹果++
			// 或者明天过期
			if nowEatPair.count > 0 && nowEatPair.expire > i {
				heap.Push(&pairHeap, nowEatPair)
			}
			// fmt.Println(nowEatPair, pairHeap)
			break
		}
		i++
	}
	return ans
}

type Pair struct {
	count  int
	expire int
}

type PairHeap []Pair

func (h PairHeap) Len() int {
	return len(h)
}

// 每次差不多要过期的苹果
func (h PairHeap) Less(i, j int) bool {
	if h[i].expire < h[j].expire {
		return true
	}
	return false
}

func (h PairHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PairHeap) Push(x interface{}) {
	(*h) = append((*h), x.(Pair))
}

func (h *PairHeap) Pop() interface{} {
	n := len((*h))
	x := (*h)[n-1]
	(*h) = (*h)[:n-1]
	return x
}

func main() {

}
