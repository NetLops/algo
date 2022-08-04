package main

import "fmt"

func getCounts(a, b int) int {
	count := 0
	temp := a
	if a > b || b > 999999 || b < 100000 {
		return 0
	}
	for temp <= b {
		ta := temp / 100000
		tb := (temp % 100000) / 10000
		tc := (temp % 10000) / 1000
		td := (temp % 1000) / 100
		te := (temp % 100) / 10
		tf := temp % 10

		//fmt.Println(ta, tb, tc, td, te, tf)
		if ta != tb && ta != tc && ta != td && ta != te && ta != tf &&
			tb != tc && tb != td && tb != te && tb != tf &&
			tc != td && tc != te && tc != tf &&
			td != te && td != tf &&
			te != tf {
			t1 := temp / 10000
			t2 := (temp % 10000) / 100
			t3 := (temp % 100)
			//fmt.Println(temp, t1, t2, t3)
			if t1+t2 == t3 {
				//fmt.Println(temp, t1, t2, t3)
				count++
			}

		}
		temp++

	}
	return count
}
func main() {
	a, b := 100000, 999999
	//fmt.Scanf("%d %d", &a, &b)
	//fmt.Println(a,b)
	fmt.Println(getCounts(a, b))
}
