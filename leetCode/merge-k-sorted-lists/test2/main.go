package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var pre, current *ListNode
	n := len(lists)
	if n < 1 {
		return nil
	}
	pre = lists[0]
	for i := 1; i < n; i++ {

		current = lists[i]
		pre = merge(pre, current)
	}
	return pre
}

func merge(left, right *ListNode) *ListNode {
	head := &ListNode{}
	current := head
	for left != nil || right != nil {
		if left != nil && right != nil {
			if left.Val < right.Val {
				current.Next = left
				left = left.Next
			} else {
				current.Next = right
				right = right.Next
			}
			current = current.Next
		} else if left != nil {
			current.Next = left
			break
		} else {
			current.Next = right
			break
		}
	}
	return head.Next
}
