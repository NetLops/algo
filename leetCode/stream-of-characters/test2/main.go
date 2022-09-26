package main

type Node struct {
	Val   byte
	Next  []*Node
	isEnd bool
}

type StreamChecker struct {
	root    *Node
	letters []byte
}

func Constructor(words []string) StreamChecker {
	s := StreamChecker{root: &Node{Val: '/'}, letters: []byte{}}
	for _, word := range words {
		s.add(word)
	}
	return s
}

func (this *StreamChecker) add(word string) {
	root := this.root

	var temp *Node
	for i := len(word) - 1; i >= 0; i-- {
		if root.Next == nil {
			root.Next = []*Node{}
		}
		temp = nil
		for _, node := range root.Next {
			if node.Val == word[i] {
				temp = node
				break
			}
		}
		if temp == nil {
			temp = &Node{
				Val:  word[i],
				Next: nil,
			}
			root.Next = append(root.Next, temp)
		}
		root = temp
	}
	root.isEnd = true
}

func (this *StreamChecker) Query(letter byte) bool {
	this.letters = append(this.letters, letter)
	root := this.root
	var temp *Node
	for i := len(this.letters) - 1; i >= 0; i-- {
		if root.Next == nil {
			return false
		}
		temp = nil
		for _, node := range root.Next {
			if node.Val == this.letters[i] {
				temp = node
				break
			}
		}
		if temp == nil {
			break
		}
		root = temp
		if root.isEnd {
			return true
		}
	}

	return false
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */
