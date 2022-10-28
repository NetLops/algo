package main

func shortestBridge(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dirs := [][]int{[]int{0, 1}, []int{0, -1}, []int{1, 0}, []int{-1, 0}}
	isValid := func(x, y int) bool {
		if x < 0 || y < 0 || x >= m || y >= n {
			return false
		}
		return true
	}
	v, v1, v2 := make([][]bool, n), make([][]bool, n), make([][]bool, n)
	for i := range v1 {
		v[i] = make([]bool, n)
		v1[i] = make([]bool, n)
		v2[i] = make([]bool, n)
	}
	var dfs func(x, y int, vited [][]bool, q *[][]int)
	dfs = func(x, y int, vited [][]bool, q *[][]int) {
		vited[x][y] = true
		v[x][y] = true
		isBoarder := false
		for _, d := range dirs {
			px, py := x+d[0], y+d[1]
			if isValid(px, py) && !v[px][py] && !vited[px][py] {
				if grid[px][py] == 1 {
					dfs(px, py, vited, q)
				} else {
					if !isBoarder {
						isBoarder = true
						*q = append(*q, []int{x, y})
					}
				}
			}
		}
	}
	q1, q2 := [][]int{}, [][]int{}
	first := true
	for i := range grid {
		for j := range grid[i] {
			if !v[i][j] && grid[i][j] == 1 && first {
				dfs(i, j, v1, &q1)
				first = false
			}
			if !v[i][j] && grid[i][j] == 1 {
				dfs(i, j, v2, &q2)
			}
		}
	}
	c1, c2 := 0, 0
	for len(q1) > 0 && len(q2) > 0 {
		t1, t2 := q1, q2
		q1, q2 = nil, nil
		for i := range t1 {
			x, y := t1[i][0], t1[i][1]
			for _, d := range dirs {
				px, py := x+d[0], y+d[1]
				if isValid(px, py) && grid[px][py] == 0 && !v1[px][py] {
					if v2[px][py] {
						return c1 + c2
					}
					v1[px][py] = true
					q1 = append(q1, []int{px, py})
				}
			}
		}
		c1++
		for i := range t2 {
			x, y := t2[i][0], t2[i][1]
			for _, d := range dirs {
				px, py := x+d[0], y+d[1]
				if isValid(px, py) && grid[px][py] == 0 && !v2[px][py] {
					if v1[px][py] {
						return c1 + c2
					}
					v2[px][py] = true
					q2 = append(q2, []int{px, py})
				}
			}
		}
		c2++
	}
	return -1
}
