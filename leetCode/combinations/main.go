package main

import "fmt"

func combind(n int, k int) [][]int {
	var list [][]int
	nums := make([]int, n, n)
	visited := []bool{}
	for i := 0; i < n; i++ {
		nums[i] = i + 1
		visited = append(visited, false)
	}
	dfs(nums, -1, k, &list, visited, n)
	return list
}

func dfs(nums []int, index int, count int, list *[][]int, visited []bool, n int) {
	if count == 0 { // k个元素已经满了
		temp := []int{}
		for i := 0; i < n; i++ {
			if visited[i] {
				temp = append(temp, i+1)
			}
			//temp = append(temp, i+1)
		}
		*list = append(*list, temp)
	} else {
		for i := index + 1; i < n; i++ {
			fmt.Println(count)
			visited[i] = true
			dfs(nums, i, count-1, list, visited, n)
			fmt.Println(count)
			visited[i] = false
		}
	}
}

func main() {
	fmt.Println(combind(4, 2))
}
