package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//func nodeLen(head *Node) (length int) {
//	l := head
//	for l != nil {
//		length++
//		l = l.Next
//	}
//	length++
//	return
//}
//func copyRandomList(head *Node) *Node {
//	length := nodeLen(head)
//	l := head
//	//nodes := make([]int, length)
//	nodePtr := make(map[*Node]*Node, length)
//	copy := &Node{
//		Val:    l.Val,
//		Next:   nil,
//		Random: l.Random,
//	}
//	//if l.Random != nil {
//	//	nodes[0] = 0
//	//} else {
//	//	nodes[0] = math.MinInt
//	//}
//	nodePtr[l] = copy
//	i := 1
//	l = l.Next
//	prev := copy
//	for l != nil {
//		now := &Node{
//			Val:    l.Val,
//			Next:   nil,
//			Random: l.Random,
//		}
//		//if l.Random != nil {
//		//	nodes[i] = i
//		//} else {
//		//	nodes[i] = math.MinInt
//		//}
//		nodePtr[l] = now
//		prev.Next = now
//		prev = now
//		l = l.Next
//		i++
//	}
//	l = copy
//	for l != nil {
//		if l.Random != nil {
//			l.Random = nodePtr[l.Random]
//		}
//		l = l.Next
//	}
//	return copy
//}

//func copyRandomList(head *Node) *Node {
//	if head == nil {
//		return nil
//	}
//	for node := head; node != nil; node = node.Next.Next {
//		node.Next = &Node{Val: node.Val, Next: node.Next, Random: nil}
//	}
//
//	for node := head; node != nil; node = node.Next.Next {
//		if node.Random != nil {
//			node.Next.Random = node.Random.Next
//		}
//	}
//	copyHead := head.Next
//	for node := head; node != nil; node = node.Next {
//		copyValue := node.Next
//		node.Next = node.Next.Next
//		if copyValue.Next != nil {
//			copyValue.Next = copyHead.Next.Next
//		}
//	}
//	return copyHead
//}
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

//func copyRandomList(head *Node) *Node {
//	if head == nil {
//		return nil
//	}
//	for node := head; node != nil; node = node.Next.Next {
//		node.Next = &Node{Val: node.Val, Next: node.Next, Random: nil}
//	}
//	for node := head; node != nil; node = node.Next.Next {
//		if node.Random != nil {
//			node.Next.Random = node.Random.Next
//		}
//	}
//	copyHead := head.Next
//	for node := head; node != nil; node = node.Next {
//		copyValue := node.Next
//		node.Next = copyValue.Next
//		if node.Next.Next != nil {
//			copyValue.Next = node.Next.Next
//		}
//	}
//	return copyHead
//}

// 原地算法
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	m := head
	// 每隔一个串起来一个
	for m != nil {
		newNode := Node{
			Val:    m.Val,
			Next:   m.Next,
			Random: nil,
		}
		m.Next = &newNode
		m = newNode.Next
	}
	// 接下来就是random了	// 隔两个就是一个
	for m = head; m != nil; m = m.Next.Next {
		if m.Random != nil {
			m.Next.Random = m.Random.Next
		}
	}
	m = head
	copyNode := m.Next
	temp := copyNode
	// 接下来就是 分离
	for m != nil {
		// fmt.Println(m.Val)
		m.Next = m.Next.Next
		if m.Next != nil {
			temp.Next = m.Next.Next
		} else {
			temp.Next = nil
		}
		m = m.Next
		temp = temp.Next
	}
	//
	//m = head
	//for m != nil {
	//	fmt.Println(m.Val)
	//	if m.Random != nil {
	//		fmt.Println(m.Random.Val)
	//	} else {
	//		fmt.Println("nil")
	//	}
	//
	//	m = m.Next
	//}
	return copyNode
}
func main() {

}
