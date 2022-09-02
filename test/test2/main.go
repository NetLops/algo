package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	var m int
	floats := []float64{}
	ints := []int{}
	for {
		scanln, err := fmt.Scanln(&n, &m)
		if scanln != 2 {
			break
		}
		if err != nil {
			break
		}
		floats = append(floats, n)
		ints = append(ints, m)
	}
	ans := []float64{}
	var temp float64
	for i := 0; i < len(ints); i++ {
		n = floats[i]
		m = ints[i]
		temp = n
		for i := 0; i < m-1; i++ {
			n = math.Sqrt(n)
			temp += n
		}
		ans = append(ans, temp)
	}
	for i := range ans {
		fmt.Printf("%.2f\n", ans[i])
	}

}
