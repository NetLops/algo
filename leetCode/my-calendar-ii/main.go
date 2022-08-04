package main

import "fmt"

//type MyCalendarTwo struct {
//	m map[[2]int]int // v为 int 代表次数
//}
//
//func Constructor() MyCalendarTwo {
//	return MyCalendarTwo{m: map[[2]int]int{}}
//}
//
//func (this *MyCalendarTwo) Book(start int, end int) bool {
//	if start > end {
//		return false
//	}
//	temp := [2]int{start, end}
//	for k, _ := range this.m {
//		//fmt.Println(k)
//		//if this.m[k] != 0 {
//		//	fmt.Println(k, this.m[k], temp)
//		//}
//		if this.m[k] != 0 && start < k[1] && end > k[0] {
//			return false
//		}
//	}
//
//	for k, _ := range this.m {
//		if this.m[k] == 0 && start < k[1] && end > k[0] {
//			left, right := 0, 0
//			if start > k[0] {
//				left = start
//			} else {
//				left = k[0]
//			}
//			if end < k[1] {
//				right = end
//			} else {
//				right = k[1]
//			}
//			this.m[[2]int{left, right}] = 1
//		}
//	}
//	_, has := this.m[temp]
//	if !has {
//		this.m[temp] = 0
//	}
//	//fmt.Println(this.m)
//	return true
//}

type MyCalendarTwo struct {
	m, repeat []int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{m: []int{}}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	if start > end {
		return false
	}
	for i := range this.repeat {
		if i%2 == 1 {
			if start < this.repeat[i] && end > this.repeat[i-1] {
				return false
			}
		}
	}

	for i := range this.m {
		if i%2 == 1 {
			if start < this.m[i] && end > this.m[i-1] {

				this.repeat = append(this.repeat, max(start, this.m[i-1]), min(end, this.m[i]))
			}
		}
	}
	this.m = append(this.m, start, end)
	//fmt.Println(this)
	return true
}

func main() {
	myCalendarTwo := Constructor()
	//fmt.Println(myCalendarTwo.Book(10, 20))
	//fmt.Println(myCalendarTwo.Book(50, 60))
	//fmt.Println(myCalendarTwo.Book(10, 40))
	//fmt.Println(myCalendarTwo.Book(5, 15))
	//fmt.Println(myCalendarTwo.Book(5, 10))
	//fmt.Println(myCalendarTwo.Book(25, 55))
	fmt.Println(myCalendarTwo.Book(26, 35))
	fmt.Println(myCalendarTwo.Book(26, 32))
	fmt.Println(myCalendarTwo.Book(25, 32))
	fmt.Println(myCalendarTwo.Book(18, 26))
}
