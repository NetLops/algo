package main

import (
	"fmt"
	"testing"
)

func combine(n int, count int) [][]int {
	var list [][]int
	visited := []bool{}
	nums := make([]int, n, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
		visited = append(visited, false)
	}
	dfs2(&list, nums, visited, 0, count, n)
	return list
}

func dfs2(list *[][]int, nums []int, visited []bool, index int, count int, n int) {
	if count == 0 {
		temp := []int{}
		for i := 0; i < n; i++ {
			if visited[i] {
				temp = append(temp, nums[i])
			}
		}
		*list = append(*list, temp)
	} else {
		// 从0开始
		for i := index; i < n; i++ {
			if i == 3 && count == 2 {
				fmt.Println(i)
			}
			visited[i] = true
			dfs2(list, nums, visited, i+1, count-1, n)
			visited[i] = false
		}
	}

}

func TestName(t *testing.T) {
	fmt.Println(combine(4, 2))
}
