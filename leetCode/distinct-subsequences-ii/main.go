package main

func distinctSubseqII(s string) int {
	MOD := int(1e9 + 7)
	ans := 0
	// dict
	letters := make([]int, 26)
	for _, char := range s {
		c := char - 'a'
		prev := letters[c]
		letters[c] = (ans + 1) % MOD
		ans = (ans + letters[c]) % MOD
		ans = (ans - prev + MOD) % MOD
	}
	return ans
}
