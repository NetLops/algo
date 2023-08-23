package main

import (
	"strconv"
	"strings"
)

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
type Codec struct {
}

func Constructor() (_ Codec) {
	return
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	s := &strings.Builder{}
	var dfs func(node *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			s.WriteString("NULL,")
			return
		}
		s.WriteString(strconv.Itoa(root.Val))
		s.WriteByte(',')
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return s.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	s := strings.Split(data, ",")
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if s[0] == "NULL" {
			s = s[1:]
			return nil
		}
		val, _ := strconv.Atoi(s[0])
		s = s[1:]
		return &TreeNode{val, dfs(), dfs()}
	}
	return dfs()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
