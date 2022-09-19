package main

import (
	"fmt"
)

func flipLights(n int, presses int) int {
	if presses == 0 {
		return 1
	}
	if n == 1 {
		return 2
	} else if n == 2 {
		if presses == 1 {
			return 3
		} else {
			return 4
		}
	} else {
		if presses == 1 {
			return 4
		} else if presses == 2 {
			return 7
		} else {
			return 8
		}
	}

}
func main() {
	i := 0
	//ctx, cancelFunc := context.WithCancel(context.TODO())
	for j := 0; j < 100; j++ {
		go func(i int) {
			i++
		}(i)

	}
	fmt.Println(i)
}
