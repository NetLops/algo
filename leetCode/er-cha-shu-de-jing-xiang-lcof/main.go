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

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	ans := &TreeNode{
		Val:   root.Val,
		Left:  nil,
		Right: nil,
	}

	//temp := ans
	queue := []*TreeNode{root}
	reverseQueue := []*TreeNode{ans}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		q := reverseQueue[0]
		reverseQueue = reverseQueue[1:]
		if p.Right != nil {
			queue = append(queue, p.Right)
			q.Left = &TreeNode{
				Val:   p.Right.Val,
				Left:  nil,
				Right: nil,
			}
			reverseQueue = append(reverseQueue, q.Left)
		}
		if p.Left != nil {
			queue = append(queue, p.Left)
			q.Right = &TreeNode{
				Val:   p.Left.Val,
				Left:  nil,
				Right: nil,
			}
			reverseQueue = append(reverseQueue, q.Right)
		}
	}
	return ans
}
func main() {

}
