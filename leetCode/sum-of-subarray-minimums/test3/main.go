package main

func sumSubarrayMins(arr []int) int {
	const MAX int = 1e+9 + 7
	n := len(arr)
	lefts := make([]int, n)
	rights := make([]int, n)

	for i := 0; i < n; i++ {
		lefts[i] = -1
	}

	for i := 0; i < n; i++ {
		rights[i] = n
	}

	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && arr[stack[len(stack)-1]] >= arr[i] {
			rights[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = []int{}

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > arr[i] {
			lefts[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans += arr[i] * (i - lefts[i]) % MAX * (rights[i] - i) % MAX
	}
	return ans % MAX
}

//func sumSubarrayMins(arr []int) int {
//	ret := 0
//	const M int = 1e9 + 7
//	n := len(arr)
//	lefts := make([]int, n)
//	rights := make([]int, n)
//
//	for i := 0; i < n; i++ {
//		lefts[i] = -1
//	}
//	for i := 0; i < n; i++ {
//		rights[i] = n
//	}
//
//	stack := []int{}
//	for i := 0; i < n; i++ {
//		for len(stack) > 0 && arr[stack[len(stack)-1]] >= arr[i] {
//			rights[stack[len(stack)-1]] = i
//			stack = stack[:len(stack)-1]
//		}
//		stack = append(stack, i)
//	}
//	stack = []int{}
//	for i := n - 1; i >= 0; i-- {
//		for len(stack) > 0 && arr[stack[len(stack)-1]] > arr[i] {
//			lefts[stack[len(stack)-1]] = i
//			stack = stack[:len(stack)-1]
//		}
//		stack = append(stack, i)
//	}
//
//	for i := 0; i < n; i++ {
//		ret += arr[i] * (i - lefts[i]) % M * (rights[i] - i) % M
//	}
//	return ret % M
//}
