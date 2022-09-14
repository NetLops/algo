package main

import (
	"fmt"
	"unsafe"
)

func reverseParentheses(s string) string {
	ans := []byte{}
	m := map[int][]byte{}
	index := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			index++
			m[index] = []byte{}
		} else if s[i] == ')' {
			if index-1 < 0 {
				for j := len(m[index]) - 1; j >= 0; j-- {
					ans = append(ans, m[index][j])
				}
				return string(*(*string)(unsafe.Pointer(&ans)))
			}
			if m[index-1] == nil {
				m[index-1] = []byte{}
			}
			for j := len(m[index]) - 1; j >= 0; j-- {
				m[index-1] = append(m[index-1], m[index][j])
			}
			index--
		} else {

			m[index] = append(m[index], s[i])
		}
	}
	//fmt.Println(m)
	//
	//for k, v := range m {
	//	fmt.Println(k, *(*string)(unsafe.Pointer(&v)))
	//}

	if m[0] != nil {
		return string(m[0])
	}
	return ""
}
func main() {
	//reverseParentheses("(ed(et(oc))el)")
	fmt.Println(reverseParentheses("sxmdll(q)eki(x)"))
	//fmt.Println(reverseParentheses("a(bcdefghijkl(mno)p)q"))
}
