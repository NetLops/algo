package main

import (
	"fmt"
	"math"
)

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	state := make([][]int, n)
	length := 0
	state[0] = []int{triangle[0][0]}
	left := 0  // 青龙
	right := 0 // 白虎
	for i := 1; i < n; i++ {
		length = len(triangle[i])
		state[i] = make([]int, length)
		for j := 0; j < length; j++ {
			if j == 0 { // 在第一个位置
				state[i][j] = state[i-1][j] + triangle[i][j]
			} else if j == length-1 { // 在最后一个位置
				state[i][j] = state[i-1][j-1] + triangle[i][j]
			} else { // 卡在中间
				left = state[i-1][j-1]
				right = state[i-1][j]
				if left > right {
					left = right
				}
				state[i][j] = triangle[i][j] + left
			}
		}
	}
	min := math.MaxInt
	for i := 0; i < len(state[n-1]); i++ {
		if min > state[n-1][i] {
			min = state[n-1][i]
		}
	}
	fmt.Println(state)
	return min
}

func minimumTotalw(triangle [][]int) int {
	n := len(triangle)
	stateLength := len(triangle[n-1])
	state := make([]int, stateLength)
	temp := make([]int, stateLength)
	length := 0
	state[0] = triangle[0][0]
	left := 0  // 青龙
	right := 0 // 白虎
	for i := 1; i < n; i++ {
		copy(temp, state)
		length = len(triangle[i])
		for j := 0; j < length; j++ {
			if j == 0 { // 在第一个位置
				state[j] = temp[j] + triangle[i][j]
			} else if j == length-1 { // 在最后一个位置
				state[j] = temp[j-1] + triangle[i][j]
			} else { // 卡在中间
				left = temp[j-1]
				right = temp[j]
				if left > right {
					left = right
				}
				state[j] = triangle[i][j] + left
			}
		}
	}
	min := math.MaxInt
	for i := 0; i < stateLength; i++ {
		if min > state[i] {
			min = state[i]
		}
	}
	return min
}

//func main() {
//	nums := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
//	fmt.Println(minimumTotalw(nums))
//}
