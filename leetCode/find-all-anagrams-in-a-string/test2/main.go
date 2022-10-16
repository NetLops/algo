package main

import "fmt"

/**
 * Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

// 分治
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if lists == nil && n == 0 {
		return nil
	}
	return merge(lists, 0, n-1)
}

func merge(lists []*ListNode, left int, right int) *ListNode {
	if left == right {
		fmt.Println(left, right)
		return lists[left]
	}
	mid := left + (right-left)>>1
	l1 := merge(lists, left, mid)
	l2 := merge(lists, mid+1, right)
	return mergeTwoLists(l1, l2)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return nil
	}
	if l1.Val < l2.Val {
		l2.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}
