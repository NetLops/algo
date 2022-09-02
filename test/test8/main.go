package main

import "fmt"

func main() {
	var n int // 订单数
	var t int // 正常 所需要的送达订单的时间
	_, _ = fmt.Scanln(&n, &t)
	orders := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &orders[i])
	}

	canTo := 0
	//pre := 0
	for i := 0; i < n; i++ {
		// 没时间做了 害了
		if orders[i]%5 == 0 {
			canTo++
		} else {
			//pre = orders[i] % 5
		}

	}
	//fmt.Println(orders)
	fmt.Println(n - canTo)
}
