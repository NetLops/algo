package main

import "reflect"

func findAnagrams(s string, p string) []int {
	n, m := len(s), len(p)
	ret := []int{}
	if n < m {
		return ret
	}
	pCnt, sCnt := make([]int, 26), make([]int, 26)
	for i := 0; i < m; i++ {
		pCnt[p[i]-'a']++
		sCnt[s[i]-'a']++
	}

	if reflect.DeepEqual(pCnt, sCnt) {
		ret = append(ret, 0)
	}

	for i := m; i < n; i++ {
		sCnt[s[i-m]-'a']--
		sCnt[s[i]-'a']++
		if reflect.DeepEqual(sCnt, pCnt) {
			ret = append(ret, i-m+1)
		}
	}

	return ret
}
