package main

type StockSpanner struct {
	deque []Stock
}

type Stock struct {
	price, day int
}

func Constructor() StockSpanner {
	return StockSpanner{
		deque: []Stock{},
	}
}

func (this *StockSpanner) Next(price int) int {
	day := 1
	for len(this.deque) != 0 && this.deque[len(this.deque)-1].price <= price {
		day += this.deque[len(this.deque)-1].day
		this.deque = this.deque[:len(this.deque)-1]
	}
	this.deque = append(this.deque, Stock{price: price, day: day})
	return day
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */

func largestRectangleArea(heights []int) (res int) {
	// 参考：
	stack := []int{-1}
	heights = append(heights, 0)
	for i := 0; i < len(heights); i++ {
		// 哨兵
		for len(stack) > 1 && heights[stack[len(stack)-1]] > heights[i] {
			current := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			res = max((i-left-1)*heights[current], res)
		}
		stack = append(stack, i)
	}
	return
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
