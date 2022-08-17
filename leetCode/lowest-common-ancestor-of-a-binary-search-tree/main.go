package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	comp := []*TreeNode{root}
	comq := []*TreeNode{root}
	commonDad := root
	//queue := []*TreeNode{}
	//for len(queue) != 0 {
	//	temp := queue[0]
	//	queue = queue[1:]
	//	if temp.Left != nil {
	//		if temp.Left.Val == p.Val {
	//
	//		}
	//		queue = append(queue, temp.Left)
	//	}
	//	if temp.Right != nil {
	//		queue = append(queue, temp.Right)
	//	}
	//
	//}
	var dfs func(root, v *TreeNode, comp *[]*TreeNode)
	dfs = func(root, v *TreeNode, com *[]*TreeNode) {
		*com = append(*com, root)
		if root == nil {
			*com = []*TreeNode{}
			return
		}
		if v.Val == root.Val {
			return
		} else if v.Val < root.Val {
			dfs(root.Left, v, com)
		} else if v.Val > root.Val {
			dfs(root.Right, v, com)
		}
	}
	dfs(root, p, &comp)
	dfs(root, q, &comq)
	min := 0
	if len(comq) < len(comp) {
		min = len(comq)
	} else {
		min = len(comp)
	}
	for i := 0; i < min; i++ {
		if comp[i] == comq[i] {
			commonDad = comp[i]
		} else {
			break
		}

	}
	return commonDad

}
func main() {

}
