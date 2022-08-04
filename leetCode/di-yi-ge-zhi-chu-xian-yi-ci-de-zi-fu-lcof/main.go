package main

import "fmt"

//func firstUniqChar(s string) byte {
//	////temp := (1 << 28) - 2
//	//temp := 0
//	////fmt.Printf("%b\n", temp)
//	////fmt.Printf("%b\n", 1<<26)
//	//for i := range s {
//	//	pre := temp
//	//	if temp&(1<<(int(s[i])-'a'+1)) > 0 {
//	//	}
//	//	//
//	//	//}
//	//	temp = temp ^ (1 << (int(s[i]) - 'a' + 1))
//	//	fmt.Printf("%26b\n", temp)
//	//	fmt.Printf("%26b\n", pre)
//	//	//fmt.Printf("%26b\n", temp&^pre)
//	//	fmt.Println()
//	//}
//	//return s[0]
//	//pos := 0
//	m := make(map[uint8]int, length)
//	now := 0
//	for i := length - 1; i >= 0; i-- {
//		m[s[i]]++
//		if m[s[i]] == 1 {
//			now = i
//		}
//	}
//	//fmt.Println(now, pre)
//	return s[now]
//}
//func firstUniqChar(s string) byte {
//	length := len(s)
//	if length == 0 {
//		return ' '
//	}
//	//pos := 0
//	m := make(map[uint8]int, length)
//	for i := length - 1; i >= 0; i-- {
//		m[s[i]]++
//	}
//	//fmt.Println(now, pre)
//
//	for i := 0; i < length; i++ {
//		if m[s[i]] == 1 {
//			return s[i]
//		}
//	}
//	return ' '
//}
func firstUniqChar(s string) byte {
	chs, chs2 := 0, 0
	for i := range s {
		abs := s[i] - 'a'
		if chs&(1<<abs) == 0 {
			chs |= (1 << abs)
		} else {
			chs2 |= (1 << abs)
		}
	}
	fmt.Printf("%26b\n", chs)
	fmt.Printf("%26b\n", chs2)
	for i := range s {
		abs := s[i] - 'a'
		if chs&(1<<abs) != 0 && chs2&(1<<abs) == 0 {
			return s[i]
		}
	}
	return ' '
}
func main() {

	firstUniqChar("abaccdefff")
}
