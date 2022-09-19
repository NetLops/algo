package main

import (
	"fmt"
	"sort"
)

func countRangeSum(nums []int, lower int, upper int) int {
	tr := make([]int, 100010*3)
	m := 0
	lowbit := func(x int) int {
		return x & -x
	}

	add := func(x, v int) {
		for i := x; i <= m; i += lowbit(i) {
			tr[i] += v
		}
	}

	query := func(x int) int {
		ans := 0
		for i := x; i > 0; i -= lowbit(i) {
			ans += tr[i]
		}
		return ans
	}
	// 假装set
	set := map[int64]struct{}{}

	var s int64
	set[s] = struct {
	}{}
	for _, num := range nums {
		s += int64(num)
		set[s] = struct{}{}
		set[s-int64(lower)] = struct{}{}
		set[s-int64(upper)] = struct{}{}
	}

	list := make([]int64, len(set))
	i := 0
	for s2, _ := range set {
		list[i] = s2
		i++
	}
	mi := map[int64]int{}
	sort.Slice(list, func(i, j int) bool {
		if list[i] < list[j] {
			return true
		}
		return false
	})
	for _, x := range list {
		m++
		mi[x] = m
	}
	s = 0
	ans := 0
	add(mi[s], 1)
	for _, i2 := range nums {
		s += int64(i2)
		a := mi[s-int64(lower)]
		b := mi[s-int64(upper)] - 1
		ans += query(a) - query(b)
		add(mi[s], 1)
	}
	return ans
}

func main() {
	fmt.Println(countRangeSum([]int{-2, 5, -1}, -2, 2))
}
