package main

import "math"

//func max(a, b int) int {
//	if a < b {
//		return b
//	}
//	return a
//}
//
//func maximizeXor(nums []int, queries [][]int) []int {
//	sort.Ints(nums)
//	minInt := nums[0]
//	ans := []int{}
//	for _, query := range queries {
//		x := query[0]
//		maxInt := -1
//		if query[1] < minInt {
//			continue
//		}
//		for i := 0; i < len(nums) && nums[i] <= query[1]; i++ {
//			maxInt = max(nums[i]|x, maxInt)
//		}
//		ans = append(ans, maxInt)
//	}
//	return ans
//}

//// 离线版本
//const L = 30
//
//type Trie struct {
//	children [2]*Trie
//}
//
//func (t *Trie) Insert(val int) {
//	root := t
//	for i := L - 1; i >= 0; i-- {
//		bit := val >> i & 1 // 取它第i位是0/1
//		if root.children[bit] == nil {
//			root.children[bit] = &Trie{}
//		}
//		root = root.children[bit]
//	}
//}
//func (t *Trie) Query(val int) (ans int) {
//	root := t
//	for i := L - 1; i >= 0; i-- {
//		bit := val >> i & 1
//		if root.children[bit^1] != nil {
//			ans |= 1 << i
//			bit ^= 1
//		}
//		root = root.children[bit]
//	}
//	return
//}
//
//func maximizeXor(nums []int, queries [][]int) []int {
//	for i := range queries {
//		queries[i] = append(queries[i], i) // 找到其索引的位置 并存起来
//	}
//	sort.Ints(nums)
//	sort.Slice(queries, func(i, j int) bool {
//		return queries[i][1] < queries[j][1]
//	})
//	t := &Trie{}
//	ans := make([]int, len(queries))
//	index := 0
//	for _, query := range queries {
//		x, m, i := query[0], query[1], query[2]
//		fmt.Println(index, nums[index], queries)
//		for j := index; j < len(nums) && nums[j] <= m; j++ {
//			t.Insert(nums[j])
//			index++
//		}
//		if index == 0 {
//			ans[i] = -1
//		} else {
//			ans[i] = t.Query(x)
//		}
//
//	}
//	return ans
//}

// 在线版本
const L = 30

type Trie struct {
	children [2]*Trie
	min      int
}

func (t *Trie) insert(val int) {
	root := t
	if val < root.min {
		root.min = val
	}
	for i := L - 1; i >= 0; i-- {
		bit := val >> i & 1
		if root.children[bit] == nil {
			root.children[bit] = &Trie{min: val}
		}
		root = root.children[bit]
		if val < root.min {
			root.min = val
		}
	}
}

func (t *Trie) query(val, limit int) (ans int) {
	root := t
	if root.min > limit {
		return -1
	}
	for i := L - 1; i >= 0; i-- {
		bit := val >> i & 1
		if root.children[bit^1] != nil && root.children[bit^1].min <= limit {
			ans |= 1 << i
			bit ^= 1
		}
		root = root.children[bit]
	}
	return
}
func maximizeXor(nums []int, queries [][]int) []int {
	t := &Trie{min: math.MaxInt32}
	for _, num := range nums {
		t.insert(num)
	}
	ans := make([]int, len(queries))
	for i, query := range queries {
		ans[i] = t.query(query[0], query[1])
	}
	return ans
}
func main() {
}
