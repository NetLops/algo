package main

import "fmt"

type Node struct {
	val   byte
	next  []*Node
	isEnd bool
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{
		root: &Node{
			val:  '/',
			next: nil,
		},
	}
}

func (this *Trie) Insert(word string) {
	root := this.root
	var temp *Node = nil
	for i := 0; i < len(word); i++ {
		if root.next == nil {
			root.next = []*Node{}
		}
		temp = nil
		for j := range root.next {
			if root.next[j].val == word[i] {
				temp = root.next[j]
				break
			}
		}
		if temp == nil {
			temp = &Node{
				val: word[i],
			}
			root.next = append(root.next, temp)

		}
		root = temp
	}

	root.isEnd = true
	fmt.Println(root)
}

func (this *Trie) Search(word string) bool {
	root := this.root

	var temp *Node = nil
	for i := range word {
		if root == nil {
			return false
		}
		temp = nil
		for j := range root.next {
			if root.next[j].val == word[i] {
				temp = root.next[j]
				break
			}
		}
		if temp == nil {
			return false
		}
		root = temp
	}
	if !root.isEnd {
		return false
	}

	return true
}

func (this *Trie) StartsWith(prefix string) bool {
	root := this.root

	var temp *Node = nil
	for i := range prefix {
		if root == nil {
			return false
		}
		temp = nil
		for j := range root.next {
			if root.next[j].val == prefix[i] {
				temp = root.next[j]
				break
			}
		}
		if temp == nil {
			return false
		}
		root = temp
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
func main() {

}
