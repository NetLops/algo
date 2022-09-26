package main

func sumSubarrayMins(arr []int) int {
	n := len(arr)
	leftBound := make([]int, n)
	rightBound := make([]int, n)

	for i := 0; i < n; i++ {
		leftBound[i] = -1
	}
	for i := 0; i < n; i++ {
		rightBound[i] = n
	}

	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && arr[stack[len(stack)-1]] >= arr[i] {
			rightBound[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > arr[i] {
			leftBound[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	ret := 0
	var M int = 1e9 + 7
	for i := 0; i < n; i++ {
		ret += arr[i] * (i - leftBound[i]) % M * (rightBound[i] - i) % M
	}
	return ret % M
}

func main() {
	sumSubarrayMins([]int{11, 81, 94, 43, 43, 8, 8, 43, 3})
}
