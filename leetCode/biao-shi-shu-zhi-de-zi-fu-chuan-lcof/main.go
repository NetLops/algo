package main

import (
	"fmt"
)

//func isNumber(s string) bool {
//	s = strings.Trim(s, " ")
//	matchString, _ := regexp.MatchString("^[+-]?(\\d+(\\.\\d*)?|\\.\\d+)([eE][+-]?\\d+)?$", s)
//	return matchString
//}

// 1.开始的空格
// 2.幂符号前的正负号
// 3.小数点前的数字
// 4.小数点、小数点后的数字
// 5.当小数点前为空格时，小数点、小数点后的数字
// 6.幂符号
// 7.幂符号后的正负号
// 8.幂符号后的数字
// 9.结尾的空格
func isNumber(s string) bool {
	states := []map[byte]int{
		{' ': 1, 's': 2, 'd': 3, '.': 5}, // 1. start with 'blank'
		{'d': 3, '.': 5},                 // 2. 'sign' before 'e'
		{'d': 3, '.': 4, 'e': 6, ' ': 9}, // 3. 'digit' before 'dot'
		{'d': 4, 'e': 6, ' ': 9},         // 4. 'digit' after 'dot'
		{'d': 4},                         // 5. 'digit' after 'dot' ('blank' before 'dot')
		{'s': 7, 'd': 8},                 // 6. 'e'
		{'d': 8},                         // 7. 'sign' after 'e'
		{'d': 8, ' ': 9},                 // 8. 'digit' after 'e'
		{' ': 9},                         // 9. end with 'blank'
	}
	p := 1 // start with state 0
	t := '?'
	for _, c := range s {
		if '0' <= c && c <= '9' { // digit
			t = 'd'
		} else if c == '+' || c == '-' { // sign
			t = 's'
		} else if c == 'e' || c == 'E' { // e or E
			t = 'e'
		} else if c == '.' || c == ' ' { // dot, blank
			t = c
		} else { // unknown
			t = '?'
		}
		if states[p-1][byte(t)] == 0 {
			return false
		}
		p = states[p-1][byte(t)]
	}
	return p == 3 || p == 4 || p == 8 || p == 9
}
func main() {
	//fmt.Println(isNumber("E9"))
	for _, v := range []string{"12e", "1a3.14", "1.2.3", "+-5", "12e+5.4"} {
		//for _, v := range []string{"+100", "5e2", "-123", "3.1416", "-1E-16", "0123"} {
		fmt.Println(v)
		fmt.Println(isNumber(v))
	}

}
