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

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	m := head
	p := head.Next
	m.Next = nil
	for p != nil {
		q := p.Next
		p.Next = m
		m = p
		p = q
	}
	return m
}
func main() {

}
