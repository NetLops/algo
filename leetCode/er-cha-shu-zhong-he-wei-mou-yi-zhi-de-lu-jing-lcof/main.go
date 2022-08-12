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

func pathSum(root *TreeNode, target int) [][]int {
	ans := [][]int{}

	temp := []int{root.Val}
	dfs(root.Left, target, temp, &ans)
	dfs(root.Right, target, temp, &ans)
	return ans
}

func dfs(root *TreeNode, target int, temp []int, ans *[][]int) {

	if root == nil {
		tempNow := make([]int, len(temp))
		//copy(temp, tempNow)
		copy(tempNow, temp)
		q := 0
		for i := range tempNow {
			q += temp[i]
		}
		if q == target {
			(*ans) = append((*ans), tempNow)
		}
	} else {
		temp = append(temp, root.Val)
		dfs(root.Left, target, temp, ans)
		dfs(root.Right, target, temp, ans)
	}
}

func main() {

}
