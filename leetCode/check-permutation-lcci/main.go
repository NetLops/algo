package main

func CheckPermutation(s1 string, s2 string) bool {
	ms1 := map[byte]int{}
	ms2 := map[byte]int{}

	for i := range s1 {
		ms1[s1[i]]++
	}
	for i := range s2 {
		ms2[s2[i]]++
	}
	if len(ms1) != len(ms2) {
		return false
	}
	for k, v := range ms1 {
		if ms2[k] != v {
			return false
		}
	}
	return true
}
