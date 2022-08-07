package main

import "fmt"

//func findRepeatNumber(nums []int) int {
//	temp := 0
//	//temp2 := 0
//	for _, i := range nums {
//		if temp&(1<<i) > 0 {
//			return i
//		}
//		temp = temp ^ (1 << i)
//	}
//	//fmt.Println(temp)
//	//fmt.Printf("%b", temp)
//	return nums[0]
//}

//
//func findRepeatNumber(nums []int) int {
//	temp := make([]int, len(nums))
//	for _, i := range nums {
//		if temp[i] == 1 {
//			return i
//		}
//		temp[i] = 1
//	}
//	//fmt.Println(temp)
//	//fmt.Printf("%b", temp)
//	return nums[0]
//}
func findRepeatNumber(nums []int) int {
	m := map[int]int{}
	for _, num := range nums {
		m[num]++
		if m[num] > 1 {
			return num
		}
	}
	return nums[len(nums)-1]
}
func main() {
	//fmt.Println(findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))
	fmt.Println(findRepeatNumber([]int{0, 1, 2, 3, 4, 11, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}))
	//fmt.Println(0 & 1)
}
