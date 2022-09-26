package main

type Node struct {
	prev, next *Node
	val        int
}

type MyLinkedList struct {
	head, tail *Node
	n          int // 存的链表大小
}

func Constructor() MyLinkedList {
	head := &Node{
		val: -1,
	}
	tail := &Node{
		val: -1,
	}
	head.next = tail
	tail.prev = head
	return MyLinkedList{
		head: head,
		tail: tail,
		n:    0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	i := -1
	temp := this.head
	for i < this.n && i < index {
		temp = temp.next
		i++
	}
	if i == index {
		return temp.val
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{
		val: val,
	}
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
	this.n++
	//this.print("head: ")
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{
		val: val,
	}
	node.next = this.tail
	node.prev = this.tail.prev
	this.tail.prev.next = node
	this.tail.prev = node
	this.n++
	//this.print("tail: ")
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	//this.print("atIndex before: ")
	if index > this.n {
		//this.print("atIndex after > n: ")
		return
	}
	if index < 0 {
		this.AddAtHead(val)
		//this.print("atIndex after < n: ")
		return
	}
	if index == this.n {
		this.AddAtTail(val)
		//this.print("atIndex after = n: ")
		return
	}
	i := -1
	temp := this.head
	for i < this.n && i < index {
		temp = temp.next
		i++
	}
	node := &Node{
		val: val,
	}

	if i == index && temp != this.tail {
		node.next = temp
		node.prev = temp.prev
		temp.prev.next = node
		temp.prev = node
		this.n++
	}
	//this.print("atIndex after: ")
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	i := -1
	temp := this.head
	for i < this.n && i < index {
		temp = temp.next
		i++
	}
	if i == index && temp != this.tail {
		temp.prev.next = temp.next
		temp.next.prev = temp.prev
		temp.next = nil
		temp.prev = nil
		temp = nil
		this.n--
	}
	//this.print("atIndex after: ")
}

//func (this *MyLinkedList) print(str string) {
//	fmt.Print(str)
//	temp := this.head.next
//	for temp != this.tail {
//		if temp.next != this.tail {
//			fmt.Print(temp.val, "=>")
//		} else {
//			fmt.Print(temp.val)
//		}
//		temp = temp.next
//	}
//	fmt.Println(",", this.n)
//}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
func main() {

}
