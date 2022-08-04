package main

import "fmt"

//
//func lengthOfLongestSubstring(s string) int {
//	//position := 0 // 用来计算 重复出现的概率
//	maxLen := 0
//	length := 1
//	m := map[byte]int{}
//	h := []byte{}
//	for i := range s {
//		// 判断是否有重复
//		if m[s[i]] > 0 {
//			if m[s[i]]+1 == length {
//				h = []byte{}
//				m = map[byte]int{}
//				length = 1
//			} else {
//				// 截取部分
//				//deleteTemp := h[:m[s[i]]+1]
//				h = h[m[s[i]] : length-1]
//				length = 1
//				m = map[byte]int{}
//				for i2 := range h {
//					m[h[i2]] = length
//					length++
//				}
//			}
//		}
//		h = append(h, s[i])
//		m[s[i]] = length // 对应h的下标
//		if length > maxLen {
//			maxLen = length
//		}
//		length++
//
//		// fmt.Printf("%v %b\n", length, position)
//
//	}
//	return maxLen
//}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{} // 记录最后一次出现的位置
	// 全部初始化为-1
	for i := range s {
		m[s[i]] = -1
	}
	if s == "" {
		return 0
	}
	ans := 1
	pre := 0
	for j := range s {
		i := m[s[j]]   // 获取索引	// 最近最后一次出现的索引
		if pre < j-i { // s[j] 不在 m 中就代表可以继续往前延伸
			pre = pre + 1
		} else { // 找到重复的点的位置，然后取两者之间的
			pre = j - i
		}
		m[s[j]] = j // 取最后出现的索引
		ans = max(pre, ans)
	}
	return ans
}
func main() {
	//s := []int{1, 2, 3, 4}
	//fmt.Println(s[:2])
	//lengthOfLongestSubstring("dvdf")
	fmt.Println(lengthOfLongestSubstring("ggububgvfk"))
}
