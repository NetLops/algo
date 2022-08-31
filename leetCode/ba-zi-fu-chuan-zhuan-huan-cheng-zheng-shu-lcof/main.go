package main

import (
	"fmt"
	"math"
	"unsafe"
)

func strToInt(str string) int {
	// Go默认默认值为0，害难受😣
	states := []map[byte]int{
		{' ': 1, 'f': 2, 'd': 3, 'o': 5}, // 第一位可能永远是符号
		{'d': 3, 'f': 5, ' ': 5, 'o': 5}, // 符号位
		{'d': 3, ' ': 5, 'f': 5, 'o': 5},
		{}, // 到这儿来就代表无
	}
	index := 0
	flag := 1
	num := []byte{}
	for i := 0; i < len(str); i++ {
		sign := '0'
		if str[i] == ' ' {
			sign = ' '
		} else if '0' <= str[i] && str[i] <= '9' {
			sign = 'd'
		} else if str[i] == '+' || str[i] == '-' {
			sign = 'f'
		} else {
			sign = 'o'
		}
		index = states[index][byte(sign)] - 1
		if index == 4 {
			break
		}
		if str[i] == '-' {
			flag = -1
		} else if '0' <= str[i] && str[i] <= '9' {
			num = append(num, str[i])
		}
	}

	fmt.Println(*(*string)(unsafe.Pointer(&num)))
	ans := 0
	for i := 0; i < len(num); i++ {
		ans = ans*10 + int(num[i]-'0')
		if flag == 1 && (ans>>31 > 0) {
			return math.MaxInt32
		} else if flag == -1 && ((-1*ans)&math.MinInt32 < math.MinInt32) {
			return math.MinInt32
		}
	}
	return flag * ans
}
func main() {
	fmt.Println(math.MinInt32)
	fmt.Println(math.MinInt32 & -2147483648)
	fmt.Printf("%b\n", math.MinInt32)
	fmt.Printf("%b\n", -2147483649)
	fmt.Printf("%b\n", math.MaxInt32)
	fmt.Printf("%b, %d\n", math.MinInt32&-2147483649, math.MinInt32&-2147483649)

	fmt.Println(math.MaxInt32)
	fmt.Println(strToInt("2147483648"))
}
