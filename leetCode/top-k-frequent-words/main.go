package main

import "sort"

func topKFrequent(words []string, k int) []string {
	ans := []string{}
	max := 0

	m := map[string]int{}
	for _, word := range words {
		m[word]++
		if m[word] > max {
			max = m[word]
		}
	}
	countMap := make([][]string, max)
	for key, v := range m {

		if countMap[v-1] == nil {
			countMap[v-1] = []string{}
		}
		countMap[v-1] = append(countMap[v-1], key)
	}

	for i := max - 1; i >= 0; i-- {
		if countMap[i] != nil {

			if len(countMap[i]) > 1 {
				sort.Strings(countMap[i])
			}
			for _, s := range countMap[i] {
				ans = append(ans, s)
				k--
				if k == 0 {
					return ans
				}
			}

		}
	}

	return ans

}
func main() {

}
