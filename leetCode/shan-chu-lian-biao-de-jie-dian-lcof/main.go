package main

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

func deleteNode(head *ListNode, val int) *ListNode {
	pre := head
	// 看看是不是头部
	if pre.Val == val {
		pre = pre.Next
		return pre
	}
	//pre := ans
	m := pre.Next
	for m != nil {
		if m.Val == val {
			pre = m.Next
			return head
		}
		pre = m
		m = m.Next
	}
	return pre
}
func main() {

}
