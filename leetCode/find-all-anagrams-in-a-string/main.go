package main

func findAnagrams(s string, p string) []int {
	ans := []int{}
	letters := make([]int, 26)
	// example aabc, a: 2,b: 1,c: 1
	for i := 0; i < len(p); i++ {
		letters[p[i]-'a']++
	}
	for l, r := 0, 0; r < len(s); r++ {
		letters[s[r]-'a']--
		for letters[s[r]-'a'] < 0 { // 比如出现不相干的 字母全部移到不相干字母的位置(毕竟不用包含它) /小于0，就l往右边移动（毕竟不能超过最大出现次数），letters[s[l] - 'a'] 复位
			letters[s[l]-'a']++
			l++
		}

		if r-l+1 == len(p) { // 全部置0 此处代表找到了
			ans = append(ans, l)
		}
	}
	return ans
}

func main() {
	findAnagrams("cbaebabacd", "aabc")
}
