package main

import "container/heap"

/**
 * Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// x 为层数，y为纵向位置
func verticalTraversal(root *TreeNode) [][]int {
	var dfs func(root *TreeNode, x, y int)
	pairHeap := PairHeap{}
	dfs = func(root *TreeNode, x, y int) {
		if root == nil {
			return
		}
		heap.Push(&pairHeap, Pair{x: x, y: y, val: root.Val})
		dfs(root.Left, x+1, y-1)
		dfs(root.Right, x+1, y+1)
	}

	dfs(root, 0, 0)
	ans := [][]int{}
	y := 0
	temp := []int{}
	for len(pairHeap) != 0 {
		pair := heap.Pop(&pairHeap).(Pair)
		if pair.y != y {
			if len(temp) > 0 {
				ans = append(ans, temp)
			}
			temp = []int{}
			y = pair.y
			temp = append(temp, pair.val)
		} else {
			temp = append(temp, pair.val)
		}
	}
	if len(temp) > 0 {
		ans = append(ans, temp)
	}
	return ans
}

type Pair struct {
	x, y int
	val  int
}

type PairHeap []Pair

func (h PairHeap) Len() int {
	return len(h)
}

func (h PairHeap) Less(i, j int) bool {
	if h[i].y < h[j].y || (h[i].y == h[j].y && h[i].x < h[j].x) || (h[i].y == h[j].y && h[i].x == h[j].x && h[i].val < h[j].val) {
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
	*h = (*h)[:n-1]
	return x
}
