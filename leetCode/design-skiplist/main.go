package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	MAX_LEVEL  = 10
	SKIPLIST_P = 0.25
)

type Skiplist struct {
	head       *Node
	levelCount int
}
type Node struct {
	Val int
	//level    int // 层数高度
	forwards [MAX_LEVEL]*Node
	counts   int // 出现次数
}

//
func Constructor() Skiplist {
	rand.Seed(time.Now().UnixNano())
	return Skiplist{
		head: &Node{
			Val: math.MaxInt,
		},
		levelCount: 1,
	}
}

// 搜索
func (this *Skiplist) Search(target int) bool {
	p := this.head
	for i := this.levelCount - 1; i >= 0; i-- { // 往下找
		for p.forwards[i] != nil && p.forwards[i].Val < target { //
			p = p.forwards[i]
		}
		if p.forwards[i] != nil && p.forwards[i].Val == target {
			return true
		}
	}

	return false
}
func (this *Skiplist) Add(num int) {
	p := this.head
	// 先随机 看看 自己在哪儿层
	randomLevel := this.randomLevel()
	// 查找更新的层
	updates := make([]*Node, randomLevel)
	for i := 0; i < randomLevel; i++ {
		updates[i] = this.head
	}
	// 找找自己在哪儿插入合适
	for i := randomLevel - 1; i >= 0; i-- { // 往下找
		for p.forwards[i] != nil && p.forwards[i].Val < num { // 往右找
			p = p.forwards[i]
		}
		if p.forwards[i] != nil && p.forwards[i].Val == num { // 重复就直接跳出去
			p.forwards[i].counts++
			return
		}
		updates[i] = p
	}

	newNode := &Node{
		Val:    num,
		counts: 1,
	}
	for i := 0; i < randomLevel; i++ {

		newNode.forwards[i] = updates[i].forwards[i]
		updates[i].forwards[i] = newNode
	}

	if this.levelCount < randomLevel {
		this.levelCount = randomLevel
	}

}

//
func (this *Skiplist) Erase(num int) bool {
	//updates := make([]*Node, this.levelCount)
	flge := false
	p := this.head
	for i := this.levelCount - 1; i >= 0; i-- {
		for p.forwards[i] != nil && p.forwards[i].Val < num {
			p = p.forwards[i]
		}
		if p.forwards[i] != nil && p.forwards[i].Val == num {
			flge = true
			if p.forwards[i].counts > 1 {
				p.forwards[i].counts--
				break
			}
			p.forwards[i] = p.forwards[i].forwards[i]
		}
	}
	//if p.forwards[0] != nil && p.forwards[0].Val == num {
	//	if p.forwards[0].counts > 1 {
	//		p.forwards[0].counts--
	//		return true
	//	}
	//	for i := this.levelCount - 1; i >= 0; i-- {
	//		if updates[i].forwards[i] != nil && updates[i].forwards[i].Val == num {
	//			updates[i].forwards[i] = updates[i].forwards[i].forwards[i]
	//		}
	//	}
	//} else {
	//	return false
	//}
	for this.levelCount > 1 && this.head.forwards[this.levelCount-1] == nil { // 有可能连续两层只有一个
		this.levelCount--
	}
	return flge
}
func (this *Skiplist) randomLevel() int {
	level := 1
	// 当level < MAX_LEVEL，且随机数小于设定的晋升概率时，level+1

	for rand.Float32() < SKIPLIST_P && level < MAX_LEVEL {
		//fmt.Println(rand.Float32())
		level += 1
	}
	return level
}
func main() {
	skiplist := Constructor()
	//m := map[int]int{}
	//for i := 0; i < 1000000; i++ {
	//	m[skiplist.randomLevel()]++
	//
	//}
	//fmt.Println(m)
	skiplist.Add(16)
	skiplist.Add(5)
	skiplist.Add(14)
	skiplist.Add(13)
	skiplist.Add(0)
	skiplist.Add(3)
	skiplist.Add(12)
	skiplist.Add(9)
	skiplist.Add(12)
	fmt.Println(skiplist.Erase(3))
	fmt.Println(skiplist.Search(6))
	skiplist.Add(7)
	fmt.Println(skiplist.Erase(0))
	fmt.Println(skiplist.Erase(1))
	fmt.Println(skiplist.Erase(10))
	skiplist.Add(5)
	fmt.Println(skiplist.Search(12))
	fmt.Println(skiplist.Search(7))
	fmt.Println(skiplist.Search(16))
	fmt.Println(skiplist.Erase(7))
	fmt.Println(skiplist.Search(0))
	skiplist.Add(9)
	skiplist.Add(16)
	skiplist.Add(3)
	fmt.Println(skiplist.Erase(2))
	fmt.Println(skiplist.Search(17))
	skiplist.Add(2)
	fmt.Println(skiplist.Search(17))
	fmt.Println(skiplist.Erase(0))
	fmt.Println(skiplist.Search(9))
	fmt.Println(skiplist.Search(14))
	fmt.Println(skiplist.Erase(1))
	fmt.Println(skiplist.Erase(6))
	skiplist.Add(1)
	fmt.Println(skiplist.Erase(16))
	fmt.Println(skiplist.Search(9))
	fmt.Println(skiplist.Erase(10))
	fmt.Println(skiplist.Erase(9))
	fmt.Println(skiplist.Search(2))
	skiplist.Add(3)
	skiplist.Add(16)
	fmt.Println(skiplist.Erase(15))
	fmt.Println(skiplist.Erase(12))
	fmt.Println(skiplist.Erase(7))
	skiplist.Add(4)
	fmt.Println(skiplist.Erase(3))
	skiplist.Add(2)
	fmt.Println(skiplist.Erase(1))
	fmt.Println(skiplist.Erase(14))
	skiplist.Add(13)
	skiplist.Add(12)
	skiplist.Add(3)
	fmt.Println(skiplist.Search(6))
	fmt.Println(skiplist.Search(17))
	skiplist.Add(2)
	fmt.Println(skiplist.Erase(3))
	fmt.Println(skiplist.Search(14))
	skiplist.Add(11)
	skiplist.Add(0)
	fmt.Println(skiplist.Search(13))
	skiplist.Add(2)
	fmt.Println(skiplist.Search(1))
	fmt.Println(skiplist.Erase(10))
	fmt.Println(skiplist.Erase(17))
	fmt.Println(skiplist.Search(0))
	fmt.Println(skiplist.Search(5))
	fmt.Println(skiplist.Erase(8))
	fmt.Println(skiplist.Search(9))
	skiplist.Add(8)
	fmt.Println(skiplist.Erase(11))
	fmt.Println(skiplist.Search(10))
	fmt.Println(skiplist.Erase(11))
	fmt.Println(skiplist.Search(10))
	fmt.Println(skiplist.Erase(9))
	fmt.Println(skiplist.Erase(8))
	fmt.Println(skiplist.Search(15))
	fmt.Println(skiplist.Search(14))
	skiplist.Add(1)
	skiplist.Add(6)
	skiplist.Add(17)
	skiplist.Add(16)
	fmt.Println(skiplist.Search(13))
	fmt.Println(skiplist.Search(4))
	fmt.Println(skiplist.Search(5))
	fmt.Println(skiplist.Search(4))
	fmt.Println(skiplist.Search(17))
	fmt.Println(skiplist.Search(16))
	fmt.Println(skiplist.Search(7))
	fmt.Println(skiplist.Search(14))
	fmt.Println(skiplist.Search(1))

}
