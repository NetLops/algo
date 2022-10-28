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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 剩掉 等于0的情况
	ans := &ListNode{0, head}
	pre := ans
	for i := 0; i < n; i++ {
		head = head.Next
	}
	for head != nil {
		pre = pre.Next
		head = head.Next
	}
	pre.Next = pre.Next.Next
	return ans.Next
}
