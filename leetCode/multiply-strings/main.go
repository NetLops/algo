package main

import (
	"fmt"
)

func multiply(num1 string, num2 string) int32 {
	s := ""
	for _, i := range num1 {
		for _, j := range num2 {
			i - j
		}
	}

}
func main() {
	num1 := "123"
	num2 := "456"
	fmt.Println(multiply(num1, num2))
}
