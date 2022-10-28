package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
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
	isLoop := false
	for fast.Next != nil && fast.Next.Next != nil && slow.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			isLoop = true
			break
		}
	}
	if !isLoop {
		return nil
	}
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
