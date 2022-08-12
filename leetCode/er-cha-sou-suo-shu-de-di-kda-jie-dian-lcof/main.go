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

func kthLargest(root *TreeNode, k int) int {
	ans := 0
	if root == nil {
		return ans
	}
	isStop := false
	dfs(root, 0, k, &ans, &isStop)
	return ans
}

func dfs(root *TreeNode, now, k int, ans *int, isStop *bool) {
	if *isStop {
		return
	}
	if root == nil {
		return
	}
	dfs(root.Right, now, k, ans, isStop)
	now++
	if k == now {
		*isStop = true
		*ans = root.Val
	}
	dfs(root.Left, now, k, ans, isStop)
}

func main() {

}
