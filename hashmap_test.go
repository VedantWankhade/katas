package dsa_test

import (
	"testing"

	"github.com/vedantwankhade/katas/dsa"
)

func newIntStringMap() *dsa.HashMap[int, string] {
	return dsa.NewHashMap[int, string]()
}

// ---------- TESTS ----------

func TestHashMap_Empty(t *testing.T) {
	hm := newIntStringMap()

	if hm.Size() != 0 {
		t.Fatalf("expected size=0, got %d", hm.Size())
	}

	if _, ok := hm.Get(1); ok {
		t.Fatal("Get on empty map should fail")
	}
}

func TestHashMap_PutGet(t *testing.T) {
	hm := newIntStringMap()

	hm.Put(1, "one")
	hm.Put(2, "two")

	v, ok := hm.Get(1)
	if !ok || v != "one" {
		t.Fatalf("expected (one, true), got (%v, %v)", v, ok)
	}

	v, ok = hm.Get(2)
	if !ok || v != "two" {
		t.Fatalf("expected (two, true), got (%v, %v)", v, ok)
	}

	if hm.Size() != 2 {
		t.Fatalf("expected size=2, got %d", hm.Size())
	}
}

func TestHashMap_PutOverwrite(t *testing.T) {
	hm := newIntStringMap()

	hm.Put(1, "one")
	hm.Put(1, "uno") // overwrite

	v, ok := hm.Get(1)
	if !ok || v != "uno" {
		t.Fatalf("expected overwrite to 'uno', got (%v, %v)", v, ok)
	}

	if hm.Size() != 1 {
		t.Fatalf("expected size=1 after overwrite, got %d", hm.Size())
	}
}

func TestHashMap_Delete(t *testing.T) {
	hm := newIntStringMap()

	hm.Put(1, "one")
	hm.Put(2, "two")

	hm.Delete(1)

	if _, ok := hm.Get(1); ok {
		t.Fatal("expected key 1 to be deleted")
	}

	if hm.Size() != 1 {
		t.Fatalf("expected size=1 after delete, got %d", hm.Size())
	}
}

func TestHashMap_DeleteNonExistent(t *testing.T) {
	hm := newIntStringMap()

	hm.Put(1, "one")
	hm.Delete(2) // should not panic or change size

	if hm.Size() != 1 {
		t.Fatalf("expected size=1, got %d", hm.Size())
	}
}

func TestHashMap_MixedOperations(t *testing.T) {
	hm := newIntStringMap()

	hm.Put(1, "one")
	hm.Put(2, "two")
	hm.Put(3, "three")

	hm.Delete(2)
	hm.Put(4, "four")
	hm.Put(1, "uno") // overwrite

	if hm.Size() != 3 {
		t.Fatalf("expected size=3, got %d", hm.Size())
	}

	tests := map[int]string{
		1: "uno",
		3: "three",
		4: "four",
	}

	for k, expected := range tests {
		v, ok := hm.Get(k)
		if !ok || v != expected {
			t.Fatalf("key %d: expected %s, got (%v, %v)", k, expected, v, ok)
		}
	}

	if _, ok := hm.Get(2); ok {
		t.Fatal("key 2 should be deleted")
	}
}

func TestHashMap_Stress(t *testing.T) {
	hm := newIntStringMap()

	const n = 100_000

	for i := 0; i < n; i++ {
		hm.Put(i, "x")
	}

	if hm.Size() != n {
		t.Fatalf("expected size=%d, got %d", n, hm.Size())
	}

	for i := 0; i < n; i++ {
		v, ok := hm.Get(i)
		if !ok || v != "x" {
			t.Fatalf("key %d missing or incorrect", i)
		}
	}

	for i := 0; i < n; i += 2 {
		hm.Delete(i)
	}

	if hm.Size() != n/2 {
		t.Fatalf("expected size=%d after deletes, got %d", n/2, hm.Size())
	}
}

func TestHashMap_NonIntKey(t *testing.T) {
	hm := dsa.NewHashMap[string, int]()

	hm.Put("a", 1)
	hm.Put("b", 2)

	if v, ok := hm.Get("a"); !ok || v != 1 {
		t.Fatalf("expected (1, true), got (%v, %v)", v, ok)
	}

	hm.Delete("a")

	if _, ok := hm.Get("a"); ok {
		t.Fatal("expected key 'a' to be deleted")
	}
}
