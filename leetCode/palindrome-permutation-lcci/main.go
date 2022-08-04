package main

import "fmt"

func canPermutePalindrome(s string) bool {
	h := map[rune]int{}
	flag := 0
	for _, v := range s {
		h[v]++
	}
	for _, v := range h {
		if v%2 == 1 {
			flag++
		}
		if flag > 1 {
			return false
		}
	}
	return true
}

func canPermutePalindrome2(s string) bool {
	var level1 int32
	var level2 int32
	var level3 int32
	var level4 int32

	for _, v := range s {
		v = v - 1
		if v < 32 {
			level1 ^= 1 << v
		} else if v >= 32 && v < 64 {
			level2 ^= 1 << (v - 32)
		} else if v >= 64 && v < 96 {
			level3 ^= 1 << (v - 64)
		} else if v >= 96 && v < 128 {
			level4 ^= 1 << (v - 96)
		}
	}

	if (level1&(level1-1)) <= 1 && level2 == 0 && level3 == 0 && level4 == 0 {
		return true
	} else if (level2&(level2-1)) <= 1 && level1 == 0 && level3 == 0 && level4 == 0 {
		return true
	} else if (level3&(level3-1)) <= 1 && level2 == 0 && level1 == 0 && level4 == 0 {
		return true
	} else if (level4&(level4-1)) <= 1 && level2 == 0 && level3 == 0 && level1 == 0 {
		return true
	} else {
		return false
	}

}

// 32 位数的
func canPermutePalindrome3(s string) bool {
	var level1 int64
	var level2 int64

	for _, v := range s {
		v = v - 1
		if v < 64 {
			level1 ^= 1 << v
		} else {
			level2 ^= 1 << (v - 64)
		}
	}

	fmt.Println(level1&(level1-1), level2&(level2-1))
	if (level1&(level1-1)) <= 1 && level2 == 0 {
		return true
	} else if (level2&(level2-1)) <= 1 && level1 == 0 {
		return true
	} else {
		return false
	}

}

func main() {
	//s := "tactcoa7"
	s := "codecode"
	fmt.Println(canPermutePalindrome3(s))
	fmt.Println(6 ^ 1)
}
