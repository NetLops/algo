package main

import (
	"fmt"
)

func mathch(ints []int) int {
	visited := make([]bool, len(ints))
	flag := false
	var dfs func(sum int, num int, index int)

	dfs = func(sum int, num int, index int) {
		//fmt.Println(sum, index, visited)
		if index > len(ints) {
			flag = false
			return
		}
		if sum == num {
			flag = true
			return
		}
		if flag {
			//fmt.Println(sum)
			return
		}
		for i := 0; i < len(ints); i++ {
			if !visited[i] {
				visited[i] = true
				sum += ints[i]
				dfs(sum, num, index+1)
				if flag {
					//fmt.Println(sum)
					return
				}
				sum -= ints[i]
				visited[i] = false
			}
		}
	}

	ans := 0
	for {
		flag = false
		visited = make([]bool, len(ints))

		ans++
		dfs(0, ans, 0)
		//fmt.Println(ans)
		if !flag {
			break
		}
	}
	//dfs(0, 8, 0)
	//fmt.Println(flag)
	return ans
}
func main() {
	sum := []int{}
	for {
		i := 0
		_, err := fmt.Scanf("%d", &i)
		if err != nil {
			//fmt.Println(err)
			break
		}
		sum = append(sum, i)
	}
	fmt.Println(mathch(sum))
}

//1 4 10 3 1
