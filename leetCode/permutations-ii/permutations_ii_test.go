package main

import "testing"

func permuteUnique2(nums []int) [][]int {
	length := len(nums)
	flags := make([]bool, length, length)
	for i := 0; i < length; i++ {
		flags[i] = false
	}
	collect := []int{}
	start := 0 // 其实位置
	list := [][]int{}
	dfs2(&list, flags, nums, collect, start, length-1)
	return list
}

func dfs2(list *[][]int, flags []bool, nums []int, collect []int, start int, end int) {
	if start == end+1 {
		temp := make([]int, end+1, end+1)
		copy(temp, collect)
		*list = append(*list, temp)
	}
	for i, num := range nums {
		if flags[i] || (i > 0 && nums[i] == nums[i-1] && !flags[i-1]) {
			continue
		}
		flags[i] = true
		collect = append(collect, num)
		dfs2(list, flags, nums, collect, start+1, end)
		flags[i] = false
		collect = append(collect[:start], collect[start+1:]...)
	}
}

func TestName(t *testing.T) {
	t.Log(permuteUnique2([]int{3, 3, 0, 3}))
	//t.Log(permuteUnique2([]int{1, 1, 1, 3}))
}
