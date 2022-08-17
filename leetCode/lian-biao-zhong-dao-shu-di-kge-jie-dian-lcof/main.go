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

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	i := 1
	low, fast := head, head
	for fast.Next != nil {
		if i%k == 0 {
			low = low.Next
		} else {
			i++
		}
		fast = fast.Next
	}
	if i%k != 0 {
		return nil
	}
	return low
}
func main() {

}
