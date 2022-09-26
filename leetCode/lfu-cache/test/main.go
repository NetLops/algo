package main

type Item struct {
	key, value  int
	left, right *Item
}

type Bucket struct {
	left, right *Bucket
	idx         int   // 调用次数
	head, tail  *Item // 初始化 占坑 可以让移动 Item 轻松点
	m           map[int]*Item
}

func (bucket *Bucket) Put(key, value int) {
	var item *Item
	var ok bool
	if item, ok = bucket.m[key]; ok {
		// 已经存在更新值
		item.value = value
		// 在原来的双向链表位置中删除
		item.left.right = item.right
		item.right.left = item.left
	} else {
		item = &Item{
			key:   key,
			value: value,
		}
		// 添加到哈希表中
		bucket.m[key] = item
	}
	//增加到双向链表头部
	item.right = bucket.head.right
	item.left = bucket.head
	bucket.head.right.left = item
	bucket.head.right = item
}

func (Bucket *Bucket) Remove(key int) *Item {
	var item *Item
	var isExists bool
	if item, isExists = Bucket.m[key]; isExists {
		// 从双向链表中删除
		item.left.right = item.right
		item.right.left = item.left
		// 从哈希表中删除
		delete(Bucket.m, key)
		return item
	}
	return nil
}

func (bucket *Bucket) Clear() *Item {
	// 双向链表找到待删除的节点
	item := bucket.tail.left
	item.left.right = item.right
	item.right.left = item.left
	// 从哈希表中删除
	delete(bucket.m, item.key)
	return item
}

// 判断是否为空，然后删除它
func (bucket *Bucket) isEmpty() bool {
	return len(bucket.m) == 0
}

func NewBucket(idx int) Bucket {
	bucket := Bucket{
		idx: idx,
		head: &Item{
			key:   -1,
			value: -1,
			left:  nil,
			right: nil,
		},
		tail: &Item{
			key:   -1,
			value: -1,
			left:  nil,
			right: nil,
		},
		m: map[int]*Item{},
	}
	bucket.tail.left = bucket.head
	bucket.head.right = bucket.tail
	return bucket
}

type LFUCache struct {
	m          map[int]*Bucket
	head, tail *Bucket
	capacity   int // 容量
	cnt        int // 计数器
}

func Constructor(capacity int) LFUCache {
	head := NewBucket(-1)
	tail := NewBucket(-1)
	head.right = &tail
	tail.left = &head
	return LFUCache{
		capacity: capacity,
		cnt:      0,
		head:     &head,
		tail:     &tail,
		m:        map[int]*Bucket{},
	}
}

func (this *LFUCache) Get(key int) int {
	var cur *Bucket
	var isExsits bool
	if cur, isExsits = this.m[key]; isExsits {
		var target *Bucket
		// 看看能不能找到将要移动的桶
		if cur.right.idx != cur.idx+1 {
			// 目标桶缺失
			bucket := NewBucket(cur.idx + 1)
			target = &bucket
			target.right = cur.right
			target.left = cur
			cur.right.left = target
			cur.right = target
		} else {
			target = cur.right
		}
		// 将对前键值对才能够前桶移除，并加入到新的桶中
		remove := cur.Remove(key)
		target.Put(remove.key, remove.value)
		// 更新键值对
		this.m[key] = target

		// 如果在移除当前键值对后，当前桶为空，则将当前桶销毁（确保当前空间是O(n)的）
		// 也确保调用编号最小的桶的clear 方法，能够有效的移除一个元素
		this.deleteIfEmpty(cur)
		return remove.value
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 { // 有k可能capacity为空
		return
	}
	var cur *Bucket
	var isExists bool
	if cur, isExists = this.m[key]; isExists {
		// 元素已经存在直接修改值
		cur.Put(key, value)
		this.Get(key) // 调用一下key, 实现 调用次数 + 1
	} else {
		// 容器满了 需要删除元素
		if this.cnt == this.capacity {
			// 从第一个桶[编号最小]中进行清除	// 此时的cur 是使用次数最小的桶
			cur = this.head.right
			item := cur.Clear()      // 删除使用次数最少的桶，并且访问次数也是最小的item
			delete(this.m, item.key) // 外层的m,也要删除
			this.cnt--
			// 如果在移除到键值对后，档案值为空，则将当前桶删除（确保空间是 O(n)的）
			// 也确保调用后最小的桶的clear 方法，能够有效移除掉一个元素
			this.deleteIfEmpty(cur)
		}
		// 需要将当前键值对增加到1号桶
		var first *Bucket

		// 1 号桶不存在则创建
		if this.head.right.idx != 1 {
			Bucket := NewBucket(1)
			first = &Bucket
			first.right = this.head.right
			first.left = this.head
			this.head.right.left = first
			this.head.right = first
		} else {
			first = this.head.right
		}

		// 将键值对添加到1号桶
		first.Put(key, value)
		// 更新键值对所在的信息
		this.m[key] = first
		// 计数器加1
		this.cnt++
	}
}

func (this *LFUCache) deleteIfEmpty(cur *Bucket) {
	if cur.isEmpty() {
		cur.left.right = cur.right
		cur.right.left = cur.left
		cur = nil // GC happy
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
func main() {

}
