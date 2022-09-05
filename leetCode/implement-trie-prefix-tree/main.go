package main

import "fmt"

type Node struct {
	Val   byte
	Next  []*Node
	isEnd bool
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{
		root: &Node{
			Val:  '/', // 占一个位置
			Next: nil,
		},
	}
}

func (this *Trie) Insert(word string) {
	root := this.root
	//if root == nil {
	//	root := []*Node{}
	//}
	var temp *Node
	for _, w := range word {
		if root == nil {
			root.Next = []*Node{}
		}

		temp = nil
		for _, node := range root.Next {
			if node.Val == byte(w) {
				temp = node
				break
			}
		}
		if temp == nil {
			temp = &Node{
				Val: byte(w),
			}
			root.Next = append(root.Next, temp)
		}
		root = temp
	}
	root.isEnd = true
}

func (this *Trie) Search(word string) bool {
	root := this.root
	var temp *Node
	for _, w := range word {
		if root.Next == nil {
			return false
		}
		temp = nil
		for _, node := range root.Next {
			if node.Val == byte(w) {
				temp = node
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
	var temp *Node
	for _, w := range prefix {
		if root.Next == nil {
			return false
		}
		temp = nil
		for _, node := range root.Next {
			if node.Val == byte(w) {
				temp = node
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
	obj := Constructor()
	obj.Insert("app")
	obj.Insert("apple")
	obj.Insert("beer")
	obj.Insert("beer")
	obj.Insert("add")
	obj.Insert("jam")
	obj.Insert("rental")
	fmt.Println(obj.Search("apps"))
	//fmt.Println(obj.Search("zhangshuns"))
	//fmt.Println(obj.Search("t"))
	//fmt.Println(obj.StartsWith("the"))
	//fmt.Println(obj)
}

/*
["Trie","insert","insert","insert","insert","insert","insert","search","search","search","search","search","search","search","search","search","startsWith","startsWith","startsWith","startsWith","startsWith","startsWith","startsWith","startsWith","startsWith"]
[[],["app"],["apple"],["beer"],["add"],["jam"],["rental"],["apps"],["app"],["ad"],["applepie"],["rest"],["jan"],["rent"],["beer"],["jam"],["apps"],["app"],["ad"],["applepie"],["rest"],["jan"],["rent"],["beer"],["jam"]]
*/
