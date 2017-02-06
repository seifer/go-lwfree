package lwfree

import (
	"sync"
	"testing"
)

type StackLock struct {
	sync.Mutex
	head *stackLockNode
}

type stackLockNode struct {
	val  interface{}
	prev *stackLockNode
}

func NewStackLock() *StackLock {
	return &StackLock{}
}

func (s *StackLock) Push(x interface{}) {
	s.Lock()
	s.head = &stackLockNode{x, s.head}
	s.Unlock()
}

func (s *StackLock) Pop() (x interface{}) {
	s.Lock()
	if s.head != nil {
		x = s.head.val
		s.head = s.head.prev
	}
	s.Unlock()

	return
}

func TestStackLock(t *testing.T) {
	s := NewStackLock()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if v := s.Pop(); v == nil || v.(int) != 3 {
		t.FailNow()
	}

	if v := s.Pop(); v == nil || v.(int) != 2 {
		t.FailNow()
	}

	if v := s.Pop(); v == nil || v.(int) != 1 {
		t.FailNow()
	}

	if v := s.Pop(); v != nil {
		t.FailNow()
	}
}

func TestStackLockFree(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if v := s.Pop(); v == nil || v.(int) != 3 {
		t.FailNow()
	}

	if v := s.Pop(); v == nil || v.(int) != 2 {
		t.FailNow()
	}

	if v := s.Pop(); v == nil || v.(int) != 1 {
		t.FailNow()
	}

	if v := s.Pop(); v != nil {
		t.FailNow()
	}
}

func BenchmarkStackLockIn75Out25(b *testing.B) {
	s := NewStackLock()
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			s.Push(i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			s.Push(i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			s.Push(i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			s.Pop()
		}
	}()
	b.ResetTimer()
	wg.Wait()
}

func BenchmarkStackLockFreeIn75Out25(b *testing.B) {
	s := NewStack()
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			s.Push(i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			s.Push(i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			s.Push(i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			s.Pop()
		}
	}()
	b.ResetTimer()
	wg.Wait()
}

func BenchmarkStackLockIn50Out50(b *testing.B) {
	s := NewStackLock()
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			s.Push(i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			s.Push(i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			s.Pop()
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			s.Pop()
		}
	}()
	b.ResetTimer()
	wg.Wait()
}

func BenchmarkStackLockFreeIn50Out50(b *testing.B) {
	s := NewStack()
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			s.Push(i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			s.Push(i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			s.Pop()
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			s.Pop()
		}
	}()
	b.ResetTimer()
	wg.Wait()
}

func BenchmarkStackLockIn25Out75(b *testing.B) {
	s := NewStackLock()
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			s.Push(i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			s.Pop()
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			s.Pop()
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			s.Pop()
		}
	}()
	b.ResetTimer()
	wg.Wait()
}

func BenchmarkStackLockFreeIn25Out75(b *testing.B) {
	s := NewStack()
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			s.Push(i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			s.Pop()
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			s.Pop()
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			s.Pop()
		}
	}()
	b.ResetTimer()
	wg.Wait()
}
