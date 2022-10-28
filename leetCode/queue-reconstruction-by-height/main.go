package main

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	ans := [][]int{}
	for _, v := range people {
		ans = append(ans, v)
		copy(ans[v[1]+1:], ans[v[1]:])
		ans[v[1]] = v
	}
	return ans
}
