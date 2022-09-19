package main

import (
	"fmt"
)

func sumSubarrayMins(arr []int) int {
	fmt.Println(arr)
	n := len(arr)
	leftBound := make([]int, n)
	for i := range leftBound {
		leftBound[i] = -1
	}
	//leftBound[0] = -1
	rightBound := make([]int, n+1)
	for i := range rightBound {
		rightBound[i] = n
	}
	stack := []int{} // 单调栈
	// 范围的 比如 34456788645442  拿它的范围是   344[5678864544]2 右边的倒数第三个4 可以等于 左边的不行
	for i := 0; i < n; i++ {
		for len(stack) != 0 && arr[stack[len(stack)-1]] > arr[i] {
			//rightBound[i] = stack[len(stack)-1]
			rightBound[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = []int{} // 清空一下 有残留

	for i := n - 1; i >= 0; i-- {
		for len(stack) != 0 && arr[stack[len(stack)-1]] >= arr[i] {

			leftBound[stack[len(stack)-1]] = i
			//leftBound[i] = stack[len(stack)-1]

			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	ret := 0
	var M int
	fmt.Println(leftBound)
	fmt.Println(rightBound)
	M = 1e9 + 7
	for i := 0; i < n; i++ {
		ret += arr[i] * ((i - leftBound[i]) % M) * ((rightBound[i] - i) % M)
	}

	return ret % M
}
func main() {
	fmt.Println(sumSubarrayMins([]int{3, 1, 2, 4}))
}

//[1,2,3,4]
