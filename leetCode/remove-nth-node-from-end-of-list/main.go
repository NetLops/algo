package main

import "runtime"

/**
Definition for singly-linked list.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	m := map[int]*ListNode{}
	headNode, next := head, head
	i := 0
	for next != nil {
		m[i] = next
		next = next.Next
		i++
	}
	n = i - n
	if n > 0 {
		m[n-1].Next = m[n].Next
	} else {
		headNode = m[n].Next
	}
	m[n].Next = nil
	delete(m, n)
	m = nil
	next = nil
	runtime.GC()
	return headNode
}
