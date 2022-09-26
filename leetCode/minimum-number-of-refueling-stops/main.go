package main

import "container/heap"

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	stationHeap := StationHeap{}
	ans := 0
	cur, idx, fuel := 0, 0, startFuel
	for cur < target {
		cur += fuel
		// 将当前油 能跑的最远距离所遇到 的加油站 全部找到，然后通过优先队列找到加油最多
		for idx < len(stations) && stations[idx][0] <= cur {
			heap.Push(&stationHeap, stations[idx][1]) // 将最多油存进去
			idx++
		}
		if cur < target {
			if len(stationHeap) == 0 { // 没加油站 G了害
				return -1
			}
			ans++
			fuel = heap.Pop(&stationHeap).(int) // 找到最多的油
		}
	}
	return ans
}

type station struct {
	mile, lol int
}

type StationHeap []int

func (h StationHeap) Len() int {
	return len(h)
}

func (h StationHeap) Less(i, j int) bool {
	if h[i] > h[j] {
		return true
	}
	return false
}

func (h StationHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *StationHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *StationHeap) Pop() interface{} {
	n := len((*h))
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func main() {

}
