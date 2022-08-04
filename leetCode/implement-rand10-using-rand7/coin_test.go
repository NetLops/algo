package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func coin() int {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(10)%10 > 5 {
		return 1
	} else {
		return 0
	}
}

func coin_new() int {
	var a int
	for true {
		a = coin()
		if coin() != a {
			return a
		}
	}
	return a
}

func TestCoin1(t *testing.T) {

	count0 := 0
	count1 := 0
	cycleCount := 100
	for i := 0; i < cycleCount; i++ {
		if coin_new() == 0 {
			count0++
		} else {
			count1++
		}
	}
	fmt.Printf("0的概率是：%f，次数：%d\n", float64(count0)/float64(cycleCount), count0)
	fmt.Printf("1的概率是：%f，次数：%d\n", float64(count1)/float64(cycleCount), count1)
}

func coin2() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2)
}

func coin_new2() int {
	if (coin2() | coin2()) == 1 {
		return 1
	} else {
		return 0
	}
}

func TestCoin2(t *testing.T) {
	count0 := 0
	count1 := 0
	cycleCount := 100
	for i := 0; i < cycleCount; i++ {
		if coin_new2() == 0 {
			count0++
		} else {
			count1++
		}
	}
	fmt.Printf("0的概率是：%f，次数：%d\n", float64(count0)/float64(cycleCount), count0)
	fmt.Printf("1的概率是：%f，次数：%d\n", float64(count1)/float64(cycleCount), count1)
}

func coin3() int {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(10)%10 > 4 {
		return 0
	} else {
		return 1
	}
}

func coin_new3() int {
	//for true {
	//	a, b := coin3(), coin3()
	//	//if a == 0 && b == 0 {
	//	//	continue
	//	//}
	//	if a&b == 1 {
	//		return 0
	//	}
	//	if a|b == 1 {
	//		return 1
	//	}
	//	//if a == 0 && b == 0 {
	//	//	fmt.Println("空")
	//	//}
	//
	//}
	//return 0
	a, b := coin3(), coin3()
	//if a == 0 && b == 0 {
	//	continue
	//}
	if a&b == 1 {
		return 0
	}
	if a|b == 1 {
		return 1
	}
	return coin_new3()
}

func TestCoin3(t *testing.T) {
	count0 := 0
	count1 := 0
	cycleCount := 10000
	var temp int
	for i := 0; i < cycleCount; i++ {
		//temp = coin_new3()
		temp = coin2()
		if temp == 0 {
			count0++
		}
		if temp == 1 {
			count1++
		}
	}
	fmt.Printf("0的概率是：%f，次数：%d\n", float64(count0)/float64(cycleCount), count0)
	fmt.Printf("1的概率是：%f，次数：%d\n", float64(count1)/float64(cycleCount), count1)

}

func TestA(t *testing.T) {
	count00 := 0
	count11 := 0
	count01 := 0
	count10 := 0
	cycleCount := 10000
	var a, b int
	for i := 0; i < cycleCount; i++ {
		a = coin3()
		b = coin3()
		if a == 0 && b == 0 {
			count00++
		}
		if a == 1 && b == 1 {
			count11++
		}
		if a == 0 && b == 1 {
			count01++
		}
		if a == 1 && b == 0 {
			count10++
		}
	}
	fmt.Printf("00的概率是：%f，次数：%d\n", float64(count00)/float64(cycleCount), count00)
	fmt.Printf("11的概率是：%f，次数：%d\n", float64(count11)/float64(cycleCount), count11)
	fmt.Printf("01的概率是：%f，次数：%d\n", float64(count01)/float64(cycleCount), count01)
	fmt.Printf("10的概率是：%f，次数：%d\n", float64(count10)/float64(cycleCount), count10)
}

func coin_new4() int {
	temp := 0
	for true {
		temp = 0
		for i := 0; i < 4; i++ {
			temp = temp<<1 + coin3()
		}
		//return temp
		if temp <= 2 {

			return 0
		}
		if temp <= 9 {
			return 1
		}
	}
	return temp
}

func TestCoin4(t *testing.T) {
	count0 := 0
	count1 := 0
	cycleCount := 100
	var temp int
	for i := 0; i < cycleCount; i++ {
		temp = coin_new4()
		if temp == 0 {
			count0++
		}
		if temp == 1 {
			count1++
		}
	}
	fmt.Printf("0的概率是：%f，次数：%d\n", float64(count0)/float64(cycleCount), count0)
	fmt.Printf("1的概率是：%f，次数：%d\n", float64(count1)/float64(cycleCount), count1)

}

func TestB(t *testing.T) {
	for i := 0; i < 10000; i++ {
		fmt.Println(coin_new4())
	}
}

func coin_new5() int {
	temp := 0
	for true {
		temp = 0
		for i := 0; i < 4; i++ {
			temp = temp<<1 + coin3()
		}
		if temp >= 1 && temp <= 10 {
			return temp
		}
	}
	return temp
}

func TestCoin5(t *testing.T) {
	m := map[int]int{}
	var temp int
	for i := 0; i < 1000; i++ {
		temp = coin_new5()
		m[temp]++
	}
	fmt.Println(m)
}
