package main

//
//type MinStack struct {
//	stack  []int
//	length int
//	min    int
//}
//
///** initialize your data structure here. */
//func Constructor() MinStack {
//	return MinStack{
//		make([]int, 0, 5000),
//		0,
//		math.MaxInt,
//	}
//}
//
//func (this *MinStack) Push(x int) {
//
//	this.stack = append(this.stack, x)
//	if x < this.min {
//		this.min = x
//	}
//	this.length++
//}
//
//func (this *MinStack) Pop() {
//	ret := this.stack[len(this.stack)-1]
//	this.stack = this.stack[:len(this.stack)-1]
//	this.length--
//	if this.min == ret {
//		this.min = math.MinInt8
//		for i := range this.stack {
//			fmt.Println(this.min, this.stack[i])
//			if this.min > this.stack[i] {
//				this.min = this.stack[i]
//			}
//		}
//	}
//}
//
//func (this *MinStack) Top() int {
//	ret := this.stack[len(this.stack)-1]
//	return ret
//}
//
//func (this *MinStack) Min() int {
//	return this.min
//}

type MinStack struct {
	MinData []int
	Data    []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		[]int{},
		[]int{},
	}
}

func (this *MinStack) Push(x int) {
	this.Data = append(this.Data, x)
	min := x
	if len(this.MinData) > 0 {
		if min > this.MinData[len(this.MinData)-1] {
			min = this.MinData[len(this.MinData)-1]
		}
	}
	this.MinData = append(this.MinData, min)
}

func (this *MinStack) Pop() {
	if len(this.Data) < 0 {
		return
	}
	this.Data = this.Data[:len(this.Data)-1]
	this.MinData = this.MinData[:len(this.MinData)-1]
}

func (this *MinStack) Top() int {
	if len(this.Data) > 0 {
		return this.Data[len(this.Data)-1]
	}
	return -1
}

func (this *MinStack) Min() int {
	if len(this.MinData) > 0 {
		return this.MinData[len(this.MinData)-1]
	}
	return -1
}

func main() {

}
