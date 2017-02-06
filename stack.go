package lwfree

import (
	"runtime"
	"sync/atomic"
	"unsafe"
)

// Stack ...
type Stack struct {
	head unsafe.Pointer
}

type stackNode struct {
	val  interface{}
	prev unsafe.Pointer
}

// NewStack create new stack
func NewStack() *Stack {
	return &Stack{}
}

// Push ...
func (s *Stack) Push(v interface{}) {
	n := &stackNode{val: v}

	for i := 0; ; i++ {
		n.prev = atomic.LoadPointer(&s.head)

		if atomic.CompareAndSwapPointer(&s.head, n.prev, unsafe.Pointer(n)) {
			break
		}

		backoff(i)
	}
}

// Pop ...
func (s *Stack) Pop() interface{} {
	var head *stackNode

	for i := 0; ; i++ {
		if head = (*stackNode)(atomic.LoadPointer(&s.head)); head == nil {
			return nil
		}

		if atomic.CompareAndSwapPointer(&s.head, unsafe.Pointer(head), head.prev) {
			return head.val
		}

		backoff(i)
	}

	return nil
}

func backoff(i int) {
	if i&1 != 0 {
		runtime.Gosched()
	} else {
		for j := 0; j < 1000; j++ {
		}
	}
}
