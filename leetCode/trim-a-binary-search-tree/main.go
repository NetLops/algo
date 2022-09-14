package main

/**
 * Definition for a binary tree node.

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	var dfs func(root *TreeNode) *TreeNode

	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		//if low <= root.Val && root.Val <= high {
		//	return root
		//} else {
		//
		//}
		if root.Val < low {
			return dfs(root.Right)
		} else if root.Val > high {
			//} else {
			return dfs(root.Left)
		}
		root.Left = dfs(root.Left)
		root.Right = dfs(root.Right)

		return root

	}
	dfs(root)
	return root
}
func main() {

}
