package dsa_test

import (
	"strings"
	"testing"

	"github.com/vedantwankhade/katas/dsa"
)

func newIntSequence() dsa.Sequence[int] {
	return dsa.NewDynamicArray[int](8)
}

func TestAddAndLen(t *testing.T) {
	seq := newIntSequence()

	if seq.Len() != 0 {
		t.Fatalf("expected empty sequence, got len=%d", seq.Len())
	}

	seq.Add(10)
	seq.Add(20)

	if seq.Len() != 2 {
		t.Fatalf("expected len=2, got %d", seq.Len())
	}
}

func TestGet(t *testing.T) {
	seq := newIntSequence()
	seq.Add(1)
	seq.Add(2)

	tests := []struct {
		index    int
		expected int
		ok       bool
	}{
		{0, 1, true},
		{1, 2, true},
		{-1, 0, false},
		{2, 0, false},
	}

	for _, tt := range tests {
		v, ok := seq.Get(tt.index)
		if ok != tt.ok {
			t.Fatalf("Get(%d): expected ok=%v, got %v", tt.index, tt.ok, ok)
		}
		if ok && v != tt.expected {
			t.Fatalf("Get(%d): expected %d, got %d", tt.index, tt.expected, v)
		}
	}
}

func TestSet(t *testing.T) {
	seq := newIntSequence()
	seq.Add(5)

	if !seq.Set(0, 42) {
		t.Fatal("Set valid index returned false")
	}

	v, ok := seq.Get(0)
	if !ok || v != 42 {
		t.Fatalf("expected value=42, got %d (ok=%v)", v, ok)
	}

	if seq.Set(1, 99) {
		t.Fatal("Set out-of-bounds index should return false")
	}
}

func TestAddAt(t *testing.T) {
	seq := newIntSequence()
	seq.Add(1)
	seq.Add(3)

	if !seq.AddAt(1, 2) {
		t.Fatal("AddAt valid index returned false")
	}

	expected := []int{1, 2, 3}
	for i, exp := range expected {
		v, ok := seq.Get(i)
		if !ok || v != exp {
			t.Fatalf("index %d: expected %d, got %d (ok=%v)", i, exp, v, ok)
		}
	}

	if seq.AddAt(-1, 0) {
		t.Fatal("AddAt with negative index should fail")
	}

	if seq.AddAt(10, 0) {
		t.Fatal("AddAt out-of-range index should fail")
	}
}

func TestIter(t *testing.T) {
	seq := newIntSequence()
	values := []int{1, 2, 3, 4}
	for _, v := range values {
		seq.Add(v)
	}

	var result []int
	for _, v := range seq.Iter() {
		result = append(result, v)
	}

	if len(result) != len(values) {
		t.Fatalf("expected %d items, got %d", len(values), len(result))
	}

	for i := range values {
		if result[i] != values[i] {
			t.Fatalf("index %d: expected %d, got %d", i, values[i], result[i])
		}
	}
}

func TestString(t *testing.T) {
	seq := newIntSequence()
	seq.Add(1)
	seq.Add(2)

	s := seq.String()
	if s == "" {
		t.Fatal("String() returned empty string")
	}

	// Loose validation: should at least contain values
	if !strings.Contains(s, "1") || !strings.Contains(s, "2") {
		t.Fatalf("String() output does not contain elements: %q", s)
	}
}

func TestGrowthStress(t *testing.T) {
	seq := newIntSequence()

	const n = 10_000
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
