/*
Sequence is a collection of items that has some order.
*/
package sequence

import "iter"

type Sequence[T any] interface {
	Add(t T)
	AddAt(index int, t T) bool
	Len() int
	Get(index int) (T, bool)
	Set(index int, t T) bool
	Iter() iter.Seq2[int, T]
	String() string
	Remove(index int) (T, bool)
	// Reverse(func(T, T) int)
}
