package main

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		left[i] = -1
		right[i] = n
	}

	q := []int{}

	for i := 0; i < n; i++ {
		for len(q) != 0 && heights[q[len(q)-1]] > heights[i] {
			x := q[len(q)-1]
			right[x] = i
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	q = []int{}

	for i := n - 1; i >= 0; i-- {
		for len(q) != 0 && heights[q[len(q)-1]] > heights[i] {
			x := q[len(q)-1]
			left[x] = i
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	ans := 0
	for i := 0; i < n; i++ {
		nowHeight := heights[i]
		ans = max(ans, nowHeight*(right[i]-left[i]-1))
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
