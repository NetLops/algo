package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func nextGreaterElement(n int) int {
	source := n
	result := -1
	itoa := strconv.Itoa(n)
	length := len(itoa)
	visited := make([]bool, length, length)
	nums := make([]int, length, length)
	//list := [][]int{}
	isFlag := false
	collect := []int{}
	for i := range nums {
		nums[i] = int(itoa[i]) - 48
		visited[i] = false
	}
	sort.Ints(nums)
	return dfs(nums, collect, visited, 0, length, source, &isFlag, &result)
	//fmt.Println(list)
	//return 0
}

func dfs(nums []int, collect []int, visited []bool, index int, length int, source int, isFlag *bool, result *int) int {
	if *isFlag {
		return *result
	}
	if index == length {
		fmt.Println(collect)
		now := 0
		for i := range collect {
			now += collect[i] * int(math.Pow10(length-i-1))
		}
		if now > math.MaxInt32 {
			*result = -1
			*isFlag = true
			//fmt.Println(now, 1<<31-)
			return *result
		}
		if now > source {

			fmt.Println(now)
			*result = now
			*isFlag = true
		}

	} else {
		for i := 0; i < length; i++ {
			if visited[i] || (i > 0 && nums[i] == nums[i-1] && !visited[i-1]) || *isFlag {
				continue
			}
			visited[i] = true
			collect = append(collect, nums[i])
			dfs(nums, collect, visited, index+1, length, source, isFlag, result)
			visited[i] = false
			collect = append(collect[:index], collect[index+1:]...)
		}
	}
	return *result
}

func main() {
	nextGreaterElement(2147483648)
}
