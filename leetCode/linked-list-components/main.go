package main

/**
 * Definition for singly-linked list.

 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, nums []int) int {
	indexs := make([]int, 10001)
	for _, num := range nums {
		indexs[num] = 1
	}
	ans, pre := 0, 0
	for head != nil {
		if pre == 0 && indexs[head.Val] == 1 {
			ans++
		}
		pre = indexs[head.Val]
		head = head.Next
	}
	return ans
}
