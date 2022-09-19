package main

type UnionFind struct {
	parents, size []int // size 平衡优化用的
	count         int   // 分量
}

func (uf *UnionFind) Find(i int) int {
	// 路径压缩， 这会让一条路径上的所有节点都指向的根节点
	if uf.parents[i] != i {
		uf.parents[i] = uf.Find(uf.parents[i])

	}
	return uf.parents[i]
}

func (uf *UnionFind) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *UnionFind) Union(p, q int) {
	rootP := uf.Find(p)
	rootQ := uf.Find(q)

	// 平衡优化
	if uf.size[rootP] < uf.size[rootQ] {
		uf.parents[rootP] = rootQ
		uf.size[rootQ] += uf.size[rootP]
	} else {
		uf.parents[rootQ] = rootP
		uf.size[rootP] += uf.size[rootQ]
	}
	uf.count--
}

func (uf *UnionFind) Count() int {
	return uf.count
}

func NewUnionFind(n int) UnionFind {
	unionFind := UnionFind{
		parents: make([]int, n),
		size:    make([]int, n),
		count:   n,
	}
	for i := range unionFind.parents {
		unionFind.parents[i] = i
		unionFind.size[i] = 1 //  默认都为一 指向它自己
	}

	return unionFind
}

func solve(board [][]byte) {
	rows := len(board)
	columns := len(board[0])
	var node func(i, j int) int
	node = func(i, j int) int {
		return i*columns + j
	}
	// 加一个 假数据
	uf := NewUnionFind(rows*columns + 1)
	dummyNode := rows * columns

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if board[i][j] == 'O' {
				// 边界
				if i == 0 || j == 0 || i == rows-1 || j == columns-1 {
					uf.Union(dummyNode, node(i, j))
				} else {
					// 上
					if i > 0 && board[i-1][j] == 'O' {
						uf.Union(node(i, j), node(i-1, j))
					}
					// 下
					if i < rows-1 && board[i+1][j] == 'O' {
						uf.Union(node(i, j), node(i+1, j))
					}
					// 左
					if j > 0 && board[i][j-1] == 'O' {
						uf.Union(node(i, j), node(i, j-1))
					}
					// 右
					if j < columns-1 && board[i][j+1] == 'O' {
						uf.Union(node(i, j), node(i, j+1))
					}
				}
			}

		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if uf.Connected(dummyNode, node(i, j)) {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}

}
func main() {

}
