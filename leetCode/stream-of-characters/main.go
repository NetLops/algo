package main

type Node struct {
	Val   bool
	Next  [26]*Node
	isEnd bool
}
type StreamChecker struct {
	root    *Node
	letters []byte
}

func Constructor(words []string) StreamChecker {
	root := &Node{
		Next: [26]*Node{},
	}
	rootTemp := root
	for _, word := range words {
		for i := len(word) - 1; i >= 0; i-- {
			index := word[i] - 'a'
			if rootTemp.Next[index] == nil {
				rootTemp.Next[index] = &Node{
					Next: [26]*Node{},
				}
			}
			rootTemp.Next[index].Val = true
			rootTemp = rootTemp.Next[index]
		}
		rootTemp.isEnd = true
		rootTemp = root
	}
	return StreamChecker{
		root:    root,
		letters: []byte{},
	}
}

func (this *StreamChecker) Query(letter byte) bool {
	this.letters = append(this.letters, letter)
	root := this.root
	for i := len(this.letters) - 1; i >= 0; i-- {
		index := this.letters[i] - 'a'
		if root.Next[index] == nil {
			return false
		}
		if root.Next[index].Val && root.Next[index].isEnd {
			return true
		}
		root = root.Next[index]
	}
	return false
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */
func main() {

}
