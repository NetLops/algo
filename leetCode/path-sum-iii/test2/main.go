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

func pathSum(root *TreeNode, targetSum int) (ans int) {
	// 前序遍历
	var dfsPrev func(root *TreeNode)
	var dfsVal func(root *TreeNode, val int)
	dfsPrev = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfsVal(root, root.Val)
		dfsPrev(root.Left)
		dfsPrev(root.Right)
	}
	dfsVal = func(root *TreeNode, val int) {
		if val == targetSum {
			ans++
		}
		if root.Left != nil {
			dfsVal(root.Left, val+root.Left.Val)
		}
		if root.Right != nil {
			dfsVal(root.Right, val+root.Right.Val)
		}
	}
	dfsPrev(root)
	return

}
