package dsa_test

import (
	"testing"

	"github.com/vedantwankhade/katas/dsa"
)

func TestDeque_Empty(t *testing.T) {
	d := dsa.NewDeque[int]()

	if d.Len() != 0 {
		t.Fatalf("expected empty deque, got len=%d", d.Len())
	}

	if _, ok := d.PopFront(); ok {
		t.Fatal("PopFront on empty deque should fail")
	}

	if _, ok := d.PopBack(); ok {
		t.Fatal("PopBack on empty deque should fail")
	}

	if _, ok := d.GetFront(); ok {
		t.Fatal("GetFront on empty deque should fail")
	}

	if _, ok := d.GetBack(); ok {
		t.Fatal("GetBack on empty deque should fail")
	}
}

func TestDeque_PushFrontBack(t *testing.T) {
	d := dsa.NewDeque[int]()

	d.PushBack(2)
	d.PushFront(1)
	d.PushBack(3)

	assertDeque(t, d, []int{1, 2, 3})
}

func TestDeque_PopFront(t *testing.T) {
	d := dsa.NewDeque[int]()

	d.PushBack(1)
	d.PushBack(2)

	v, ok := d.PopFront()
	if !ok || v != 1 {
		t.Fatalf("expected PopFront=1, got %d (ok=%v)", v, ok)
	}

	assertDeque(t, d, []int{2})
}

func TestDeque_PopBack(t *testing.T) {
	d := dsa.NewDeque[int]()

	d.PushBack(1)
	d.PushBack(2)

	v, ok := d.PopBack()
	if !ok || v != 2 {
		t.Fatalf("expected PopBack=2, got %d (ok=%v)", v, ok)
	}

	assertDeque(t, d, []int{1})
}

func TestDeque_GettersDoNotMutate(t *testing.T) {
	d := dsa.NewDeque[int]()
	d.PushBack(1)
	d.PushBack(2)

	f, _ := d.GetFront()
	b, _ := d.GetBack()

	if f != 1 || b != 2 {
		t.Fatalf("expected front=1 back=2, got %d %d", f, b)
	}

	if d.Len() != 2 {
		t.Fatal("GetFront/GetBack should not mutate deque")
	}
}

func TestDeque_MixedOperations(t *testing.T) {
	d := dsa.NewDeque[int]()

	d.PushFront(2)
	d.PushFront(1)
	d.PushBack(3)
	d.PushBack(4)

	v, _ := d.PopFront()
	if v != 1 {
		t.Fatalf("expected 1, got %d", v)
	}

	v, _ = d.PopBack()
	if v != 4 {
		t.Fatalf("expected 4, got %d", v)
	}

	assertDeque(t, d, []int{2, 3})
}

func TestDeque_Stress(t *testing.T) {
	d := dsa.NewDeque[int]()

	const n = 50_000
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			d.PushFront(i)
		} else {
			d.PushBack(i)
		}
	}

	count := 0
	for d.Len() > 0 {
		if count%2 == 0 {
			_, _ = d.PopFront()
		} else {
			_, _ = d.PopBack()
		}
		count++
	}

	if d.Len() != 0 {
		t.Fatal("deque should be empty after stress test")
	}
}

func assertDeque(t *testing.T, d dsa.Deque[int], expected []int) {
	t.Helper()

	if d.Len() != len(expected) {
		t.Fatalf("expected len=%d, got %d", len(expected), d.Len())
	}

	for i, v := range expected {
		var got int
		var ok bool

		if i == 0 {
			got, ok = d.GetFront()
		} else if i == len(expected)-1 {
			got, ok = d.GetBack()
		} else {
			continue
		}

		if !ok || got != v {
			t.Fatalf("expected %d, got %d (ok=%v)", v, got, ok)
		}
	}
}
