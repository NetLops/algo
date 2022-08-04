package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := []int{root.Val}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		if node.Left != nil {
			queue = append(queue, node.Left)
			ans = append(ans, node.Left.Val)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			ans = append(ans, node.Right.Val)
		}
		queue = queue[1:]
	}
	return ans
}
func main() {

}
