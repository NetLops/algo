package main

import "fmt"

func move(n int, A, B, C *[]int) {
	if n > 0 {
		move(n-1, A, C, B)
		*C = append(*C, (*A)[len(*A)-1])
		*A = (*A)[:len(*A)-1]
		move(n-1, B, A, C)
	}

}
func hanota(A []int, B []int, C []int) []int {
	n := len(A)
	move(n, &A, &B, &C)
	return C
}
func main() {
	A := []int{2, 1, 0}
	B := []int{}
	C := []int{}
	fmt.Println(hanota(A, B, C))
}
