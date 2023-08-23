package main

type Node struct {
	key, val  int
	pre, next *Node
}

type LRUCache struct {
	head, tail *Node
	capacity   int
	n          int // 长度
	m          map[int]*Node
}

func Constructor(capacity int) LRUCache {
	temp := new(Node)
	temp.next = temp
	temp.pre = temp
	return LRUCache{
		head:     temp,
		tail:     temp,
		capacity: capacity,
		n:        0,
		m:        map[int]*Node{},
	}
}

func (this *LRUCache) Get(key int) int {
	// 出现了的话 就需要移动一手
	if node, exists := this.m[key]; exists {
		// 先将节点从原来的位置删除掉
		node.pre.next = node.next
		node.next.pre = node.pre
		// 然后再插入到链表的头结点
		node.next = this.head.next
		node.pre = this.head
		// 头插一手
		this.head.next.pre = node
		this.head.next = node
		return node.val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	// head -> next => head -> node -> next
	if node, exists := this.m[key]; exists {
		node.val = value
		// 先将节点从原来的位置删除掉
		node.pre.next = node.next
		node.next.pre = node.pre
		// 然后再插入到链表的头结点
		node.next = this.head.next
		node.pre = this.head
		// 头插一手
		this.head.next.pre = node
		this.head.next = node
		return
	}
	this.n++
	if this.n > this.capacity {
		// 从map中删除元素
		oldKey := this.tail.pre.key
		delete(this.m, oldKey)
		// 获取尾结点
		node := this.tail.pre
		// 删除尾结点
		node.pre.next = node.next
		node.next.pre = node.pre

		this.n--
	}

	newNode := new(Node)
	this.m[key] = newNode
	newNode.val = value
	newNode.key = key

	//将新节点插入到头结点
	newNode.pre = this.head
	newNode.next = this.head.next
	this.head.next.pre = newNode
	this.head.next = newNode
}
