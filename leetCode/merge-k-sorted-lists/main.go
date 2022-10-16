package main

import "container/heap"

/**
 * Definition for singly-linked list.

 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	nodeHeap := NodeHeap{}
	for i := 0; i < len(lists); i++ {
		temp := lists[i]
		for temp != nil {
			node := temp
			temp = temp.Next
			node.Next = nil
			heap.Push(&nodeHeap, node)
			// nodeHeap.Push(node)
		}
	}
	if len(nodeHeap) == 0 {
		return nil
	}
	// fmt.Println(nodeHeap)
	ans := heap.Pop(&nodeHeap).(*ListNode)
	temp := ans
	for len(nodeHeap) != 0 {
		temp.Next = heap.Pop(&nodeHeap).(*ListNode)
		temp = temp.Next
	}
	return ans
}

type NodeHeap []*ListNode

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *NodeHeap) Pop() interface{} {
	n := h.Len()
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}
