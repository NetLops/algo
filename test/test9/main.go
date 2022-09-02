package main

import "fmt"

func main() {
	n, a, b := 0, "", ""
	fmt.Scanln(&n)
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	ans := []byte{}
	for i := 0; i < n; i++ {
		ans = append(ans, a[i])
		ans = append(ans, b[i])
	}
	fmt.Println(string(ans))
}
