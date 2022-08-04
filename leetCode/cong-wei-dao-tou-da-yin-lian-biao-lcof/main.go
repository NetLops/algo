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

func reversePrint(head *ListNode) []int {
	res := []int{}
	m := head
	for m != nil {
		res = append(res, m.Val)
		m = m.Next
	}
	length := len(res)
	for i := 0; i < length; i++ {
		if i < length-1-i {
			break
		}
		res[i], res[length-1-i] = res[length-1-i], res[i]
	}
	return res
}
func main() {

}
