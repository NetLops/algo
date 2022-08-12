package main

import "fmt"

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

// 伴生子树
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	// 中间匹配/ b 在左子树 / b 右子树
	return recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func recur(A, B *TreeNode) bool {
	// B 为nil 说明支部匹配完成
	if B == nil {
		return true
	}
	// 超过A树的范围 / 不等的情况 返回 false
	if A == nil || A.Val != B.Val {
		return false
	}
	// 继续先序遍历
	return recur(A.Left, B.Left) && recur(A.Right, B.Right)
}
func main() {
	s := []int{1}
	fmt.Println(s[1:])
}
