package main

import "fmt"

/**
Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	fmt.Println(pre)
	for node := head; node != nil; {
		temp := node.Next
		node.Next = pre
		pre = node
		node = temp
	}
	return pre
}
func main() {
	reverseList(nil)
}
