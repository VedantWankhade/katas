package sequence_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/vedantwankhade/katas/dsa/sequence"
)

func benchmarkAdd(b *testing.B, newSeq func() sequence.Sequence[int]) {
	seq := newSeq()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		seq.Add(i)
	}
}

func benchmarkAddAtHead(b *testing.B, newSeq func() sequence.Sequence[int]) {
	seq := newSeq()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		seq.AddAt(0, i)
	}
}

func benchmarkAddAtMiddle(b *testing.B, newSeq func() sequence.Sequence[int]) {
	seq := newSeq()
	for i := 0; i < 1_000; i++ {
		seq.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		seq.AddAt(seq.Len()/2, i)
	}
}

func benchmarkGetRandom(b *testing.B, newSeq func() sequence.Sequence[int]) {
	seq := newSeq()
	const size = 10_000
	for i := 0; i < size; i++ {
		seq.Add(i)
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = seq.Get(rng.Intn(size))
	}
}

func benchmarkIter(b *testing.B, newSeq func() sequence.Sequence[int]) {
	seq := newSeq()
	const size = 10_000
	for i := 0; i < size; i++ {
		seq.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for range seq.Iter() {
		}
	}
}

// ---------- BENCHMARKS ----------

func BenchmarkArray_Add(b *testing.B) {
	benchmarkAdd(b, newIntSequence)
}

func BenchmarkList_Add(b *testing.B) {
	benchmarkAdd(b, newIntLinkedList)
}

func BenchmarkArray_AddAtHead(b *testing.B) {
	benchmarkAddAtHead(b, newIntSequence)
}

func BenchmarkList_AddAtHead(b *testing.B) {
	benchmarkAddAtHead(b, newIntLinkedList)
}

func BenchmarkArray_AddAtMiddle(b *testing.B) {
	benchmarkAddAtMiddle(b, newIntSequence)
}

func BenchmarkList_AddAtMiddle(b *testing.B) {
	benchmarkAddAtMiddle(b, newIntLinkedList)
}

func BenchmarkArray_GetRandom(b *testing.B) {
	benchmarkGetRandom(b, newIntSequence)
}

func BenchmarkList_GetRandom(b *testing.B) {
	benchmarkGetRandom(b, newIntLinkedList)
}

func BenchmarkArray_Iter(b *testing.B) {
	benchmarkIter(b, newIntSequence)
}

func BenchmarkList_Iter(b *testing.B) {
	benchmarkIter(b, newIntLinkedList)
}
