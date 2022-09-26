package main

import "container/heap"

type StockPrice struct {
	MaxTimestamp       int
	TimeMap            map[int]int
	maxPrice, minPrice MinHeap
}

func Constructor() StockPrice {
	return StockPrice{
		TimeMap: map[int]int{},
	}
}

func (this *StockPrice) Update(timestamp int, price int) {
	if this.MaxTimestamp < timestamp {
		this.MaxTimestamp = timestamp
	}
	this.TimeMap[timestamp] = price
	heap.Push(&this.maxPrice, Pair{-price, timestamp})
	heap.Push(&this.minPrice, Pair{price, timestamp})
}

func (this *StockPrice) Current() int {
	return this.TimeMap[this.MaxTimestamp]
}

func (this *StockPrice) Maximum() int {
	for {
		cur := heap.Pop(&this.maxPrice).(Pair)
		v := this.TimeMap[cur.timestamp]
		if v == -cur.Price {
			heap.Push(&this.maxPrice, cur)
			return v
		}
	}
}

func (this *StockPrice) Minimum() int {
	for {
		cur := heap.Pop(&this.minPrice).(Pair)
		v := this.TimeMap[cur.timestamp]
		if v == cur.Price {
			heap.Push(&this.minPrice, cur)
			return v
		}
	}
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */

type Pair struct {
	Price, timestamp int
}

type MinHeap []Pair

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	return m[i].Price < m[j].Price
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(x interface{}) {
	*m = append(*m, x.(Pair))
}

func (m *MinHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}
