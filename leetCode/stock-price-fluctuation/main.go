package main

import (
	"container/heap"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type StockPrice struct {
	m map[int]int
	//current  int
	maxPrice IntHeap
	minPrice IntHeap
	newTime  int
}

func Constructor() StockPrice {
	return StockPrice{
		m:        map[int]int{},
		minPrice: IntHeap{},
		maxPrice: IntHeap{},
	}
}

func (this *StockPrice) Update(timestamp int, price int) {
	//oldPrice := this.m[timestamp]
	this.m[timestamp] = price
	if timestamp > this.newTime {
		this.newTime = timestamp
	}
	heap.Push(&this.maxPrice, pair{-price, timestamp})
	heap.Push(&this.minPrice, pair{price, timestamp})
	//if oldPrice == this.maxPrice {
	//	this.maxPrice = 0
	//}
	//if oldPrice == this.minPrice {
	//	this.minPrice = math.MaxInt32
	//}
	//for _, p := range this.m {
	//	this.maxPrice = max(p, this.maxPrice)
	//	this.minPrice = min(p, this.minPrice)
	//}
	//if oldPrice == this.maxPrice {
	//
	//}
	//if oldPrice == this.minPrice {
	//
	//}
	//
	//this.maxPrice = max(price, this.maxPrice)
	//this.minPrice = min(price, this.minPrice)
}

func (this *StockPrice) Current() int {
	return this.m[this.newTime]
}

func (this *StockPrice) Maximum() int {
	for {
		cur := heap.Pop(&this.maxPrice).(pair)
		if this.m[cur.timeStamp] == -cur.price {
			heap.Push(&this.maxPrice, cur)
			return -cur.price
		}
	}
}

func (this *StockPrice) Minimum() int {
	for {
		cur := heap.Pop(&this.minPrice).(pair)
		if this.m[cur.timeStamp] == cur.price {
			heap.Push(&this.minPrice, cur)
			return cur.price
		}

	}
}

type pair struct {
	price     int
	timeStamp int
}

type IntHeap []pair

func (h *IntHeap) Len() int {
	return len((*h))
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i].price < (*h)[j].price
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x interface{}) {
	(*h) = append((*h), x.(pair))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	*h = (*h)[:len(*h)-1]
	return old[len(old)-1]
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
func main() {

}
