package main

import "fmt"

//func finalPrices(prices []int) []int {
//	ans := make([]int, len(prices))
//	for i := range prices {
//		if i != len(prices)-1 {
//			j := i + 1
//			for j != len(prices) {
//				if prices[j] <= prices[i] {
//					ans[i] = prices[i] - prices[j]
//					break
//				}
//				j++
//			}
//			if j == len(prices) {
//				ans[i] = prices[i]
//			}
//
//		} else {
//			ans[i] = prices[i]
//		}
//	}
//	return ans
//}
type Deque []int

func (d *Deque) peekLast() int {
	if !d.isEmpty() {
		return (*d)[len(*d)-1]
	}
	return -1
}

func (d *Deque) isEmpty() bool {
	return len(*d) == 0
}

func (d *Deque) pollLast() int {
	if !d.isEmpty() {
		res := (*d)[len(*d)-1]
		*d = (*d)[:len(*d)-1]
		return res
	}
	return -1
}

func (d *Deque) addLast(val int) {
	(*d) = append((*d), val)
}

func (d *Deque) String() string {
	return fmt.Sprintf("%v\n", *d)
}

func finalPrices(prices []int) []int {
	ans := make([]int, len(prices))
	deque := Deque{} // 记录索引

	for i := 0; i < len(prices); i++ {
		// 判断是否为空，队尾元素是否有第一个大于的值出现
		for !deque.isEmpty() && prices[deque.peekLast()] >= prices[i] {
			last := deque.pollLast()
			ans[last] = prices[last] - prices[i]
		}
		ans[i] = prices[i]
		deque.addLast(i)
	}
	return ans
}
func main() {
	fmt.Println(finalPrices([]int{8, 4, 6, 2, 3}))
	fmt.Println(finalPrices([]int{1, 2, 3, 4, 5}))
	fmt.Println(finalPrices([]int{10, 1, 1, 6}))
}
