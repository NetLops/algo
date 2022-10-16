package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scanln(&n)
	c := make([]int, n)
	i := 0
	for {
		var temp int
		if _, err := fmt.Scanf("%d", &temp); err != nil {
			//fmt.Println(err)
			break
		} else {
			c[i] = temp
			i++
		}
	}

}
