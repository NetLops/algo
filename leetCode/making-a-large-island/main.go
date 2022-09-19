package main

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func largestIsland(grid [][]int) int {
	N := 510
	n := len(grid)
	parents := make([]int, N*N)
	sz := make([]int, N*N)

	var find func(x int) int
	find = func(x int) int {
		// 路径压缩
		if parents[x] != x {
			parents[x] = find(parents[x])
		}
		return parents[x]
	}

	var union = func(p, q int) {
		rootP := find(p)
		rootQ := find(q)
		// 相等就返回个寂寞吧
		if rootQ == rootP {
			return
		}
		if sz[rootP] < sz[rootQ] {
			//parents[rootP] = rootQ
			parents[rootP] = rootQ
			sz[rootQ] += sz[rootP]
		} else {
			parents[rootQ] = rootP
			sz[rootP] += sz[rootQ]
		}
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for i := 1; i < n*n; i++ {
		parents[i] = i
		sz[i] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			for _, dir := range dirs {
				x := i + dir[0]
				y := j + dir[1]
				if x < 0 || x >= n || y >= n || y < 0 || grid[x][y] == 0 {
					continue
				}
				union(i*n+j+1, x*n+y+1)
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				ans = max(ans, sz[find(i*n+j+1)])
			} else {
				tot := 1
				set := map[int]struct{}{}
				for _, dir := range dirs {
					x := i + dir[0]
					y := j + dir[1]
					if x < 0 || x >= n || y < 0 || y >= n || grid[x][y] == 0 {
						continue
					}
					root := find(x*n + y + 1)
					if _, ok := set[root]; ok {
						continue
					}
					tot += sz[root]
					set[root] = struct{}{}
				}
				ans = max(ans, tot)
			}
		}
	}
	return ans
}
func main() {

}
