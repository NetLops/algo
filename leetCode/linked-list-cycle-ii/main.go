package main

/**
 * Definition for singly-linked list.

 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	hasLoop := false
	for fast.Next != nil && fast.Next.Next != nil && slow.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			hasLoop = true
			break
		}
	}
	if hasLoop {
		fast = head
		for fast != slow {
			fast = fast.Next
			slow = slow.Next
		}
		return fast
	} else {
		return nil
	}
}
