package main

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	var minStack MinStack
	return minStack
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, val)
	} else {
		min := this.minStack[len(this.minStack)-1]
		if min > val {
			min = val
		}
		this.minStack = append(this.minStack, min)
	}
}

func (this *MinStack) Pop() {
	n := len(this.stack)
	this.stack = this.stack[:n-1]
	this.minStack = this.minStack[:n-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.stack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
