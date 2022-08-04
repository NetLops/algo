package main

import (
	"fmt"
	"math"
)

var (
	minDist = math.MaxInt
)

func minDistBT(i, j, dist int, w [][]int, n int) {
	// 到达了 n-1, n-1 这个位置了，这里看着有点奇怪哈，你自己举个例子看下
	if i == n-1 && j == n-1 {
		if dist < minDist {
			minDist = dist
		}
		return
	}
	if i < n-1 { // 往下走，更新 i=i+1, j=j
		minDistBT(i+1, j, dist+w[i+1][j], w, n)
	}
	if j < n-1 { // 往右走，更新 i=i, j=j+1
		minDistBT(i, j+1, dist+w[i][j+1], w, n)
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minDistDP(matrix [][]int, n int) int {
	state := make([][]int, n)
	for i := 0; i < n; i++ {
		state[i] = make([]int, n)
	}
	sum := 0
	// 初始化行
	for i := 0; i < n; i++ {
		sum += matrix[0][i]
		state[0][i] = sum
	}
	// 初始化列
	sum = 0
	for i := 0; i < n; i++ {
		sum += matrix[i][0]
		state[i][0] = sum
	}
	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			state[i][j] = matrix[i][j] + min(state[i][j-1], state[i-1][j])
		}
	}
	fmt.Println(state)
	return state[n-1][n-1]
}

func minDict2(matrix, mem [][]int, i, j int) int {
	if i == 0 && j == 0 {
		return matrix[0][0]
	}
	if mem[i][j] > 0 {
		return mem[i][j]
	}
	minLeft := math.MaxInt
	if j-1 >= 0 {
		minLeft = minDict2(matrix, mem, i, j-1)
	}
	minUp := math.MaxInt
	if i-1 >= 0 {
		minUp = minDict2(matrix, mem, i-1, j)
	}
	currentMinDist := matrix[i][j] + min(minLeft, minUp)
	mem[i][j] = currentMinDist
	return currentMinDist
}

//func main() {
//	matrix := [][]int{{1, 3, 5, 9}, {2, 1, 3, 4}, {5, 2, 6, 7}, {6, 8, 4, 3}}
//	length := len(matrix)
//	mem := make([][]int, length)
//	for i := 0; i < length; i++ {
//		mem[i] = make([]int, length)
//	}
//
//	minDistBT(0, 0, matrix[0][0], matrix, 4)
//	fmt.Println(minDist)
//
//	fmt.Println(minDistDP(matrix, 4))
//
//	fmt.Println(minDict2(matrix, mem, 3, 3))
//}
