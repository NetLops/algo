package main

func uniqueLetterString(s string) int {
	m := [26][]int{}
	n := len(s)
	// 数组都灌好
	for i := range m {
		m[i] = []int{}
	}

	// 虚拟构造范围 -1
	for i := 0; i < 26; i++ {
		m[i] = append(m[i], -1)
	}
	// 找区间
	for i := 0; i < n; i++ {
		m[s[i]-'A'] = append(m[s[i]-'A'], i)
	}
	// 虚拟构造范围 n
	for i := 0; i < 26; i++ {
		m[i] = append(m[i], n)
	}

	ret := 0
	// 然后简单的遍历一手
	for i := 0; i < 26; i++ {
		// [1,n - 1] 的范围
		for j := 1; j < len(m[i])-1; j++ {
			ret += (m[i][j] - m[i][j-1]) * (m[i][j+1] - m[i][j])
		}

	}
	return ret
}

/**
x x A [x x x x A x x x]  A  x  x
0 1 2  3 4 5 6 7 8 9 10 11 12  13
A 贡献的分：(7 -2) * (11 - 7)
   x C x x x
-1 0 1 2 3 4 n	// n 为数组长度， 虚拟构造 -1,n
c 贡献的分：(1--1) * n - 1
*/
func main() {

}
