package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//
//func levelOrder(root *TreeNode) [][]int {
//	if root == nil {
//		return nil
//	}
//	ans := [][]int{{root.Val}}
//	queue := []*TreeNode{root}
//	layers := map[int]int{}
//	nowLayer := 0
//	layers[0] = 2
//	q := []int{}
//	reduce := 0
//	for len(queue) > 0 {
//		node := queue[0]
//		if node.Left != nil {
//			queue = append(queue, node.Left)
//			q = append(q, node.Left.Val)
//			reduce += 2
//			check(&layers, &nowLayer, &reduce, &q, &ans)
//		} else {
//			check(&layers, &nowLayer, &reduce, &q, &ans)
//		}
//		if node.Right != nil {
//			queue = append(queue, node.Right)
//			q = append(q, node.Right.Val)
//			reduce += 2
//			check(&layers, &nowLayer, &reduce, &q, &ans)
//		} else {
//			check(&layers, &nowLayer, &reduce, &q, &ans)
//		}
//		queue = queue[1:]
//	}
//	return ans
//}
//
//func check(layers *map[int]int, nowLayer *int, reduce *int, q *[]int, ans *[][]int) {
//	(*layers)[*nowLayer]--
//	if (*layers)[*nowLayer] == 0 {
//		(*nowLayer)++
//		(*layers)[*nowLayer] = *reduce
//		*reduce = 0
//		if len(*q) >= 1 {
//			*ans = append(*ans, *q)
//		}
//		*q = []int{}
//	}
//}
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := [][]int{{root.Val}}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		temp := []int{}
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
				temp = append(temp, node.Left.Val)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				temp = append(temp, node.Right.Val)
			}
		}
		if len(temp) > 0 {
			ans = append(ans, temp)
		}
	}
	return ans
}
func main() {

}
