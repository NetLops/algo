package head

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	// 一些元素以及它们的优先级
	items := map[string]int{
		"banana": 3,
		"apple":  2,
		"pear":   4,
	}
	// 创建一个优先队列，并将上述元素放入到队列里面
	// 然后对队列进行初始化以满足优先队列（堆）的不变性
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priotity := range items {
		pq[i] = &Item{
			value:    value,
			priority: priotity,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// 插入一个新元素，然后修改它的优先级
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	//// 以降序形式取出并打印队列中的所有元素
	//for len(pq) > 0 {
	//	item := heap.Pop(&pq).(*Item)
	//	fmt.Printf("%.2d:%s ", item.priority, item.value)
	//
	//}
	fmt.Println()
	pq.update(item, item.value, 5)

	// 以降序形式取出并打印队列中的所有元素
	for len(pq) > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s ", item.priority, item.value)

	}

}
