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

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	left, right := head, head

	for left != nil && right != nil && right.Next != nil && right.Next.Next != nil {
		left = left.Next
		right = right.Next.Next
		if left == right {
			return true
		}
	}
	return false
}
