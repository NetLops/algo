package main

import (
	"fmt"
	"sort"
	"testing"
)

func subsetsWithDup2(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	visited := make([]bool, length, length)
	list := [][]int{}
	dfs2(&list, nums, visited, 0, length)
	return list
}

func dfs2(list *[][]int, nums []int, visited []bool, index int, length int) {
	temp := []int{}
	for i := 0; i < length; i++ {
		if visited[i] {
			temp = append(temp, nums[i])
		}
	}
	*list = append(*list, temp)
	for i := index; i < length; i++ {
		if i == 0 || (nums[i] != nums[i-1]) || (i > 0 && visited[i-1] && nums[i] == nums[i-1]) {
			visited[i] = true
			dfs2(list, nums, visited, i+1, length)
			visited[i] = false
		}
	}
}

func TestName(t *testing.T) {
	fmt.Println(subsetsWithDup2([]int{1, 2, 3}))
}
