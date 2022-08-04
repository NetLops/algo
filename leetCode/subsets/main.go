package main

import (
	"fmt"
	"sort"
)

func subsets(nums []int) [][]int {
	sort.Ints(nums)
	var list [][]int
	length := len(nums)
	visited := make([]bool, length, length)
	dfs(&list, nums, 0, length, visited)
	return list
}

func dfs(list *[][]int, nums []int, index int, length int, visited []bool) {
	temp := []int{}
	for i := 0; i < length; i++ {
		if visited[i] {
			temp = append(temp, nums[i])
		}
	}
	*list = append((*list), temp)

	for i := index; i < length; i++ {
		visited[i] = true
		dfs(list, nums, i+1, length, visited)
		visited[i] = false
	}
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	//fmt.Println(subsets([]int{0}))
}
