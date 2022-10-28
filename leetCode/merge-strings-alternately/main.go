package main

import "unsafe"

func mergeAlternately(word1 string, word2 string) string {
	ans := make([]byte, len(word1)+len(word2))
	i, j := 0, 0
	index := 0
	for ; i < len(word1) && j < len(word2); i, j = i+1, j+1 {
		ans[index] = word1[i]
		index++
		ans[index] = word2[j]
		index++
	}
	for i < len(word1) {
		ans[index] = word1[i]
		index++
		i++
	}
	for j < len(word2) {
		ans[index] = word2[j]
		index++
		j++
	}
	return *(*string)(unsafe.Pointer(&ans))
}
