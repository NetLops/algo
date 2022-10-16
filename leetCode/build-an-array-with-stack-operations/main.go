package main

import "fmt"

func buildArray(target []int, n int) []string {
	index := 1
	ans := []string{}

	for i := 0; i < len(target); i++ {
		if index < target[i] {
			for index != target[i] {
				ans = append(ans, "Push", "Pop")
				index++
			}
		}
		if index == target[i] {
			ans = append(ans, "Push")
			index++
		}
	}

	return ans
}

func main() {
	fmt.Println(buildArray([]int{1, 3}, 3))
}
