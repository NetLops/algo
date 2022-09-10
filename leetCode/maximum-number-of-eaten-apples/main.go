package main

import (
	"container/heap"
	"fmt"
)

// 核心思想 每次吃 离保质期最近的 （近视快要过期的 那种爷爷奶奶级别思想）
func eatenApples(apples []int, days []int) int {
	ans := 0
	hp := heapPair{}
	// 存活的天数
	i := 0
	for i < len(apples) || hp.Len() != 0 {
		// 判断下有没有烂掉的苹果
		for hp.Len() != 0 && hp[0].date <= i { // 就是烂掉的苹果
			heap.Pop(&hp)
		}
		// 新增的小苹果，比如是有苹果 和时间才行
		if i < len(apples) && apples[i] > 0 && days[i] > 0 {
			heap.Push(&hp, pair{
				date:  i + days[i],
				count: apples[i],
			})
		}
		// 开始吃苹果了
		if hp.Len() != 0 {
			//temp := heap.Pop(hp).(pair)
			//temp.count--
			//if temp.count != 0 {
			//	heap.Push(hp, temp)
			//}
			if hp[0].count > 1 {
				hp[0].count--
			} else {
				heap.Pop(&hp)
			}
			ans++
		}

		i++
	}
	return ans
}

type pair struct {
	date  int
	count int
}

type heapPair []pair

func (h heapPair) Len() int {
	return len(h)
}

func (h heapPair) Less(i, j int) bool {
	if h[i].date < h[j].date {
		return true
	}
	return false
}

func (h heapPair) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *heapPair) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *heapPair) Pop() interface{} {
	length := h.Len()
	if length <= 0 {
		return nil
	}
	x := (*h)[length-1]
	(*h) = (*h)[:length-1]
	return x
}

func main() {
	h := heapPair{}
	p1 := pair{
		date:  33,
		count: 1,
	}
	p2 := pair{
		date:  44,
		count: 5,
	}
	p3 := pair{
		date:  66,
		count: 5,
	}

	p4 := pair{
		date:  6,
		count: 5,
	}

	heap.Push(&h, p2)
	heap.Push(&h, p1)

	heap.Push(&h, p3)
	heap.Push(&h, p4)
	fmt.Println(heap.Pop(&h))
}
