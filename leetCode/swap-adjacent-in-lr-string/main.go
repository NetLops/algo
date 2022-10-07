package main

import "fmt"

func canTransform(start string, end string) bool {
	n := len(start)
	i, j := 0, 0
	for i < n || j < n {
		for i < n && start[i] == 'X' {
			i++
		}
		for j < n && end[j] == 'X' {
			j++
		}
		if i == n || j == n { // 遍历完了也就无了
			return i == j
		}
		if start[i] != end[j] {
			return false // 无法交换
		}

		if start[i] == 'L' && i < j {
			return false
		}
		if start[i] == 'R' && i > j {
			return false
		}
		i++
		j++
	}
	return true
}

func main() {
	var s fmt.Stringer

	fmt.Println(s)

}
