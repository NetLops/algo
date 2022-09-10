package main

import (
	"container/heap"
	"fmt"
	"github.com/emirpasic/gods/trees/redblacktree"
)

//var servers map[int]*server = map[int]*server{}

//var nowTime, length int

//func busiestServers(k int, arrival []int, load []int) []int {
//	var servers map[int]*server = map[int]*server{}
//	var nowTime, length int
//	ans := []int{}
//	nowTime = 1
//	length = k
//	var getRestSerer func(index int) int
//	getRestSerer = func(index int) int {
//		pos := index % length
//		for i := 0; i < length; i++ {
//			if nowTime >= servers[(pos+i)&(length-1)].next {
//				return (pos + i) & (length - 1)
//			}
//		}
//		return -1
//	}
//	for i := 0; i < k; i++ {
//		servers[i] = &server{
//			now:  i,
//			next: 1,
//		}
//	}
//	for i := 0; i < len(arrival); i++ {
//		nowTime = arrival[i]
//		index := getRestSerer(i)
//		if index != -1 { // -1 就代表没有闲着的服务器
//			servers[index].next = nowTime + load[i]
//			servers[index].handle++
//		}
//	}
//	nowTime = nowTime
//	length = length
//	max := -1
//	for _, s := range servers {
//		if s.handle == max {
//			ans = append(ans, s.now)
//			continue
//		}
//		if s.handle > max {
//			max = s.handle
//			ans = ans[:0]
//			ans = append(ans, s.now)
//		}
//	}
//	return ans
//}
func busiestServers(k int, arrival []int, load []int) []int {
	ans := []int{}
	usedTree := redblacktree.NewWithIntComparator()
	for i := 0; i < k; i++ {
		usedTree.Put(i, nil)
	}
	busyHeap := ListHeap{}
	maxReq := 0
	req := map[int]int{}
	for i, time := range arrival {
		for len(busyHeap) > 0 && busyHeap[0].next <= time {
			usedTree.Put(busyHeap[0].now, nil)
			heap.Pop(&busyHeap)
		}
		if usedTree.Size() == 0 { // 没有空闲的
			continue
		}
		node, _ := usedTree.Ceiling(i % k)
		if node == nil {
			node = usedTree.Left()
		}
		id := node.Key.(int)
		req[id]++
		if req[id] > maxReq {
			maxReq = req[id]
			ans = ans[:0]
			ans = append(ans, id)
		} else if req[id] == maxReq {
			ans = append(ans, id)
		}
		heap.Push(&busyHeap, server{
			now:  id,
			next: time + load[i],
		})
		usedTree.Remove(id)
	}
	return ans
}

type server struct {
	now    int
	next   int
	handle int
}
type ListHeap []server

func (h ListHeap) Len() int           { return len(h) }
func (h ListHeap) Less(i, j int) bool { return h[i].next < h[j].next }
func (h ListHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *ListHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *ListHeap) Push(x interface{}) {
	*h = append(*h, x.(server))
}

//type IntHeap []int

//func (h IntHeap) Len() int           { return len(h) }
//func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
//func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
//func (h *IntHeap) Pop() interface{} {
//	old := *h
//	n := len(old)
//	x := old[n-1]
//	*h = old[0 : n-1]
//	return x
//}
//func (h *IntHeap) Push(x interface{}) {
//	*h = append(*h, x.(int))
//}
func main() {
	//fmt.Println(busiestServers(3, []int{1, 2, 3, 4, 8, 9, 10}, []int{5, 2, 10, 3, 1, 2, 2}))
	fmt.Println(busiestServers(8, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{5, 2, 3, 3, 3, 5, 7, 8, 9, 10}))
	//fmt.Println(busiestServers(3, []int{1}, []int{1}))

}

//7
//[1,3,4,5,6,11,12,13,15,19,20,21,23,25,31,32]
//[9,16,14,1,5,15,6,10,1,1,7,5,11,4,4,6]

//3
//[1,2,3,4,8,9,10]
//[5,2,10,3,1,2,2]
