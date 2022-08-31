package main

import (
	"fmt"
	"strconv"
)

/*
0123456789101112131415161718192021222324252627282930

num		bit		start		totalBitCount
0~9  	1		0			9*1 = 9
10~99	2		10			90*2 = 180
100~999	3		100			900*3 = 2700
....
*/
func findNthDigit(n int) int {
	if n <= 9 {
		return n
	}
	bit, start, totalBitCount := 1, 0, 9
	for n > totalBitCount {
		n -= totalBitCount
		//fmt.Println(n, totalBitCount)
		bit++
		if start == 0 {
			start = 1 * 10
		} else {
			start *= 10
		}
		totalBitCount = 9 * start * bit
	}
	fmt.Println(n, totalBitCount)
	res := start + (n-1)/bit                       // 找到那个数	// (n - 1)/bit 就是 第多少个的数 比如 100  那  (n - 1) / bit 范围为 100 ~ 999,n的索引为 bit ~ bit * 900
	return int(strconv.Itoa(res)[(n-1)%bit] - '0') //
}
func main() {
	findNthDigit(2889)
	findNthDigit(2890)
	findNthDigit(2899)
	findNthDigit(3)
}
