package main

import (
	"fmt"
	"math"
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
	} else {
		return a
	}
}

// 一：将target数组下标记录如map，
// 二：在arr中寻找值相等，并且按照寻找顺序将下标加入下标数组
// 三：调用最长递增子序列函数s
func minOperations(target []int, arr []int) int {
	n, m := len(target), len(arr)
	dict := map[int]int{}
	// 获取target 的所有下标
	for i := 0; i < n; i++ {
		dict[target[i]] = i
	}
	// 获取 list 相对于 target 的映射关系
	list := []int{}
	// 看看 arr 中是否有存在的
	for i := 0; i < m; i++ {
		x := arr[i]
		if _, ok := dict[x]; ok { // 避免默认 map[key] 为 "零值"
			list = append(list, dict[x])
		}
	}
	length := len(list)
	// f 代表最长上生子序列的长度
	// g 为贪心数组 g[len] =x 代表上升子序列长度为 len 的上升子序列的[最小结尾元素] 为x,   g数组代替线性遍历
	f, g := make([]int, length), make([]int, length+1)
	for i := range g {
		g[i] = math.MaxInt32
	}
	maxAns := 0
	for i := 0; i < length; i++ {
		l, r := 0, length
		for l < r {
			mid := (l + r + 1) >> 1
			if g[mid] < list[i] {
				l = mid
			} else {
				r = mid - 1
			}
		}
		clean := r + 1
		f[i] = clean
		g[clean] = min(g[clean], list[i])
		maxAns = max(maxAns, clean)
	}
	fmt.Println("target: ", target)
	fmt.Println("arr: ", arr)
	fmt.Println("m: ", dict)
	fmt.Println("list: ", list)
	fmt.Println("f: ", f)
	fmt.Println("g: ", g)
	fmt.Println("length: ", length)
	fmt.Println("m: ", m)
	fmt.Println("n: ", n)
	return n - maxAns

}
func main() {
	target := []int{6, 4, 8, 1, 3, 2}
	//target := []int{5, 1, 3}
	arr := []int{4, 7, 6, 2, 3, 8, 6, 1}
	//arr := []int{9, 4, 2, 3, 4}
	fmt.Println(minOperations(target, arr))
}
