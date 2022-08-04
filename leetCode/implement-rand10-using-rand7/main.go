package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rand7() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(7)%7 + 1
}

func rand10() int {
	temp := 0
	for true {
		temp = 0
		//for i := 0; i < 2; i++ {
		//	temp = temp*7 + rand7() - 1
		//}
		temp = (rand7()-1)*7 + (rand7() - 1)
		if temp >= 0 && temp <= 39 { // 0 ~ 39
			return temp%10 + 1
		}
		// 剩余 40 41 42 43 44 45 46 47 48
		temp = (temp%40)*7 + rand7() - 1 // 0 ~ 62
		if temp >= 0 && temp <= 59 {
			return temp%10 + 1
		}
		// 剩余 20
		temp = (temp%60)*7 + rand7() - 1 // 0 ~ 20
		if temp >= 0 && temp <= 19 {
			return temp%10 + 1
		}
	}
	return temp
}

func main() {
	m := map[int]int{}
	temp := 0
	for i := 0; i < 10000; i++ {
		temp = rand10()
		m[temp]++
	}
	fmt.Println(m)
}
