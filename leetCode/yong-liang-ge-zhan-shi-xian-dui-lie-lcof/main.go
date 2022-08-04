package main

//type CQueue struct {
//	m      []int
//	length int
//}
//
//func Constructor() CQueue {
//	return CQueue{
//		m:      []int{},
//		length: 0,
//	}
//}
//
//func (this *CQueue) AppendTail(value int) {
//	this.m = append(this.m, value)
//	this.length++
//}
//
//func (this *CQueue) DeleteHead() int {
//	if this.length <= 0 {
//		return -1
//	}
//	num := this.m[0]
//	this.m = this.m[1:]
//	this.length--
//	return num
//}

type CQueue struct {
	InStack  []int
	OutStack []int
}

func Constructor() CQueue {
	return CQueue{
		InStack:  make([]int, 0, 5000),
		OutStack: make([]int, 0, 5000),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.InStack = append(this.InStack, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.OutStack) > 0 {
		ret := this.OutStack[len(this.OutStack)-1]
		this.OutStack = this.OutStack[:len(this.OutStack)-1]
		return ret
	} else {
		if len(this.InStack) == 0 {
			return -1
		} else {
			for i := len(this.InStack) - 1; i >= 0; i-- {
				this.OutStack = append(this.OutStack, this.InStack[i])
			}
			this.InStack = this.InStack[:0]
			ret := this.OutStack[len(this.OutStack)-1]
			this.OutStack = this.OutStack[:len(this.OutStack)-1]
			return ret
		}
	}
}
func main() {

}
