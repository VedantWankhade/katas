package sequence_test

import (
	"testing"

	"github.com/vedantwankhade/katas/dsa/sequence"
)

func newIntLinkedList() sequence.Sequence[int] {
	return sequence.NewLinkedList[int]()
}

// ---------- TESTS ----------

func TestLinkedListAddAndLen(t *testing.T) {
	seq := newIntLinkedList()

	for i := 0; i < 5; i++ {
		seq.Add(i)
	}

	if seq.Len() != 5 {
		t.Fatalf("expected len=5, got %d", seq.Len())
	}
}

func TestLinkedListGetTraversal(t *testing.T) {
	seq := newIntLinkedList()
	values := []int{10, 20, 30, 40}

	for _, v := range values {
		seq.Add(v)
	}

	for i, expected := range values {
		v, ok := seq.Get(i)
		if !ok || v != expected {
			t.Fatalf("Get(%d): expected %d, got %d (ok=%v)", i, expected, v, ok)
		}
	}
}

func TestLinkedListSetMiddle(t *testing.T) {
	seq := newIntLinkedList()
	for i := 0; i < 5; i++ {
		seq.Add(i)
	}

	if !seq.Set(2, 99) {
		t.Fatal("Set failed at valid index")
	}

	v, _ := seq.Get(2)
	if v != 99 {
		t.Fatalf("expected value 99 at index 2, got %d", v)
	}
}

func TestLinkedListAddAtHead(t *testing.T) {
	seq := newIntLinkedList()
	seq.Add(2)
	seq.Add(3)

	if !seq.AddAt(0, 1) {
		t.Fatal("AddAt head failed")
	}

	expected := []int{1, 2, 3}
	for i, exp := range expected {
		v, _ := seq.Get(i)
		if v != exp {
			t.Fatalf("index %d: expected %d, got %d", i, exp, v)
		}
	}
}

func TestLinkedListAddAtTail(t *testing.T) {
	seq := newIntLinkedList()
	seq.Add(1)
	seq.Add(2)

	if !seq.AddAt(seq.Len(), 3) {
		t.Fatal("AddAt tail failed")
	}

	v, _ := seq.Get(2)
	if v != 3 {
		t.Fatalf("expected 3 at tail, got %d", v)
	}
}

func TestLinkedListAddAtMiddle(t *testing.T) {
	seq := newIntLinkedList()
	seq.Add(1)
	seq.Add(3)

	if !seq.AddAt(1, 2) {
		t.Fatal("AddAt middle failed")
	}

	expected := []int{1, 2, 3}
	for i, exp := range expected {
		v, _ := seq.Get(i)
		if v != exp {
			t.Fatalf("index %d: expected %d, got %d", i, exp, v)
		}
	}
}

func TestLinkedListInvalidIndexes(t *testing.T) {
	seq := newIntLinkedList()
	seq.Add(1)

	if seq.AddAt(-1, 0) {
		t.Fatal("AddAt negative index should fail")
	}

	if seq.AddAt(2, 0) {
		t.Fatal("AddAt out-of-range index should fail")
	}

	if seq.Set(1, 0) {
		t.Fatal("Set out-of-range index should fail")
	}

	if _, ok := seq.Get(1); ok {
		t.Fatal("Get out-of-range index should fail")
	}
}

func TestLinkedListIterOrder(t *testing.T) {
	seq := newIntLinkedList()
	values := []int{1, 2, 3, 4, 5}

	for _, v := range values {
		seq.Add(v)
	}

	i := 0
	for _, v := range seq.Iter() {
		if v != values[i] {
			t.Fatalf("iteration %d: expected %d, got %d", i, values[i], v)
		}
		i++
	}

	if i != len(values) {
		t.Fatalf("iterator yielded %d values, expected %d", i, len(values))
	}
}

func TestLinkedListStressTraversal(t *testing.T) {
	seq := newIntLinkedList()

	const n = 5_000
	for i := 0; i < n; i++ {
		seq.Add(i)
	}

	if seq.Len() != n {
		t.Fatalf("expected len=%d, got %d", n, seq.Len())
	}

	for i := 0; i < n; i++ {
		v, ok := seq.Get(i)
		if !ok || v != i {
			t.Fatalf("index %d: expected %d, got %d (ok=%v)", i, i, v, ok)
		}
	}
}
