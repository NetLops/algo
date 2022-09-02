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

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func longestUnivaluePath(root *TreeNode) int {
	var dfs func(root *TreeNode) int
	ans := 0
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		i, j := 0, 0 // 如果 没左/右节点 那 左/右连续 就连不上
		if root.Left != nil && root.Val == root.Left.Val {
			i = left + 1
		}
		if root.Right != nil && root.Val == root.Right.Val {
			j = right + 1
		}
		ans = maxInt(ans, i+j)
		return maxInt(i, j)
	}
	dfs(root)
	return ans
}
func main() {

}
