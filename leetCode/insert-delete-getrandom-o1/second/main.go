package main

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	m    map[int]int
	nums []int
}

func Constructor() RandomizedSet {
	rand.Seed(time.Now().UnixNano())
	return RandomizedSet{
		nums: []int{},
		m:    map[int]int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; !ok {
		this.m[val] = len(this.nums)
		this.nums = append(this.nums, val)
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if oldIndex, ok := this.m[val]; ok {
		oldKey := this.nums[len(this.nums)-1]
		this.m[oldKey] = oldIndex
		this.nums[oldIndex] = this.nums[len(this.nums)-1]
		this.nums = this.nums[:len(this.nums)-1]
		delete(this.m, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
func main() {
	//constructor := Constructor()
	//constructor.r.Intn(10)
}
