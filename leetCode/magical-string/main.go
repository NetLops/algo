package main

import "bytes"

//参考:https://leetcode.cn/problems/magical-string/solution/by-endlesscheng-z8o1/
func magicalString(n int) (count int) {
	str := make([]byte, 0, n+1)
	str = append(str, 1, 2, 2)
	convert := []byte{2}
	for i := 2; len(str) < n; i++ {
		convert[0] ^= 3 // 1 ^ 3 = 2, 2 ^ 3 = 1	// 快速转换
		str = append(str, bytes.Repeat(convert, int(str[i]))...)
	}

	str = str[:n]

	for i := 0; i < len(str); i++ {
		if str[i] == 1 {
			count++
		}
	}
	return
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	nodes := make([]int, numCourses)
	q := make([]int, 0, numCourses)

	for _, requisites := range prerequisites {
		edges[requisites[1]] = append(edges[requisites[1]], requisites[0])
		nodes[requisites[0]]++
	}

	for i, node := range nodes {
		if node == 0 {
			q = append(q, i)
		}
	}
	for ; len(q) > 0 && numCourses > 0; numCourses, q = numCourses-1, q[1:] {
		node := q[0]
		for _, v := range edges[node] {
			if nodes[v]--; nodes[v] == 0 {
				q = append(q, v)
			}
		}
	}
	if len(q) == 0 && numCourses == 0 {
		return true
	}
	return false
}

func maxSlidingWindow(nums []int, k int) (ret []int) {
	if k == 1 {
		return nums
	}
	q := []int{0}
	for i := 1; i < len(nums); i++ {
		for i-q[0] >= k {
			q = q[1:]
		}
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		if i >= k-1 {
			ret = append(ret, nums[q[0]])
		}
	}
	return ret
}
