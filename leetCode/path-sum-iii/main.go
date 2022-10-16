package main

/**
 * Definition for a binary tree node.

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	var dfsPrev func(root *TreeNode)
	var dfsRoot func(root *TreeNode, val int)
	ans := 0
	// 前序遍历
	dfsPrev = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfsRoot(root, root.Val)
		dfsPrev(root.Left)
		dfsPrev(root.Right)
	}
	dfsRoot = func(root *TreeNode, val int) {
		if val == targetSum {
			ans++
		}
		if root.Left != nil {
			dfsRoot(root.Left, val+root.Left.Val)
		}
		if root.Right != nil {
			dfsRoot(root.Right, val+root.Right.Val)
		}
	}
	dfsPrev(root)
	return ans

}
