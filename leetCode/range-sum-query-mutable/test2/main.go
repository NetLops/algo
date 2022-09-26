package main

import "fmt"

func lowBit(x int) int {
	// x & -x == x & (^x + 1)
	// 因为 (^x - 1) == -1
	// 1110 & 0010 = 10  ,1100 & 0100 = 100 (快速找到父节点)
	// 父 = 子 + lowBit(子)
	return x & -x
}

type NumArray struct {
	nums, collection []int
	n                int
}

func Constructor(nums []int) NumArray {
	numArray := NumArray{
		nums:       nums,
		collection: make([]int, len(nums)+1),
		n:          len(nums),
	}
	for i := range nums {
		numArray.Add(i+1, nums[i])
	}
	return numArray
}

func (this *NumArray) Add(x, y int) {
	for ; x <= this.n; x += lowBit(x) {
		this.collection[x] += y
	}
}

func (this *NumArray) Query(x int) int {
	ans := 0
	// 区间关系 假如c 是 this.collection
	// x = 15 = 1111 那他的区间和为
	// sum[x] = c[x] + c[1110] + c[1100] + c[1000]
	// sum[x] = c[15] + c[14] + c[12] + c[8]
	for ; x > 0; x -= lowBit(x) {
		ans += this.collection[x]
	}
	return ans
}

func (this *NumArray) Update(index int, val int) {
	this.Add(index+1, val-this.nums[index])
	this.nums[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.Query(right+1) - this.Query(left)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
func main() {
	fmt.Println()
}
