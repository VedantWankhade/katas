package dsa_test

import (
	"testing"

	"github.com/vedantwankhade/katas/dsa"
)

func TestQueue_Empty(t *testing.T) {
	q := dsa.NewQueue[int]()

	if q.Len() != 0 {
		t.Fatalf("expected empty queue, got len=%d", q.Len())
	}

	if _, ok := q.PopFront(); ok {
		t.Fatal("PopFront on empty queue should fail")
	}

	if _, ok := q.GetFront(); ok {
		t.Fatal("GetFront on empty queue should fail")
	}
}

func TestQueue_PushPop_FIFO(t *testing.T) {
	q := dsa.NewQueue[int]()

	values := []int{1, 2, 3, 4}
	for _, v := range values {
		q.PushBack(v)
	}

	for _, expected := range values {
		v, ok := q.PopFront()
		if !ok || v != expected {
			t.Fatalf("expected %d, got %d (ok=%v)", expected, v, ok)
		}
	}

	if q.Len() != 0 {
		t.Fatalf("expected empty queue, got len=%d", q.Len())
	}
}

func TestQueue_GetFront(t *testing.T) {
	q := dsa.NewQueue[int]()
	q.PushBack(10)
	q.PushBack(20)

	v, ok := q.GetFront()
	if !ok || v != 10 {
		t.Fatalf("expected front=10, got %d (ok=%v)", v, ok)
	}

	// GetFront must not remove
	if q.Len() != 2 {
		t.Fatal("GetFront should not modify queue")
	}
}

func TestQueue_InterleavedOps(t *testing.T) {
	q := dsa.NewQueue[int]()

	q.PushBack(1)
	q.PushBack(2)

	v, _ := q.PopFront()
	if v != 1 {
		t.Fatalf("expected 1, got %d", v)
	}

	q.PushBack(3)

	v, _ = q.PopFront()
	if v != 2 {
		t.Fatalf("expected 2, got %d", v)
	}

	v, _ = q.PopFront()
	if v != 3 {
		t.Fatalf("expected 3, got %d", v)
	}
}

func TestQueue_LenIntegrity(t *testing.T) {
	q := dsa.NewQueue[int]()

	for i := 0; i < 100; i++ {
		q.PushBack(i)
		if q.Len() != i+1 {
			t.Fatalf("after push: expected len=%d, got %d", i+1, q.Len())
		}
	}

	for i := 99; i >= 0; i-- {
		_, _ = q.PopFront()
		if q.Len() != i {
			t.Fatalf("after pop: expected len=%d, got %d", i, q.Len())
		}
	}
}
