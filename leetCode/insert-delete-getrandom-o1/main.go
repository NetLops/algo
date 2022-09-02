package main

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	m map[int]int // key: val, value: index_pos
	p []int
}

func Constructor() RandomizedSet {
	rand.Seed(time.Now().Unix())
	return RandomizedSet{
		m: map[int]int{},
		p: []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.m[val] = len(this.p)
	this.p = append(this.p, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.m[val]; !ok {
		return false
	}
	last := len(this.p) - 1 // 最后一个数的位置,替补上
	// 以下两行是顶替
	this.p[this.m[val]] = this.p[this.m[last]]
	this.m[this.m[val]] = this.m[val]
	this.p = this.p[:last]
	delete(this.m, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.p[rand.Intn(len(this.p))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 *
 */
func main() {

}
