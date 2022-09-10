package main

import (
	"strconv"
	"unsafe"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var dfs func(root *TreeNode) []byte

	m := map[string]int{}
	ans := []*TreeNode{}
	dfs = func(root *TreeNode) []byte {
		if root == nil {
			return []byte{}
		}
		str := []byte{}
		//str.WriteString(strconv.Itoa(root.Val))
		str = append(str, strconv.Itoa(root.Val)...)
		str = append(str, '_')
		//str.WriteString(dfs(root.Left))
		str = append(str, dfs(root.Left)...)
		//str.WriteString("_")
		//str.WriteString(dfs(root.Right))
		str = append(str, '_')
		str = append(str, dfs(root.Right)...)
		//s :=
		m[*(*string)(unsafe.Pointer(&str))]++
		if m[*(*string)(unsafe.Pointer(&str))] == 2 {
			ans = append(ans, root)
		}
		return str
	}
	dfs(root)
	return ans
}
func main() {

}
