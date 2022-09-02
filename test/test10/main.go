package main

import (
	"fmt"
)

func main() {
	n := 0
	xy := make([][]int, 3)
	for i := range xy {
		xy[i] = make([]int, 2)
	}
	distance := make([]int, 3)

	fmt.Scanln(&n)
	fmt.Scanln(&xy[0][0], &xy[0][1])
	fmt.Scanln(&xy[1][0], &xy[1][1])
	fmt.Scanln(&xy[2][0], &xy[2][1])
	fmt.Scanln(&distance[0], &distance[1], &distance[2])
	collection := [][][]int{}
	for i := range distance {
		collection = append(collection, getSign(xy[i][0], xy[i][1], distance[i], n))
	}
	u := collection[0]
	for i := 1; i < len(collection); i++ {
		u = union(u, collection[i])
	}
	// 求出最小值
	min := []int{u[0][0], u[0][1]}
	for i := 1; i < len(u); i++ {
		min = getMin(min, u[i])
	}

	fmt.Printf("%d %d", min[0], min[1])

}

// 求出信标

func getSign(a, b, distance, n int) [][]int {
	result := [][]int{}
	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			if abs(a-i)+abs(b-j) == distance {
				temp := []int{i, j}
				result = append(result, temp)
			}
		}
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 求交集
func union(a, b [][]int) [][]int {
	result := [][]int{}
	if len(a) < len(b) {
		for i := 0; i < len(a); i++ {
			for j := 0; j < len(b); j++ {
				if a[i][0] == b[j][0] && a[i][1] == b[j][1] {
					temp := []int{a[i][0], a[i][1]}
					result = append(result, temp)
				}
			}
		}
	} else {
		for i := 0; i < len(b); i++ {
			for j := 0; j < len(a); j++ {
				if b[i][0] == a[j][0] && b[i][1] == a[j][1] {
					temp := []int{b[i][0], b[i][1]}
					result = append(result, temp)
				}
			}
		}
	}
	return result

}

func getMin(a, b []int) []int {
	if a[0] < b[0] || (a[0] == b[0] && a[1] < b[1]) {
		return a
	}
	return b
}
