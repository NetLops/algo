package main

import (
	"fmt"
	"testing"
)

func permute(nums []int) [][]int {
	var list [][]int
	start := 0
	end := len(nums) - 1
	allArrange(&list, &nums, start, end)
	return list
}

func allArrange(list *[][]int, nums *[]int, start int, end int) {
	if start == end {
		//temp := make([]int, end+1, end+1)
		//for i := 0; i <= end; i++ {
		//	temp[i] = (*nums)[i]
		//}
		temp := []int{}
		for _, num := range *nums {
			temp = append(temp, num)
		}
		*list = append(*list, temp)
	}
	for i := start; i <= end; i++ {
		(*nums)[i], (*nums)[start] = (*nums)[start], (*nums)[i]
		allArrange(list, nums, start+1, end)
		(*nums)[i], (*nums)[start] = (*nums)[start], (*nums)[i]

	}
}

func TestTest1(t *testing.T) {
	fmt.Println(permute([]int{1, 2, 3, 4}))
}
