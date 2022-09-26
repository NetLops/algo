package main

func lowBit(x int) int {
	return x & -x // x & (^x + 1) = 1110 & 0010 = 10 = 2s
}

type NumArray struct {
	nums, collection []int
	n                int // 长度
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
	i := x
	for ; i <= this.n; i += lowBit(i) {
		this.collection[i] += y
	}
}

func (this *NumArray) Query(x int) int {
	i := x
	ans := 0
	for ; i > 0; i -= lowBit(i) {
		ans += this.collection[i]
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

}
