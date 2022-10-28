package main

/*
type ListNode struct{
  Val int
  Next *ListNode
}
*/

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param a ListNode类
 * @param b ListNode类
 * @return ListNode类
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func xorList(a *ListNode, b *ListNode) *ListNode {
	// write code here
	aArr := []int{}
	tArr := []int{}
	bArr := []int{}
	temp := a
	for temp != nil {
		aArr = append(aArr, temp.Val)
		temp = temp.Next
	}
	temp = b
	for temp != nil {
		tArr = append(aArr, temp.Val)
		temp = temp.Next
	}
	index := 0
	for i := len(tArr) - 1; i >= 0; i-- {
		bArr[index] = tArr[i]
		index++
	}
	mArr := make([]int, len(aArr))
	for i := len(aArr) - 1; i >= 0 && len(bArr)-i > 0; i-- {
		if aArr[i] == bArr[i] {
			mArr[i] = 1
		} else {
			mArr[i] = 0
		}
	}
	pre := false
	var l, tmp *ListNode
	for i := 0; i < len(mArr); i++ {
		if !pre {
			if mArr[i] != 0 {
				pre = true
				l = &ListNode{Val: mArr[i]}
				tmp = l
			}
		} else {
			tmp.Next = &ListNode{Val: mArr[i]}
		}
	}
	return l

}
