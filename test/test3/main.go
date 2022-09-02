package main

/**
Definition for a binary tree node.
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

//
//func isBalanced(root *TreeNode) bool {
//	if root == nil {
//		return false
//	}
//	return abs(depth(root.Left)-depth(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
//}
//
////  测试深度
//func depth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	return max(depth(root.Left), depth(root.Right)) + 1
//}
//
//func max(a, b int) int {
//	if a < b {
//		return b
//	}
//	return a
//
//}
//func abs(a int) int {
//	if a < 0 {
//		return -a
//	}
//	return a
//
//}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a

}

func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftH := height(root.Left)
	rightH := height(root.Right)
	if leftH == -1 {
		return -1
	}
	if rightH == -1 {
		return -1
	}
	if abs(leftH-rightH) <= 1 {
		return max(leftH, rightH) + 1
	} else {
		return -1
	}
}

func main() {

}
