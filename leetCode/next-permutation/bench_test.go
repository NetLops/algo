package main

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkName(b *testing.B) {
	//b.StartTimer()
	for i := 0; i < b.N; i++ {
		nums := []int{6, 7, 5, 3, 5, 6, 2, 9, 1, 2, 7, 0, 9}
		//nums := []int{4, 3, 2, 1}
		nextPermutation(nums)
	}
	//b.ResetTimer()
}

func TestName(t *testing.T) {
	collect := []int{1, 2, 4, 6}
	s := make([]byte, 4, 4)
	for i := range collect {
		s[i] = byte(strconv.Itoa(collect[i])[0])
	}
	fmt.Println(s)
	atoi, _ := strconv.Atoi(string(s))
	fmt.Println(atoi)
}
