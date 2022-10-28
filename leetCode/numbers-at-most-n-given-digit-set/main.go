package main

import "strconv"

func atMostNGivenDigitSet(digits []string, n int) int {
	ans, prev := 1, 1
	min := -1
	chs := []byte(strconv.Itoa(n))
	for i, m := 1, 1; i <= len(chs); i, m = i+1, m*len(digits) {
		prev, ans = ans, 0
		ch := int(chs[len(chs)-i] - '0')
		for j := 0; j < len(digits); j++ {
			atoi, _ := strconv.Atoi(digits[j])
			if atoi < ch {
				ans += m
			} else if atoi == ch {
				ans += prev
			}
		}
		min += m
	}
	return ans + min
}
