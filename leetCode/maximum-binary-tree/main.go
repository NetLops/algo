package main

/**
* Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	max := 0

	for i, num := range nums {
		if max < num {
			max = i
		}
	}
	return &TreeNode{Val: nums[max], Left: constructMaximumBinaryTree(nums[:max]), Right: constructMaximumBinaryTree(nums[max+1:])}
}

func main() {

}
