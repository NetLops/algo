package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	m := map[int]struct{}{} // 假装自己是个map
	for _, num := range nums {
		m[num] = struct{}{}
	}
	ans := []int{}
	for i := 1; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	m := 1
	m |= 1 << 4
	fmt.Println(m)
	m |= 1 << 4
	fmt.Println(m)
}
