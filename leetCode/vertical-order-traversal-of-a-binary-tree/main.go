package main

import (
	"math"
	"sort"
)

/**
 * Definition for a binary tree node.

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func verticalTraversal(root *TreeNode) [][]int {
	m := map[[2]int][]int{} // map[(y,x)]:[]int{root.val}

	list := [][2]int{}
	min_bround, max_bround := math.MaxInt, math.MinInt
	var dfs func(root *TreeNode, x, y int)
	// 前序遍历
	dfs = func(root *TreeNode, x, y int) {
		if root == nil {
			return
		}
		min_bround = min(min_bround, x)
		max_bround = max(max_bround, x)

		if m[[2]int{y, x}] == nil {
			m[[2]int{y, x}] = []int{}
			list = append(list, [2]int{y, x})
		}
		m[[2]int{y, x}] = append(m[[2]int{y, x}], root.Val)
		dfs(root.Left, x-1, y+1)
		dfs(root.Right, x+1, y+1)
	}
	dfs(root, 0, 0)
	//for i := range list {
	//	sort.Ints(ans[list[i]])
	//	ret = append(ret, ans[list[i]])
	//}
	// 让y轴排个序
	sort.Slice(list, func(i, j int) bool {
		if list[i][0] < list[j][0] {
			return true
		}
		return false
	})
	// 同 y 轴 同x 轴排个序	// 比如 (2,0):[6,5]， 同 y 轴 x 轴 应该从小到大排序
	for i := range list {
		if len(m[list[i]]) > 1 {
			sort.Ints(m[list[i]])
		}
	}
	ret := make([][]int, max_bround-min_bround+1) // 取得 x 轴上区间范围
	for i := range list {                         // 一不小心 根据y轴顺序遍历map了
		k := list[i]
		v := m[k]
		index := k[1] - min_bround // 让 index 成为 将x转化为  0,1,2,3,4,5... 这种
		//if len(ret) != index+1 {
		//	ret = append(ret, []int{})
		//}

		ret[index] = append(ret[index], v...)
	}
	return ret
}
func main() {

}
