package main

import (
	"unsafe"
)

func isDigit(v byte) bool {
	return '0' <= v && v <= '9'

}
func reformatNumber(number string) string {
	str := []byte{}

	for i := range number {
		if isDigit(number[i]) {
			str = append(str, number[i])
		}
	}

	ans := []byte{}
	for len(str) != 0 {
		if len(str) == 4 {
			ans = append(append(append(ans, str[0:2]...), '-'), str[2:]...)
			str = nil
		} else if len(str) > 3 {
			ans = append(append(ans, str[0:3]...), '-')
			str = str[3:]
		} else {
			ans = append(ans, str[:]...)
			str = nil
		}
	}
	return *(*string)(unsafe.Pointer(&ans))
}

//func main() {
//	str := []byte{}
//
//	str = nil
//	fmt.Println(len(str))
//}
