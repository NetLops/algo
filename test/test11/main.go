package main

import (
	"fmt"
	"sort"
)

func main() {
	total, revice := 0, 0
	fmt.Scanln(&total, &revice)
	p := make([]int, total)
	score := make([]int, total)
	for i := 0; i < total; i++ {
		_, err := fmt.Scanf("%d", &p[i])
		if err != nil {
			i = -1
		}
	}
	fmt.Scanln()
	for i := 0; i < total; i++ {
		_, err := fmt.Scanf("%d", &score[i])
		if err != nil {
			i = -1
		}
	}
	m := map[int]int{}
	for i := 0; i < total; i++ {
		m[score[i]] = p[i]
	}
	sort.Ints(score)
	maxScore := 0
	for i := range score[:len(score)-revice] {
		maxScore += m[score[i]] * score[i]
	}
	for i := range score[len(score)-revice:] {
		maxScore += score[len(score)-revice+i] * 100
	}

	fmt.Printf("%.2f", float64(maxScore)/100.0)
}
