package main

func isLetter(v byte) bool {
	if v >= 'a' && v <= 'z' {
		return true
	}
	if v >= 'A' && v <= 'Z' {
		return true
	}

	return false
}

func getIndexForLetter(v byte) int {
	if v < 'a' {
		return int(v - 'A')
	}
	return int(v - 'a')
}

// 参考：https://leetcode.cn/problems/letter-case-permutation/solution/by-ac_oier-x7f0/
func letterCasePermutation(s string) []string {
	str := make([]byte, len(s))
	ans := []string{}
	var dfs func(index int)
	dfs = func(index int) {
		if index == len(s) {
			ans = append(ans, string(str))
			return
		}
		str[index] = s[index]
		dfs(index + 1)
		if isLetter(s[index]) {
			str[index] = s[index] ^ 32
			dfs(index + 1)
		}
	}
	dfs(0)
	return ans
}
