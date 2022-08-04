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
type CBTInserter struct {
	root     *TreeNode
	candiate []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	node := []*TreeNode{root}
	candiate := []*TreeNode{}
	// BFS
	for len(node) > 0 {
		if node[0].Left != nil {
			node = append(node, node[0].Left)
		}
		if node[0].Right != nil {
			node = append(node, node[0].Right)
		}
		if node[0].Left == nil || node[0].Right == nil { // 判断缺左右结点的进组
			candiate = append(candiate, node[0])
		}
		node = node[1:]
	}
	return CBTInserter{root: root, candiate: candiate}
}

func (this *CBTInserter) Insert(val int) int {
	node := &TreeNode{Val: val, Left: nil, Right: nil}
	pre := this.candiate[0]
	if pre.Left == nil {
		pre.Left = node
	} else {
		pre.Right = node
		this.candiate = this.candiate[1:]
	}
	this.candiate = append(this.candiate, node)
	return pre.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}
func main() {

}
