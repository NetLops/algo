package main

import (
	"fmt"
	"sort"
	"unsafe"
)

func countRangeSum(nums []int, lower int, upper int) int {
	tr := make([]int, 100010*3)
	m := 0
	lowbit := func(x int) int {
		return x & -x // 找父下标
	}

	add := func(x, v int) {
		for i := x; i <= m; i += lowbit(i) {
			tr[i] += v
		}
	}

	query := func(x int) int {
		ans := 0
		// 因为从1开始
		for i := x; i > 0; i -= lowbit(i) {
			ans += tr[i]
		}
		return ans
	}

	// 做个set
	set := map[int64]struct{}{}
	var s int64
	set[s] = struct{}{} // 将0 填充一手

	for _, num := range nums {
		s += *(*int64)(unsafe.Pointer(&num))
		set[s] = struct{}{}
		set[s-*(*int64)(unsafe.Pointer(&lower))] = struct{}{}
		set[s-*(*int64)(unsafe.Pointer(&upper))] = struct{}{}
	}

	list := make([]int64, len(set))
	i := 0
	for s2, _ := range set {
		list[i] = s2
		i++
	}
	mi := map[int64]int{}
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
	for _, x := range list {
		m++
		mi[x] = m
	}
	s = 0

	ans := 0
	add(mi[s], 1)
	for _, num := range nums {
		s += *(*int64)(unsafe.Pointer(&num))
		a := mi[s-*(*int64)(unsafe.Pointer(&lower))]
		b := mi[s-*(*int64)(unsafe.Pointer(&upper))] - 1
		ans += query(a) - query(b)
		add(mi[s], 1)
	}
	return ans
}
func init() {
	fmt.Println("为能被调用")
}
func main() {
	//fmt.Println(countRangeSum([]int{-2, 5, -1}, -2, 2))

	//s := []int{}
	//s2 := map[int]int{}
	//s3 := make(chan string, 2)
	//
	//fmt.Println(cap(s))
	//fmt.Println(len(s2))
	//fmt.Println(cap(s3))

	//s := make([]int, 0)
	//s2 := [2]int{1, 2}
	//fmt.Println(reflect.TypeOf(s))
	//fmt.Println(reflect.TypeOf(s2))
	//fmt.Println(reflect.TypeOf(s3))
	s := map[string]int{}

	s["li"] = 1
	s["is"] = 1
	s["sad"] = 1
	s["3ee"] = 1
	s["lrer"] = 1
	s["trei"] = 1
	s["pli"] = 1
	for k, v := range s {
		fmt.Println(k, v)
	}

}
