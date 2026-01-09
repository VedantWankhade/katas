package sequence_test

import (
	"testing"

	"github.com/vedantwankhade/katas/dsa/sequence"
)

// ---------- HELPERS ----------

func runRemoveTests(t *testing.T, name string, newSeq func() sequence.Sequence[int]) {
	t.Run(name, func(t *testing.T) {

		t.Run("RemoveFromEmpty", func(t *testing.T) {
			seq := newSeq()
			if _, ok := seq.Remove(0); ok {
				t.Fatal("expected remove from empty to fail")
			}
		})

		t.Run("RemoveInvalidIndex", func(t *testing.T) {
			seq := newSeq()
			seq.Add(1)

			if _, ok := seq.Remove(-1); ok {
				t.Fatal("negative index should fail")
			}
			if _, ok := seq.Remove(1); ok {
				t.Fatal("out-of-range index should fail")
			}
		})

		t.Run("RemoveHead", func(t *testing.T) {
			seq := newSeq()
			seq.Add(1)
			seq.Add(2)
			seq.Add(3)

			v, ok := seq.Remove(0)
			if !ok || v != 1 {
				t.Fatalf("expected to remove 1, got %d (ok=%v)", v, ok)
			}

			assertSequence(t, seq, []int{2, 3})
		})

		t.Run("RemoveMiddle", func(t *testing.T) {
			seq := newSeq()
			seq.Add(1)
			seq.Add(2)
			seq.Add(3)

			v, ok := seq.Remove(1)
			if !ok || v != 2 {
				t.Fatalf("expected to remove 2, got %d (ok=%v)", v, ok)
			}

			assertSequence(t, seq, []int{1, 3})
		})

		t.Run("RemoveTail", func(t *testing.T) {
			seq := newSeq()
			seq.Add(1)
			seq.Add(2)
			seq.Add(3)

			v, ok := seq.Remove(2)
			if !ok || v != 3 {
				t.Fatalf("expected to remove 3, got %d (ok=%v)", v, ok)
			}

			assertSequence(t, seq, []int{1, 2})
		})

		t.Run("RemoveUntilEmpty", func(t *testing.T) {
			seq := newSeq()
			for i := 0; i < 5; i++ {
				seq.Add(i)
			}

			for seq.Len() > 0 {
				_, ok := seq.Remove(0)
				if !ok {
					t.Fatal("expected successful remove")
				}
			}

			if seq.Len() != 0 {
				t.Fatalf("expected empty sequence, got len=%d", seq.Len())
			}
		})
	})
}

// ---------- ASSERT HELPERS ----------

func assertSequence(t *testing.T, seq sequence.Sequence[int], expected []int) {
	t.Helper()

	if seq.Len() != len(expected) {
		t.Fatalf("expected len=%d, got %d", len(expected), seq.Len())
	}

	for i, exp := range expected {
		v, ok := seq.Get(i)
		if !ok || v != exp {
			t.Fatalf("index %d: expected %d, got %d (ok=%v)", i, exp, v, ok)
		}
	}
}

// ---------- ENTRY POINT ----------

func TestRemove_Array(t *testing.T) {
	runRemoveTests(t, "DynamicArray", newIntSequence)
}

func TestRemove_LinkedList(t *testing.T) {
	runRemoveTests(t, "LinkedList", newIntLinkedList)
}
