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

func reverse(ints *[]int) *[]int {
	for i := 0; i < len(*ints)/2; i++ {
		(*ints)[i], (*ints)[len(*ints)-1-i] = (*ints)[len(*ints)-1-i], (*ints)[i]
	}
	return ints
}

func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	deep := 2 //  记一个层数

	queue := []*TreeNode{root}
	ans = append(ans, []int{root.Val})
	for len(queue) != 0 {
		size := len(queue)
		temp := queue
		queue = []*TreeNode{}
		ansLayer := []int{}
		for size > 0 {
			q := temp[0]
			if size > 1 {
				temp = temp[1:]
			}
			if q.Left != nil {
				queue = append(queue, q.Left)
				ansLayer = append(ansLayer, q.Left.Val)
			}

			if q.Right != nil {
				queue = append(queue, q.Right)
				ansLayer = append(ansLayer, q.Right.Val)
			}

			size--
		}

		if len(ansLayer) != 0 {
			if deep%2 == 0 {
				reverse(&ansLayer)
			}
			ans = append(ans, ansLayer)
		}
		deep++
	}
	return ans
}
func main() {

}
