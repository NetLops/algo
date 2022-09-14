package main

import "math"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type MinStack struct {
	stack []int
	min   []int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		min:   []int{math.MaxInt},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	this.min = append(this.min, min(this.min[len(this.min)-1], val))
}

func (this *MinStack) Pop() {
	n := len(this.stack)
	this.stack = this.stack[:n-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
func main() {

}
