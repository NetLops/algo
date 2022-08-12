package main

import "fmt"

//// dfs
//func movingCount(m int, n int, k int) int {
//	var dfs func(i, j, si, sj int) int
//	// 初始化 辅助矩阵
//	visited := make([][]bool, m)
//	for i := range visited {
//		visited[i] = make([]bool, n)
//	}
//	dfs = func(i, j, si, sj int) int {
//		if i >= m || j >= n || k < si+sj || visited[i][j] {
//			return 0
//		}
//		visited[i][j] = true
//		ans := 1
//		//  向下走
//		if (i+1)%10 == 0 { // 19 20 = 10,2
//			ans += dfs(i+1, j, si-8, sj)
//		} else {
//			ans += dfs(i+1, j, si+1, sj)
//		}
//		// 下右走
//		if (j+1)%10 == 0 {
//			ans += dfs(i, j+1, si, sj-8)
//		} else {
//			ans += dfs(i, j+1, si, sj+1)
//		}
//		return ans
//	}
//	return dfs(0, 0, 0, 0)
//}
//
//// BFS
//func movingCount(m int, n int, k int) int {
//	// 初始化 visited
//	visited := make([][]bool, m)
//	for i := range visited {
//		visited[i] = make([]bool, n)
//	}
//	ans := 0
//	queue := list.List{}
//	queue.PushBack([]int{0, 0, 0, 0}) // 放入起点
//	for queue.Len() > 0 {
//		element := *queue.Back()
//		queue.Remove(&element)
//		q := element.Value.([]int)
//		i, j, si, sj := q[0], q[1], q[2], q[3]
//		if i >= m || j >= n || visited[i][j] || si+sj > k {
//			continue
//		}
//		visited[i][j] = true
//		ans++
//		if (i+1)%10 == 0 { // 19,20 = 10,2
//			queue.PushBack([]int{i + 1, j, si - 8, sj})
//		} else {
//			queue.PushBack([]int{i + 1, j, si + 1, sj})
//		}
//		if (j+1)%10 == 0 {
//			queue.PushBack([]int{i, j + 1, si, sj - 8})
//		} else {
//			queue.PushBack([]int{i, j + 1, si, sj + 1})
//		}
//	}
//	return ans
//}

//// BFS
//func movingCount(m int, n int, k int) int {
//	// 初始化 visited
//	visited := make([][]bool, m)
//	for i := range visited {
//		visited[i] = make([]bool, n)
//	}
//	ans := 0
//	queue := [][]int{}
//	queue = append(queue, []int{0, 0, 0, 0})
//	for len(queue) > 0 {
//		q := queue[len(queue)-1]
//		queue = queue[:len(queue)-1]
//		i, j, si, sj := q[0], q[1], q[2], q[3]
//		if i >= m || j >= n || visited[i][j] || si+sj > k {
//			continue
//		}
//		visited[i][j] = true
//		ans++
//		if (i+1)%10 == 0 { // 19,20 = 10,2
//			queue = append(queue, []int{i + 1, j, si - 8, sj})
//		} else {
//			queue = append(queue, []int{i + 1, j, si + 1, sj})
//		}
//		if (j+1)%10 == 0 {
//			queue = append(queue, []int{i, j + 1, si, sj - 8})
//		} else {
//			queue = append(queue, []int{i, j + 1, si, sj + 1})
//		}
//	}
//	return ans
//}

// dp
func movingCount(m int, n int, k int) int {
	if k == 0 {
		return 1
	}
	ans := 1

	visited := make([][]int, m)
	for i := range visited {
		visited[i] = make([]int, n)
	}
	visited[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 && j == 0) || get(i)+get(j) > k {
				continue
			}
			// 边界判断
			if i-1 >= 0 { // 从上面来的
				visited[i][j] |= visited[i-1][j]
			}
			if j-1 >= 0 { // 从左侧来
				visited[i][j] |= visited[i][j-1]
			}
			// 以上两个任意一个即可

			ans += visited[i][j]
		}
	}
	//fmt.Println()
	for i := range visited {
		fmt.Println(visited[i])
	}
	return ans
}
func get(x int) int {
	res := 0
	for ; x > 0; x = x / 10 {
		res += x % 10
	}
	return res
}

func main() {
	//fmt.Println(get(224234))
	//fmt.Println(2 | 3)

	movingCount(38, 15, 9)
}
