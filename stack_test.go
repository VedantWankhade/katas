package dsa_test

import (
	"testing"

	katas "github.com/vedantwankhade/katas/dsa"
)

func newIntStack() katas.Stack[int] {
	return katas.NewStack[int]()
}

// ---------- TESTS ----------

func TestStack_Empty(t *testing.T) {
	s := newIntStack()

	if s.Len() != 0 {
		t.Fatalf("expected empty stack, got len=%d", s.Len())
	}

	if _, ok := s.Pop(); ok {
		t.Fatal("Pop on empty stack should fail")
	}

	if _, ok := s.Peek(); ok {
		t.Fatal("Peek on empty stack should fail")
	}
}

func TestStack_PushPeek(t *testing.T) {
	s := newIntStack()

	s.Push(10)
	s.Push(20)

	if s.Len() != 2 {
		t.Fatalf("expected len=2, got %d", s.Len())
	}

	v, ok := s.Peek()
	if !ok || v != 20 {
		t.Fatalf("Peek: expected 20, got %d (ok=%v)", v, ok)
	}

	// Peek must not remove
	if s.Len() != 2 {
		t.Fatal("Peek should not modify stack")
	}
}

func TestStack_PushPop_LIFO(t *testing.T) {
	s := newIntStack()

	values := []int{1, 2, 3, 4}
	for _, v := range values {
		s.Push(v)
	}

	for i := len(values) - 1; i >= 0; i-- {
		v, ok := s.Pop()
		if !ok || v != values[i] {
			t.Fatalf("Pop: expected %d, got %d (ok=%v)", values[i], v, ok)
		}
	}

	if s.Len() != 0 {
		t.Fatalf("expected empty stack after pops, got len=%d", s.Len())
	}
}

func TestStack_InterleavedOps(t *testing.T) {
	s := newIntStack()

	s.Push(1)
	s.Push(2)

	v, _ := s.Pop()
	if v != 2 {
		t.Fatalf("expected 2, got %d", v)
	}

	s.Push(3)

	v, _ = s.Pop()
	if v != 3 {
		t.Fatalf("expected 3, got %d", v)
	}

	v, _ = s.Pop()
	if v != 1 {
		t.Fatalf("expected 1, got %d", v)
	}

	if _, ok := s.Pop(); ok {
		t.Fatal("Pop on empty stack should fail")
	}
}

func TestStack_LenIntegrity(t *testing.T) {
	s := newIntStack()

	for i := 0; i < 100; i++ {
		s.Push(i)
		if s.Len() != i+1 {
			t.Fatalf("after push %d: expected len=%d, got %d", i, i+1, s.Len())
		}
	}

	for i := 99; i >= 0; i-- {
		_, _ = s.Pop()
		if s.Len() != i {
			t.Fatalf("after pop: expected len=%d, got %d", i, s.Len())
		}
	}
}

func TestStack_Stress(t *testing.T) {
	s := newIntStack()

	const n = 100_000
	for i := 0; i < n; i++ {
		s.Push(i)
	}

	if s.Len() != n {
		t.Fatalf("expected len=%d, got %d", n, s.Len())
	}

	for i := n - 1; i >= 0; i-- {
		v, ok := s.Pop()
		if !ok || v != i {
			t.Fatalf("expected %d, got %d (ok=%v)", i, v, ok)
		}
	}
}
