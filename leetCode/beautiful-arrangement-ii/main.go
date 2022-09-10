package main

import "fmt"

func constructArray(n int, k int) []int {
	ans := []int{}
	i := 0
	even := 1
	odd := n
	for k > 1 && i < k {
		if i%2 == 0 {
			ans = append(ans, even)
			even++

		} else {
			ans = append(ans, odd)
			odd--
		}
		i++
	}
	if k%2 == 0 {
		for i < n {
			ans = append(ans, odd)
			odd--
			i++
		}
	} else {
		for i < n {
			ans = append(ans, even)
			even++
			i++
		}
	}
	//fmt.Println(ans)
	return ans

}

func main() {
	fmt.Println(constructArray(5, 2))
}
