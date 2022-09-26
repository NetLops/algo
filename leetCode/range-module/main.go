package main

type RangeModule struct {
	root *Node
}

func Constructor() RangeModule {
	return RangeModule{new(Node)}
}

func (this *RangeModule) AddRange(left int, right int) {
	this.root.update(1, 1e9, left, right-1, true)
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	return this.root.query(1, 1e9, left, right-1)
}

func (this *RangeModule) RemoveRange(left int, right int) {
	this.root.update(1, 1e9, left, right-1, false)
}

type Node struct {
	lc, rc   *Node
	Val, Add bool
}

func (t *Node) update(L, R, left, right int, v bool) {
	if left <= L && R <= right {
		t.Val, t.Add = v, true
		return
	}
	t.pushdown()
	mid := (L + R) >> 1
	if left <= mid {
		t.lc.update(L, mid, left, right, v)
	}
	if right > mid {
		t.rc.update(mid+1, R, left, right, v)
	}
	t.pushup()
}

func (t *Node) query(L, R, left, right int) bool {
	if left <= L && R <= right {
		return t.Val
	} else if t.Val {
		return t.Val
	}
	t.pushdown()
	mid, ans := (L+R)>>1, true
	if left <= mid {
		ans = ans && t.lc.query(L, mid, left, right)
	}
	if right > mid {
		ans = ans && t.rc.query(mid+1, R, left, right)
	}
	return ans
}

func (t *Node) pushup() { t.Val = t.lc.Val && t.rc.Val }

func (t *Node) pushdown() {
	if t.lc == nil {
		t.lc = new(Node)
	}
	if t.rc == nil {
		t.rc = new(Node)
	}
	if t.Add {
		t.lc.Val = t.Val
		t.lc.Add = true
		t.rc.Val = t.Val
		t.rc.Add = true
		t.Add = false
	}
}

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */
func main() {

}
