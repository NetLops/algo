package main

import "strconv"

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool {
	return false
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int {
	return 0
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {
	return
}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger {
	return nil
}
func deserialize(s string) *NestedInteger {
	// 如果第一个位置不是 '[' 就直接暴力返回了，搞个锤子潜逃
	if s[0] != '[' {
		num, _ := strconv.Atoi(s)
		n := &NestedInteger{}
		n.SetInteger(num)
		return n
	}
	// sign 符号位
	stack, cur, sign := []*NestedInteger{}, 0, 1
	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			stack = append(stack, &NestedInteger{})
		} else if s[i] == ',' || s[i] == ']' {
			if s[i-1] >= '0' && s[i-1] <= '9' {
				n := &NestedInteger{}
				n.SetInteger(cur * sign)
				stack[len(stack)-1].Add(*n)
				// 重置一下
				cur, sign = 0, 1
			}
			if l := len(stack); s[i] == ']' && l > 1 {
				stack[l-2].Add(*stack[l-1])
				stack = stack[:l-1]
			}
		} else if s[i] == '-' { // 符号位该改变
			sign = -1
		} else {
			cur = 10*cur + int(s[i]-'0')
		}
	}
	return stack[0]
}
func main() {

}
