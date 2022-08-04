package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	visited := []bool{}
	for i := 0; i < length; i++ {
		visited = append(visited, false)
	}
	list := [][]int{}
	dfs(nums, -1, &list, visited, length)
	return list
}

func dfs(nums []int, index int, list *[][]int, visited []bool, length int) {
	temp := []int{}
	for i := 0; i < length; i++ {
		if visited[i] {
			temp = append(temp, nums[i])
		}
	}
	*list = append(*list, temp)
	for i := index + 1; i < length; i++ {
		if i == 0 || (nums[i] != nums[i-1]) || (i > 0 && visited[i-1] && nums[i] == nums[i-1]) {
			visited[i] = true
			dfs(nums, i, list, visited, length)
			visited[i] = false
		}
	}
}

func main() {
	//fmt.Println(subsetsWithDup([]int{1, 2, 3}))
	fmt.Println(subsetsWithDup([]int{4, 4, 4, 1, 4}))
}
