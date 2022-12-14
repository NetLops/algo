package head

import "container/heap"

// Item 是优先队列中包含的元素
type Item struct {
	value    string // 元素的值，可以是任意字符串
	priority int    // 元素在队列中的优先级
	// 元素的索引可以用于更新操作，它由heap.Interface 定义的方法维护
	index int // 元素在堆中的索引
}

// 一个实现了 heap.Interface  接口的优先队列，队列中包含任意多个 Item 结构
type PriorityQueue []*Item

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // 为了安全性考虑而做的设置
	*pq = old[0 : n-1]
	return item
}

// 更新函数会修改队列中制定元素的优先级以及值
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// 希望 pop 返回的是最大值而不是最小值
	// 因此此处使用大于号进行对比
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
