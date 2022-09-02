package head

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	h := &IntHeap{2, 1, 5}
	fmt.Printf("%p\n", h)
	h.Len()
	heap.Init(h)
	heap.Push(h, 3)
	t.Logf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		t.Logf("%d ", heap.Pop(h))
	}
	t.Log()
}

func f() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func TestClosure(t *testing.T) {
	b := f()
	b()
	b()
	b()
	i := b()
	fmt.Println(i)
}
