package main

import "fmt"

func rotatedDigits(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		flag := true
		temp := i
		has := false
		for temp != 0 {
			now := temp % 10
			if now == 2 || now == 5 || now == 6 || now == 9 {
				has = true
			} else if i > 10 && (now == 1 || now == 0 || now == 8) {

			} else {
				flag = false
				break
			}
			temp = temp / 10
		}
		if flag && has {
			ans++
		}
	}
	return ans
}
func main() {
	fmt.Println(rotatedDigits(10))
}
