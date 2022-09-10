package main

type Node struct {
	key, val  int
	pre, next *Node
}

type LRUCache struct {
	m        map[int]*Node
	head     *Node
	tail     *Node
	capacity int
}

func (this *LRUCache) moveToTail(node *Node) {
	if this.capacity == 0 {
		return
	}
	isExists := false
	// 判断一手是否存在
	if _, isExists = this.m[node.key]; isExists {
		// 存在的话 移花接木

		if node.next != nil {
			node.pre.next = node.next
			node.next.pre = node.pre
		} else {
			// 那就说明它是最后一个节点
			this.tail = node.pre
			this.tail.next = nil
		}
		node.pre = nil
		node.next = nil
		delete(this.m, node.key)
	}
	// 判断一手容量是否够
	nowCapacity := len(this.m)
	if isExists {
		nowCapacity -= 1
	}
	// fmt.Println(nowCapacity,this.capacity)
	if nowCapacity >= this.capacity { // 触发删除机制
		delNode := this.head.next
		if delNode != nil {
			if delNode.next != nil {
				delNode.next.pre = this.head
			}
			this.head.next = delNode.next
			delNode.pre = nil
			delNode.next = nil
			delete(this.m, delNode.key)
			// 有可能见底了
			if this.head.next == nil {
				this.tail = this.head
			}
		}
	}
	//fmt.Println(this.head, this.tail)
	node.pre = this.tail
	node.next = nil
	this.tail.next = node
	this.tail = this.tail.next
	this.m[node.key] = node
	//// 有可能 this.head 被删完了
	//if this.head.next == nil {
	//	this.head.next = this.tail
	//}

}

func Constructor(capacity int) LRUCache {
	node := &Node{
		key:  -1,
		val:  -1,
		pre:  nil,
		next: nil,
	}
	return LRUCache{
		m:        map[int]*Node{},
		capacity: capacity,
		head:     node,
		tail:     node,
	}
}

func (this *LRUCache) Get(key int) int {
	// fmt.Println(this.m)
	if node, ok := this.m[key]; !ok {
		return -1
	} else {
		this.moveToTail(node)
		return node.val
	}
}

func (this *LRUCache) Put(key int, value int) {
	var node *Node
	var ok bool
	// fmt.Println(this.m,key)
	if node, ok = this.m[key]; !ok {
		node = &Node{
			key:  key,
			val:  value,
			pre:  nil,
			next: nil,
		}
	} else {
		node.val = value
	}
	this.moveToTail(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 *
 */
func main() {

}
