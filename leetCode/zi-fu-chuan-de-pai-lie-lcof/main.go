package main

import (
	"fmt"
	"sort"
)

func permutation(s string) []string {
	collect := make([]byte, len(s))
	ans := []string{}
	visited := map[int]bool{}
	// 先排序一波
	temp := []byte(s)
	sort.Slice(temp, func(i, j int) bool {
		if temp[i] < temp[j] {
			return true
		}
		return false
	})
	//fmt.Println(temp)
	var dfs func(collect []byte, index int)

	dfs = func(collect []byte, index int) {
		if index == len(collect) {
			ans = append(ans, string(collect))
			return
		}
		for i := 0; i < len(temp); i++ {
			if visited[i] || (i > 0 && temp[i-1] == temp[i] && !visited[i-1]) {

			} else {
				visited[i] = true
				collect[index] = temp[i]
				dfs(collect, index+1)
				visited[i] = false

			}
			//fmt.Println(visited)
		}
	}
	dfs(collect, 0)
	return ans
}
func main() {
	//fmt.Println(permutation("abc"))
	//fmt.Println(permutation("aab"))
	//strs := permutation("abcc")
	strs := permutation("suvyls")
	sort.Strings(strs)

	fmt.Println(strs)
}
