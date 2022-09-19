package main

import "fmt"

type UnionFind struct {
	parents, size []int
	count         int // 分量
}

func (uf *UnionFind) Find(i int) int {
	// 路径压缩
	if uf.parents[i] != i {
		uf.parents[i] = uf.Find(uf.parents[i])
	}
	return uf.parents[i]
}

func (uf *UnionFind) Union(p, q int) {
	rootP := uf.Find(p)
	rootQ := uf.Find(q)
	if rootP < rootQ {
		uf.parents[rootP] = rootQ
		uf.size[rootQ] += uf.size[rootP]
	} else {
		uf.parents[rootQ] = rootP
		uf.size[rootP] += uf.size[rootQ]
	}
	uf.count--
}

func (uf *UnionFind) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *UnionFind) Count() int {
	return uf.count
}

func NewUnionFind(n int) UnionFind {
	uf := UnionFind{
		parents: make([]int, n),
		size:    make([]int, n),
		count:   n,
	}
	for i := range uf.parents {
		uf.parents[i] = i
		uf.size[i] = 1
	}
	return uf
}

func pacificAtlantic(heights [][]int) [][]int {
	rows := len(heights)
	cols := len(heights[0])
	var node func(i, j int) int
	node = func(i, j int) int {
		return i*cols + j
	}
	uf1 := NewUnionFind(rows*cols + 1)
	uf2 := NewUnionFind(rows*cols + 1)
	dummy := rows * cols

	var dfs func(i, j int, uf *UnionFind)

	dirs := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	dfs = func(i, j int, uf *UnionFind) {
		uf.Union(dummy, node(i, j))
		//res[i][j] = true
		for _, dir := range dirs {
			nx := i + dir[0]
			ny := j + dir[1]
			if nx < 0 || nx >= rows || ny < 0 || ny >= cols {
				continue
			}
			if uf.Connected(dummy, node(nx, ny)) || heights[nx][ny] < heights[i][j] {
				continue
			}
			dfs(nx, ny, uf)
		}
	}

	ans := [][]int{}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == 0 || j == 0 {
				if !uf1.Connected(dummy, node(i, j)) {
					dfs(i, j, &uf1)
				}
			}

			if i == rows-1 || j == cols-1 {
				if !uf2.Connected(dummy, node(i, j)) {
					dfs(i, j, &uf2)
				}
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if uf1.Connected(dummy, node(i, j)) && uf2.Connected(dummy, node(i, j)) {
				ans = append(ans, []int{i, j})
			}
		}
	}
	//fmt.Println(heights)
	return ans
}
func main() {
	//fmt.Println(pacificAtlantic([][]int{
	//	{1, 2, 2, 3, 5},
	//	{3, 2, 3, 4, 4},
	//	{2, 4, 5, 3, 1},
	//	{6, 7, 1, 4, 5},
	//	{5, 1, 1, 2, 4},
	//}))
	fmt.Println(pacificAtlantic([][]int{
		{1, 2, 3},
		{8, 9, 4},
		{7, 6, 5},
	}))
}
