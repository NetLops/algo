package main

import "fmt"

func kSimilarity(s1 string, s2 string) int {
	var next func(s string) []string
	next = func(s string) []string {
		res := []string{}
		i := 0
		for ; i < len(s); i++ {
			if s[i] != s2[i] {
				break
			}
		}
		for j := i + 1; j < len(s); j++ {
			// 保证 s[j] == s2[j] 的时候 不随便动
			if s[j] == s2[i] && s[j] != s2[j] {
				// 换的位置
				res = append(res, s[:i]+string(s[j])+s[i+1:j]+string(s[i])+s[j+1:])
			}
		}
		return res
	}

	visited := map[string]bool{s1: true}
	ans := 0

	// 默认队列中 存的就是s1
	q := []string{s1}
	for {
		for i := len(q); i > 0; i-- {
			s := q[0]
			q = q[1:]
			if s == s2 {
				return ans
			}
			for _, nxt := range next(s) {
				if !visited[nxt] {
					visited[nxt] = true
					q = append(q, nxt)
				}
			}
		}
		ans++
	}
	return ans
}
func main() {
	fmt.Println(kSimilarity("abcdadc", "aabccdd"))
}
