package main

import "fmt"

func frequencySort(s string) string {
	m := map[byte]int{}
	max := 0
	for i := range s {
		m[s[i]]++
		if m[s[i]] > max {
			max = m[s[i]]
		}
	}
	n := map[int][]byte{}
	for i, b := range m {
		if n[b] == nil {
			n[b] = []byte{}
		}
		n[b] = append(n[b], i)
	}
	//fmt.Println(m)
	//fmt.Println(n)
	//fmt.Println(max)
	ans := []byte{}
	for i := max; i >= 0; i-- {
		if n[i] != nil {
			for _, value := range n[i] {
				for j := 0; j < i; j++ {
					ans = append(ans, value)
				}
			}

		}
	}
	return string(ans)
}
func main() {
	fmt.Println(frequencySort("cccaaa"))
	fmt.Println(frequencySort("tree"))
	fmt.Println(frequencySort("Aabb"))
}
