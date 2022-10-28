package main

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	collection := make([]int, n)
	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) != 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			x := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			collection[x] = i - x
		}

		stack = append(stack, i)
	}
	return collection
}
