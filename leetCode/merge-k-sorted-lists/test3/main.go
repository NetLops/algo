package main

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

func mergeKLists(lsts []*ListNode) *ListNode {
	if lsts == nil || len(lsts) == 0 {
		return nil
	}
	pre := lsts[0]
	//temp := s
	for i := 0; i < len(lsts); i++ {
		pre = merge(pre, lsts[i])
	}
	return pre
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		l1.Next = merge(l1.Next, l2)
		return l1
	}
	l2.Next = merge(l2.Next, l1)
	return l2
}
