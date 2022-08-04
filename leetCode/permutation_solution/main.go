package main

import "fmt"

func solution(nums []int) [][]int {
	list := [][]int{}
	arrange(nums, 0, len(nums)-1, &list)
	return list
}

func arrange(nums []int, start int, end int, list *[][]int) {
	if start == end { // 到最后一个，添加到结果中
		var temp []int
		for _, num := range nums {
			temp = append(temp, num)
		}
		*list = append(*list, temp)
	}
	for i := start; i <= end; i++ {
		nums[i], nums[start] = nums[start], nums[i]
		arrange(nums, start+1, end, list)
		nums[i], nums[start] = nums[start], nums[i] // 还原
	}
}

func main() {
	fmt.Println(solution([]int{1, 2, 3, 4}))
}
