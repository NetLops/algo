package main

import (
	"fmt"
)

func main() {
	var left, right int64 = 0, 0

	_, _ = fmt.Scan(&left, &right)
	var ans, index int64 = 1, 1
	for {
		temp := ans // 1
		ans <<= index
		ans |= ((index<<1 - 1) ^ temp)
		index <<= 1
		fmt.Println(ans)

		if index > right {
			break
		}
	}
	ans = ans >> (index - right)
	res := ans & 1
	for right != left {
		ans = ans >> 1
		if ans&1 == 1 {
			res++
		}
		right--
	}
	fmt.Println(res)
	//
	//len := math.pow
	//ans := []int{}
	//
	//for {
	//	ans = {1}
	//}

	//fmt.Println((4 - 1) ^ 2)

}
