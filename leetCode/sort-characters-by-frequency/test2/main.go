package main

import (
	"sort"
	"unsafe"
)

func frequencySort(s string) string {
	m := map[byte]int{}
	for i := range s {
		m[s[i]]++
	}
	mcount := map[int][]byte{}
	counts := []int{}
	for k, v := range m {
		if _, ok := mcount[v]; !ok {
			mcount[v] = []byte{}
		}
		mcount[v] = append(mcount[v], k)
	}

	for k, _ := range mcount {
		counts = append(counts, k)
	}
	sort.Slice(counts, func(i, j int) bool {
		if counts[i] > counts[j] {
			return true
		}
		return false
	})
	// fmt.Println(counts,mcount)

	ans := []byte{}
	for _, count := range counts {
		for i := range mcount[count] {
			for j := 0; j < count; j++ {
				ans = append(ans, mcount[count][i])
			}
		}
	}
	return *(*string)(unsafe.Pointer(&ans))
}
func main() {

}
