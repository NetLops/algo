package main

func lowBit(x int) int {
	return x & -x
}

func (this *NumArray) query(x int) int {
	ans := 0
	for i := x; i > 0; {
		ans += this.tree[i]
		i -= lowBit(i)
	}
	return ans
}

func (this *NumArray) add(x, y int) {
	for i := x; i <= this.length; {
		this.tree[i] += y
		i += lowBit(i)
	}
}

type NumArray struct {
	nums   []int
	tree   []int
	length int
}

func Constructor(nums []int) NumArray {
	na := NumArray{
		nums:   nums,
		length: len(nums),
		tree:   make([]int, len(nums)+1),
	}
	for i := 0; i < len(nums); i++ {
		na.add(i+1, nums[i])
	}
	return na
}

func (this *NumArray) Update(index int, val int) {
	this.add(index+1, val-this.nums[index])
	this.nums[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.query(right+1) - this.query(left)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 *
 */
func main() {

}
