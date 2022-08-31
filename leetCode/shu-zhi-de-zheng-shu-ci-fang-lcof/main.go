package main

import "fmt"

/*
n = 10 ==> 1010
1 * 2^3 + 0 * 2^2 + 1 * 2^1 + 0 * 2 ^ 0

x^10 = x^(1 * 2^3)* x^(0 * 2^2) * x^(1 * 2^1) * x^(0 * 2 ^ 0)
	= x^(1 * 2^3)* 1 * x^(1 * 2^1)
*/
func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	ans := 1.0
	if n < 0 {
		x = 1 / x
		n = -n
	}
	for n != 0 {
		// 判断末位是否为1
		if n&1 == 1 {
			ans *= x
		}
		x *= x  // x^2、x^4、x^8.....
		n >>= 1 // 1010、 101、 10、 1、 0
	}
	return ans
}
func main() {
	fmt.Println(myPow(2.00000, -2))
}
