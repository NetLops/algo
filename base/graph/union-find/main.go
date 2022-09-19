package main

type UnionFind struct {
	parent, size []int
	count        int // 记录连通分量
}

func (uf *UnionFind) Union(p, q int) {
	rootP := uf.Find(p)
	rootQ := uf.Find(q)
	// 平衡优化
	if uf.size[rootP] < uf.size[rootQ] {
		uf.parent[rootP] = rootQ
		uf.size[rootQ] += uf.size[rootP]
	} else {
		uf.parent[rootQ] = rootP
		uf.size[rootP] += uf.size[rootQ]
	}
	uf.count--
}

// Connected 是否连通
func (uf *UnionFind) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *UnionFind) Find(i int) int {
	// 路径压缩
	if uf.parent[i] != i {
		uf.parent[i] = uf.Find(uf.parent[i])
	}
	// 直接找到祖先节点
	return uf.parent[i]
}

func (uf *UnionFind) Count() int {
	return uf.count
}

func NewUnionFind(n int) UnionFind {
	unionFind := UnionFind{
		parent: make([]int, n),
		size:   make([]int, n),
		count:  n,
	}
	for i := range unionFind.parent {
		unionFind.parent[i] = i
		unionFind.size[i] = 1
	}
	return unionFind
}
