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

func addOneRow(root *TreeNode, val, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{val, root, nil}
	}
	q := []*TreeNode{root}
	for i := 1; i < depth-1; i++ {
		tmp := q
		q = nil

		for _, node := range tmp {
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	for _, node := range q {
		node.Left = &TreeNode{
			Val:   val,
			Left:  node.Left,
			Right: nil,
		}
		node.Right = &TreeNode{
			Val:   val,
			Right: node.Right,
			Left:  nil,
		}
	}
	return root
}
func main() {

}
