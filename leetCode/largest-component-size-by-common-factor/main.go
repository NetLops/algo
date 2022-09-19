package main

import "fmt"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func largestComponentSize(nums []int) int {
	const N = 20020
	p := [N]int{}  // UnionFind Parents
	sz := [N]int{} // UnionFind ParentSize
	ans := 1
	var find func(x int) int
	find = func(x int) int {
		// 路径压缩
		if p[x] != x {
			p[x] = find(p[x])
		}
		return p[x]
	}

	union := func(a, b int) {
		rootA := find(a)
		rootB := find(b)
		if rootA == rootB {
			return
		}
		if sz[rootA] < sz[rootB] {
			p[rootA] = rootB
			sz[rootB] += sz[rootA]
			ans = max(ans, sz[rootB])
		} else {
			p[rootB] = rootA
			sz[rootA] += sz[rootB]
			ans = max(ans, sz[rootA])
		}
	}
	n := len(nums)
	m := map[int][]int{}

	add := func(key, val int) {
		if _, ok := m[key]; !ok {
			m[key] = []int{}
		}
		m[key] = append(m[key], val)
	}
	for i := 0; i < n; i++ {
		cur := nums[i]
		for j := 2; j*j <= cur; j++ {
			if cur%j == 0 {
				add(j, i)
			}
			for cur%j == 0 {
				cur /= j
			}
		}
		if cur > 1 { // 此时就代表 cur
			add(cur, i)
		}
	}
	// 初始化UnionFind
	for i := 0; i <= n; i++ {
		p[i] = i
		sz[i] = 1
	}
	for _, valList := range m {
		for i := 1; i < len(valList); i++ {
			union(valList[0], valList[i])
		}
	}
	fmt.Println(m)
	return ans
}
func main() {
	largestComponentSize([]int{2, 4, 8, 16, 3, 5, 15})
}
